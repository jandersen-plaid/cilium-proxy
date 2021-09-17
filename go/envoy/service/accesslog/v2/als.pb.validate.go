// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/service/accesslog/v2/als.proto

package envoy_service_accesslog_v2

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on StreamAccessLogsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *StreamAccessLogsResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// StreamAccessLogsResponseValidationError is the validation error returned by
// StreamAccessLogsResponse.Validate if the designated constraints aren't met.
type StreamAccessLogsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamAccessLogsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamAccessLogsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamAccessLogsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamAccessLogsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamAccessLogsResponseValidationError) ErrorName() string {
	return "StreamAccessLogsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e StreamAccessLogsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamAccessLogsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamAccessLogsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamAccessLogsResponseValidationError{}

// Validate checks the field values on StreamAccessLogsMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *StreamAccessLogsMessage) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetIdentifier()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StreamAccessLogsMessageValidationError{
				field:  "Identifier",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.LogEntries.(type) {

	case *StreamAccessLogsMessage_HttpLogs:

		if v, ok := interface{}(m.GetHttpLogs()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StreamAccessLogsMessageValidationError{
					field:  "HttpLogs",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *StreamAccessLogsMessage_TcpLogs:

		if v, ok := interface{}(m.GetTcpLogs()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StreamAccessLogsMessageValidationError{
					field:  "TcpLogs",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return StreamAccessLogsMessageValidationError{
			field:  "LogEntries",
			reason: "value is required",
		}

	}

	return nil
}

// StreamAccessLogsMessageValidationError is the validation error returned by
// StreamAccessLogsMessage.Validate if the designated constraints aren't met.
type StreamAccessLogsMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamAccessLogsMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamAccessLogsMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamAccessLogsMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamAccessLogsMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamAccessLogsMessageValidationError) ErrorName() string {
	return "StreamAccessLogsMessageValidationError"
}

// Error satisfies the builtin error interface
func (e StreamAccessLogsMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamAccessLogsMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamAccessLogsMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamAccessLogsMessageValidationError{}

// Validate checks the field values on StreamAccessLogsMessage_Identifier with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *StreamAccessLogsMessage_Identifier) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetNode() == nil {
		return StreamAccessLogsMessage_IdentifierValidationError{
			field:  "Node",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StreamAccessLogsMessage_IdentifierValidationError{
				field:  "Node",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(m.GetLogName()) < 1 {
		return StreamAccessLogsMessage_IdentifierValidationError{
			field:  "LogName",
			reason: "value length must be at least 1 bytes",
		}
	}

	return nil
}

// StreamAccessLogsMessage_IdentifierValidationError is the validation error
// returned by StreamAccessLogsMessage_Identifier.Validate if the designated
// constraints aren't met.
type StreamAccessLogsMessage_IdentifierValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamAccessLogsMessage_IdentifierValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamAccessLogsMessage_IdentifierValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamAccessLogsMessage_IdentifierValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamAccessLogsMessage_IdentifierValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamAccessLogsMessage_IdentifierValidationError) ErrorName() string {
	return "StreamAccessLogsMessage_IdentifierValidationError"
}

// Error satisfies the builtin error interface
func (e StreamAccessLogsMessage_IdentifierValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamAccessLogsMessage_Identifier.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamAccessLogsMessage_IdentifierValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamAccessLogsMessage_IdentifierValidationError{}

// Validate checks the field values on
// StreamAccessLogsMessage_HTTPAccessLogEntries with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetLogEntry()) < 1 {
		return StreamAccessLogsMessage_HTTPAccessLogEntriesValidationError{
			field:  "LogEntry",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetLogEntry() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StreamAccessLogsMessage_HTTPAccessLogEntriesValidationError{
					field:  fmt.Sprintf("LogEntry[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// StreamAccessLogsMessage_HTTPAccessLogEntriesValidationError is the
// validation error returned by
// StreamAccessLogsMessage_HTTPAccessLogEntries.Validate if the designated
// constraints aren't met.
type StreamAccessLogsMessage_HTTPAccessLogEntriesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamAccessLogsMessage_HTTPAccessLogEntriesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamAccessLogsMessage_HTTPAccessLogEntriesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamAccessLogsMessage_HTTPAccessLogEntriesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamAccessLogsMessage_HTTPAccessLogEntriesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamAccessLogsMessage_HTTPAccessLogEntriesValidationError) ErrorName() string {
	return "StreamAccessLogsMessage_HTTPAccessLogEntriesValidationError"
}

// Error satisfies the builtin error interface
func (e StreamAccessLogsMessage_HTTPAccessLogEntriesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamAccessLogsMessage_HTTPAccessLogEntries.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamAccessLogsMessage_HTTPAccessLogEntriesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamAccessLogsMessage_HTTPAccessLogEntriesValidationError{}

// Validate checks the field values on
// StreamAccessLogsMessage_TCPAccessLogEntries with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *StreamAccessLogsMessage_TCPAccessLogEntries) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetLogEntry()) < 1 {
		return StreamAccessLogsMessage_TCPAccessLogEntriesValidationError{
			field:  "LogEntry",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetLogEntry() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StreamAccessLogsMessage_TCPAccessLogEntriesValidationError{
					field:  fmt.Sprintf("LogEntry[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// StreamAccessLogsMessage_TCPAccessLogEntriesValidationError is the validation
// error returned by StreamAccessLogsMessage_TCPAccessLogEntries.Validate if
// the designated constraints aren't met.
type StreamAccessLogsMessage_TCPAccessLogEntriesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamAccessLogsMessage_TCPAccessLogEntriesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamAccessLogsMessage_TCPAccessLogEntriesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamAccessLogsMessage_TCPAccessLogEntriesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamAccessLogsMessage_TCPAccessLogEntriesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamAccessLogsMessage_TCPAccessLogEntriesValidationError) ErrorName() string {
	return "StreamAccessLogsMessage_TCPAccessLogEntriesValidationError"
}

// Error satisfies the builtin error interface
func (e StreamAccessLogsMessage_TCPAccessLogEntriesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamAccessLogsMessage_TCPAccessLogEntries.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamAccessLogsMessage_TCPAccessLogEntriesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamAccessLogsMessage_TCPAccessLogEntriesValidationError{}
