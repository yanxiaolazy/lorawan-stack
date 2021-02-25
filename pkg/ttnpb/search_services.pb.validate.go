// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

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

// define the regex for a UUID once up-front
var _search_services_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// ValidateFields checks the field values on SearchApplicationsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SearchApplicationsRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = SearchApplicationsRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "id_contains":

			if utf8.RuneCountInString(m.GetIDContains()) > 50 {
				return SearchApplicationsRequestValidationError{
					field:  "id_contains",
					reason: "value length must be at most 50 runes",
				}
			}

		case "name_contains":

			if utf8.RuneCountInString(m.GetNameContains()) > 50 {
				return SearchApplicationsRequestValidationError{
					field:  "name_contains",
					reason: "value length must be at most 50 runes",
				}
			}

		case "description_contains":

			if utf8.RuneCountInString(m.GetDescriptionContains()) > 50 {
				return SearchApplicationsRequestValidationError{
					field:  "description_contains",
					reason: "value length must be at most 50 runes",
				}
			}

		case "attributes_contain":

			if len(m.GetAttributesContain()) > 10 {
				return SearchApplicationsRequestValidationError{
					field:  "attributes_contain",
					reason: "value must contain no more than 10 pair(s)",
				}
			}

			for key, val := range m.GetAttributesContain() {
				_ = val

				if utf8.RuneCountInString(key) > 36 {
					return SearchApplicationsRequestValidationError{
						field:  fmt.Sprintf("attributes_contain[%v]", key),
						reason: "value length must be at most 36 runes",
					}
				}

				if !_SearchApplicationsRequest_AttributesContain_Pattern.MatchString(key) {
					return SearchApplicationsRequestValidationError{
						field:  fmt.Sprintf("attributes_contain[%v]", key),
						reason: "value does not match regex pattern \"^[a-z0-9](?:[-]?[a-z0-9]){2,}$\"",
					}
				}

				if utf8.RuneCountInString(val) > 50 {
					return SearchApplicationsRequestValidationError{
						field:  fmt.Sprintf("attributes_contain[%v]", key),
						reason: "value length must be at most 50 runes",
					}
				}

			}

		case "field_mask":

			if v, ok := interface{}(&m.FieldMask).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return SearchApplicationsRequestValidationError{
						field:  "field_mask",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "order":
			// no validation rules for Order
		case "limit":

			if m.GetLimit() > 1000 {
				return SearchApplicationsRequestValidationError{
					field:  "limit",
					reason: "value must be less than or equal to 1000",
				}
			}

		case "page":
			// no validation rules for Page
		case "deleted":
			// no validation rules for Deleted
		default:
			return SearchApplicationsRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// SearchApplicationsRequestValidationError is the validation error returned by
// SearchApplicationsRequest.ValidateFields if the designated constraints
// aren't met.
type SearchApplicationsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SearchApplicationsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SearchApplicationsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SearchApplicationsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SearchApplicationsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SearchApplicationsRequestValidationError) ErrorName() string {
	return "SearchApplicationsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SearchApplicationsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchApplicationsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SearchApplicationsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SearchApplicationsRequestValidationError{}

var _SearchApplicationsRequest_AttributesContain_Pattern = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){2,}$")

// ValidateFields checks the field values on SearchClientsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SearchClientsRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = SearchClientsRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "id_contains":

			if utf8.RuneCountInString(m.GetIDContains()) > 50 {
				return SearchClientsRequestValidationError{
					field:  "id_contains",
					reason: "value length must be at most 50 runes",
				}
			}

		case "name_contains":

			if utf8.RuneCountInString(m.GetNameContains()) > 50 {
				return SearchClientsRequestValidationError{
					field:  "name_contains",
					reason: "value length must be at most 50 runes",
				}
			}

		case "description_contains":

			if utf8.RuneCountInString(m.GetDescriptionContains()) > 50 {
				return SearchClientsRequestValidationError{
					field:  "description_contains",
					reason: "value length must be at most 50 runes",
				}
			}

		case "attributes_contain":

			if len(m.GetAttributesContain()) > 10 {
				return SearchClientsRequestValidationError{
					field:  "attributes_contain",
					reason: "value must contain no more than 10 pair(s)",
				}
			}

			for key, val := range m.GetAttributesContain() {
				_ = val

				if utf8.RuneCountInString(key) > 36 {
					return SearchClientsRequestValidationError{
						field:  fmt.Sprintf("attributes_contain[%v]", key),
						reason: "value length must be at most 36 runes",
					}
				}

				if !_SearchClientsRequest_AttributesContain_Pattern.MatchString(key) {
					return SearchClientsRequestValidationError{
						field:  fmt.Sprintf("attributes_contain[%v]", key),
						reason: "value does not match regex pattern \"^[a-z0-9](?:[-]?[a-z0-9]){2,}$\"",
					}
				}

				if utf8.RuneCountInString(val) > 50 {
					return SearchClientsRequestValidationError{
						field:  fmt.Sprintf("attributes_contain[%v]", key),
						reason: "value length must be at most 50 runes",
					}
				}

			}

		case "field_mask":

			if v, ok := interface{}(&m.FieldMask).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return SearchClientsRequestValidationError{
						field:  "field_mask",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "order":
			// no validation rules for Order
		case "limit":

			if m.GetLimit() > 1000 {
				return SearchClientsRequestValidationError{
					field:  "limit",
					reason: "value must be less than or equal to 1000",
				}
			}

		case "page":
			// no validation rules for Page
		case "deleted":
			// no validation rules for Deleted
		default:
			return SearchClientsRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// SearchClientsRequestValidationError is the validation error returned by
// SearchClientsRequest.ValidateFields if the designated constraints aren't met.
type SearchClientsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SearchClientsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SearchClientsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SearchClientsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SearchClientsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SearchClientsRequestValidationError) ErrorName() string {
	return "SearchClientsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SearchClientsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchClientsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SearchClientsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SearchClientsRequestValidationError{}

var _SearchClientsRequest_AttributesContain_Pattern = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){2,}$")

// ValidateFields checks the field values on SearchGatewaysRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SearchGatewaysRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = SearchGatewaysRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "id_contains":

			if utf8.RuneCountInString(m.GetIDContains()) > 50 {
				return SearchGatewaysRequestValidationError{
					field:  "id_contains",
					reason: "value length must be at most 50 runes",
				}
			}

		case "name_contains":

			if utf8.RuneCountInString(m.GetNameContains()) > 50 {
				return SearchGatewaysRequestValidationError{
					field:  "name_contains",
					reason: "value length must be at most 50 runes",
				}
			}

		case "description_contains":

			if utf8.RuneCountInString(m.GetDescriptionContains()) > 50 {
				return SearchGatewaysRequestValidationError{
					field:  "description_contains",
					reason: "value length must be at most 50 runes",
				}
			}

		case "attributes_contain":

			if len(m.GetAttributesContain()) > 10 {
				return SearchGatewaysRequestValidationError{
					field:  "attributes_contain",
					reason: "value must contain no more than 10 pair(s)",
				}
			}

			for key, val := range m.GetAttributesContain() {
				_ = val

				if utf8.RuneCountInString(key) > 36 {
					return SearchGatewaysRequestValidationError{
						field:  fmt.Sprintf("attributes_contain[%v]", key),
						reason: "value length must be at most 36 runes",
					}
				}

				if !_SearchGatewaysRequest_AttributesContain_Pattern.MatchString(key) {
					return SearchGatewaysRequestValidationError{
						field:  fmt.Sprintf("attributes_contain[%v]", key),
						reason: "value does not match regex pattern \"^[a-z0-9](?:[-]?[a-z0-9]){2,}$\"",
					}
				}

				if utf8.RuneCountInString(val) > 50 {
					return SearchGatewaysRequestValidationError{
						field:  fmt.Sprintf("attributes_contain[%v]", key),
						reason: "value length must be at most 50 runes",
					}
				}

			}

		case "eui_contains":

			if utf8.RuneCountInString(m.GetEUIContains()) > 16 {
				return SearchGatewaysRequestValidationError{
					field:  "eui_contains",
					reason: "value length must be at most 16 runes",
				}
			}

		case "field_mask":

			if v, ok := interface{}(&m.FieldMask).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return SearchGatewaysRequestValidationError{
						field:  "field_mask",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "order":
			// no validation rules for Order
		case "limit":

			if m.GetLimit() > 1000 {
				return SearchGatewaysRequestValidationError{
					field:  "limit",
					reason: "value must be less than or equal to 1000",
				}
			}

		case "page":
			// no validation rules for Page
		case "deleted":
			// no validation rules for Deleted
		default:
			return SearchGatewaysRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// SearchGatewaysRequestValidationError is the validation error returned by
// SearchGatewaysRequest.ValidateFields if the designated constraints aren't met.
type SearchGatewaysRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SearchGatewaysRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SearchGatewaysRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SearchGatewaysRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SearchGatewaysRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SearchGatewaysRequestValidationError) ErrorName() string {
	return "SearchGatewaysRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SearchGatewaysRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchGatewaysRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SearchGatewaysRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SearchGatewaysRequestValidationError{}

var _SearchGatewaysRequest_AttributesContain_Pattern = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){2,}$")

// ValidateFields checks the field values on SearchOrganizationsRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *SearchOrganizationsRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = SearchOrganizationsRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "id_contains":

			if utf8.RuneCountInString(m.GetIDContains()) > 50 {
				return SearchOrganizationsRequestValidationError{
					field:  "id_contains",
					reason: "value length must be at most 50 runes",
				}
			}

		case "name_contains":

			if utf8.RuneCountInString(m.GetNameContains()) > 50 {
				return SearchOrganizationsRequestValidationError{
					field:  "name_contains",
					reason: "value length must be at most 50 runes",
				}
			}

		case "description_contains":

			if utf8.RuneCountInString(m.GetDescriptionContains()) > 50 {
				return SearchOrganizationsRequestValidationError{
					field:  "description_contains",
					reason: "value length must be at most 50 runes",
				}
			}

		case "attributes_contain":

			if len(m.GetAttributesContain()) > 10 {
				return SearchOrganizationsRequestValidationError{
					field:  "attributes_contain",
					reason: "value must contain no more than 10 pair(s)",
				}
			}

			for key, val := range m.GetAttributesContain() {
				_ = val

				if utf8.RuneCountInString(key) > 36 {
					return SearchOrganizationsRequestValidationError{
						field:  fmt.Sprintf("attributes_contain[%v]", key),
						reason: "value length must be at most 36 runes",
					}
				}

				if !_SearchOrganizationsRequest_AttributesContain_Pattern.MatchString(key) {
					return SearchOrganizationsRequestValidationError{
						field:  fmt.Sprintf("attributes_contain[%v]", key),
						reason: "value does not match regex pattern \"^[a-z0-9](?:[-]?[a-z0-9]){2,}$\"",
					}
				}

				if utf8.RuneCountInString(val) > 50 {
					return SearchOrganizationsRequestValidationError{
						field:  fmt.Sprintf("attributes_contain[%v]", key),
						reason: "value length must be at most 50 runes",
					}
				}

			}

		case "field_mask":

			if v, ok := interface{}(&m.FieldMask).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return SearchOrganizationsRequestValidationError{
						field:  "field_mask",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "order":
			// no validation rules for Order
		case "limit":

			if m.GetLimit() > 1000 {
				return SearchOrganizationsRequestValidationError{
					field:  "limit",
					reason: "value must be less than or equal to 1000",
				}
			}

		case "page":
			// no validation rules for Page
		case "deleted":
			// no validation rules for Deleted
		default:
			return SearchOrganizationsRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// SearchOrganizationsRequestValidationError is the validation error returned
// by SearchOrganizationsRequest.ValidateFields if the designated constraints
// aren't met.
type SearchOrganizationsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SearchOrganizationsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SearchOrganizationsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SearchOrganizationsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SearchOrganizationsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SearchOrganizationsRequestValidationError) ErrorName() string {
	return "SearchOrganizationsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SearchOrganizationsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchOrganizationsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SearchOrganizationsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SearchOrganizationsRequestValidationError{}

var _SearchOrganizationsRequest_AttributesContain_Pattern = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){2,}$")

// ValidateFields checks the field values on SearchUsersRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SearchUsersRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = SearchUsersRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "id_contains":

			if utf8.RuneCountInString(m.GetIDContains()) > 50 {
				return SearchUsersRequestValidationError{
					field:  "id_contains",
					reason: "value length must be at most 50 runes",
				}
			}

		case "name_contains":

			if utf8.RuneCountInString(m.GetNameContains()) > 50 {
				return SearchUsersRequestValidationError{
					field:  "name_contains",
					reason: "value length must be at most 50 runes",
				}
			}

		case "description_contains":

			if utf8.RuneCountInString(m.GetDescriptionContains()) > 50 {
				return SearchUsersRequestValidationError{
					field:  "description_contains",
					reason: "value length must be at most 50 runes",
				}
			}

		case "attributes_contain":

			if len(m.GetAttributesContain()) > 10 {
				return SearchUsersRequestValidationError{
					field:  "attributes_contain",
					reason: "value must contain no more than 10 pair(s)",
				}
			}

			for key, val := range m.GetAttributesContain() {
				_ = val

				if utf8.RuneCountInString(key) > 36 {
					return SearchUsersRequestValidationError{
						field:  fmt.Sprintf("attributes_contain[%v]", key),
						reason: "value length must be at most 36 runes",
					}
				}

				if !_SearchUsersRequest_AttributesContain_Pattern.MatchString(key) {
					return SearchUsersRequestValidationError{
						field:  fmt.Sprintf("attributes_contain[%v]", key),
						reason: "value does not match regex pattern \"^[a-z0-9](?:[-]?[a-z0-9]){2,}$\"",
					}
				}

				if utf8.RuneCountInString(val) > 50 {
					return SearchUsersRequestValidationError{
						field:  fmt.Sprintf("attributes_contain[%v]", key),
						reason: "value length must be at most 50 runes",
					}
				}

			}

		case "field_mask":

			if v, ok := interface{}(&m.FieldMask).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return SearchUsersRequestValidationError{
						field:  "field_mask",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "order":
			// no validation rules for Order
		case "limit":

			if m.GetLimit() > 1000 {
				return SearchUsersRequestValidationError{
					field:  "limit",
					reason: "value must be less than or equal to 1000",
				}
			}

		case "page":
			// no validation rules for Page
		case "deleted":
			// no validation rules for Deleted
		default:
			return SearchUsersRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// SearchUsersRequestValidationError is the validation error returned by
// SearchUsersRequest.ValidateFields if the designated constraints aren't met.
type SearchUsersRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SearchUsersRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SearchUsersRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SearchUsersRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SearchUsersRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SearchUsersRequestValidationError) ErrorName() string {
	return "SearchUsersRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SearchUsersRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchUsersRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SearchUsersRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SearchUsersRequestValidationError{}

var _SearchUsersRequest_AttributesContain_Pattern = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){2,}$")

// ValidateFields checks the field values on SearchEndDevicesRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SearchEndDevicesRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = SearchEndDevicesRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "application_ids":

			if v, ok := interface{}(&m.ApplicationIdentifiers).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return SearchEndDevicesRequestValidationError{
						field:  "application_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "id_contains":

			if utf8.RuneCountInString(m.GetIDContains()) > 50 {
				return SearchEndDevicesRequestValidationError{
					field:  "id_contains",
					reason: "value length must be at most 50 runes",
				}
			}

		case "name_contains":

			if utf8.RuneCountInString(m.GetNameContains()) > 50 {
				return SearchEndDevicesRequestValidationError{
					field:  "name_contains",
					reason: "value length must be at most 50 runes",
				}
			}

		case "description_contains":

			if utf8.RuneCountInString(m.GetDescriptionContains()) > 50 {
				return SearchEndDevicesRequestValidationError{
					field:  "description_contains",
					reason: "value length must be at most 50 runes",
				}
			}

		case "attributes_contain":

			if len(m.GetAttributesContain()) > 10 {
				return SearchEndDevicesRequestValidationError{
					field:  "attributes_contain",
					reason: "value must contain no more than 10 pair(s)",
				}
			}

			for key, val := range m.GetAttributesContain() {
				_ = val

				if utf8.RuneCountInString(key) > 36 {
					return SearchEndDevicesRequestValidationError{
						field:  fmt.Sprintf("attributes_contain[%v]", key),
						reason: "value length must be at most 36 runes",
					}
				}

				if !_SearchEndDevicesRequest_AttributesContain_Pattern.MatchString(key) {
					return SearchEndDevicesRequestValidationError{
						field:  fmt.Sprintf("attributes_contain[%v]", key),
						reason: "value does not match regex pattern \"^[a-z0-9](?:[-]?[a-z0-9]){2,}$\"",
					}
				}

				if utf8.RuneCountInString(val) > 50 {
					return SearchEndDevicesRequestValidationError{
						field:  fmt.Sprintf("attributes_contain[%v]", key),
						reason: "value length must be at most 50 runes",
					}
				}

			}

		case "dev_eui_contains":

			if utf8.RuneCountInString(m.GetDevEUIContains()) > 16 {
				return SearchEndDevicesRequestValidationError{
					field:  "dev_eui_contains",
					reason: "value length must be at most 16 runes",
				}
			}

		case "join_eui_contains":

			if utf8.RuneCountInString(m.GetJoinEUIContains()) > 16 {
				return SearchEndDevicesRequestValidationError{
					field:  "join_eui_contains",
					reason: "value length must be at most 16 runes",
				}
			}

		case "dev_addr_contains":

			if utf8.RuneCountInString(m.GetDevAddrContains()) > 8 {
				return SearchEndDevicesRequestValidationError{
					field:  "dev_addr_contains",
					reason: "value length must be at most 8 runes",
				}
			}

		case "field_mask":

			if v, ok := interface{}(&m.FieldMask).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return SearchEndDevicesRequestValidationError{
						field:  "field_mask",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "order":
			// no validation rules for Order
		case "limit":

			if m.GetLimit() > 1000 {
				return SearchEndDevicesRequestValidationError{
					field:  "limit",
					reason: "value must be less than or equal to 1000",
				}
			}

		case "page":
			// no validation rules for Page
		default:
			return SearchEndDevicesRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// SearchEndDevicesRequestValidationError is the validation error returned by
// SearchEndDevicesRequest.ValidateFields if the designated constraints aren't met.
type SearchEndDevicesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SearchEndDevicesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SearchEndDevicesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SearchEndDevicesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SearchEndDevicesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SearchEndDevicesRequestValidationError) ErrorName() string {
	return "SearchEndDevicesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SearchEndDevicesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchEndDevicesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SearchEndDevicesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SearchEndDevicesRequestValidationError{}

var _SearchEndDevicesRequest_AttributesContain_Pattern = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){2,}$")
