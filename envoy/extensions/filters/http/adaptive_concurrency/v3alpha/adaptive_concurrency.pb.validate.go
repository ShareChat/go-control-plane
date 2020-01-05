// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/http/adaptive_concurrency/v3alpha/adaptive_concurrency.proto

package envoy_extensions_filters_http_adaptive_concurrency_v3alpha

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

// define the regex for a UUID once up-front
var _adaptive_concurrency_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on GradientControllerConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GradientControllerConfig) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetSampleAggregatePercentile()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GradientControllerConfigValidationError{
				field:  "SampleAggregatePercentile",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetConcurrencyLimitParams() == nil {
		return GradientControllerConfigValidationError{
			field:  "ConcurrencyLimitParams",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetConcurrencyLimitParams()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GradientControllerConfigValidationError{
				field:  "ConcurrencyLimitParams",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetMinRttCalcParams() == nil {
		return GradientControllerConfigValidationError{
			field:  "MinRttCalcParams",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetMinRttCalcParams()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GradientControllerConfigValidationError{
				field:  "MinRttCalcParams",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GradientControllerConfigValidationError is the validation error returned by
// GradientControllerConfig.Validate if the designated constraints aren't met.
type GradientControllerConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GradientControllerConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GradientControllerConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GradientControllerConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GradientControllerConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GradientControllerConfigValidationError) ErrorName() string {
	return "GradientControllerConfigValidationError"
}

// Error satisfies the builtin error interface
func (e GradientControllerConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGradientControllerConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GradientControllerConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GradientControllerConfigValidationError{}

// Validate checks the field values on AdaptiveConcurrency with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AdaptiveConcurrency) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetEnabled()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AdaptiveConcurrencyValidationError{
				field:  "Enabled",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.ConcurrencyControllerConfig.(type) {

	case *AdaptiveConcurrency_GradientControllerConfig:

		if m.GetGradientControllerConfig() == nil {
			return AdaptiveConcurrencyValidationError{
				field:  "GradientControllerConfig",
				reason: "value is required",
			}
		}

		if v, ok := interface{}(m.GetGradientControllerConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AdaptiveConcurrencyValidationError{
					field:  "GradientControllerConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return AdaptiveConcurrencyValidationError{
			field:  "ConcurrencyControllerConfig",
			reason: "value is required",
		}

	}

	return nil
}

// AdaptiveConcurrencyValidationError is the validation error returned by
// AdaptiveConcurrency.Validate if the designated constraints aren't met.
type AdaptiveConcurrencyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AdaptiveConcurrencyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AdaptiveConcurrencyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AdaptiveConcurrencyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AdaptiveConcurrencyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AdaptiveConcurrencyValidationError) ErrorName() string {
	return "AdaptiveConcurrencyValidationError"
}

// Error satisfies the builtin error interface
func (e AdaptiveConcurrencyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAdaptiveConcurrency.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AdaptiveConcurrencyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AdaptiveConcurrencyValidationError{}

// Validate checks the field values on
// GradientControllerConfig_ConcurrencyLimitCalculationParams with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) Validate() error {
	if m == nil {
		return nil
	}

	if wrapper := m.GetMaxConcurrencyLimit(); wrapper != nil {

		if wrapper.GetValue() <= 0 {
			return GradientControllerConfig_ConcurrencyLimitCalculationParamsValidationError{
				field:  "MaxConcurrencyLimit",
				reason: "value must be greater than 0",
			}
		}

	}

	if m.GetConcurrencyUpdateInterval() == nil {
		return GradientControllerConfig_ConcurrencyLimitCalculationParamsValidationError{
			field:  "ConcurrencyUpdateInterval",
			reason: "value is required",
		}
	}

	if d := m.GetConcurrencyUpdateInterval(); d != nil {
		dur, err := ptypes.Duration(d)
		if err != nil {
			return GradientControllerConfig_ConcurrencyLimitCalculationParamsValidationError{
				field:  "ConcurrencyUpdateInterval",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		gt := time.Duration(0*time.Second + 0*time.Nanosecond)

		if dur <= gt {
			return GradientControllerConfig_ConcurrencyLimitCalculationParamsValidationError{
				field:  "ConcurrencyUpdateInterval",
				reason: "value must be greater than 0s",
			}
		}

	}

	return nil
}

// GradientControllerConfig_ConcurrencyLimitCalculationParamsValidationError is
// the validation error returned by
// GradientControllerConfig_ConcurrencyLimitCalculationParams.Validate if the
// designated constraints aren't met.
type GradientControllerConfig_ConcurrencyLimitCalculationParamsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GradientControllerConfig_ConcurrencyLimitCalculationParamsValidationError) Field() string {
	return e.field
}

// Reason function returns reason value.
func (e GradientControllerConfig_ConcurrencyLimitCalculationParamsValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e GradientControllerConfig_ConcurrencyLimitCalculationParamsValidationError) Cause() error {
	return e.cause
}

// Key function returns key value.
func (e GradientControllerConfig_ConcurrencyLimitCalculationParamsValidationError) Key() bool {
	return e.key
}

// ErrorName returns error name.
func (e GradientControllerConfig_ConcurrencyLimitCalculationParamsValidationError) ErrorName() string {
	return "GradientControllerConfig_ConcurrencyLimitCalculationParamsValidationError"
}

// Error satisfies the builtin error interface
func (e GradientControllerConfig_ConcurrencyLimitCalculationParamsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGradientControllerConfig_ConcurrencyLimitCalculationParams.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GradientControllerConfig_ConcurrencyLimitCalculationParamsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GradientControllerConfig_ConcurrencyLimitCalculationParamsValidationError{}

// Validate checks the field values on
// GradientControllerConfig_MinimumRTTCalculationParams with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GradientControllerConfig_MinimumRTTCalculationParams) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetInterval() == nil {
		return GradientControllerConfig_MinimumRTTCalculationParamsValidationError{
			field:  "Interval",
			reason: "value is required",
		}
	}

	if d := m.GetInterval(); d != nil {
		dur, err := ptypes.Duration(d)
		if err != nil {
			return GradientControllerConfig_MinimumRTTCalculationParamsValidationError{
				field:  "Interval",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		gt := time.Duration(0*time.Second + 0*time.Nanosecond)

		if dur <= gt {
			return GradientControllerConfig_MinimumRTTCalculationParamsValidationError{
				field:  "Interval",
				reason: "value must be greater than 0s",
			}
		}

	}

	if wrapper := m.GetRequestCount(); wrapper != nil {

		if wrapper.GetValue() <= 0 {
			return GradientControllerConfig_MinimumRTTCalculationParamsValidationError{
				field:  "RequestCount",
				reason: "value must be greater than 0",
			}
		}

	}

	if v, ok := interface{}(m.GetJitter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GradientControllerConfig_MinimumRTTCalculationParamsValidationError{
				field:  "Jitter",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if wrapper := m.GetMinConcurrency(); wrapper != nil {

		if wrapper.GetValue() <= 0 {
			return GradientControllerConfig_MinimumRTTCalculationParamsValidationError{
				field:  "MinConcurrency",
				reason: "value must be greater than 0",
			}
		}

	}

	if v, ok := interface{}(m.GetBuffer()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GradientControllerConfig_MinimumRTTCalculationParamsValidationError{
				field:  "Buffer",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GradientControllerConfig_MinimumRTTCalculationParamsValidationError is the
// validation error returned by
// GradientControllerConfig_MinimumRTTCalculationParams.Validate if the
// designated constraints aren't met.
type GradientControllerConfig_MinimumRTTCalculationParamsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GradientControllerConfig_MinimumRTTCalculationParamsValidationError) Field() string {
	return e.field
}

// Reason function returns reason value.
func (e GradientControllerConfig_MinimumRTTCalculationParamsValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e GradientControllerConfig_MinimumRTTCalculationParamsValidationError) Cause() error {
	return e.cause
}

// Key function returns key value.
func (e GradientControllerConfig_MinimumRTTCalculationParamsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GradientControllerConfig_MinimumRTTCalculationParamsValidationError) ErrorName() string {
	return "GradientControllerConfig_MinimumRTTCalculationParamsValidationError"
}

// Error satisfies the builtin error interface
func (e GradientControllerConfig_MinimumRTTCalculationParamsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGradientControllerConfig_MinimumRTTCalculationParams.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GradientControllerConfig_MinimumRTTCalculationParamsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GradientControllerConfig_MinimumRTTCalculationParamsValidationError{}