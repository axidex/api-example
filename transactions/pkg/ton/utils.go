package ton

import (
	"strings"
	"unicode/utf8"
)

// ValidateUserID validates and sanitizes the userId string
func ValidateUserID(userId string) (string, bool) {
	// Check if the string is valid UTF-8
	if !utf8.ValidString(userId) {
		return "", false
	}

	// Additional validation rules (customize based on your requirements)
	userId = strings.TrimSpace(userId)

	// Check if empty after trimming
	if userId == "" {
		return "", false
	}

	// Optional: Check length limits
	if len(userId) > 255 {
		return "", false
	}

	// Optional: Check for printable characters only
	for _, r := range userId {
		if !utf8.ValidRune(r) || r < 32 || r == 127 {
			return "", false
		}
	}

	return userId, true
}
