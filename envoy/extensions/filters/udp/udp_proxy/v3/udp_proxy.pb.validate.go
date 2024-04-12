// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/udp/udp_proxy/v3/udp_proxy.proto

package udp_proxyv3

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

// Validate checks the field values on UdpProxyConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UdpProxyConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UdpProxyConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UdpProxyConfigMultiError,
// or nil if none found.
func (m *UdpProxyConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *UdpProxyConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetStatPrefix()) < 1 {
		err := UdpProxyConfigValidationError{
			field:  "StatPrefix",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetIdleTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UdpProxyConfigValidationError{
					field:  "IdleTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UdpProxyConfigValidationError{
					field:  "IdleTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetIdleTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UdpProxyConfigValidationError{
				field:  "IdleTimeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for UseOriginalSrcIp

	if len(m.GetHashPolicies()) > 1 {
		err := UdpProxyConfigValidationError{
			field:  "HashPolicies",
			reason: "value must contain no more than 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetHashPolicies() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UdpProxyConfigValidationError{
						field:  fmt.Sprintf("HashPolicies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UdpProxyConfigValidationError{
						field:  fmt.Sprintf("HashPolicies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UdpProxyConfigValidationError{
					field:  fmt.Sprintf("HashPolicies[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetUpstreamSocketConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UdpProxyConfigValidationError{
					field:  "UpstreamSocketConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UdpProxyConfigValidationError{
					field:  "UpstreamSocketConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpstreamSocketConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UdpProxyConfigValidationError{
				field:  "UpstreamSocketConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for UsePerPacketLoadBalancing

	for idx, item := range m.GetAccessLog() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UdpProxyConfigValidationError{
						field:  fmt.Sprintf("AccessLog[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UdpProxyConfigValidationError{
						field:  fmt.Sprintf("AccessLog[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UdpProxyConfigValidationError{
					field:  fmt.Sprintf("AccessLog[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetProxyAccessLog() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UdpProxyConfigValidationError{
						field:  fmt.Sprintf("ProxyAccessLog[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UdpProxyConfigValidationError{
						field:  fmt.Sprintf("ProxyAccessLog[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UdpProxyConfigValidationError{
					field:  fmt.Sprintf("ProxyAccessLog[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetSessionFilters() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UdpProxyConfigValidationError{
						field:  fmt.Sprintf("SessionFilters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UdpProxyConfigValidationError{
						field:  fmt.Sprintf("SessionFilters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UdpProxyConfigValidationError{
					field:  fmt.Sprintf("SessionFilters[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	oneofRouteSpecifierPresent := false
	switch v := m.RouteSpecifier.(type) {
	case *UdpProxyConfig_Cluster:
		if v == nil {
			err := UdpProxyConfigValidationError{
				field:  "RouteSpecifier",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofRouteSpecifierPresent = true

		if utf8.RuneCountInString(m.GetCluster()) < 1 {
			err := UdpProxyConfigValidationError{
				field:  "Cluster",
				reason: "value length must be at least 1 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	case *UdpProxyConfig_Matcher:
		if v == nil {
			err := UdpProxyConfigValidationError{
				field:  "RouteSpecifier",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofRouteSpecifierPresent = true

		if all {
			switch v := interface{}(m.GetMatcher()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UdpProxyConfigValidationError{
						field:  "Matcher",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UdpProxyConfigValidationError{
						field:  "Matcher",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetMatcher()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UdpProxyConfigValidationError{
					field:  "Matcher",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofRouteSpecifierPresent {
		err := UdpProxyConfigValidationError{
			field:  "RouteSpecifier",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UdpProxyConfigMultiError(errors)
	}

	return nil
}

// UdpProxyConfigMultiError is an error wrapping multiple validation errors
// returned by UdpProxyConfig.ValidateAll() if the designated constraints
// aren't met.
type UdpProxyConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UdpProxyConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UdpProxyConfigMultiError) AllErrors() []error { return m }

// UdpProxyConfigValidationError is the validation error returned by
// UdpProxyConfig.Validate if the designated constraints aren't met.
type UdpProxyConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UdpProxyConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UdpProxyConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UdpProxyConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UdpProxyConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UdpProxyConfigValidationError) ErrorName() string { return "UdpProxyConfigValidationError" }

// Error satisfies the builtin error interface
func (e UdpProxyConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUdpProxyConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UdpProxyConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UdpProxyConfigValidationError{}

// Validate checks the field values on UdpProxyConfig_HashPolicy with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UdpProxyConfig_HashPolicy) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UdpProxyConfig_HashPolicy with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UdpProxyConfig_HashPolicyMultiError, or nil if none found.
func (m *UdpProxyConfig_HashPolicy) ValidateAll() error {
	return m.validate(true)
}

func (m *UdpProxyConfig_HashPolicy) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	oneofPolicySpecifierPresent := false
	switch v := m.PolicySpecifier.(type) {
	case *UdpProxyConfig_HashPolicy_SourceIp:
		if v == nil {
			err := UdpProxyConfig_HashPolicyValidationError{
				field:  "PolicySpecifier",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofPolicySpecifierPresent = true

		if m.GetSourceIp() != true {
			err := UdpProxyConfig_HashPolicyValidationError{
				field:  "SourceIp",
				reason: "value must equal true",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	case *UdpProxyConfig_HashPolicy_Key:
		if v == nil {
			err := UdpProxyConfig_HashPolicyValidationError{
				field:  "PolicySpecifier",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofPolicySpecifierPresent = true

		if utf8.RuneCountInString(m.GetKey()) < 1 {
			err := UdpProxyConfig_HashPolicyValidationError{
				field:  "Key",
				reason: "value length must be at least 1 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofPolicySpecifierPresent {
		err := UdpProxyConfig_HashPolicyValidationError{
			field:  "PolicySpecifier",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UdpProxyConfig_HashPolicyMultiError(errors)
	}

	return nil
}

// UdpProxyConfig_HashPolicyMultiError is an error wrapping multiple validation
// errors returned by UdpProxyConfig_HashPolicy.ValidateAll() if the
// designated constraints aren't met.
type UdpProxyConfig_HashPolicyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UdpProxyConfig_HashPolicyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UdpProxyConfig_HashPolicyMultiError) AllErrors() []error { return m }

// UdpProxyConfig_HashPolicyValidationError is the validation error returned by
// UdpProxyConfig_HashPolicy.Validate if the designated constraints aren't met.
type UdpProxyConfig_HashPolicyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UdpProxyConfig_HashPolicyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UdpProxyConfig_HashPolicyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UdpProxyConfig_HashPolicyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UdpProxyConfig_HashPolicyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UdpProxyConfig_HashPolicyValidationError) ErrorName() string {
	return "UdpProxyConfig_HashPolicyValidationError"
}

// Error satisfies the builtin error interface
func (e UdpProxyConfig_HashPolicyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUdpProxyConfig_HashPolicy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UdpProxyConfig_HashPolicyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UdpProxyConfig_HashPolicyValidationError{}

// Validate checks the field values on UdpProxyConfig_SessionFilter with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UdpProxyConfig_SessionFilter) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UdpProxyConfig_SessionFilter with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UdpProxyConfig_SessionFilterMultiError, or nil if none found.
func (m *UdpProxyConfig_SessionFilter) ValidateAll() error {
	return m.validate(true)
}

func (m *UdpProxyConfig_SessionFilter) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := UdpProxyConfig_SessionFilterValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	switch v := m.ConfigType.(type) {
	case *UdpProxyConfig_SessionFilter_TypedConfig:
		if v == nil {
			err := UdpProxyConfig_SessionFilterValidationError{
				field:  "ConfigType",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetTypedConfig()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UdpProxyConfig_SessionFilterValidationError{
						field:  "TypedConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UdpProxyConfig_SessionFilterValidationError{
						field:  "TypedConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetTypedConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UdpProxyConfig_SessionFilterValidationError{
					field:  "TypedConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return UdpProxyConfig_SessionFilterMultiError(errors)
	}

	return nil
}

// UdpProxyConfig_SessionFilterMultiError is an error wrapping multiple
// validation errors returned by UdpProxyConfig_SessionFilter.ValidateAll() if
// the designated constraints aren't met.
type UdpProxyConfig_SessionFilterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UdpProxyConfig_SessionFilterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UdpProxyConfig_SessionFilterMultiError) AllErrors() []error { return m }

// UdpProxyConfig_SessionFilterValidationError is the validation error returned
// by UdpProxyConfig_SessionFilter.Validate if the designated constraints
// aren't met.
type UdpProxyConfig_SessionFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UdpProxyConfig_SessionFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UdpProxyConfig_SessionFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UdpProxyConfig_SessionFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UdpProxyConfig_SessionFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UdpProxyConfig_SessionFilterValidationError) ErrorName() string {
	return "UdpProxyConfig_SessionFilterValidationError"
}

// Error satisfies the builtin error interface
func (e UdpProxyConfig_SessionFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUdpProxyConfig_SessionFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UdpProxyConfig_SessionFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UdpProxyConfig_SessionFilterValidationError{}
