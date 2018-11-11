// Package twofer implements a simple function to return
// specific strings given specific inputs
package twofer

import (
	"fmt"
)

// ShareWith takes a name of type string and returns
// a string in the form of 'One for <name>, one for me.
// If passed an empty string, the function will replace
// <name> with 'you'.
func ShareWith(name string) string {
	buf := name
	if len(buf) == 0 {
		buf = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", buf)
}
