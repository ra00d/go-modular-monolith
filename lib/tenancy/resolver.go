package tenancy

import (
	"strings"

	"github.com/google/uuid"
)

// Resolve parses a tenant identifier (the X-TENANT-ID header value) into a
// Tenant. It performs format validation only; existence and status checks
// belong to the tenancy module's store once persistence lands.
func Resolve(raw string) (*Tenant, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil, MissingTenantError()
	}
	id, err := uuid.Parse(raw)
	if err != nil {
		return nil, InvalidTenantError(raw)
	}
	return &Tenant{ID: id}, nil
}
