// Package twofer provides a function to generate dialogue for sharing cookies.

package twofer
import "fmt"

const defaultName = "you"

// The ShareWith function determines the appropriate dialogue based on a given name.
func ShareWith(name string) string {
	if name == "" {
		name = defaultName
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}

