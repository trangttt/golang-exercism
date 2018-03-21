// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import "strings"
import "fmt"

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
    fmt.Println(s)
    s = strings.Replace(s, "-", " ", -1)
    parts := strings.Split(s, " ")
    for i := 0; i < len(parts); i++ {
        fmt.Println(parts[i])
        parts[i] = strings.ToUpper(string(parts[i][0]))
        fmt.Println(i, parts[i])
    }
	return strings.Join(parts, "")
}
