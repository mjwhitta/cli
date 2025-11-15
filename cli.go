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
	var exit int = 128
	var f *cliFlag

	if len(args) == 0 {
		return
	}

	if f, e = newFlag(args...); e != nil {
		fmt.Fprintln(os.Stderr, e.Error())
		os.Exit(exit)
	}

	if e = f.validate(); e != nil {
		fmt.Fprintln(os.Stderr, e.Error())
		os.Exit(exit)
	}

	flags = append(flags, f)

	if e = f.enable(f.short); e != nil {
		fmt.Fprintln(os.Stderr, e.Error())
		os.Exit(exit)
	}

	if e = f.enable(f.long); e != nil {
		fmt.Fprintln(os.Stderr, e.Error())
		os.Exit(exit)
	}

	f.updateMaxWidth()
}

func getAuthors(md bool) string {
	var sb strings.Builder

	if len(Authors) > 0 {
		if md {
			sb.WriteString("\n## Authors\n\n")

			for _, author := range Authors {
				sb.WriteString(author + "\n")
			}
		} else {
			sb.WriteString("AUTHORS\n")

			for _, author := range Authors {
				for range TabWidth {
					sb.WriteString(" ")
				}

				sb.WriteString(author + "\n")
			}
		}
	}

	return sb.String()
}

func getBugEmail(md bool) string {
	var lines []string
	var sb strings.Builder

	if BugEmail != "" {
		if md {
			lines = wrap(
				"Email bug reports to <"+BugEmail+">.",
				MaxWidth,
			)

			sb.WriteString("\n## Reporting bugs\n\n")

			for _, line := range lines {
				sb.WriteString(line + "\n")
			}
		} else {
			lines = wrap(
				"Email bug reports to <"+BugEmail+">.",
				MaxWidth-TabWidth,
			)

			sb.WriteString("\nBUG REPORTS\n")

			for _, line := range lines {
				for range TabWidth {
					sb.WriteString(" ")
				}

				sb.WriteString(line + "\n")
			}
		}
	}

	return sb.String()
}

func getCustomSections(md bool) string {
	var sb strings.Builder

	for _, s := range sections {
		s.md = md
		sb.WriteString(s.String())
	}

	return sb.String()
}

func getExitStatus(md bool) string {
	var sb strings.Builder

	if exitStatus != "" {
		if md {
			sb.WriteString("\n## Exit status\n\n")

			for _, line := range wrap(exitStatus, MaxWidth) {
				sb.WriteString(line + "\n")
			}
		} else {
			sb.WriteString("\nEXIT STATUS\n")

			for _, line := range wrap(exitStatus, MaxWidth-TabWidth) {
				for range TabWidth {
					sb.WriteString(" ")
				}

				sb.WriteString(line + "\n")
			}
		}
	}

	return sb.String()
}

func getSeeAlso(md bool) string {
	var sb strings.Builder
	var tmp string

	if len(SeeAlso) == 0 {
		return ""
	}

	tmp = strings.Join(SeeAlso, ", ")

	if md {
		sb.WriteString("\n## See also\n\n")

		for _, ln := range wrap(tmp, MaxWidth) {
			sb.WriteString(ln + "\n")
		}
	} else {
		sb.WriteString("\nSEE ALSO\n")

		for _, ln := range wrap(tmp, MaxWidth) {
			for range TabWidth {
				sb.WriteString(" ")
			}

			sb.WriteString(ln + "\n")
		}
	}

	return sb.String()
}

// Info sets the description of how the program works.
func Info(text ...string) {
	info = strings.Join(text, " ")
}

func init() {
	var exit int = 127

	flag.Usage = func() { Usage(exit) }

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
	var sb strings.Builder

	for _, line := range wrap("Usage: "+Banner, MaxWidth) {
		sb.WriteString(line + "\n")
	}

	sb.WriteString("\nDESCRIPTION\n")

	for _, line := range wrap(info, MaxWidth-TabWidth) {
		for range TabWidth {
			sb.WriteString(" ")
		}

		sb.WriteString(line + "\n")
	}

	sb.WriteString("\nOPTIONS\n")

	fmt.Fprint(os.Stderr, sb.String())
}

// Readme will attempt to print out a basic README.md file based on
// the provided details.
func Readme() {
	var sb strings.Builder

	// Title
	sb.WriteString("# " + Title + "\n")

	// Synopsis
	sb.WriteString("\n## Synopsis\n\n")

	for _, line := range wrap(Banner, MaxWidth) {
		sb.WriteString("`" + line + "`\n")
	}

	// Description
	sb.WriteString("\n## Description\n\n")

	for _, line := range wrap(info, MaxWidth) {
		sb.WriteString(line + "\n")
	}

	// Options and descriptions
	sb.WriteString("\n## Options\n\n")

	if !sort.SliceIsSorted(flags, less) {
		sort.SliceStable(flags, less)
	}

	sb.WriteString("Option | Args | Description\n")
	sb.WriteString("------ | ---- | -----------\n")

	for _, f := range flags {
		if !f.hidden {
			sb.WriteString(f.table())
		}
	}

	sb.WriteString(getCustomSections(true))
	sb.WriteString(getAuthors(true))
	sb.WriteString(getBugEmail(true))
	sb.WriteString(getExitStatus(true))
	sb.WriteString(getSeeAlso(true))

	fmt.Print(sb.String())
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
			switch {
			case len(line)+len(word)+1 > width:
				lines = append(lines, line)
				line = word
			case line == "":
				line = word
			default:
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
