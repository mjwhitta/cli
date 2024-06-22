package main

import (
	"flag"
	"fmt"
	"os"
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
	aBool   bool
	aFloat  float64
	aInt    int
	aString string
	aUint   uint
}

func init() {
	// Parse cli flags
	flag.BoolVar(&flags.aBool, "b", false, "Sample boolean flag.")
	flag.BoolVar(&flags.aBool, "bool", false, "Sample boolean flag.")
	flag.Float64Var(&flags.aFloat, "f", 0.0, "Sample float flag.")
	flag.Float64Var(&flags.aFloat, "float", 0.0, "Sample float flag.")
	flag.IntVar(&flags.aInt, "i", 0, "Sample int flag.")
	flag.StringVar(
		&flags.aString,
		"s",
		"",
		"Mandatory sample string flag.",
	)
	flag.UintVar(&flags.aUint, "u", 0, "Sample uint flag.")
	flag.UintVar(&flags.aUint, "uint", 0, "Sample uint flag.")
	flag.Parse()

	// Validate cli args
	if flag.NArg() == 0 {
		os.Exit(MissingArgument)
	} else if flag.NArg() > 1 {
		os.Exit(ExtraArgument)
	} else if flags.aString == "" {
		os.Exit(MissingOption)
	}
}

func main() {
	fmt.Printf("bool: %t\n", flags.aBool)
	fmt.Printf("float: %f\n", flags.aFloat)
	fmt.Printf("int: %d\n", flags.aInt)
	fmt.Printf("string: %s\n", flags.aString)
	fmt.Printf("uint: %d\n", flags.aUint)
	fmt.Printf("args: %d - %s\n", flag.NArg(), flag.Args())
}
