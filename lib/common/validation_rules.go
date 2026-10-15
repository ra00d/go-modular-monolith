package common

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gvalid"
)

// RegisterConditionalRules registers a set of Laravel-style conditional
// validation rules for GoFrame/gvalid.
//
// Registered rules:
//
//   - required-if
//   - required-if-not
//   - exclude-if
//   - exclude-unless
//   - prohibited-if
//   - prohibited-unless
//
// This function should be called once during application startup.
func RegisterConditionalRules() {
	register("required-if-not", requiredIfNot)
	register("exclude-if", excludeIf)
	register("exclude-unless", excludeUnless)
	register("prohibited-if", prohibitedIf)
	register("prohibited-unless", prohibitedUnless)
	register("skip-if", skipIf)
	register("skip-unless", skipUnless)
	register("skip-with", skipWith)
	register("skip-without", skipWithout)
	register("skip-any", skipAny)
	register("skip-all", skipAll)
}

func register(name string, fn gvalid.RuleFunc) {
	gvalid.RegisterRule(name, fn)
}

// required-if-not
//
// Makes the current field required when another field does not equal a specified value.
//
// Format:
//
//	required-if-not:Field,Value
//
// Examples:
//
//	v:"required-if-not:Identifier,username"
//
// Rules:
//
//   - If Identifier != "username", the current field must be present and non-empty.
//   - Otherwise, validation passes.
//
// Similar to Laravel's required_unless rule.
func requiredIfNot(ctx context.Context, in gvalid.RuleFuncInput) error {
	field, expected := parseParams(in.Rule)

	if getFieldValue(in.Data, field) == expected {
		return nil
	}

	if isEmpty(in.Value.Val()) {
		return gerror.Newf("%s is required", in.Field)
	}

	return nil
}

// exclude-if
//
// Excludes the current field from further validation when another field equals
// a specified value.
//
// Format:
//
//	exclude-if:Field,Value
//
// Examples:
//
//	v:"exclude-if:Identifier,username|email"
//
// Rules:
//
//   - If Identifier == "username", all remaining validators for the current
//     field are skipped.
//   - Otherwise, validation continues normally.
//
// Similar to Laravel's exclude_if rule.
func excludeIf(ctx context.Context, in gvalid.RuleFuncInput) error {
	field, expected := parseParams(in.Rule)

	if getFieldValue(in.Data, field) != expected {
		return nil
	}

	// Skip all following rules.
	return gerror.New("")
}

// exclude-unless
//
// Excludes the current field from further validation unless another field
// equals a specified value.
//
// Format:
//
//	exclude-unless:Field,Value
//
// Examples:
//
//	v:"exclude-unless:Type,email|email"
//
// Rules:
//
//   - If Type != "email", all remaining validators for the current field
//     are skipped.
//   - Otherwise, validation continues normally.
//
// Similar to Laravel's exclude_unless rule.
func excludeUnless(ctx context.Context, in gvalid.RuleFuncInput) error {
	field, expected := parseParams(in.Rule)

	if getFieldValue(in.Data, field) == expected {
		return nil
	}

	// Skip all following rules.
	return nil
}

// prohibited-if
//
// Prohibits the current field when another field equals a specified value.
//
// Format:
//
//	prohibited-if:Field,Value
//
// Examples:
//
//	v:"prohibited-if:Identifier,email"
//
// Rules:
//
//   - If Identifier == "email", the current field must be empty or absent.
//   - Otherwise, validation passes.
//
// Similar to Laravel's prohibited_if rule.
func prohibitedIf(ctx context.Context, in gvalid.RuleFuncInput) error {
	field, expected := parseParams(in.Rule)

	if getFieldValue(in.Data, field) != expected {
		return nil
	}

	if !isEmpty(in.Value.Val()) {
		return gerror.Newf("%s is prohibited", in.Field)
	}

	return nil
}

// prohibited-unless
//
// Prohibits the current field unless another field equals a specified value.
//
// Format:
//
//	prohibited-unless:Field,Value
//
// Examples:
//
//	v:"prohibited-unless:Type,admin"
//
// Rules:
//
//   - If Type != "admin", the current field must be empty or absent.
//   - Otherwise, validation passes.
//
// Similar to Laravel's prohibited_unless rule.
func prohibitedUnless(ctx context.Context, in gvalid.RuleFuncInput) error {
	field, expected := parseParams(in.Rule)

	if getFieldValue(in.Data, field) == expected {
		return nil
	}

	if !isEmpty(in.Value.Val()) {
		return gerror.Newf("%s is prohibited", in.Field)
	}

	return nil
}

func parseParams(pattern string) (field string, value string) {
	parts := strings.SplitN(pattern, ",", 2)

	if len(parts) != 2 {
		panic(fmt.Sprintf("invalid rule pattern: %s", pattern))
	}

	return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
}

func getFieldValue(data any, field string) string {
	m := gconv.Map(data)
	return gconv.String(m[field])
}

func isEmpty(v any) bool {
	if v == nil {
		return true
	}

	switch x := v.(type) {
	case string:
		return strings.TrimSpace(x) == ""
	case []byte:
		return len(x) == 0
	case []string:
		return len(x) == 0
	case []int:
		return len(x) == 0
	case []any:
		return len(x) == 0
	}

	return gconv.String(v) == ""
}

// skip-if:Field=value
// skip-if:Amount>=10
// skip-if:Role~^(admin|root)$
func skipIf(ctx context.Context, in gvalid.RuleFuncInput) error {
	if evaluate(in.Data, in.Rule) {
		return gerror.New("")
	}
	return nil
}

// skip-unless:Field=value
func skipUnless(ctx context.Context, in gvalid.RuleFuncInput) error {
	if !evaluate(in.Data, in.Rule) {
		return gerror.New("")
	}
	return nil
}

// skip-with:Field
func skipWith(ctx context.Context, in gvalid.RuleFuncInput) error {
	if hasValue(getField(in.Data, in.Rule)) {
		return gerror.New("")
	}
	return nil
}

// skip-without:Field
func skipWithout(ctx context.Context, in gvalid.RuleFuncInput) error {
	if !hasValue(getField(in.Data, in.Rule)) {
		return gerror.New("")
	}
	return nil
}

// skip-any:Field1,Field2,Field3
func skipAny(ctx context.Context, in gvalid.RuleFuncInput) error {
	fields := splitFields(in.Rule)

	for _, field := range fields {
		if hasValue(getField(in.Data, field)) {
			return gerror.New("")
		}
	}

	return nil
}

// skip-all:Field1,Field2,Field3
func skipAll(ctx context.Context, in gvalid.RuleFuncInput) error {
	fields := splitFields(in.Rule)

	for _, field := range fields {
		if !hasValue(getField(in.Data, field)) {
			return nil
		}
	}

	return gerror.New("")
}

func evaluate(data any, expr string) bool {
	for _, op := range []string{
		"!~",
		">=",
		"<=",
		"!=",
		">",
		"<",
		"=",
		"~",
	} {
		if idx := strings.Index(expr, op); idx > 0 {
			field := strings.TrimSpace(expr[:idx])
			expected := strings.TrimSpace(expr[idx+len(op):])

			actual := getField(data, field)

			switch op {
			case "=":
				return actual == expected

			case "!=":
				return actual != expected

			case "~":
				ok, _ := regexp.MatchString(expected, actual)
				return ok

			case "!~":
				ok, _ := regexp.MatchString(expected, actual)
				return !ok

			case ">", ">=", "<", "<=":
				return compareNumeric(actual, expected, op)
			}
		}
	}

	panic(gerror.Newf("invalid skip expression: %s", expr))
}

func compareNumeric(actual, expected, op string) bool {
	a, err1 := strconv.ParseFloat(actual, 64)
	b, err2 := strconv.ParseFloat(expected, 64)

	if err1 != nil || err2 != nil {
		return false
	}

	switch op {
	case ">":
		return a > b
	case ">=":
		return a >= b
	case "<":
		return a < b
	case "<=":
		return a <= b
	}

	return false
}

func getField(data any, field string) string {
	m := gconv.Map(data)
	return gconv.String(m[field])
}

func splitFields(fields string) []string {
	parts := strings.Split(fields, ",")

	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	return parts
}

func hasValue(v string) bool {
	return strings.TrimSpace(v) != ""
}
