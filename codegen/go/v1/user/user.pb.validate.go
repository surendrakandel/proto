// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: user/user.proto

package user

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

// Validate checks the field values on All with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *All) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Kind

	// no validation rules for Path

	return nil
}

// AllValidationError is the validation error returned by All.Validate if the
// designated constraints aren't met.
type AllValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AllValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AllValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AllValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AllValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AllValidationError) ErrorName() string { return "AllValidationError" }

// Error satisfies the builtin error interface
func (e AllValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAll.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AllValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AllValidationError{}

// Validate checks the field values on ModifyRightsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ModifyRightsRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for UserId

	if v, ok := interface{}(m.GetBanReview()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ModifyRightsRequestValidationError{
				field:  "BanReview",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ModifyRightsRequestValidationError is the validation error returned by
// ModifyRightsRequest.Validate if the designated constraints aren't met.
type ModifyRightsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ModifyRightsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ModifyRightsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ModifyRightsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ModifyRightsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ModifyRightsRequestValidationError) ErrorName() string {
	return "ModifyRightsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ModifyRightsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sModifyRightsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ModifyRightsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ModifyRightsRequestValidationError{}

// Validate checks the field values on ShortUsers with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *ShortUsers) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetUsers() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ShortUsersValidationError{
					field:  fmt.Sprintf("Users[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ShortUsersValidationError is the validation error returned by
// ShortUsers.Validate if the designated constraints aren't met.
type ShortUsersValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ShortUsersValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ShortUsersValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ShortUsersValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ShortUsersValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ShortUsersValidationError) ErrorName() string { return "ShortUsersValidationError" }

// Error satisfies the builtin error interface
func (e ShortUsersValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sShortUsers.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ShortUsersValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ShortUsersValidationError{}

// Validate checks the field values on ShortUser with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *ShortUser) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for TtId

	// no validation rules for FirstName

	// no validation rules for MiddleName

	// no validation rules for LastName

	// no validation rules for Photo

	// no validation rules for PhotoRec

	// no validation rules for BanReview

	return nil
}

// ShortUserValidationError is the validation error returned by
// ShortUser.Validate if the designated constraints aren't met.
type ShortUserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ShortUserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ShortUserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ShortUserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ShortUserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ShortUserValidationError) ErrorName() string { return "ShortUserValidationError" }

// Error satisfies the builtin error interface
func (e ShortUserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sShortUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ShortUserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ShortUserValidationError{}

// Validate checks the field values on GetByIdRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetByIdRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for UserId

	return nil
}

// GetByIdRequestValidationError is the validation error returned by
// GetByIdRequest.Validate if the designated constraints aren't met.
type GetByIdRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetByIdRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetByIdRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetByIdRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetByIdRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetByIdRequestValidationError) ErrorName() string { return "GetByIdRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetByIdRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetByIdRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetByIdRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetByIdRequestValidationError{}

// Validate checks the field values on GetByIdsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetByIdsRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// GetByIdsRequestValidationError is the validation error returned by
// GetByIdsRequest.Validate if the designated constraints aren't met.
type GetByIdsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetByIdsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetByIdsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetByIdsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetByIdsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetByIdsRequestValidationError) ErrorName() string { return "GetByIdsRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetByIdsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetByIdsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetByIdsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetByIdsRequestValidationError{}

// Validate checks the field values on AuthRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *AuthRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Token

	return nil
}

// AuthRequestValidationError is the validation error returned by
// AuthRequest.Validate if the designated constraints aren't met.
type AuthRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthRequestValidationError) ErrorName() string { return "AuthRequestValidationError" }

// Error satisfies the builtin error interface
func (e AuthRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthRequestValidationError{}

// Validate checks the field values on MicrosoftAuthRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MicrosoftAuthRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	return nil
}

// MicrosoftAuthRequestValidationError is the validation error returned by
// MicrosoftAuthRequest.Validate if the designated constraints aren't met.
type MicrosoftAuthRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MicrosoftAuthRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MicrosoftAuthRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MicrosoftAuthRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MicrosoftAuthRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MicrosoftAuthRequestValidationError) ErrorName() string {
	return "MicrosoftAuthRequestValidationError"
}

// Error satisfies the builtin error interface
func (e MicrosoftAuthRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMicrosoftAuthRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MicrosoftAuthRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MicrosoftAuthRequestValidationError{}

// Validate checks the field values on YahooAuthRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *YahooAuthRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	return nil
}

// YahooAuthRequestValidationError is the validation error returned by
// YahooAuthRequest.Validate if the designated constraints aren't met.
type YahooAuthRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e YahooAuthRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e YahooAuthRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e YahooAuthRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e YahooAuthRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e YahooAuthRequestValidationError) ErrorName() string { return "YahooAuthRequestValidationError" }

// Error satisfies the builtin error interface
func (e YahooAuthRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sYahooAuthRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = YahooAuthRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = YahooAuthRequestValidationError{}

// Validate checks the field values on FacebookAuthRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *FacebookAuthRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	return nil
}

// FacebookAuthRequestValidationError is the validation error returned by
// FacebookAuthRequest.Validate if the designated constraints aren't met.
type FacebookAuthRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FacebookAuthRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FacebookAuthRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FacebookAuthRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FacebookAuthRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FacebookAuthRequestValidationError) ErrorName() string {
	return "FacebookAuthRequestValidationError"
}

// Error satisfies the builtin error interface
func (e FacebookAuthRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFacebookAuthRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FacebookAuthRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FacebookAuthRequestValidationError{}

// Validate checks the field values on AppleAuthRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *AppleAuthRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	return nil
}

// AppleAuthRequestValidationError is the validation error returned by
// AppleAuthRequest.Validate if the designated constraints aren't met.
type AppleAuthRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AppleAuthRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AppleAuthRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AppleAuthRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AppleAuthRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AppleAuthRequestValidationError) ErrorName() string { return "AppleAuthRequestValidationError" }

// Error satisfies the builtin error interface
func (e AppleAuthRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAppleAuthRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AppleAuthRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AppleAuthRequestValidationError{}

// Validate checks the field values on GoogleAuthRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GoogleAuthRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	return nil
}

// GoogleAuthRequestValidationError is the validation error returned by
// GoogleAuthRequest.Validate if the designated constraints aren't met.
type GoogleAuthRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GoogleAuthRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GoogleAuthRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GoogleAuthRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GoogleAuthRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GoogleAuthRequestValidationError) ErrorName() string {
	return "GoogleAuthRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GoogleAuthRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGoogleAuthRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GoogleAuthRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GoogleAuthRequestValidationError{}

// Validate checks the field values on SelfUser with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *SelfUser) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for VkId

	// no validation rules for Token

	// no validation rules for FirstName

	// no validation rules for MiddleName

	// no validation rules for LastName

	// no validation rules for Photo

	// no validation rules for PhotoRec

	return nil
}

// SelfUserValidationError is the validation error returned by
// SelfUser.Validate if the designated constraints aren't met.
type SelfUserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SelfUserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SelfUserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SelfUserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SelfUserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SelfUserValidationError) ErrorName() string { return "SelfUserValidationError" }

// Error satisfies the builtin error interface
func (e SelfUserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSelfUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SelfUserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SelfUserValidationError{}

// Validate checks the field values on SignUpUserData with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SignUpUserData) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for FirstName

	// no validation rules for MiddleName

	// no validation rules for LastName

	if utf8.RuneCountInString(m.GetUserName()) < 8 {
		return SignUpUserDataValidationError{
			field:  "UserName",
			reason: "value length must be at least 8 runes",
		}
	}

	if utf8.RuneCountInString(m.GetEmail()) < 10 {
		return SignUpUserDataValidationError{
			field:  "Email",
			reason: "value length must be at least 10 runes",
		}
	}

	if err := m._validateEmail(m.GetEmail()); err != nil {
		return SignUpUserDataValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
	}

	if utf8.RuneCountInString(m.GetPassword()) < 8 {
		return SignUpUserDataValidationError{
			field:  "Password",
			reason: "value length must be at least 8 runes",
		}
	}

	if utf8.RuneCountInString(m.GetConfirmPassword()) < 8 {
		return SignUpUserDataValidationError{
			field:  "ConfirmPassword",
			reason: "value length must be at least 8 runes",
		}
	}

	if len(m.GetRoles()) > 0 {

		if len(m.GetRoles()) != 5 {
			return SignUpUserDataValidationError{
				field:  "Roles",
				reason: "value must contain exactly 5 item(s)",
			}
		}

		for idx, item := range m.GetRoles() {
			_, _ = idx, item

			if val := item; val < 0 || val >= 5 {
				return SignUpUserDataValidationError{
					field:  fmt.Sprintf("Roles[%v]", idx),
					reason: "value must be inside range [0, 5)",
				}
			}

		}

	}

	// no validation rules for PhoneNumber

	// no validation rules for Version

	return nil
}

func (m *SignUpUserData) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *SignUpUserData) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// SignUpUserDataValidationError is the validation error returned by
// SignUpUserData.Validate if the designated constraints aren't met.
type SignUpUserDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignUpUserDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignUpUserDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignUpUserDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignUpUserDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignUpUserDataValidationError) ErrorName() string { return "SignUpUserDataValidationError" }

// Error satisfies the builtin error interface
func (e SignUpUserDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignUpUserData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignUpUserDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignUpUserDataValidationError{}

// Validate checks the field values on UserAccount with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UserAccount) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for FirstName

	// no validation rules for MiddleName

	// no validation rules for LastName

	if utf8.RuneCountInString(m.GetUserName()) < 8 {
		return UserAccountValidationError{
			field:  "UserName",
			reason: "value length must be at least 8 runes",
		}
	}

	if utf8.RuneCountInString(m.GetEmail()) < 10 {
		return UserAccountValidationError{
			field:  "Email",
			reason: "value length must be at least 10 runes",
		}
	}

	if err := m._validateEmail(m.GetEmail()); err != nil {
		return UserAccountValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
	}

	if utf8.RuneCountInString(m.GetPassword()) < 8 {
		return UserAccountValidationError{
			field:  "Password",
			reason: "value length must be at least 8 runes",
		}
	}

	if len(m.GetRoles()) > 0 {

		if l := len(m.GetRoles()); l < 5 || l > 7 {
			return UserAccountValidationError{
				field:  "Roles",
				reason: "value must contain between 5 and 7 items, inclusive",
			}
		}

		for idx, item := range m.GetRoles() {
			_, _ = idx, item

			if val := item; val < 0 || val >= 5 {
				return UserAccountValidationError{
					field:  fmt.Sprintf("Roles[%v]", idx),
					reason: "value must be inside range [0, 5)",
				}
			}

		}

	}

	// no validation rules for PhoneNumber

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserAccountValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserAccountValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetLastLoginAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserAccountValidationError{
				field:  "LastLoginAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDeletedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserAccountValidationError{
				field:  "DeletedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Version

	return nil
}

func (m *UserAccount) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *UserAccount) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// UserAccountValidationError is the validation error returned by
// UserAccount.Validate if the designated constraints aren't met.
type UserAccountValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserAccountValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserAccountValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserAccountValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserAccountValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserAccountValidationError) ErrorName() string { return "UserAccountValidationError" }

// Error satisfies the builtin error interface
func (e UserAccountValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserAccount.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserAccountValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserAccountValidationError{}

// Validate checks the field values on SignInRequestData with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *SignInRequestData) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateEmail(m.GetEmail()); err != nil {
		return SignInRequestDataValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
	}

	if utf8.RuneCountInString(m.GetPassword()) < 8 {
		return SignInRequestDataValidationError{
			field:  "Password",
			reason: "value length must be at least 8 runes",
		}
	}

	return nil
}

func (m *SignInRequestData) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *SignInRequestData) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// SignInRequestDataValidationError is the validation error returned by
// SignInRequestData.Validate if the designated constraints aren't met.
type SignInRequestDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignInRequestDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignInRequestDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignInRequestDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignInRequestDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignInRequestDataValidationError) ErrorName() string {
	return "SignInRequestDataValidationError"
}

// Error satisfies the builtin error interface
func (e SignInRequestDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignInRequestData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignInRequestDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignInRequestDataValidationError{}

// Validate checks the field values on SignInResponseData with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SignInResponseData) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Token

	return nil
}

// SignInResponseDataValidationError is the validation error returned by
// SignInResponseData.Validate if the designated constraints aren't met.
type SignInResponseDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignInResponseDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignInResponseDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignInResponseDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignInResponseDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignInResponseDataValidationError) ErrorName() string {
	return "SignInResponseDataValidationError"
}

// Error satisfies the builtin error interface
func (e SignInResponseDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignInResponseData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignInResponseDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignInResponseDataValidationError{}

// Validate checks the field values on SignUpResponseData with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SignUpResponseData) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for UserId

	// no validation rules for Success

	return nil
}

// SignUpResponseDataValidationError is the validation error returned by
// SignUpResponseData.Validate if the designated constraints aren't met.
type SignUpResponseDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignUpResponseDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignUpResponseDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignUpResponseDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignUpResponseDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignUpResponseDataValidationError) ErrorName() string {
	return "SignUpResponseDataValidationError"
}

// Error satisfies the builtin error interface
func (e SignUpResponseDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignUpResponseData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignUpResponseDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignUpResponseDataValidationError{}
