package utils

import (
	"mime"
	"net/http"
	"reflect"
)

type PatchMediaType string

const (
	JSONPatch  PatchMediaType = "json-patch"
	MergePatch PatchMediaType = "merge-patch"
)

func DetectPatchMediaType(headers http.Header) *PatchMediaType {
	ct := headers.Get("Content-Type")
	mediaType, _, err := mime.ParseMediaType(ct)
	if err != nil {
		return nil
	}

	switch mediaType {
	case "application/json-patch+json":
		pt := JSONPatch
		return &pt
	case "application/merge-patch+json":
		pt := MergePatch
		return &pt
	default:
		return nil
	}
}

func ValidatePaginationParams(offset *int64, limit *int64) (*int64, *int64) {
	// Handle nil values for limit and offset
	var offsetDefault int64 = 0 // Default offset
	var limitDefault int64 = 25 // Default limit
	var limitMax int64 = 500    // Max limit

	if offset == nil {
		offset = &offsetDefault
	}
	if limit == nil {
		limit = &limitDefault
	}
	if limit != nil && *limit > limitMax {
		limit = &limitMax
	}

	return offset, limit
}

// OnlyFieldSet returns true if only the specified field is set and all other fields are zero/nil/empty.
// `obj` must be a pointer to a struct, `fieldName` is the field expected to be set.
func IsOnlyFieldSet(obj interface{}, fieldName string) bool {
	if obj == nil {
		return false
	}

	val := reflect.ValueOf(obj)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		// not a pointer to struct
		return false
	}

	val = val.Elem() // get struct value
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		f := val.Field(i)
		name := typ.Field(i).Name

		if name == fieldName {
			// Check that the target field is set (non-zero)
			if isZeroValue(f) {
				return false
			}
			continue
		}

		// All other fields must be zero/nil/empty
		if !isZeroValue(f) {
			return false
		}
	}

	return true
}

// isZeroValue returns true if v is nil, empty slice, or zero value.
func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	case reflect.Slice, reflect.Array, reflect.Map:
		return v.Len() == 0
	default:
		zero := reflect.Zero(v.Type())
		return reflect.DeepEqual(v.Interface(), zero.Interface())
	}
}
