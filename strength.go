package strength

import (
	"errors"
	"fmt"
	"strings"
)

// Validate returns nil if the password has more than or equal to the minimum entropy.
// If not, an error is returned explaining how the password can be strengthened.
// This error is safe to show to the client.
func Validate(password string, minEntropy float64) error {
	entropy := getEntropy(password)
	if entropy >= minEntropy {
		return nil
	}

	hasReplace := false
	hasSep := false
	hasOtherSpecial := false
	hasLower := false
	hasUpper := false
	hasDigits := false

	for _, c := range password {
		switch {
		case strings.ContainsRune(replaceChars, c):
			hasReplace = true
		case strings.ContainsRune(sepChars, c):
			hasSep = true
		case strings.ContainsRune(otherSpecialChars, c):
			hasOtherSpecial = true
		case strings.ContainsRune(lowerChars, c):
			hasLower = true
		case strings.ContainsRune(upperChars, c):
			hasUpper = true
		case strings.ContainsRune(digitsChars, c):
			hasDigits = true
		}
	}

	allMessages := []string{}

	if !hasOtherSpecial || !hasSep || !hasReplace {
		allMessages = append(allMessages, "including more special characters")
	}
	if !hasLower {
		allMessages = append(allMessages, "using lowercase letters")
	}
	if !hasUpper {
		allMessages = append(allMessages, "using uppercase letters")
	}
	if !hasDigits {
		allMessages = append(allMessages, "using numbers")
	}

	if len(allMessages) > 0 {
		return fmt.Errorf(
			"insecure password, try %v or using a longer password",
			strings.Join(allMessages, ", "),
		)
	}

	return errors.New("insecure password, try using a longer password")
}
