package cli

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

type columnWidth struct {
	desc  int
	left  int
	long  int
	short int
}

type section struct {
	text  string
	title string
}

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
	updateMaxWidth(f)
}

func flagToString(f *cliFlag) string {
	var enoughRoom = ((TabWidth + colWidth.left) <= (MaxWidth / 2))
	var fillto int
	var lines []string
	var str string

	// Leading space
	for i := 0; i < TabWidth; i++ {
		str += " "
	}

	// Flags
	str += getFlagColumn(f, Align && enoughRoom)

	// Description
	if Align && enoughRoom {
		// Filler
		fillto = TabWidth + colWidth.left - len(str)
		for i := 0; i < fillto; i++ {
			str += " "
		}

		lines = wrap(f.desc, colWidth.desc)
		for i, line := range lines {
			if i > 0 {
				// Leading space plus filler
				for j := 0; j < (TabWidth + colWidth.left); j++ {
					str += " "
				}
			}

			// Alignment
			for j := 0; j < TabWidth; j++ {
				str += " "
			}

			// Actual description
			str += line + "\n"
		}
	} else {
		// New line b/c colWidth.left is too big
		str += "\n"

		lines = wrap(f.desc, MaxWidth-(2*TabWidth))
		for _, line := range lines {
			// Leading space plus filler
			for i := 0; i < (2 * TabWidth); i++ {
				str += " "
			}

			// Actual description
			str += line + "\n"
		}
		str += "\n"
	}

	return str
}

func flagToTable(f *cliFlag) string {
	var str string

	// Option
	if f.short != "" {
		str += "`-" + f.short + "`"
	}
	if (f.short != "") && (f.long != "") {
		str += ", "
	}
	if f.long != "" {
		str += "`--" + f.long + "`"
	}

	// Separator
	str += " | "

	// Args
	if f.thetype != "" {
		str += "`" + f.thetype + "`"
	}

	// Separator
	str += " | "

	// Description
	str += f.desc + "\n"

	return str
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
				for i := 0; i < TabWidth; i++ {
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
				for i := 0; i < TabWidth; i++ {
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
		if md {
			ret += "\n## " + s.title + "\n\n"

			for _, line := range wrap(s.text, MaxWidth) {
				ret += line + "\n"
			}
		} else {
			ret += s.title + "\n"

			for _, line := range wrap(s.text, MaxWidth-TabWidth) {
				for i := 0; i < TabWidth; i++ {
					ret += " "
				}

				ret += line + "\n"
			}

			ret += "\n"
		}
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
				for i := 0; i < TabWidth; i++ {
					ret += " "
				}

				ret += line + "\n"
			}
		}
	}

	return
}

func getFlagColumn(f *cliFlag, align bool) string {
	var fillto int
	var sep string
	var str string

	// Short flag
	if f.short != "" {
		str += "-" + f.short
		if (f.thetype != "") && (f.long == "") {
			str += " " + f.thetype
		}
	}

	// Separator
	sep = "  "
	if (f.short != "") && (f.long != "") {
		sep = ", "
	}
	if f.short != "" {
		str += sep
	}

	// Alignment
	if align {
		fillto = colWidth.short + len(sep) - len(str)
		for i := 0; i < fillto; i++ {
			str += " "
		}
	}

	// Long flag
	if f.long != "" {
		str += "--" + f.long
		if f.thetype != "" {
			str += "=" + f.thetype
		}
	}

	return str
}

func getSeeAlso(md bool) (ret string) {
	if len(SeeAlso) > 0 {
		if md {
			ret += "\n## See also\n\n"

			for _, see := range SeeAlso {
				ret += see + "\n"
			}
		} else {
			ret += "\nSEE ALSO\n"

			for _, see := range SeeAlso {
				for i := 0; i < TabWidth; i++ {
					ret += " "
				}
				ret += see + "\n"
			}
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

// PrintDefaults will print the configured flags for Usage(). It
// ignores --readme and other hidden flags.
func PrintDefaults() {
	if !sort.SliceIsSorted(flags, less) {
		sort.SliceStable(flags, less)
	}

	for _, f := range flags {
		if !f.hidden {
			fmt.Fprint(os.Stderr, flagToString(f))
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
		for j := 0; j < TabWidth; j++ {
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
			readme += flagToTable(f)
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

func updateMaxWidth(f *cliFlag) {
	var sw = 2
	var lw = 0
	var dw int

	if (f.short != "") && (f.long == "") {
		sw += len(f.thetype)
		if f.thetype != "" {
			sw++
		}
	}

	if f.long != "" {
		lw = 2 + len(f.long) + len(f.thetype)
		if f.thetype != "" {
			lw++
		}
	}

	dw = MaxWidth - TabWidth - sw - 2 - lw - TabWidth

	if sw > colWidth.short {
		colWidth.short = sw
	}
	if lw > colWidth.long {
		colWidth.long = lw
	}
	if dw < colWidth.desc {
		colWidth.desc = dw
	}

	colWidth.left = colWidth.short + 2 + colWidth.long
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
	var strs = strings.Split(input, "\n")
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
