package rbac

import (
	"context"

	"modular-monolith/internal/modules/auth/domain/access"
	"modular-monolith/internal/modules/auth/infra/.gen/models"
	"modular-monolith/internal/modules/auth/infra/.gen/table"
	"modular-monolith/lib/db"
	"modular-monolith/lib/utils"

	pg "github.com/go-jet/jet/v2/postgres"
	"github.com/jmoiron/sqlx"
)

var _ access.AuditRepository = (*AuditRepository)(nil)

type AuditRepository struct {
	db.BaseRepository
}

func NewAuditRepository(DB *sqlx.DB) *AuditRepository {
	return &AuditRepository{
		BaseRepository: db.NewRepo(DB),
	}
}

// Create implements [access.AuditRepository].
func (a *AuditRepository) Create(ctx context.Context, in *access.AuditLog) error {
	al := table.AuditLog
	stmt := al.
		INSERT(
			al.CreatedAt, al.ActorType, al.ActorID, al.ActorRole,
			al.Action, al.Category, al.ResourceType, al.ResourceID,
			al.IPAddress, al.UserAgent, al.SessionID, al.RequestID,
			al.BeforeState, al.AfterState, al.Metadata,
			al.Status, al.ErrorMessage, al.Checksum,
		).
		MODEL(auditModel(in)).
		RETURNING(al.ID)

	return stmt.QueryContext(ctx, a.DB(ctx), &in.ID)
}

// GetByID implements [access.AuditRepository].
func (a *AuditRepository) GetByID(
	ctx context.Context,
	id access.AuditID,
) (*access.AuditLog, error) {
	al := table.AuditLog
	var row models.AuditLog
	if err := al.SELECT(al.AllColumns).WHERE(al.ID.EQ(pg.Int(int64(id)))).
		LIMIT(1).QueryContext(ctx, a.DB(ctx), &row); err != nil {
		if db.IsNoRowError(err) {
			return nil, access.ErrAuditEntryNotFound
		}
		return nil, err
	}
	return toAuditLog(row), nil
}

// List implements [access.AuditRepository].
func (a *AuditRepository) List(
	ctx context.Context,
	filter access.AuditFilter,
	limit, offset int,
) ([]*access.AuditLog, int, error) {
	al := table.AuditLog
	entries := []*access.AuditLog{}

	// Newest first: the read model is a timeline, not a directory listing.
	stmt := al.SELECT(al.AllColumns).
		WHERE(auditFilters(filter)).
		ORDER_BY(al.CreatedAt.DESC(), al.ID.DESC())
	rows, count, err := a.Paginate[[]models.AuditLog](ctx, stmt, limit, offset)
	if err != nil {
		return entries, 0, err
	}
	for _, v := range rows {
		entries = append(entries, toAuditLog(v))
	}
	return entries, count, nil
}

// Count implements [access.AuditRepository].
func (a *AuditRepository) Count(ctx context.Context, filter access.AuditFilter) (int, error) {
	var count int
	stmt := table.AuditLog.
		SELECT(pg.COUNT(pg.Int(1))).
		WHERE(auditFilters(filter))
	if err := a.QueryOne(ctx, stmt, &count); err != nil {
		return 0, err
	}
	return count, nil
}

func auditFilters(filter access.AuditFilter) pg.BoolExpression {
	al := table.AuditLog
	conds := []pg.BoolExpression{
		eqText(al.ActorType, filter.ActorType),
		eqText(al.ActorID, filter.ActorID),
		eqText(al.Action, filter.Action),
		eqText(al.Category, filter.Category),
		eqText(al.ResourceType, filter.ResourceType),
		eqText(al.ResourceID, filter.ResourceID),
		eqText(al.Status, filter.Status),
		eqText(al.RequestID, filter.RequestID),
		eqText(al.SessionID, filter.SessionID),
	}
	if filter.Since != nil {
		conds = append(conds, al.CreatedAt.GT_EQ(pg.TimestampT(*filter.Since)))
	}
	if filter.Until != nil {
		conds = append(conds, al.CreatedAt.LT_EQ(pg.TimestampT(*filter.Until)))
	}
	return pg.AND(conds...)
}

// eqText skips the predicate when the caller left the filter empty, so an
// unset filter string means "any" rather than "equals empty string".
func eqText(col pg.ColumnString, value string) pg.BoolExpression {
	if value == "" {
		return nil
	}
	return col.EQ(pg.String(value))
}

func auditModel(in *access.AuditLog) models.AuditLog {
	return models.AuditLog{
		CreatedAt:    in.CreatedAt,
		ActorType:    in.ActorType,
		ActorID:      utils.Nullable(in.ActorID),
		ActorRole:    utils.Nullable(in.ActorRole),
		Action:       in.Action,
		Category:     in.Category,
		ResourceType: in.ResourceType,
		ResourceID:   utils.Nullable(in.ResourceID),
		IPAddress:    utils.Nullable(in.IPAddress),
		UserAgent:    utils.Nullable(in.UserAgent),
		SessionID:    utils.Nullable(in.SessionID),
		RequestID:    utils.Nullable(in.RequestID),
		BeforeState:  utils.Nullable(in.BeforeState),
		AfterState:   utils.Nullable(in.AfterState),
		Metadata:     utils.Nullable(in.Metadata),
		Status:       in.Status,
		ErrorMessage: utils.Nullable(in.ErrorMessage),
		Checksum:     utils.Nullable(in.Checksum),
	}
}

func toAuditLog(row models.AuditLog) *access.AuditLog {
	return &access.AuditLog{
		ID:           row.ID,
		CreatedAt:    row.CreatedAt,
		ActorType:    row.ActorType,
		ActorID:      utils.Deref(row.ActorID),
		ActorRole:    utils.Deref(row.ActorRole),
		Action:       row.Action,
		Category:     row.Category,
		ResourceType: row.ResourceType,
		ResourceID:   utils.Deref(row.ResourceID),
		IPAddress:    utils.Deref(row.IPAddress),
		UserAgent:    utils.Deref(row.UserAgent),
		SessionID:    utils.Deref(row.SessionID),
		RequestID:    utils.Deref(row.RequestID),
		BeforeState:  utils.Deref(row.BeforeState),
		AfterState:   utils.Deref(row.AfterState),
		Metadata:     utils.Deref(row.Metadata),
		Status:       row.Status,
		ErrorMessage: utils.Deref(row.ErrorMessage),
		Checksum:     utils.Deref(row.Checksum),
	}
}
