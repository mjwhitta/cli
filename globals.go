package cli

import (
	"fmt"
	"os"
	"strings"
)

// Align determines if short and long cli flags are aligned or not.
var Align = false

// Authors is the configured list of authors.
var Authors []string

// Banner is the initial Usage() line.
var Banner string = fmt.Sprintf("%s [OPTIONS]", os.Args[0])

// BugEmail is the configured email to send bug reports to.
var BugEmail string

var colWidth = columnWidth{
	desc:  1024,
	left:  0,
	long:  0,
	short: 0,
}

// ExitStatus is the description of program exit status.
var ExitStatus string

var flags []*cliFlag
var help bool

// Info is the description of how the program works.
var Info string

var less = func(i, j int) bool {
	// Sort by long flag
	var left = flags[i].long
	if left == "" {
		left = flags[i].short
	}

	var right = flags[j].long
	if right == "" {
		right = flags[j].short
	}

	// Fallback to short flag if comparing same long flag (should not
	// happen)
	if strings.EqualFold(left, right) {
		if flags[i].short != "" {
			left = flags[i].short
		}

		if flags[j].short != "" {
			right = flags[j].short
		}
	}

	// Check again b/c values may have changed
	if strings.EqualFold(left, right) {
		return left < right
	}

	return strings.ToLower(left) < strings.ToLower(right)
}

// MaxWidth is how wide the Usage() message should be.
var MaxWidth int = 80

var readme bool
var sections []section

// SeeAlso is a list of related tools.
var SeeAlso []string

// TabWidth determines the indentation size.
var TabWidth int = 4

// Title is the title for the README.md generated with the --readme
// flag.
var Title string

// Version is the package version.
const Version = "1.8.2"
