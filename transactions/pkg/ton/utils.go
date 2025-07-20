package ton

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

// ValidateUserID validates and sanitizes the userId string
func ValidateUserID(userId string) (*int64, bool) {
	// Check if the string is valid UTF-8
	if !utf8.ValidString(userId) {
		return nil, false
	}

	// Additional validation rules (customize based on your requirements)
	userId = strings.TrimSpace(userId)

	// Check if empty after trimming
	if userId == "" {
		return nil, false
	}

	// Optional: Check length limits
	if len(userId) > 255 {
		return nil, false
	}

	// Optional: Check for printable characters only
	for _, r := range userId {
		if !utf8.ValidRune(r) || r < 32 || r == 127 {
			return nil, false
		}
	}

	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return nil, false
	}

	return &userIdInt, true
}
