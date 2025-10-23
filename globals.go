package cli

//go:generate go run ./tools/gentypes

import (
	"os"
	"strings"
)

// Version is the package version.
const Version string = "1.13.1"

var (
	// Align determines if short and long cli flags are aligned or
	// not.
	Align = false

	// Authors is the configured list of authors.
	Authors []string

	// Banner is the initial Usage() line.
	Banner string = os.Args[0] + " [OPTIONS]"

	// BugEmail is the configured email to send bug reports to.
	BugEmail string

	colWidth = columnWidth{
		desc:  1024, //nolint:mnd // start with big number
		left:  0,
		long:  0,
		short: 0,
	}
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

	// MaxWidth is how wide the Usage() message should be.
	MaxWidth int = 80

	readme   bool
	sections []section

	// SeeAlso is a list of related tools.
	SeeAlso []string

	// TabWidth determines the indentation size.
	TabWidth int = 4

	// Title is the title for the README.md generated with the
	// --readme flag.
	Title string
)
