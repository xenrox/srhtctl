package helpers

import (
	"errors"
	"strings"

	"git.xenrox.net/~xenrox/srhtctl/config"
	"github.com/atotto/clipboard"
)

// ValidateVisibility checks whether the visibility value is valid
func ValidateVisibility(visibility string) error {
	values := [3]string{"public", "unlisted", "private"}
	for _, value := range values {
		if visibility == value {
			return nil
		}
	}
	return errors.New("Not a valid visibility")
}

// TransformTags formats a tag string as an array
func TransformTags(tagString *string) []string {
	if tagString == nil {
		return make([]string, 0)
	}
	return strings.Split(*tagString, "/")
}

// CopyToClipboard copys text to clipboard if set by user
func CopyToClipboard(text string) {
	if config.GetConfigValue("settings", "copyToClipboard", "false") == "true" {
		clipboard.WriteAll(text)
	}
}
