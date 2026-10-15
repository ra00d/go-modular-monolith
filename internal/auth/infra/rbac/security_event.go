package rbac

import (
	"context"
	"encoding/json"
	"time"

	"modular-monolith/internal/auth/domain/access"
	"modular-monolith/internal/auth/infra/.gen/models"
	"modular-monolith/internal/auth/infra/.gen/table"
	"modular-monolith/lib/db"
	"modular-monolith/lib/utils"

	pg "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

var _ access.SecurityEventRepository = (*SecurityEventRepository)(nil)

type SecurityEventRepository struct {
	db.BaseRepository
}

func NewSecurityEventRepository(DB *sqlx.DB) *SecurityEventRepository {
	return &SecurityEventRepository{
		BaseRepository: db.NewRepo(DB),
	}
}

// Create implements [access.SecurityEventRepository].
func (s *SecurityEventRepository) Create(
	ctx context.Context,
	in *access.SecurityEvent,
) error {
	se := table.SecurityEvent
	details, err := marshalDetails(in.Details)
	if err != nil {
		return err
	}
	row := models.SecurityEvent{
		EventType:  in.EventType,
		UserID:     in.UserID,
		Severity:   in.Severity,
		Details:    details,
		IPAddress:  utils.Nullable(in.IPAddress),
		ResolvedAt: in.ResolvedAt,
		ResolvedBy: in.ResolvedBy,
	}
	stmt := se.
		INSERT(se.EventType, se.UserID, se.Severity, se.Details, se.IPAddress).
		MODEL(row).
		RETURNING(se.ID)

	return stmt.QueryContext(ctx, s.DB(ctx), &in.ID)
}

// GetByID implements [access.SecurityEventRepository].
func (s *SecurityEventRepository) GetByID(
	ctx context.Context,
	id int64,
) (access.SecurityEvent, error) {
	se := table.SecurityEvent
	var row models.SecurityEvent
	if err := se.SELECT(se.AllColumns).WHERE(se.ID.EQ(pg.Int(id))).
		LIMIT(1).QueryContext(ctx, s.DB(ctx), &row); err != nil {
		if db.IsNoRowError(err) {
			return access.SecurityEvent{}, access.ErrSecurityEventNotFound
		}
		return access.SecurityEvent{}, err
	}
	return toSecurityEvent(row)
}

// List implements [access.SecurityEventRepository].
func (s *SecurityEventRepository) List(
	ctx context.Context,
	filter access.SecurityEventFilter,
	limit, offset int,
) ([]access.SecurityEvent, int, error) {
	se := table.SecurityEvent
	events := []access.SecurityEvent{}

	// Newest first: operators triage the most recent events.
	stmt := se.SELECT(se.AllColumns).
		WHERE(securityEventFilters(filter)).
		ORDER_BY(se.CreatedAt.DESC(), se.ID.DESC())
	rows, count, err := s.Paginate[[]models.SecurityEvent](ctx, stmt, limit, offset)
	if err != nil {
		return events, 0, err
	}
	for _, v := range rows {
		event, err := toSecurityEvent(v)
		if err != nil {
			return events, 0, err
		}
		events = append(events, event)
	}
	return events, count, nil
}

// Resolve implements [access.SecurityEventRepository].
func (s *SecurityEventRepository) Resolve(
	ctx context.Context,
	id int64,
	by uuid.UUID,
) error {
	se := table.SecurityEvent
	now := time.Now()
	res, err := se.UPDATE(se.ResolvedAt, se.ResolvedBy).
		MODEL(models.SecurityEvent{ResolvedAt: &now, ResolvedBy: &by}).
		WHERE(pg.AND(
			se.ID.EQ(pg.Int(id)),
			se.ResolvedAt.IS_NULL(),
		)).
		ExecContext(ctx, s.DB(ctx))
	if err != nil {
		return err
	}
	if n, err := res.RowsAffected(); err == nil && n == 0 {
		// Either the event does not exist or it was already triaged.
		if _, err := s.GetByID(ctx, id); err != nil {
			return err
		}
		return access.ErrSecurityEventResolved
	}
	return nil
}

func securityEventFilters(filter access.SecurityEventFilter) pg.BoolExpression {
	se := table.SecurityEvent
	conds := []pg.BoolExpression{eqText(se.EventType, filter.EventType)}
	if filter.Severity != "" {
		conds = append(conds, se.Severity.EQ(pg.String(filter.Severity)))
	}
	if filter.UserID != nil {
		conds = append(conds, se.UserID.EQ(pg.UUID(*filter.UserID)))
	}
	if filter.Unresolved != nil {
		if *filter.Unresolved {
			conds = append(conds, se.ResolvedAt.IS_NULL())
		} else {
			conds = append(conds, se.ResolvedAt.IS_NOT_NULL())
		}
	}
	if filter.Since != nil {
		conds = append(conds, se.CreatedAt.GT_EQ(pg.TimestampT(*filter.Since)))
	}
	if filter.Until != nil {
		conds = append(conds, se.CreatedAt.LT_EQ(pg.TimestampT(*filter.Until)))
	}
	return pg.AND(conds...)
}

func marshalDetails(details map[string]any) (*string, error) {
	if details == nil {
		return nil, nil
	}
	raw, err := json.Marshal(details)
	if err != nil {
		return nil, err
	}
	return utils.Nullable(string(raw)), nil
}

func toSecurityEvent(row models.SecurityEvent) (access.SecurityEvent, error) {
	var details map[string]any
	if raw := utils.Deref(row.Details); raw != "" {
		if err := json.Unmarshal([]byte(raw), &details); err != nil {
			return access.SecurityEvent{}, err
		}
	}
	return access.SecurityEvent{
		ID:         int64(row.ID),
		EventType:  row.EventType,
		UserID:     row.UserID,
		Severity:   row.Severity,
		Details:    details,
		IPAddress:  utils.Deref(row.IPAddress),
		ResolvedAt: row.ResolvedAt,
		ResolvedBy: row.ResolvedBy,
		CreatedAt:  row.CreatedAt,
	}, nil
}
