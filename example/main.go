package main

import (
	"fmt"
	"os"
	"strings"

	"gitlab.com/mjwhitta/cli"
)

// Exit status
const (
	Good = iota
	InvalidOption
	MissingOption
	InvalidArgument
	MissingArgument
	ExtraArgument
	Exception
)

// Flags
var flags struct {
	aBool       bool
	aCounter    cli.Counter
	aFloat      float64
	aInt        int
	anIntList   cli.IntList
	aString     string
	aStringList cli.StringList
	aUint       uint
}

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
			fmt.Sprintf("  %d: Invalid option\n", InvalidOption),
			fmt.Sprintf("  %d: Missing option\n", MissingOption),
			fmt.Sprintf("  %d: Invalid argument\n", InvalidArgument),
			fmt.Sprintf("  %d: Missing argument\n", MissingArgument),
			fmt.Sprintf("  %d: Extra argument\n", ExtraArgument),
			fmt.Sprintf("  %d: Exception", Exception),
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
		"CUSTOM SECTION EXAMPLE",
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
	cli.Flag(&flags.aBool, "b", "bool", false, "Sample boolean flag.")
	cli.Flag(&flags.aCounter, "c", "cnt", "Sample counter flag.")
	cli.Flag(&flags.aFloat, "f", "float", 0.0, "Sample float flag.")
	cli.Flag(&flags.aInt, "i", 0, "Sample", "int flag.")
	cli.Flag(&flags.anIntList, "int", "Sample int list flag.")
	cli.Flag(&flags.aString, "s", "", "Mandatory sample string flag.")
	cli.Flag(&flags.aStringList, "string", "Sample string list flag.")
	cli.Flag(&flags.aUint, "u", "uint", 0, "Sample uint flag.")
	cli.Parse()

	// Validate cli args
	if cli.NArg() == 0 {
		cli.Usage(MissingArgument)
	} else if cli.NArg() > 1 {
		cli.Usage(ExtraArgument)
	} else if flags.aString == "" {
		cli.Usage(MissingOption)
	}
}

func main() {
	fmt.Printf("bool: %t\n", flags.aBool)
	fmt.Printf("Counter: %v\n", flags.aCounter)
	fmt.Printf("float: %f\n", flags.aFloat)
	fmt.Printf("int: %d\n", flags.aInt)
	fmt.Printf(
		"IntList: %d - %v\n",
		len(flags.anIntList),
		flags.anIntList,
	)
	fmt.Printf(
		"IntList: %d - %s\n",
		len(flags.anIntList),
		flags.anIntList.String(),
	)
	fmt.Printf("string: %s\n", flags.aString)
	fmt.Printf(
		"StringList: %d - %v\n",
		len(flags.aStringList),
		flags.aStringList,
	)
	fmt.Printf(
		"StringList: %d - %s\n",
		len(flags.aStringList),
		flags.aStringList.String(),
	)
	fmt.Printf("uint: %d\n", flags.aUint)
	fmt.Printf("args: %d - %s\n", cli.NArg(), cli.Args())
}
