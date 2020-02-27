package helpers

import "errors"

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
