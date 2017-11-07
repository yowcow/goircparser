package goircparser

import (
	"regexp"
	"strings"
)

// Row represents a parsed IRC command row
type Row struct {
	RawLine string
	Prefix  string
	Command string
	Params  []string
	Suffix  string
}

/**
 * \x00 -> NUL
 * \x20 -> Space
 * \x0A -> LF
 * \x0D -> CR
 * \x3A -> :
 */
var re = regexp.MustCompile(`^(?:\:(\S+)\s)?([A-Za-z]+|\d{3})(?:\s((?:[^\x00\x20\x0D\x0A\x3A][^\x00\x0A\x0D\x20]*)(?:\s[^\x00\x20\x0D\x0A\x3A][^\x00\x0A\x0D\x20]*)*))?(?:\s\:([^\x00\x0A\x0D]*))?$`)

// Parse parses a IRC command line and transforms into a *Row
func Parse(line string) *Row {
	matches := re.FindStringSubmatch(line)
	return &Row{
		RawLine: line,
		Prefix:  matches[1],
		Command: matches[2],
		Params:  strings.Split(matches[3], " "),
		Suffix:  matches[4],
	}
}
