package cli

import "flag"

// Arg wraps flag.Arg(i int).
func Arg(i int) string {
	return flag.Arg(i)
}

// Args wraps flag.Args().
func Args() []string {
	return flag.Args()
}

// NArg wraps flag.NArg().
func NArg() int {
	return flag.NArg()
}

// NFlag wraps flag.NFlag().
func NFlag() int {
	return flag.NFlag()
}

// Parse will run flag.Parse() and then check for the --help or
// --readme flags.
func Parse() {
	flag.Parse()
	if help {
		Usage(0)
	}
	if readme {
		Readme()
	}
}

// Parsed wraps flag.Parsed().
func Parsed() bool {
	return flag.Parsed()
}
