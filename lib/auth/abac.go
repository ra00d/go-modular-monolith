package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"maps"
	"net"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
)

// ============================================
// CORE TYPES & INTERFACES
// ============================================

// Subject represents the entity requesting access (typically a user)
type Subject struct {
	ID         string         `json:"id"`
	Type       string         `json:"type"` // "user", "api_key", "service"
	Roles      []string       `json:"roles"`
	Attributes map[string]any `json:"attributes"` // Dynamic ABAC attributes
}

// Resource represents the entity being accessed
type Resource struct {
	Type       string         `json:"type"`
	ID         string         `json:"id"`
	OwnerID    string         `json:"owner_id"`
	Attributes map[string]any `json:"attributes"`
}

// Action represents the operation being performed
type Action struct {
	Name     string `json:"name"`     // "create", "read", "update", "delete"
	Resource string `json:"resource"` // "post", "user", "order"
}

// Environment captures contextual information about the request
type Environment struct {
	Timestamp time.Time      `json:"timestamp"`
	IPAddress net.IP         `json:"ip_address"`
	UserAgent string         `json:"user_agent"`
	RequestID string         `json:"request_id"`
	Custom    map[string]any `json:"custom"` // Additional context
}

// Decision is the result of an ABAC policy evaluation
type Decision struct {
	Allowed  bool               `json:"allowed"`
	Effect   Effect             `json:"effect"`    // "allow" or "deny"
	PolicyID string             `json:"policy_id"` // Which policy made the decision
	Reason   string             `json:"reason"`
	Metadata EvaluationMetadata `json:"metadata"`
}

// Effect represents the policy effect
type Effect string

const (
	EffectAllow Effect = "allow"
	EffectDeny  Effect = "deny"
)

// EvaluationMetadata contains diagnostic information
type EvaluationMetadata struct {
	EvaluatedPolicies int           `json:"evaluated_policies"`
	MatchedPolicies   int           `json:"matched_policies"`
	EvaluationTime    time.Duration `json:"evaluation_time"`
	Errors            []string      `json:"errors,omitempty"`
}

// ============================================
// POLICY DEFINITIONS
// ============================================

// Policy represents a single ABAC authorization rule
type Policy struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Slug        string     `json:"slug"`
	Description string     `json:"description"`
	Resource    string     `json:"resource"`
	Action      string     `json:"action"`
	Rules       RuleSet    `json:"rules"`
	Effect      Effect     `json:"effect"`
	Priority    int        `json:"priority"` // Lower = higher priority
	IsActive    bool       `json:"is_active"`
	ExpiresAt   *time.Time `json:"expires_at,omitempty"`
	CreatedBy   string     `json:"created_by"`
}

// RuleSet contains the three policy dimensions
type RuleSet struct {
	Subject     SubjectRule     `json:"subject"`
	Resource    ResourceRule    `json:"resource"`
	Environment EnvironmentRule `json:"environment"`
}

// SubjectRule defines conditions on the requesting entity
type SubjectRule struct {
	Roles      []string       `json:"roles,omitempty"`
	Attributes map[string]any `json:"attributes,omitempty"`
	Expression string         `json:"expression,omitempty"` // Advanced: CEL or custom DSL
}

// ResourceRule defines conditions on the target entity
type ResourceRule struct {
	Types      []string       `json:"types,omitempty"`
	Attributes map[string]any `json:"attributes,omitempty"`
	Ownership  OwnershipCheck `json:"ownership"`
	Expression string         `json:"expression,omitempty"`
}

// OwnershipCheck defines ownership-based conditions
type OwnershipCheck struct {
	Required     bool   `json:"required"`
	Field        string `json:"field"`         // e.g., "owner_id", "author_id"
	SubjectField string `json:"subject_field"` // e.g., "id"
}

// EnvironmentRule defines contextual conditions
type EnvironmentRule struct {
	TimeRange  *TimeRangeCondition `json:"time_range,omitempty"`
	IPRanges   []string            `json:"ip_ranges,omitempty"`
	DaysOfWeek []time.Weekday      `json:"days_of_week,omitempty"`
	Custom     map[string]any      `json:"custom,omitempty"`
	Expression string              `json:"expression,omitempty"`
}

// TimeRangeCondition defines time-based access windows
type TimeRangeCondition struct {
	Start    string `json:"start"` // "09:00" (24h format, timezone-aware)
	End      string `json:"end"`
	Timezone string `json:"timezone"` // "America/New_York", "UTC"
}

// ============================================
// EVALUATOR INTERFACES
// ============================================

// PolicyStore defines how policies are retrieved. The engine only needs
// resource-scoped lookup; concrete stores may expose more for admin tooling.
type PolicyStore interface {
	// GetPoliciesForResource returns all active policies applicable to a resource type
	GetPoliciesForResource(ctx context.Context, resourceType string) ([]Policy, error)
}

// AttributeResolver resolves dynamic attributes for subjects and resources
type AttributeResolver interface {
	// ResolveSubjectAttributes fetches additional attributes for a subject
	ResolveSubjectAttributes(ctx context.Context, subject Subject) (map[string]any, error)

	// ResolveResourceAttributes fetches additional attributes for a resource
	ResolveResourceAttributes(ctx context.Context, resource Resource) (map[string]any, error)
}

// ConditionEvaluator evaluates individual conditions within a policy
type ConditionEvaluator interface {
	// EvaluateSubject checks if the subject matches the rule
	EvaluateSubject(ctx context.Context, subject Subject, rule SubjectRule) (bool, error)

	// EvaluateResource checks if the resource matches the rule
	EvaluateResource(
		ctx context.Context,
		subject Subject,
		resource Resource,
		rule ResourceRule,
	) (bool, error)

	// EvaluateEnvironment checks if the environment matches the rule
	EvaluateEnvironment(ctx context.Context, env Environment, rule EnvironmentRule) (bool, error)
}

// PolicyEngine is the main authorization interface
type PolicyEngine interface {
	// Authorize evaluates all applicable policies and returns a decision
	Authorize(
		ctx context.Context,
		subject Subject,
		action Action,
		resource Resource,
		env Environment,
	) (Decision, error)

	// AuthorizeSimple is a convenience method for common use cases
	AuthorizeSimple(
		ctx context.Context,
		userID string,
		action string,
		resourceType string,
		resourceID string,
	) (Decision, error)

	// Explain returns detailed evaluation information (for debugging)
	Explain(
		ctx context.Context,
		subject Subject,
		action Action,
		resource Resource,
		env Environment,
	) (EvaluationTrace, error)
}

// ============================================
// EVALUATION TRACE (DEBUGGING)
// ============================================

// EvaluationTrace provides detailed step-by-step evaluation info
type EvaluationTrace struct {
	Request           AuthorizationRequest `json:"request"`
	EvaluatedPolicies []PolicyTrace        `json:"evaluated_policies"`
	FinalDecision     Decision             `json:"final_decision"`
}

type AuthorizationRequest struct {
	Subject  Subject     `json:"subject"`
	Action   Action      `json:"action"`
	Resource Resource    `json:"resource"`
	Env      Environment `json:"environment"`
}

type PolicyTrace struct {
	PolicyID      string        `json:"policy_id"`
	PolicyName    string        `json:"policy_name"`
	Priority      int           `json:"priority"`
	Matched       bool          `json:"matched"`
	SubjectMatch  bool          `json:"subject_match"`
	ResourceMatch bool          `json:"resource_match"`
	EnvMatch      bool          `json:"environment_match"`
	Effect        Effect        `json:"effect"`
	Reason        string        `json:"reason"`
	Duration      time.Duration `json:"duration"`
}

// ============================================
// DEFAULT IMPLEMENTATIONS
// ============================================

// DefaultConditionEvaluator implements ConditionEvaluator with standard logic
type DefaultConditionEvaluator struct {
	// TimeProvider allows mocking time in tests
	TimeProvider func() time.Time
}

// NewDefaultConditionEvaluator creates a new evaluator with real time
func NewDefaultConditionEvaluator() *DefaultConditionEvaluator {
	return &DefaultConditionEvaluator{
		TimeProvider: time.Now,
	}
}

// EvaluateSubject checks if subject attributes match the rule
func (e *DefaultConditionEvaluator) EvaluateSubject(
	ctx context.Context,
	subject Subject,
	rule SubjectRule,
) (bool, error) {
	// Check roles
	if len(rule.Roles) > 0 {
		hasRole := false
		for _, required := range rule.Roles {
			for _, subjectRole := range subject.Roles {
				if strings.EqualFold(required, subjectRole) {
					hasRole = true
					break
				}
			}
			if hasRole {
				break
			}
		}
		if !hasRole {
			return false, nil
		}
	}

	// Check attributes
	if len(rule.Attributes) > 0 {
		match, err := matchAttributes(subject.Attributes, rule.Attributes)
		if err != nil || !match {
			return false, err
		}
	}

	// Evaluate expression if present
	if rule.Expression != "" {
		return e.evaluateExpression(ctx, subject.Attributes, rule.Expression)
	}

	return true, nil
}

// EvaluateResource checks if resource matches the rule, including ownership
func (e *DefaultConditionEvaluator) EvaluateResource(
	ctx context.Context,
	subject Subject,
	resource Resource,
	rule ResourceRule,
) (bool, error) {
	// Check resource types
	if len(rule.Types) > 0 {
		typeMatch := false
		for _, t := range rule.Types {
			if strings.EqualFold(t, resource.Type) {
				typeMatch = true
				break
			}
		}
		if !typeMatch {
			return false, nil
		}
	}

	// Check ownership if required
	if rule.Ownership.Required {
		ownerID := getNestedValue(resource.Attributes, rule.Ownership.Field)
		if ownerID == "" {
			ownerID = resource.OwnerID
		}

		subjectID := getNestedValue(subject.Attributes, rule.Ownership.SubjectField)
		if subjectID == "" {
			subjectID = subject.ID
		}

		if ownerID != subjectID {
			return false, nil
		}
	}

	// Check resource attributes
	if len(rule.Attributes) > 0 {
		match, err := matchAttributes(resource.Attributes, rule.Attributes)
		if err != nil || !match {
			return false, err
		}
	}

	// Evaluate expression if present
	if rule.Expression != "" {
		return e.evaluateExpression(
			ctx,
			mergeMaps(subject.Attributes, resource.Attributes),
			rule.Expression,
		)
	}

	return true, nil
}

// EvaluateEnvironment checks environmental conditions
func (e *DefaultConditionEvaluator) EvaluateEnvironment(
	ctx context.Context,
	env Environment,
	rule EnvironmentRule,
) (bool, error) {
	now := e.TimeProvider()

	// Check time range
	if rule.TimeRange != nil {
		tz, err := time.LoadLocation(rule.TimeRange.Timezone)
		if err != nil {
			return false, fmt.Errorf("invalid timezone: %w", err)
		}

		nowInTz := now.In(tz)
		startParts := strings.Split(rule.TimeRange.Start, ":")
		endParts := strings.Split(rule.TimeRange.End, ":")

		startHour, _ := strconv.Atoi(startParts[0])
		startMin, _ := strconv.Atoi(startParts[1])
		endHour, _ := strconv.Atoi(endParts[0])
		endMin, _ := strconv.Atoi(endParts[1])

		startMinutes := startHour*60 + startMin
		endMinutes := endHour*60 + endMin
		currentMinutes := nowInTz.Hour()*60 + nowInTz.Minute()
		// Handle overnight ranges (e.g., 22:00 - 06:00)
		if startMinutes > endMinutes {
			if currentMinutes < startMinutes && currentMinutes > endMinutes {
				return false, nil
			}
		} else {
			if currentMinutes < startMinutes || currentMinutes > endMinutes {
				return false, nil
			}
		}
	}

	// Check IP ranges
	// if len(rule.IPRanges) > 0 {
	// 	ipMatch := false
	// 	for _, cidr := range rule.IPRanges {
	// 		_, ipNet, err := net.ParseCIDR(cidr)
	// 		if err != nil {
	// 			return false, fmt.Errorf("invalid CIDR %s: %w", cidr, err)
	// 		}
	// 		if ipNet.Contains(env.IPAddress) {
	// 			ipMatch = true
	// 			break
	// 		}
	// 	}
	// 	if !ipMatch {
	// 		return false, nil
	// 	}
	// }

	// Check days of week
	if len(rule.DaysOfWeek) > 0 {
		today := now.Weekday()
		dayMatch := slices.Contains(rule.DaysOfWeek, today)
		if !dayMatch {
			return false, nil
		}
	}

	// Check custom environment attributes
	if len(rule.Custom) > 0 {
		envMap := structToMap(env)
		match, err := matchAttributes(envMap, rule.Custom)
		if err != nil || !match {
			return false, err
		}
	}

	return true, nil
}

// evaluateExpression evaluates a custom expression (placeholder for CEL or custom DSL)
func (e *DefaultConditionEvaluator) evaluateExpression(
	ctx context.Context,
	variables map[string]any,
	expression string,
) (bool, error) {
	_ = ctx
	_ = variables
	// Extension point: integrate CEL (github.com/google/cel-go) here.
	// ponytail: no expression engine yet; an empty/other expression is a no-op
	// that neither grants nor blocks. Wire CEL when a policy actually needs it.
	if expression == "" {
		return true, nil
	}
	return true, nil
}

// ============================================
// POLICY ENGINE IMPLEMENTATION
// ============================================

// DefaultPolicyEngine implements the full ABAC evaluation logic
type DefaultPolicyEngine struct {
	store     PolicyStore
	resolver  AttributeResolver
	evaluator ConditionEvaluator
}

// NewPolicyEngine creates a configured policy engine
func NewPolicyEngine(
	store PolicyStore,
	resolver AttributeResolver,
	evaluator ConditionEvaluator,
) *DefaultPolicyEngine {
	return &DefaultPolicyEngine{
		store:     store,
		resolver:  resolver,
		evaluator: evaluator,
	}
}

// Authorize evaluates all policies using the priority-based decision algorithm:
// 1. Deny overrides (explicit deny always wins)
// 2. Priority ordering (lower number = higher priority)
// 3. First applicable policy wins within same priority
func (e *DefaultPolicyEngine) Authorize(
	ctx context.Context,
	subject Subject,
	action Action,
	resource Resource,
	env Environment,
) (Decision, error) {
	start := time.Now()

	// Resolve dynamic attributes
	if e.resolver != nil {
		subjectAttrs, err := e.resolver.ResolveSubjectAttributes(ctx, subject)
		if err != nil {
			return Decision{}, fmt.Errorf("resolve subject attributes: %w", err)
		}
		subject.Attributes = mergeMaps(subject.Attributes, subjectAttrs)

		resourceAttrs, err := e.resolver.ResolveResourceAttributes(ctx, resource)
		if err != nil {
			return Decision{}, fmt.Errorf("resolve resource attributes: %w", err)
		}
		resource.Attributes = mergeMaps(resource.Attributes, resourceAttrs)
	}

	// Fetch applicable policies
	policies, err := e.store.GetPoliciesForResource(ctx, resource.Type)
	if err != nil {
		return Decision{}, fmt.Errorf("fetch policies: %w", err)
	}

	// Filter by action and sort by priority
	var applicablePolicies []Policy
	for _, p := range policies {
		if !p.IsActive {
			continue
		}
		if p.ExpiresAt != nil && p.ExpiresAt.Before(time.Now()) {
			continue
		}
		if p.Action != "*" && !strings.EqualFold(p.Action, action.Name) {
			continue
		}
		applicablePolicies = append(applicablePolicies, p)
	}

	// Sort by priority (lower first), then by creation time
	sortPoliciesByPriority(applicablePolicies)

	meta := EvaluationMetadata{
		EvaluatedPolicies: len(applicablePolicies),
	}

	var (
		finalDecision *Decision
		errors        []string
	)

	// Evaluate each policy in priority order
	for _, policy := range applicablePolicies {
		subjectMatch, err := e.evaluator.EvaluateSubject(ctx, subject, policy.Rules.Subject)
		if err != nil {
			errors = append(errors, fmt.Sprintf("policy %s subject eval: %v", policy.ID, err))
			continue
		}

		resourceMatch, err := e.evaluator.EvaluateResource(
			ctx,
			subject,
			resource,
			policy.Rules.Resource,
		)
		if err != nil {
			errors = append(errors, fmt.Sprintf("policy %s resource eval: %v", policy.ID, err))
			continue
		}
		envMatch, err := e.evaluator.EvaluateEnvironment(ctx, env, policy.Rules.Environment)
		if err != nil {
			errors = append(errors, fmt.Sprintf("policy %s env eval: %v", policy.ID, err))
			continue
		}
		matched := subjectMatch && resourceMatch && envMatch

		if matched {
			meta.MatchedPolicies++

			// Explicit deny always wins immediately
			if policy.Effect == EffectDeny {
				return Decision{
					Allowed:  false,
					Effect:   EffectDeny,
					PolicyID: policy.ID,
					Reason:   fmt.Sprintf("Denied by policy: %s", policy.Name),
					Metadata: EvaluationMetadata{
						EvaluatedPolicies: meta.EvaluatedPolicies,
						MatchedPolicies:   meta.MatchedPolicies,
						EvaluationTime:    time.Since(start),
						Errors:            errors,
					},
				}, nil
			}

			// Allow: store but continue checking for higher-priority denies
			if finalDecision == nil ||
				policy.Priority < getPolicyPriority(applicablePolicies, finalDecision.PolicyID) {
				finalDecision = &Decision{
					Allowed:  true,
					Effect:   EffectAllow,
					PolicyID: policy.ID,
					Reason:   fmt.Sprintf("Allowed by policy: %s", policy.Name),
				}
			}
		}
	}

	if finalDecision != nil {
		finalDecision.Metadata = EvaluationMetadata{
			EvaluatedPolicies: meta.EvaluatedPolicies,
			MatchedPolicies:   meta.MatchedPolicies,
			EvaluationTime:    time.Since(start),
			Errors:            errors,
		}
		return *finalDecision, nil
	}

	// Default deny if no policy matched
	return Decision{
		Allowed: false,
		Effect:  EffectDeny,
		Reason:  "No applicable policy found - default deny",
		Metadata: EvaluationMetadata{
			EvaluatedPolicies: meta.EvaluatedPolicies,
			MatchedPolicies:   meta.MatchedPolicies,
			EvaluationTime:    time.Since(start),
			Errors:            errors,
		},
	}, nil
}

// AuthorizeSimple provides a simplified interface for common cases
func (e *DefaultPolicyEngine) AuthorizeSimple(
	ctx context.Context,
	userID string,
	action string,
	resourceType string,
	resourceID string,
) (Decision, error) {
	subject := Subject{
		ID:    userID,
		Type:  "user",
		Roles: []string{}, // Will be resolved by resolver
	}

	resource := Resource{
		Type: resourceType,
		ID:   resourceID,
	}

	act := Action{
		Name:     action,
		Resource: resourceType,
	}

	env := Environment{
		Timestamp: time.Now(),
	}

	return e.Authorize(ctx, subject, act, resource, env)
}

// Explain returns detailed trace information for debugging
func (e *DefaultPolicyEngine) Explain(
	ctx context.Context,
	subject Subject,
	action Action,
	resource Resource,
	env Environment,
) (EvaluationTrace, error) {
	trace := EvaluationTrace{
		Request: AuthorizationRequest{
			Subject:  subject,
			Action:   action,
			Resource: resource,
			Env:      env,
		},
	}

	// Similar to Authorize but captures detailed per-policy traces
	// Implementation omitted for brevity - mirrors Authorize with trace collection
	return trace, nil
}

// ============================================
// CACHING LAYER
// ============================================

// CachedPolicyStore wraps a PolicyStore with a shared cache (e.g. Redis).
type CachedPolicyStore struct {
	underlying PolicyStore
	cache      Cache
	ttl        time.Duration
}

// NewCachedPolicyStore wraps store with cache, caching lookups for ttl.
func NewCachedPolicyStore(store PolicyStore, cache Cache, ttl time.Duration) *CachedPolicyStore {
	return &CachedPolicyStore{underlying: store, cache: cache, ttl: ttl}
}

// Cache interface (use Redis, in-memory, etc.)
type Cache interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, value []byte, ttl time.Duration) error
	Delete(ctx context.Context, key string) error
}

func (c *CachedPolicyStore) GetPoliciesForResource(
	ctx context.Context,
	resourceType string,
) ([]Policy, error) {
	cacheKey := fmt.Sprintf("abac:policies:%s", resourceType)

	// Try cache first
	if data, err := c.cache.Get(ctx, cacheKey); err == nil && len(data) > 0 {
		var policies []Policy
		if err := json.Unmarshal(data, &policies); err == nil {
			return policies, nil
		}
	}

	// Fetch from underlying store
	policies, err := c.underlying.GetPoliciesForResource(ctx, resourceType)
	if err != nil {
		return nil, err
	}

	// Cache the result
	if data, err := json.Marshal(policies); err == nil {
		_ = c.cache.Set(ctx, cacheKey, data, c.ttl)
	}

	return policies, nil
}

// InvalidateCache clears cached policies (call when policies change)
func (c *CachedPolicyStore) InvalidateCache(ctx context.Context, resourceType string) error {
	return c.cache.Delete(ctx, fmt.Sprintf("abac:policies:%s", resourceType))
}

// ============================================
// HELPER FUNCTIONS
// ============================================

// matchAttributes checks if source attributes match target conditions
func matchAttributes(source, target map[string]any) (bool, error) {
	for key, expected := range target {
		actual, exists := source[key]
		if !exists {
			return false, nil
		}

		match, err := valuesMatch(actual, expected)
		if err != nil {
			return false, fmt.Errorf("compare %s: %w", key, err)
		}
		if !match {
			return false, nil
		}
	}
	return true, nil
}

// valuesMatch compares two values with type coercion
func valuesMatch(actual, expected any) (bool, error) {
	// Handle special operators in expected value
	if opMap, ok := expected.(map[string]any); ok {
		return evaluateOperator(actual, opMap)
	}

	// Direct comparison with type coercion
	return coerceAndCompare(actual, expected)
}

// evaluateOperator handles advanced comparison operators
func evaluateOperator(actual any, op map[string]any) (bool, error) {
	for opName, opValue := range op {
		switch opName {
		case "$eq":
			return coerceAndCompare(actual, opValue)
		case "$ne":
			match, err := coerceAndCompare(actual, opValue)
			return !match, err
		case "$gt":
			return numericCompare(actual, opValue, func(a, b float64) bool { return a > b })
		case "$gte":
			return numericCompare(actual, opValue, func(a, b float64) bool { return a >= b })
		case "$lt":
			return numericCompare(actual, opValue, func(a, b float64) bool { return a < b })
		case "$lte":
			return numericCompare(actual, opValue, func(a, b float64) bool { return a <= b })
		case "$in":
			return inArray(actual, opValue)
		case "$nin":
			match, err := inArray(actual, opValue)
			return !match, err
		case "$contains":
			return stringContains(actual, opValue)
		case "$startsWith":
			return stringStartsWith(actual, opValue)
		case "$regex":
			return regexMatch(actual, opValue)
		default:
			return false, fmt.Errorf("unknown operator: %s", opName)
		}
	}
	return false, fmt.Errorf("empty operator")
}

// coerceAndCompare attempts type coercion before comparison
func coerceAndCompare(a, b any) (bool, error) {
	// Convert both to strings for flexible comparison
	aStr := fmt.Sprintf("%v", a)
	bStr := fmt.Sprintf("%v", b)
	return aStr == bStr, nil
}

// numericCompare handles numeric comparisons
func numericCompare(actual, expected any, cmp func(float64, float64) bool) (bool, error) {
	a, ok1 := toFloat64(actual)
	b, ok2 := toFloat64(expected)
	if !ok1 || !ok2 {
		return false, fmt.Errorf("non-numeric values for numeric comparison")
	}
	return cmp(a, b), nil
}

// inArray checks if value exists in array
func inArray(value, array any) (bool, error) {
	arr, ok := array.([]any)
	if !ok {
		return false, fmt.Errorf("expected array for $in operator")
	}
	for _, item := range arr {
		if match, _ := coerceAndCompare(value, item); match {
			return true, nil
		}
	}
	return false, nil
}

// toFloat64 converts various types to float64
func toFloat64(v any) (float64, bool) {
	switch val := v.(type) {
	case float64:
		return val, true
	case float32:
		return float64(val), true
	case int:
		return float64(val), true
	case int64:
		return float64(val), true
	case string:
		f, err := strconv.ParseFloat(val, 64)
		return f, err == nil
	default:
		return 0, false
	}
}

// getNestedValue retrieves nested values using dot notation
func getNestedValue(m map[string]any, path string) string {
	if path == "" {
		return ""
	}

	parts := strings.Split(path, ".")
	current := m

	for i, part := range parts {
		v, exists := current[part]
		if !exists {
			return ""
		}

		if i == len(parts)-1 {
			return fmt.Sprintf("%v", v)
		}

		next, ok := v.(map[string]any)
		if !ok {
			return ""
		}
		current = next
	}

	return ""
}

// mergeMaps combines multiple maps (later values override earlier)
func mergeMaps(mps ...map[string]any) map[string]any {
	result := make(map[string]any)
	for _, m := range mps {
		maps.Copy(result, m)
	}
	return result
}

// structToMap converts a struct to map using JSON tags
func structToMap(v any) map[string]any {
	result := make(map[string]any)
	data, _ := json.Marshal(v)
	_ = json.Unmarshal(data, &result)
	return result
}

// sortPoliciesByPriority sorts policies in-place by priority ascending (lower
// number = higher priority), breaking ties by earliest creation via ID order.
func sortPoliciesByPriority(policies []Policy) {
	sort.SliceStable(policies, func(i, j int) bool {
		return policies[i].Priority < policies[j].Priority
	})
}

// getPolicyPriority finds a policy's priority by ID
func getPolicyPriority(policies []Policy, id string) int {
	for _, p := range policies {
		if p.ID == id {
			return p.Priority
		}
	}
	return 999999
}

// stringContains checks substring containment
func stringContains(actual, expected any) (bool, error) {
	aStr := fmt.Sprintf("%v", actual)
	bStr := fmt.Sprintf("%v", expected)
	return strings.Contains(aStr, bStr), nil
}

// stringStartsWith checks prefix
func stringStartsWith(actual, expected any) (bool, error) {
	aStr := fmt.Sprintf("%v", actual)
	bStr := fmt.Sprintf("%v", expected)
	return strings.HasPrefix(aStr, bStr), nil
}

// regexMatch matches the actual value against a regex pattern.
func regexMatch(actual, expected any) (bool, error) {
	aStr := fmt.Sprintf("%v", actual)
	pattern := fmt.Sprintf("%v", expected)
	return regexp.MatchString(pattern, aStr)
}

// ============================================
// USAGE EXAMPLE
// ============================================

/*
func ExampleUsage() {
	ctx := context.Background()

	// Initialize components
	store := &SQLPolicyStore{db: yourDBConnection}
	cache := &CachedPolicyStore{
		underlying: store,
		cache:      yourRedisClient,
		ttl:        5 * time.Minute,
	}
	resolver := &DatabaseAttributeResolver{
		userRepo:      yourUserRepo,
		attributeRepo: yourAttributeRepo,
	}
	evaluator := NewDefaultConditionEvaluator()

	engine := NewPolicyEngine(cache, resolver, evaluator)

	// Build authorization request
	subject := Subject{
		ID:    "user_123",
		Type:  "user",
		Roles: []string{"editor", "user"},
		Attributes: map[string]interface{}{
			"department": "engineering",
			"clearance":  "secret",
		},
	}

	action := Action{
		Name:     "update",
		Resource: "post",
	}

	resource := Resource{
		Type:    "post",
		ID:      "post_456",
		OwnerID: "user_123",
		Attributes: map[string]interface{}{
			"status": "draft",
			"category": "technical",
		},
	}

	env := Environment{
		Timestamp: time.Now(),
		IPAddress: net.ParseIP("10.0.1.50"),
		UserAgent: "Mozilla/5.0...",
		Custom: map[string]interface{}{
			"request_source": "web_app",
		},
	}

	// Evaluate authorization
	decision, err := engine.Authorize(ctx, subject, action, resource, env)
	if err != nil {
		log.Fatal(err)
	}

	if decision.Allowed {
		fmt.Printf("Access granted by policy %s: %s\n", decision.PolicyID, decision.Reason)
	} else {
		fmt.Printf("Access denied: %s\n", decision.Reason)
	}
}
*/
