package helpers

import "strings"

// TransformName takes a username and returns it as a canonical name
func TransformCanonical(username string) string {
	if strings.HasPrefix(username, "~") {
		return username
	}
	return "~" + username
}
