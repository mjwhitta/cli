package cli

import (
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const Version = "1.5.1"

// Float64List allows setting a value multiple times as in:
// --flag=float1 --flag=float2
type Float64List []float64

func (list *Float64List) String() string {
	if len(*list) == 0 {
		return ""
	}
	return fmt.Sprint(*list)
}

func (list *Float64List) Set(val string) error {
	var v float64
	v, _ = strconv.ParseFloat(val, 64)
	(*list) = append(*list, v)
	return nil
}

// IntList allows setting a value multiple times as in:
// --flag=int1 --flag=int2
type IntList []int64

func (list *IntList) String() string {
	if len(*list) == 0 {
		return ""
	}
	return fmt.Sprint(*list)
}

func (list *IntList) Set(val string) error {
	var v int64
	v, _ = strconv.ParseInt(val, 0, 0)
	(*list) = append(*list, v)
	return nil
}

// Int64List allows setting a value multiple times as in:
// --flag=int1 --flag=int2
type Int64List []int64

func (list *Int64List) String() string {
	if len(*list) == 0 {
		return ""
	}
	return fmt.Sprint(*list)
}

func (list *Int64List) Set(val string) error {
	var v int64
	v, _ = strconv.ParseInt(val, 0, 64)
	(*list) = append(*list, v)
	return nil
}

// StringList allows setting a value multiple times as in:
// --flag=string1 --flag=string2
type StringList []string

func (list *StringList) String() string {
	if len(*list) == 0 {
		return ""
	}
	return fmt.Sprint(*list)
}

func (list *StringList) Set(val string) error {
	(*list) = append(*list, val)
	return nil
}

// UintList allows setting a value multiple times as in:
// --flag=uint1 --flag=uint2
type UintList []uint64

func (list *UintList) String() string {
	if len(*list) == 0 {
		return ""
	}
	return fmt.Sprint(*list)
}

func (list *UintList) Set(val string) error {
	var v uint64
	v, _ = strconv.ParseUint(val, 0, 0)
	(*list) = append(*list, v)
	return nil
}

// Uint64List allows setting a value multiple times as in:
// --flag=uint1 --flag=uint2
type Uint64List []uint64

func (list *Uint64List) String() string {
	if len(*list) == 0 {
		return ""
	}
	return fmt.Sprint(*list)
}

func (list *Uint64List) Set(val string) error {
	var v uint64
	v, _ = strconv.ParseUint(val, 0, 64)
	(*list) = append(*list, v)
	return nil
}

type columnWidth struct {
	short int
	long  int
	desc  int
	left  int
}

type flagVar struct {
	short   string
	long    string
	desc    string
	thetype string
}

var Align = false
var Authors []string
var Banner = fmt.Sprintf("%s [OPTIONS]", os.Args[0])
var BugEmail string
var colWidth = columnWidth{0, 0, math.MaxInt64, 0}
var ExitStatus string
var flags []flagVar
var help bool
var Info = ""
var less = func(i, j int) bool {
	// Sort by short flag
	var left = flags[i].short
	if len(left) == 0 {
		left = flags[i].long
	}

	var right = flags[j].short
	if len(right) == 0 {
		right = flags[j].long
	}

	// Fallback to long flag if comparing same short flag
	if strings.ToLower(left) == strings.ToLower(right) {
		if len(flags[i].long) != 0 {
			left = flags[i].long
		}
		if len(flags[j].long) != 0 {
			right = flags[j].long
		}
	}

	return (strings.ToLower(left) < strings.ToLower(right))
}
var MaxWidth = 80
var readme bool
var SeeAlso []string
var TabWidth = 4
var Title string

func init() {
	flag.Usage = func() { Usage(127) }
	Flag(&help, "h", "help", false, "Display this help message.")
	Flag(&readme, "readme", false, "Autogenerate a README.md file.")
}

func Arg(i int) string {
	return flag.Arg(i)
}

func Args() []string {
	return flag.Args()
}

func chkFlags(s string, l string, u string) (string, string, string) {
	var err = false

	if (len(s) == 0) && (len(l) == 0) {
		fmt.Fprint(os.Stderr, "No flag provided\n")
		err = true
	}

	if len(s) > 1 {
		fmt.Fprint(os.Stderr, "Invalid short flag:\n")
		fmt.Fprintf(os.Stderr, "  %s has length greater than 1\n", s)
		err = true
	}

	if len(l) == 1 {
		fmt.Fprint(os.Stderr, "Invalid long flag:\n")
		fmt.Fprintf(os.Stderr, "  %s has length equal to 1\n", l)
		err = true
	}

	if len(u) == 0 {
		fmt.Fprint(os.Stderr, "No description provided\n")
		err = true
	}

	if err {
		os.Exit(128)
	}

	return strings.TrimSpace(s),
		strings.TrimSpace(l),
		strings.TrimSpace(u)
}

func Flag(args ...interface{}) {
	var containsSpace, isString bool
	var long string
	var ptr interface{}
	var short, thetype, tmp, usage string
	var value interface{}

	// Process args
	for i := range args {
		switch args[i].(type) {
		case *bool, *float64, *int, *int64, *string, *uint, *uint64,
			*Float64List, *IntList, *Int64List, *StringList,
			*UintList, *Uint64List:
			ptr = args[i]
			switch args[i].(type) {
			case *float64, *Float64List:
				thetype = "FLOAT"
			case *int, *int64, *IntList, *Int64List:
				thetype = "INT"
			case *string:
				isString = true
				thetype = "STRING"
			case *StringList:
				thetype = "STRING"
			case *uint, *uint64, *UintList, *Uint64List:
				thetype = "UINT"
			}
		case bool, float64, int, int64, uint, uint64:
			value = args[i]
		case string:
			tmp = args[i].(string)
			containsSpace = strings.Contains(tmp, " ")
			if (short == "") && (len(tmp) == 1) {
				short = tmp
			} else if (long == "") && !containsSpace {
				long = tmp
			} else if isString && (value == nil) {
				value = args[i]
			} else {
				usage = tmp
			}
		default:
			fmt.Print("Error!\n")
			os.Exit(127)
		}
	}

	// Validate flags
	short, long, usage = chkFlags(short, long, usage)

	var f = flagVar{short, long, usage, thetype}
	flags = append(flags, f)

	switch ptr.(type) {
	case *bool:
		if len(f.short) > 0 {
			flag.BoolVar(ptr.(*bool), f.short, value.(bool), f.desc)
		}
		if len(f.long) > 0 {
			flag.BoolVar(ptr.(*bool), f.long, value.(bool), f.desc)
		}
	case *float64:
		if len(f.short) > 0 {
			flag.Float64Var(
				ptr.(*float64),
				f.short,
				value.(float64),
				f.desc,
			)
		}
		if len(f.long) > 0 {
			flag.Float64Var(
				ptr.(*float64),
				f.long,
				value.(float64),
				f.desc,
			)
		}
	case *Float64List:
		if len(f.short) > 0 {
			flag.Var(ptr.(*Float64List), f.short, f.desc)
		}
		if len(f.long) > 0 {
			flag.Var(ptr.(*Float64List), f.long, f.desc)
		}
	case *int:
		if len(f.short) > 0 {
			flag.IntVar(ptr.(*int), f.short, value.(int), f.desc)
		}
		if len(f.long) > 0 {
			flag.IntVar(ptr.(*int), f.long, value.(int), f.desc)
		}
	case *int64:
		if len(f.short) > 0 {
			flag.Int64Var(
				ptr.(*int64),
				f.short,
				int64(value.(int)),
				f.desc,
			)
		}
		if len(f.long) > 0 {
			flag.Int64Var(
				ptr.(*int64),
				f.long,
				int64(value.(int)),
				f.desc,
			)
		}
	case *Int64List:
		if len(f.short) > 0 {
			flag.Var(ptr.(*Int64List), f.short, f.desc)
		}
		if len(f.long) > 0 {
			flag.Var(ptr.(*Int64List), f.long, f.desc)
		}
	case *IntList:
		if len(f.short) > 0 {
			flag.Var(ptr.(*IntList), f.short, f.desc)
		}
		if len(f.long) > 0 {
			flag.Var(ptr.(*IntList), f.long, f.desc)
		}
	case *string:
		if len(f.short) > 0 {
			flag.StringVar(
				ptr.(*string),
				f.short,
				value.(string),
				f.desc,
			)
		}
		if len(f.long) > 0 {
			flag.StringVar(
				ptr.(*string),
				f.long,
				value.(string),
				f.desc,
			)
		}
	case *StringList:
		if len(f.short) > 0 {
			flag.Var(ptr.(*StringList), f.short, f.desc)
		}
		if len(f.long) > 0 {
			flag.Var(ptr.(*StringList), f.long, f.desc)
		}
	case *uint:
		if len(f.short) > 0 {
			flag.UintVar(
				ptr.(*uint),
				f.short,
				uint(value.(int)),
				f.desc,
			)
		}
		if len(f.long) > 0 {
			flag.UintVar(
				ptr.(*uint),
				f.long,
				uint(value.(int)),
				f.desc,
			)
		}
	case *uint64:
		if len(f.short) > 0 {
			flag.Uint64Var(
				ptr.(*uint64),
				f.short,
				uint64(value.(int)),
				f.desc,
			)
		}
		if len(f.long) > 0 {
			flag.Uint64Var(
				ptr.(*uint64),
				f.long,
				uint64(value.(int)),
				f.desc,
			)
		}
	case *Uint64List:
		if len(f.short) > 0 {
			flag.Var(ptr.(*Uint64List), f.short, f.desc)
		}
		if len(f.long) > 0 {
			flag.Var(ptr.(*Uint64List), f.long, f.desc)
		}
	case *UintList:
		if len(f.short) > 0 {
			flag.Var(ptr.(*UintList), f.short, f.desc)
		}
		if len(f.long) > 0 {
			flag.Var(ptr.(*UintList), f.long, f.desc)
		}
	}

	updateMaxWidth(f)
}

func flagToString(f flagVar) string {
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
		for i := range lines {
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
			str += lines[i] + "\n"
		}
	} else {
		// New line b/c colWidth.left is too big
		str += "\n"

		lines = wrap(f.desc, MaxWidth-(2*TabWidth))
		for i := range lines {
			// Leading space plus filler
			for j := 0; j < (2 * TabWidth); j++ {
				str += " "
			}

			// Actual description
			str += lines[i] + "\n"
		}
		str += "\n"
	}

	return str
}

func flagToTable(f flagVar) string {
	var str string

	// Option
	if len(f.short) > 0 {
		str += "`-" + f.short + "`"
	}
	if (len(f.short) > 0) && (len(f.long) > 0) {
		str += ", "
	}
	if len(f.long) > 0 {
		str += "`--" + f.long + "`"
	}

	// Separator
	str += " | "

	// Args
	if len(f.thetype) > 0 {
		str += "`" + f.thetype + "`"
	}

	// Separator
	str += " | "

	// Description
	str += f.desc + "\n"

	return str
}

func getFlagColumn(f flagVar, align bool) string {
	var fillto int
	var sep string
	var str string

	// Short flag
	if len(f.short) > 0 {
		str += "-" + f.short
		if len(f.thetype) > 0 {
			str += " " + f.thetype
		}
	}

	// Separator
	sep = "  "
	if (len(f.short) > 0) && (len(f.long) > 0) {
		sep = ", "
	}
	if len(f.short) > 0 {
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
	if len(f.long) > 0 {
		str += "--" + f.long
		if len(f.thetype) > 0 {
			str += "=" + f.thetype
		}
	}

	return str
}

func NArg() int {
	return flag.NArg()
}

func NFlag() int {
	return flag.NFlag()
}

func Parse() {
	flag.Parse()
	if help {
		Usage(0)
	}
	if readme {
		Readme()
	}
}

func Parsed() bool {
	return flag.Parsed()
}

func PrintDefaults() {
	if !sort.SliceIsSorted(flags, less) {
		sort.SliceStable(flags, less)
	}

	for i := range flags {
		fmt.Fprint(os.Stderr, flagToString(flags[i]))
	}
}

func PrintHeader() {
	var header string

	var banner = wrap("Usage: "+Banner, MaxWidth)
	for i := range banner {
		header += banner[i] + "\n"
	}

	header += "\nDESCRIPTION\n"

	var info = wrap(Info, MaxWidth-TabWidth)
	for i := range info {
		for j := 0; j < TabWidth; j++ {
			header += " "
		}
		header += info[i] + "\n"
	}

	header += "\nOPTIONS\n"

	fmt.Fprint(os.Stderr, header)
}

func Readme() {
	var readme string

	// Title
	readme += "# " + Title + "\n"

	// Synopsis
	readme += "\n## Synopsis\n\n"
	var banner = wrap(Banner, MaxWidth)
	for i := range banner {
		readme += "`" + banner[i] + "`\n"
	}

	// Description
	readme += "\n## Description\n\n"
	var info = wrap(Info, MaxWidth)
	for i := range info {
		readme += info[i] + "\n"
	}

	// Options and descriptions
	readme += "\n## Options\n\n"
	if !sort.SliceIsSorted(flags, less) {
		sort.SliceStable(flags, less)
	}
	readme += "Option | Args | Description\n"
	readme += "------ | ---- | -----------\n"
	for i := range flags {
		readme += flagToTable(flags[i])
	}

	// Author info
	if len(Authors) > 0 {
		readme += "\n## Authors\n\n"
		for i := range Authors {
			readme += Authors[i] + "\n"
		}
	}

	// Info for reporting bugs
	if len(BugEmail) > 0 {
		var bugs = wrap(
			strings.Join(
				[]string{
					"Email bug reports to the bug-reporting address ",
					"(",
					BugEmail,
					").",
				},
				"",
			),
			MaxWidth,
		)

		readme += "\n## Reporting bugs\n\n"
		for i := range bugs {
			readme += bugs[i] + "\n"
		}
	}

	// Exit status info
	if len(ExitStatus) > 0 {
		readme += "\n## Exit status\n\n"
		var exitStatus = wrap(ExitStatus, MaxWidth)
		for i := range exitStatus {
			readme += exitStatus[i] + "\n"
		}
	}

	// See also for more info
	if len(SeeAlso) > 0 {
		readme += "\n## See also\n\n"
		for i := range SeeAlso {
			readme += SeeAlso[i] + "\n"
		}
	}

	fmt.Print(readme)
	os.Exit(0)
}

func updateMaxWidth(f flagVar) {
	var sw = 0
	var lw = 0
	var dw = MaxWidth

	if len(f.short) > 0 {
		sw = 2 + len(f.thetype)
		if len(f.thetype) > 0 {
			sw += 1
		}
	}

	if len(f.long) > 0 {
		lw = 2 + len(f.long) + len(f.thetype)
		if len(f.thetype) > 0 {
			lw += 1
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

func Usage(status int) {
	PrintHeader()
	PrintDefaults()
	os.Exit(status)
}

func wrap(str string, cols int) []string {
	var line string
	var lines []string
	var words = strings.Fields(str)

	for i := range words {
		if len(line)+len(words[i]) > cols {
			lines = append(lines, line)
			line = words[i]
		} else if len(line) == 0 {
			line = words[i]
		} else {
			line += " " + words[i]
		}
	}
	if len(line) > 0 {
		lines = append(lines, line)
	}

	return lines
}
