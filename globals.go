package cli

//go:generate ./scripts/generate_go_funcs

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
var Banner = fmt.Sprintf("%s [OPTIONS]", os.Args[0])

// BugEmail is the configured email to send bug reports to.
var BugEmail string

var colWidth = columnWidth{0, 0, 1024, 0}

// ExitStatus is the description of program exit status.
var ExitStatus string

var flags []flagVar
var help bool

// Info is the description of how the program works.
var Info = ""

var less = func(i, j int) bool {
	// Sort by short flag
	var left = flags[i].short
	if len(left) == 0 {
		left = flags[i].long
	}

	var right = flags[j].short
	if len(right) == 0 {
		right = flags[j].long
	}

	// Fallback to long flag if comparing same short flag
	if strings.ToLower(left) == strings.ToLower(right) {
		if len(flags[i].long) != 0 {
			left = flags[i].long
		}
		if len(flags[j].long) != 0 {
			right = flags[j].long
		}
	}

	return (strings.ToLower(left) < strings.ToLower(right))
}

// MaxWidth is how wide the Usage() message should be.
var MaxWidth = 80

var readme bool

// SeeAlso is a list of related tools.
var SeeAlso []string

// TabWidth determines the indentation size.
var TabWidth = 4

// Title is the title for the README.md generated with the --readme
// flag.
var Title string

// Version is the package version.
const Version = "1.7.5"
