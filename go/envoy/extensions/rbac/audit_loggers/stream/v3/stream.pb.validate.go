// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/rbac/audit_loggers/stream/v3/stream.proto

package streamv3

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

// Validate checks the field values on StdoutAuditLog with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StdoutAuditLog) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StdoutAuditLog with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StdoutAuditLogMultiError,
// or nil if none found.
func (m *StdoutAuditLog) ValidateAll() error {
	return m.validate(true)
}

func (m *StdoutAuditLog) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return StdoutAuditLogMultiError(errors)
	}
	return nil
}

// StdoutAuditLogMultiError is an error wrapping multiple validation errors
// returned by StdoutAuditLog.ValidateAll() if the designated constraints
// aren't met.
type StdoutAuditLogMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StdoutAuditLogMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StdoutAuditLogMultiError) AllErrors() []error { return m }

// StdoutAuditLogValidationError is the validation error returned by
// StdoutAuditLog.Validate if the designated constraints aren't met.
type StdoutAuditLogValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StdoutAuditLogValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StdoutAuditLogValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StdoutAuditLogValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StdoutAuditLogValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StdoutAuditLogValidationError) ErrorName() string { return "StdoutAuditLogValidationError" }

// Error satisfies the builtin error interface
func (e StdoutAuditLogValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStdoutAuditLog.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StdoutAuditLogValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StdoutAuditLogValidationError{}
