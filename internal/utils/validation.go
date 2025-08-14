package utils

import (
	"mime"
	"net/http"
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
