package utils

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
