package cli

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

// ExitStatus sets the description of the program exit status.
func ExitStatus(text ...string) {
	exitStatus = strings.Join(text, " ")
}

// Flag will process the provided values to create a cli flag. Below
// are a few examples:
//
//	var short bool
//	cli.Flag(&short, "s", false, "A short bool flag.")
//
//	var long string
//	cli.Flag(&long, "long", "asdf", "A long string flag.")
//
//	var both int
//	cli.Flag(&both, "b", "both", 0, "A short/long int flag.")
//
//	var secret bool
//	cli.Flag(
//		&secret,
//		"s",
//		"secret",
//		false,
//		"A flag can be hidden by passing true as the last arg.",
//		true,
//	)
func Flag(args ...any) {
	var e error
	var f *cliFlag

	if len(args) == 0 {
		return
	}

	if f, e = newFlag(args...); e != nil {
		fmt.Fprintln(os.Stderr, e.Error())
		os.Exit(127)
	}

	if e = f.validate(); e != nil {
		fmt.Fprintln(os.Stderr, e.Error())
		os.Exit(128)
	}

	flags = append(flags, f)
	f.enable()
	f.updateMaxWidth()
}

func getAuthors(md bool) (ret string) {
	if len(Authors) > 0 {
		if md {
			ret += "\n## Authors\n\n"

			for _, author := range Authors {
				ret += author + "\n"
			}
		} else {
			ret += "AUTHORS\n"

			for _, author := range Authors {
				for range TabWidth {
					ret += " "
				}

				ret += author + "\n"
			}
		}
	}

	return
}

func getBugEmail(md bool) (ret string) {
	var lines []string

	if BugEmail != "" {
		if md {
			lines = wrap(
				"Email bug reports to <"+BugEmail+">.",
				MaxWidth,
			)
			ret += "\n## Reporting bugs\n\n"

			for _, line := range lines {
				ret += line + "\n"
			}
		} else {
			lines = wrap(
				"Email bug reports to <"+BugEmail+">.",
				MaxWidth-TabWidth,
			)
			ret += "\nBUG REPORTS\n"

			for _, line := range lines {
				for range TabWidth {
					ret += " "
				}

				ret += line + "\n"
			}
		}
	}

	return
}

func getCustomSections(md bool) (ret string) {
	for _, s := range sections {
		s.md = md
		ret += s.String()
	}

	return
}

func getExitStatus(md bool) (ret string) {
	if exitStatus != "" {
		if md {
			ret += "\n## Exit status\n\n"

			for _, line := range wrap(exitStatus, MaxWidth) {
				ret += line + "\n"
			}
		} else {
			ret += "\nEXIT STATUS\n"

			for _, line := range wrap(exitStatus, MaxWidth-TabWidth) {
				for range TabWidth {
					ret += " "
				}

				ret += line + "\n"
			}
		}
	}

	return
}

func getSeeAlso(md bool) (ret string) {
	var tmp string

	if len(SeeAlso) == 0 {
		return
	}

	tmp = strings.Join(SeeAlso, ", ")

	if md {
		ret += "\n## See also\n\n"

		for _, ln := range wrap(tmp, MaxWidth) {
			ret += ln + "\n"
		}
	} else {
		ret += "\nSEE ALSO\n"

		for _, ln := range wrap(tmp, MaxWidth) {
			for range TabWidth {
				ret += " "
			}

			ret += ln + "\n"
		}
	}

	return
}

// Info sets the description of how the program works.
func Info(text ...string) {
	info = strings.Join(text, " ")
}

func init() {
	flag.Usage = func() { Usage(127) }
	Flag(&help, "h", "help", false, "Display this help message.")
	Flag(&readme, "readme", false, "Autogenerate README.md.", true)
}

// PrintDefaults will print the configured flags for Usage(). It
// ignores --readme and other hidden flags.
func PrintDefaults() {
	if !sort.SliceIsSorted(flags, less) {
		sort.SliceStable(flags, less)
	}

	for _, f := range flags {
		if !f.hidden {
			fmt.Fprint(os.Stderr, f.String())
		}
	}
	if Align {
		fmt.Fprintln(os.Stderr)
	}
}

// PrintExtra will print the Usage() extra details.
func PrintExtra() {
	var extra string

	extra += getCustomSections(false)
	extra += getAuthors(false)
	extra += getBugEmail(false)
	extra += getExitStatus(false)
	extra += getSeeAlso(false)

	fmt.Fprint(os.Stderr, extra)
}

// PrintHeader will print the Usage() header.
func PrintHeader() {
	var header string
	var lines []string

	lines = wrap("Usage: "+Banner, MaxWidth)
	for _, line := range lines {
		header += line + "\n"
	}

	header += "\nDESCRIPTION\n"

	lines = wrap(info, MaxWidth-TabWidth)
	for _, line := range lines {
		for range TabWidth {
			header += " "
		}
		header += line + "\n"
	}

	header += "\nOPTIONS\n"

	fmt.Fprint(os.Stderr, header)
}

// Readme will attempt to print out a basic README.md file based on
// the provided details.
func Readme() {
	var lines []string
	var readme string

	// Title
	readme += "# " + Title + "\n"

	// Synopsis
	readme += "\n## Synopsis\n\n"
	lines = wrap(Banner, MaxWidth)
	for _, line := range lines {
		readme += "`" + line + "`\n"
	}

	// Description
	readme += "\n## Description\n\n"
	lines = wrap(info, MaxWidth)
	for _, line := range lines {
		readme += line + "\n"
	}

	// Options and descriptions
	readme += "\n## Options\n\n"
	if !sort.SliceIsSorted(flags, less) {
		sort.SliceStable(flags, less)
	}
	readme += "Option | Args | Description\n"
	readme += "------ | ---- | -----------\n"
	for _, f := range flags {
		if !f.hidden {
			readme += f.table()
		}
	}

	readme += getCustomSections(true)
	readme += getAuthors(true)
	readme += getBugEmail(true)
	readme += getExitStatus(true)
	readme += getSeeAlso(true)

	fmt.Print(readme)
	os.Exit(0)
}

// Section will add a new custom section with the specified title and
// text.
func Section(title string, text ...string) {
	sections = append(
		sections,
		section{
			text:  strings.Join(text, " "),
			title: title,
		},
	)
}

// SectionAligned will add a new custom section with the specified
// title and text.
func SectionAligned(title string, align string, text ...string) {
	sections = append(
		sections,
		section{
			alignOn: align,
			text:    strings.Join(text, " "),
			title:   title,
		},
	)
}

// Usage will essentially print a manpage.
func Usage(status int) {
	PrintHeader()
	PrintDefaults()
	PrintExtra()
	os.Exit(status)
}

func wrap(input string, width int) []string {
	var line string
	var lines []string
	var strs []string = strings.Split(input, "\n")
	var words []string

	for _, str := range strs {
		line = ""
		words = strings.Fields(str)

		for _, word := range words {
			if len(line)+len(word)+1 > width {
				lines = append(lines, line)
				line = word
			} else if line == "" {
				line = word
			} else {
				line += " " + word
			}
		}

		if line != "" {
			lines = append(lines, line)
		} else {
			lines = append(lines, "")
		}
	}

	if strings.HasSuffix(input, "\n") {
		lines = append(lines, "")
	}

	return lines
}
