package cli

//go:generate go run ./tools/gentypes

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

var (
	exitStatus string
	flags      []*cliFlag
	help       bool
	info       string
	less       func(i int, j int) bool = func(i int, j int) bool {
		var left string = flags[i].long
		var right string = flags[j].long

		// Sort by long flag, unless it is not defined
		if left == "" {
			left = flags[i].short
		}
		if right == "" {
			right = flags[j].short
		}

		// Fallback to short flag if comparing same long flag (should
		// not happen)
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
)

// MaxWidth is how wide the Usage() message should be.
var MaxWidth int = 80

var (
	readme   bool
	sections []section
)

// SeeAlso is a list of related tools.
var SeeAlso []string

// TabWidth determines the indentation size.
var TabWidth int = 4

// Title is the title for the README.md generated with the --readme
// flag.
var Title string

// Version is the package version.
const Version string = "1.12.4"
