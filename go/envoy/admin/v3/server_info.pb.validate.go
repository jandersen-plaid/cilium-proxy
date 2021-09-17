// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/admin/v3/server_info.proto

package envoy_admin_v3

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

// Validate checks the field values on ServerInfo with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *ServerInfo) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Version

	// no validation rules for State

	if v, ok := interface{}(m.GetUptimeCurrentEpoch()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServerInfoValidationError{
				field:  "UptimeCurrentEpoch",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUptimeAllEpochs()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServerInfoValidationError{
				field:  "UptimeAllEpochs",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for HotRestartVersion

	if v, ok := interface{}(m.GetCommandLineOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServerInfoValidationError{
				field:  "CommandLineOptions",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServerInfoValidationError{
				field:  "Node",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ServerInfoValidationError is the validation error returned by
// ServerInfo.Validate if the designated constraints aren't met.
type ServerInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServerInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServerInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServerInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServerInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServerInfoValidationError) ErrorName() string { return "ServerInfoValidationError" }

// Error satisfies the builtin error interface
func (e ServerInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServerInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServerInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServerInfoValidationError{}

// Validate checks the field values on CommandLineOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CommandLineOptions) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for BaseId

	// no validation rules for UseDynamicBaseId

	// no validation rules for BaseIdPath

	// no validation rules for Concurrency

	// no validation rules for ConfigPath

	// no validation rules for ConfigYaml

	// no validation rules for AllowUnknownStaticFields

	// no validation rules for RejectUnknownDynamicFields

	// no validation rules for IgnoreUnknownDynamicFields

	// no validation rules for AdminAddressPath

	// no validation rules for LocalAddressIpVersion

	// no validation rules for LogLevel

	// no validation rules for ComponentLogLevel

	// no validation rules for LogFormat

	// no validation rules for LogFormatEscaped

	// no validation rules for LogPath

	// no validation rules for ServiceCluster

	// no validation rules for ServiceNode

	// no validation rules for ServiceZone

	if v, ok := interface{}(m.GetFileFlushInterval()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommandLineOptionsValidationError{
				field:  "FileFlushInterval",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDrainTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommandLineOptionsValidationError{
				field:  "DrainTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for DrainStrategy

	if v, ok := interface{}(m.GetParentShutdownTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommandLineOptionsValidationError{
				field:  "ParentShutdownTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Mode

	// no validation rules for DisableHotRestart

	// no validation rules for EnableMutexTracing

	// no validation rules for RestartEpoch

	// no validation rules for CpusetThreads

	// no validation rules for BootstrapVersion

	// no validation rules for EnableFineGrainLogging

	// no validation rules for SocketPath

	// no validation rules for SocketMode

	// no validation rules for EnableCoreDump

	return nil
}

// CommandLineOptionsValidationError is the validation error returned by
// CommandLineOptions.Validate if the designated constraints aren't met.
type CommandLineOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommandLineOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommandLineOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommandLineOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommandLineOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommandLineOptionsValidationError) ErrorName() string {
	return "CommandLineOptionsValidationError"
}

// Error satisfies the builtin error interface
func (e CommandLineOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommandLineOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommandLineOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommandLineOptionsValidationError{}
