// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

// Hey should have a comment documenting it.
func Hey(remark string) string {
    ret := ""
    remark = strings.Trim(remark, " \t\n\r")
    if remark == "" {
        ret = "Fine. Be that way!"
    } else if strings.ToUpper(remark) == remark && strings.ContainsAny(remark, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
        // shouting
        if strings.HasSuffix(remark, "?") {
            ret = "Calm down, I know what I'm doing!"
        } else {
            ret = "Whoa, chill out!"
        }
    } else {
        if strings.HasSuffix(remark, "?") {
            ret = "Sure."
        } else {
            ret = "Whatever."
        }
    }
	return ret
}
