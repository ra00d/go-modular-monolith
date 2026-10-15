package services

import (
	"context"
	"strconv"
	"time"

	"modular-monolith/internal/auth/domain/access"
	"modular-monolith/lib/logger"
	"modular-monolith/lib/utils"

	"github.com/google/uuid"
)

// auditRecorder appends to the access audit trail. Every admin mutation in this
// package goes through it so the trail has one shape and one write path.
type auditRecorder struct {
	repo access.AuditRepository
}

func newAuditRecorder(repo access.AuditRepository) auditRecorder {
	return auditRecorder{repo: repo}
}

// Record writes one audit entry. A failure is logged rather than returned: the
// mutation it describes has already committed, and losing it because the audit
// insert hiccuped would be worse than a gap in the trail.
func (a auditRecorder) Record(
	ctx context.Context,
	actor uuid.UUID,
	action, resourceType, resourceID string,
	before any,
) {
	entry := &access.AuditLog{
		CreatedAt:    time.Now(),
		ActorType:    "user",
		ActorID:      actor.String(),
		Action:       action,
		Category:     access.CategoryAdmin,
		ResourceType: resourceType,
		ResourceID:   resourceID,
		Status:       string(access.AuditStatusSuccess),
	}
	if before != nil {
		entry.BeforeState, _ = utils.MarshalJSON(before)
	}
	if err := a.repo.Create(ctx, entry); err != nil {
		logger.Error("audit ", action, " for ", resourceType, " ", resourceID, ": ", err)
	}
}

// RecordFailure writes an audit entry for an action that was refused or failed.
func (a auditRecorder) RecordFailure(
	ctx context.Context,
	actor uuid.UUID,
	action, resourceType, resourceID, reason string,
) {
	entry := &access.AuditLog{
		CreatedAt:    time.Now(),
		ActorType:    "user",
		ActorID:      actor.String(),
		Action:       action,
		Category:     access.CategorySecurity,
		ResourceType: resourceType,
		ResourceID:   resourceID,
		Status:       string(access.AuditStatusDenied),
		ErrorMessage: reason,
	}
	if err := a.repo.Create(ctx, entry); err != nil {
		logger.Error("audit ", action, " for ", resourceType, " ", resourceID, ": ", err)
	}
}

// RecordRaw writes a caller-built entry, for the cases where the shared shape
// does not fit, such as recording why a ban was applied.
func (a auditRecorder) RecordRaw(ctx context.Context, entry *access.AuditLog) {
	if err := a.repo.Create(ctx, entry); err != nil {
		logger.Error("audit ", entry.Action, " for ", entry.ResourceType, ": ", err)
	}
}

func resourceID(id int64) string { return strconv.FormatInt(id, 10) }

// AuditService serves the admin read model over the audit trail and the
// high-priority security events.
type AuditService struct {
	auditRepo  access.AuditRepository
	eventsRepo access.SecurityEventRepository
	auditor    auditRecorder
}

func NewAuditService(
	auditRepo access.AuditRepository,
	eventsRepo access.SecurityEventRepository,
) *AuditService {
	return &AuditService{
		auditRepo:  auditRepo,
		eventsRepo: eventsRepo,
		auditor:    newAuditRecorder(auditRepo),
	}
}

func (s *AuditService) List(
	ctx context.Context,
	filter access.AuditFilter,
	limit, offset int,
) ([]access.AuditEntry, int, error) {
	rows, count, err := s.auditRepo.List(ctx, filter, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	entries := make([]access.AuditEntry, 0, len(rows))
	for _, row := range rows {
		entries = append(entries, access.AuditEntry{
			ID:           row.ID,
			CreatedAt:    row.CreatedAt,
			ActorType:    row.ActorType,
			ActorID:      row.ActorID,
			ActorRole:    row.ActorRole,
			Action:       row.Action,
			Category:     row.Category,
			ResourceType: row.ResourceType,
			ResourceID:   row.ResourceID,
			IPAddress:    row.IPAddress,
			UserAgent:    row.UserAgent,
			RequestID:    row.RequestID,
			SessionID:    row.SessionID,
			BeforeState:  row.BeforeState,
			AfterState:   row.AfterState,
			Metadata:     row.Metadata,
			Status:       access.AuditStatus(row.Status),
			ErrorMessage: row.ErrorMessage,
		})
	}
	return entries, count, nil
}

func (s *AuditService) GetByID(
	ctx context.Context,
	id access.AuditID,
) (access.AuditEntry, error) {
	row, err := s.auditRepo.GetByID(ctx, id)
	if err != nil {
		return access.AuditEntry{}, err
	}
	return access.AuditEntry{
		ID:           row.ID,
		CreatedAt:    row.CreatedAt,
		ActorType:    row.ActorType,
		ActorID:      row.ActorID,
		ActorRole:    row.ActorRole,
		Action:       row.Action,
		Category:     row.Category,
		ResourceType: row.ResourceType,
		ResourceID:   row.ResourceID,
		IPAddress:    row.IPAddress,
		UserAgent:    row.UserAgent,
		RequestID:    row.RequestID,
		SessionID:    row.SessionID,
		BeforeState:  row.BeforeState,
		AfterState:   row.AfterState,
		Metadata:     row.Metadata,
		Status:       access.AuditStatus(row.Status),
		ErrorMessage: row.ErrorMessage,
	}, nil
}

func (s *AuditService) ListEvents(
	ctx context.Context,
	filter access.SecurityEventFilter,
	limit, offset int,
) ([]access.SecurityEvent, int, error) {
	return s.eventsRepo.List(ctx, filter, limit, offset)
}

func (s *AuditService) GetEvent(ctx context.Context, id int64) (access.SecurityEvent, error) {
	return s.eventsRepo.GetByID(ctx, id)
}

func (s *AuditService) ResolveEvent(
	ctx context.Context,
	id int64,
	actor uuid.UUID,
) error {
	if err := s.eventsRepo.Resolve(ctx, id, actor); err != nil {
		return err
	}
	s.auditor.Record(ctx, actor, "security_event.resolve", "security_event", resourceID(id), nil)
	return nil
}
