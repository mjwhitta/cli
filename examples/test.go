package main

import (
	"fmt"
	"os"
	"strings"

	// "cli"
	"gitlab.com/mjwhitta/cli"
)

var abool bool
var afloat float64
var aint int64
var astring string
var astringlist cli.StringList
var auint uint64

func init() {
	// Configure cli package
	cli.Align = false // Default
	// cli.Align = true
	cli.Authors = []string{"Miles W <mjwhitta@some.domain>"}
	cli.Banner = fmt.Sprintf("%s [OPTIONS] <arg>", os.Args[0])
	cli.BugEmail = "bugs@some.domain"
	cli.ExitStatus = strings.Join(
		[]string{
			"Normally the exit status is 0. In the event of invalid",
			"or missing arguments, the exit status will be non-zero.",
		},
		" ",
	)
	cli.Info = strings.Join(
		[]string{
			"Lorem ipsum dolor sit amet, consectetur adipiscing",
			"elit. Mauris ut augue euismod, cursus nulla ut, semper",
			"eros. Integer pulvinar a lectus sed pretium. Cras a",
			"luctus odio, eget sagittis leo. Interdum et malesuada",
			"fames ac ante ipsum primis in faucibus. Nunc lectus",
			"metus, consectetur et tellus tempus, accumsan rhoncus",
			"arcu. Ut velit dui, aliquet eget tellus eget, commodo",
			"auctor enim. Duis blandit, metus vitae sagittis",
			"vestibulum, diam nisl fermentum lorem, non convallis",
			"enim urna rutrum mauris. Duis in scelerisque mauris.",
		},
		"",
	)
	cli.MaxWidth = 80 // Default
	cli.SeeAlso = []string{
		"flag",
	}
	cli.TabWidth = 4 // Default
	cli.Title = "Sample README.md"

	// Parse cli flags
	cli.Flag(
		&abool,
		"a",
		"praesent",
		// "praesentpraesentpraesent",
		false,
		strings.Join(
			[]string{
				"Fusce blandit nisi eu sem maximus, ut eleifend",
				"lectus semper.",
			},
			" ",
		),
	)
	cli.Flag(
		&afloat,
		"b",
		"facilisis",
		0.0,
		strings.Join(
			[]string{
				"Cras ac dolor ante. Nulla posuere non purus ac",
				"vehicula.",
			},
			" ",
		),
	)
	cli.Flag(
		&aint,
		"c",
		0,
		strings.Join(
			[]string{
				"Nullam efficitur elit vel venenatis mattis. Aenean",
				"viverra tincidunt ligula.",
			},
			" ",
		),
	)
	cli.Flag(
		&astring,
		"d",
		"ut",
		"",
		strings.Join(
			[]string{
				"Nulla a sollicitudin ex. Sed faucibus lacus congue",
				"dapibus vestibulum.",
			},
			" ",
		),
	)
	cli.Flag(
		&astringlist,
		"e",
		"placerat",
		strings.Join(
			[]string{
				"Etiam efficitur interdum leo nec luctus. Etiam est",
				"erat, bibendum.",
			},
			" ",
		),
	)
	cli.Flag(
		&auint,
		"sagittis",
		0,
		strings.Join(
			[]string{
				"Donec posuere efficitur massa, ut imperdiet mauris",
				"sodales sit amet.",
			},
			" ",
		),
	)
	cli.Parse()

	// Validate cli args
	if cli.NArg() == 0 {
		cli.Usage(1)
	} else if len(astring) == 0 {
		cli.Usage(2)
	}
}

func main() {
	fmt.Printf("%t\n", abool)
	fmt.Printf("%f\n", afloat)
	fmt.Printf("%d\n", aint)
	fmt.Printf("%s\n", astring)
	fmt.Printf("%d - %s\n", len(astringlist), astringlist)
	fmt.Printf("%d\n", auint)
	fmt.Printf("%d %s\n", cli.NArg(), cli.Args())
}
