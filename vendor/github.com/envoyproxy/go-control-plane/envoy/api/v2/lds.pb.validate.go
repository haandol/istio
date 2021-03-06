// Code generated by protoc-gen-validate
// source: envoy/api/v2/lds.proto
// DO NOT EDIT!!!

package v2

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

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// Validate checks the field values on Listener with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Listener) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	if v, ok := interface{}(m.GetAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListenerValidationError{
				Field:  "Address",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if len(m.GetFilterChains()) < 1 {
		return ListenerValidationError{
			Field:  "FilterChains",
			Reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetFilterChains() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListenerValidationError{
					Field:  fmt.Sprintf("FilterChains[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetUseOriginalDst()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListenerValidationError{
				Field:  "UseOriginalDst",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetPerConnectionBufferLimitBytes()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListenerValidationError{
				Field:  "PerConnectionBufferLimitBytes",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListenerValidationError{
				Field:  "Metadata",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDeprecatedV1()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListenerValidationError{
				Field:  "DeprecatedV1",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	// no validation rules for DrainType

	for idx, item := range m.GetListenerFilters() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListenerValidationError{
					Field:  fmt.Sprintf("ListenerFilters[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetListenerFiltersTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListenerValidationError{
				Field:  "ListenerFiltersTimeout",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTransparent()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListenerValidationError{
				Field:  "Transparent",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetFreebind()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListenerValidationError{
				Field:  "Freebind",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	for idx, item := range m.GetSocketOptions() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListenerValidationError{
					Field:  fmt.Sprintf("SocketOptions[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetTcpFastOpenQueueLength()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListenerValidationError{
				Field:  "TcpFastOpenQueueLength",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetBugfixReverseWriteFilterOrder()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListenerValidationError{
				Field:  "BugfixReverseWriteFilterOrder",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// ListenerValidationError is the validation error returned by
// Listener.Validate if the designated constraints aren't met.
type ListenerValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e ListenerValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListener.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = ListenerValidationError{}

// Validate checks the field values on Listener_DeprecatedV1 with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *Listener_DeprecatedV1) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetBindToPort()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Listener_DeprecatedV1ValidationError{
				Field:  "BindToPort",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// Listener_DeprecatedV1ValidationError is the validation error returned by
// Listener_DeprecatedV1.Validate if the designated constraints aren't met.
type Listener_DeprecatedV1ValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e Listener_DeprecatedV1ValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListener_DeprecatedV1.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = Listener_DeprecatedV1ValidationError{}
