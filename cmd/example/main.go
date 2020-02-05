package main

import (
	"fmt"
	"os"
	"strings"

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
	cli.Align = true // Defaults to false
	cli.Authors = []string{"Miles Whittaker <mj@whitta.dev>"}
	cli.Banner = fmt.Sprintf("%s [OPTIONS] <arg>", os.Args[0])
	cli.BugEmail = "cli.bugs@whitta.dev"
	cli.ExitStatus = strings.Join(
		[]string{
			"Normally the exit status is 0. In the event of an error",
			"the exit status will be one of the below:\n\n",
			"  1: Invalid option\n",
			"  2: Invalid argument\n",
			"  3: Missing argument\n",
			"  4: Extra arguments\n",
			"  5: Exception\n",
			"  6: Ambiguous argument",
		},
		" ",
	)
	cli.Info = strings.Join(
		[]string{
			"Lorem ipsum dolor sit amet, consectetur adipiscing",
			"elit. Mauris ut augue euismod, cursus nulla ut, semper",
			"eros. Integer pulvinar a lectus sed pretium. Cras a",
			"luctus odio, eget sagittis leo. Interdum et malesuada",
			"fames ac ante ipsum primis in faucibus.",
		},
		" ",
	)
	cli.MaxWidth = 80 // Defaults to 80
	cli.Section(
		"FAUCIBUS",
		strings.Join(
			[]string{
				"Nunc lectus metus, consectetur et tellus tempus,",
				"accumsan rhoncus arcu.\n\n",
				"Ut velit dui, aliquet eget tellus eget, commodo",
				"auctor enim.\n\n",
				"Duis blandit, metus vitae sagittis vestibulum, diam",
				"nisl fermentum lorem, non convallis enim urna",
				"rutrum mauris.\n\n",
				"Duis in scelerisque mauris.",
			},
			" ",
		),
	)
	cli.SeeAlso = []string{"flag"}
	cli.TabWidth = 4 // Defaults to 4
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
		cli.Usage(3)
	} else if cli.NArg() > 1 {
		cli.Usage(4)
	} else if len(astring) == 0 {
		cli.Usage(1)
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
