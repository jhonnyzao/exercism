package twofer

import (
	"fmt"
)

// ShareWith should return a string with the people names.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
