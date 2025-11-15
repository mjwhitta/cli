package cli

//go:generate go run ./tools/gentypes

import "os"

// Version is the package version.
const Version string = "1.13.2"

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

	// MaxWidth is how wide the Usage() message should be.
	MaxWidth int = 80

	// SeeAlso is a list of related tools.
	SeeAlso []string

	// TabWidth determines the indentation size.
	TabWidth int = 4

	// Title is the title for the README.md generated with the
	// --readme flag.
	Title string

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
	readme     bool
	sections   []section
)
