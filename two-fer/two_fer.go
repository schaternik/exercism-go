// Package twofer containes functions:
// - ShareWith
package twofer

import "fmt"

// ShareWith returns string
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
