// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: peasant/v1/password_service.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
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
	_ = sort.Sort
)

// define the regex for a UUID once up-front
var _password_service_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on ResetPasswordRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ResetPasswordRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ResetPasswordRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ResetPasswordRequestMultiError, or nil if none found.
func (m *ResetPasswordRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ResetPasswordRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for EmailOrUsername

	if len(errors) > 0 {
		return ResetPasswordRequestMultiError(errors)
	}

	return nil
}

// ResetPasswordRequestMultiError is an error wrapping multiple validation
// errors returned by ResetPasswordRequest.ValidateAll() if the designated
// constraints aren't met.
type ResetPasswordRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ResetPasswordRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ResetPasswordRequestMultiError) AllErrors() []error { return m }

// ResetPasswordRequestValidationError is the validation error returned by
// ResetPasswordRequest.Validate if the designated constraints aren't met.
type ResetPasswordRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResetPasswordRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResetPasswordRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResetPasswordRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResetPasswordRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResetPasswordRequestValidationError) ErrorName() string {
	return "ResetPasswordRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ResetPasswordRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResetPasswordRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResetPasswordRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResetPasswordRequestValidationError{}

// Validate checks the field values on SetPasswordRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SetPasswordRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SetPasswordRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SetPasswordRequestMultiError, or nil if none found.
func (m *SetPasswordRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SetPasswordRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	// no validation rules for Password

	if len(errors) > 0 {
		return SetPasswordRequestMultiError(errors)
	}

	return nil
}

// SetPasswordRequestMultiError is an error wrapping multiple validation errors
// returned by SetPasswordRequest.ValidateAll() if the designated constraints
// aren't met.
type SetPasswordRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SetPasswordRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SetPasswordRequestMultiError) AllErrors() []error { return m }

// SetPasswordRequestValidationError is the validation error returned by
// SetPasswordRequest.Validate if the designated constraints aren't met.
type SetPasswordRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SetPasswordRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SetPasswordRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SetPasswordRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SetPasswordRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SetPasswordRequestValidationError) ErrorName() string {
	return "SetPasswordRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SetPasswordRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSetPasswordRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SetPasswordRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SetPasswordRequestValidationError{}

// Validate checks the field values on UpdatePasswordRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdatePasswordRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdatePasswordRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdatePasswordRequestMultiError, or nil if none found.
func (m *UpdatePasswordRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdatePasswordRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetAccountId()); err != nil {
		err = UpdatePasswordRequestValidationError{
			field:  "AccountId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetOldPassword()) > 64 {
		err := UpdatePasswordRequestValidationError{
			field:  "OldPassword",
			reason: "value length must be at most 64 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetNewPassword()); l < 8 || l > 64 {
		err := UpdatePasswordRequestValidationError{
			field:  "NewPassword",
			reason: "value length must be between 8 and 64 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UpdatePasswordRequestMultiError(errors)
	}

	return nil
}

func (m *UpdatePasswordRequest) _validateUuid(uuid string) error {
	if matched := _password_service_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// UpdatePasswordRequestMultiError is an error wrapping multiple validation
// errors returned by UpdatePasswordRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdatePasswordRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdatePasswordRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdatePasswordRequestMultiError) AllErrors() []error { return m }

// UpdatePasswordRequestValidationError is the validation error returned by
// UpdatePasswordRequest.Validate if the designated constraints aren't met.
type UpdatePasswordRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdatePasswordRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdatePasswordRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdatePasswordRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdatePasswordRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdatePasswordRequestValidationError) ErrorName() string {
	return "UpdatePasswordRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdatePasswordRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdatePasswordRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdatePasswordRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdatePasswordRequestValidationError{}

// Validate checks the field values on UpdatePasswordResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdatePasswordResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdatePasswordResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdatePasswordResponseMultiError, or nil if none found.
func (m *UpdatePasswordResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdatePasswordResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetAccount()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdatePasswordResponseValidationError{
					field:  "Account",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdatePasswordResponseValidationError{
					field:  "Account",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAccount()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdatePasswordResponseValidationError{
				field:  "Account",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdatePasswordResponseMultiError(errors)
	}

	return nil
}

// UpdatePasswordResponseMultiError is an error wrapping multiple validation
// errors returned by UpdatePasswordResponse.ValidateAll() if the designated
// constraints aren't met.
type UpdatePasswordResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdatePasswordResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdatePasswordResponseMultiError) AllErrors() []error { return m }

// UpdatePasswordResponseValidationError is the validation error returned by
// UpdatePasswordResponse.Validate if the designated constraints aren't met.
type UpdatePasswordResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdatePasswordResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdatePasswordResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdatePasswordResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdatePasswordResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdatePasswordResponseValidationError) ErrorName() string {
	return "UpdatePasswordResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdatePasswordResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdatePasswordResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdatePasswordResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdatePasswordResponseValidationError{}

// Validate checks the field values on SetPasswordResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SetPasswordResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SetPasswordResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SetPasswordResponseMultiError, or nil if none found.
func (m *SetPasswordResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SetPasswordResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetAccount()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SetPasswordResponseValidationError{
					field:  "Account",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SetPasswordResponseValidationError{
					field:  "Account",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAccount()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SetPasswordResponseValidationError{
				field:  "Account",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SetPasswordResponseMultiError(errors)
	}

	return nil
}

// SetPasswordResponseMultiError is an error wrapping multiple validation
// errors returned by SetPasswordResponse.ValidateAll() if the designated
// constraints aren't met.
type SetPasswordResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SetPasswordResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SetPasswordResponseMultiError) AllErrors() []error { return m }

// SetPasswordResponseValidationError is the validation error returned by
// SetPasswordResponse.Validate if the designated constraints aren't met.
type SetPasswordResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SetPasswordResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SetPasswordResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SetPasswordResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SetPasswordResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SetPasswordResponseValidationError) ErrorName() string {
	return "SetPasswordResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SetPasswordResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSetPasswordResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SetPasswordResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SetPasswordResponseValidationError{}

// Validate checks the field values on ResetPasswordResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ResetPasswordResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ResetPasswordResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ResetPasswordResponseMultiError, or nil if none found.
func (m *ResetPasswordResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ResetPasswordResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetAccount()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ResetPasswordResponseValidationError{
					field:  "Account",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ResetPasswordResponseValidationError{
					field:  "Account",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAccount()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ResetPasswordResponseValidationError{
				field:  "Account",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ResetPasswordResponseMultiError(errors)
	}

	return nil
}

// ResetPasswordResponseMultiError is an error wrapping multiple validation
// errors returned by ResetPasswordResponse.ValidateAll() if the designated
// constraints aren't met.
type ResetPasswordResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ResetPasswordResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ResetPasswordResponseMultiError) AllErrors() []error { return m }

// ResetPasswordResponseValidationError is the validation error returned by
// ResetPasswordResponse.Validate if the designated constraints aren't met.
type ResetPasswordResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResetPasswordResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResetPasswordResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResetPasswordResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResetPasswordResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResetPasswordResponseValidationError) ErrorName() string {
	return "ResetPasswordResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ResetPasswordResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResetPasswordResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResetPasswordResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResetPasswordResponseValidationError{}
