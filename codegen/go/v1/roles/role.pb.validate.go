// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: roles/role.proto

package roles

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// Validate checks the field values on Grant with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Grant) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetRoles()) != 5 {
		return GrantValidationError{
			field:  "Roles",
			reason: "value must contain exactly 5 item(s)",
		}
	}

	for idx, item := range m.GetRoles() {
		_, _ = idx, item

		if _, ok := _Grant_Roles_InLookup[item]; !ok {
			return GrantValidationError{
				field:  fmt.Sprintf("Roles[%v]", idx),
				reason: "value must be in list [NONE STAFF CASHIER MANAGER ADMIN OWNER]",
			}
		}

	}

	return nil
}

// GrantValidationError is the validation error returned by Grant.Validate if
// the designated constraints aren't met.
type GrantValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GrantValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GrantValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GrantValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GrantValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GrantValidationError) ErrorName() string { return "GrantValidationError" }

// Error satisfies the builtin error interface
func (e GrantValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrant.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GrantValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GrantValidationError{}

var _Grant_Roles_InLookup = map[string]struct{}{
	"NONE":    {},
	"STAFF":   {},
	"CASHIER": {},
	"MANAGER": {},
	"ADMIN":   {},
	"OWNER":   {},
}

// Validate checks the field values on GetRoleOfId with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetRoleOfId) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// GetRoleOfIdValidationError is the validation error returned by
// GetRoleOfId.Validate if the designated constraints aren't met.
type GetRoleOfIdValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetRoleOfIdValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetRoleOfIdValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetRoleOfIdValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetRoleOfIdValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetRoleOfIdValidationError) ErrorName() string { return "GetRoleOfIdValidationError" }

// Error satisfies the builtin error interface
func (e GetRoleOfIdValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRoleOfId.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetRoleOfIdValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetRoleOfIdValidationError{}
