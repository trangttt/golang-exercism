// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer
import (
    "strings"
    "fmt"
)

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
    name = strings.Trim(name, " ")
    if name == "" {
        return "One for you, one for me."
    } else {
        return fmt.Sprintf("One for %s, one for me.", name)
    }
}
