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

const Version = "1.3.1"

// Float64ListVar allows setting a value multiple times as in:
// --flag=float1 --flag=float2
type Float64ListVar []float64

func (list *Float64ListVar) String() string {
	if (len(*list) == 0) {return ""}
	return fmt.Sprint(*list)
}

func (list *Float64ListVar) Set(val string) error {
    var v float64
    v, _ = strconv.ParseFloat(val, 64)
	(*list) = append(*list, v)
	return nil
}

// IntListVar allows setting a value multiple times as in:
// --flag=int1 --flag=int2
type IntListVar []int64

func (list *IntListVar) String() string {
	if (len(*list) == 0) {return ""}
	return fmt.Sprint(*list)
}

func (list *IntListVar) Set(val string) error {
    var v int64
    v, _ = strconv.ParseInt(val, 0, 0)
	(*list) = append(*list, v)
	return nil
}

// Int64ListVar allows setting a value multiple times as in:
// --flag=int1 --flag=int2
type Int64ListVar []int64

func (list *Int64ListVar) String() string {
	if (len(*list) == 0) {return ""}
	return fmt.Sprint(*list)
}

func (list *Int64ListVar) Set(val string) error {
    var v int64
    v, _ = strconv.ParseInt(val, 0, 64)
	(*list) = append(*list, v)
	return nil
}

// StringListVar allows setting a value multiple times as in:
// --flag=string1 --flag=string2
type StringListVar []string

func (list *StringListVar) String() string {
	if (len(*list) == 0) {return ""}
	return fmt.Sprint(*list)
}

func (list *StringListVar) Set(val string) error {
	(*list) = append(*list, val)
	return nil
}

// UintListVar allows setting a value multiple times as in:
// --flag=uint1 --flag=uint2
type UintListVar []uint64

func (list *UintListVar) String() string {
	if (len(*list) == 0) {return ""}
	return fmt.Sprint(*list)
}

func (list *UintListVar) Set(val string) error {
    var v uint64
    v, _ = strconv.ParseUint(val, 0, 0)
	(*list) = append(*list, v)
	return nil
}

// Uint64ListVar allows setting a value multiple times as in:
// --flag=uint1 --flag=uint2
type Uint64ListVar []uint64

func (list *Uint64ListVar) String() string {
	if (len(*list) == 0) {return ""}
	return fmt.Sprint(*list)
}

func (list *Uint64ListVar) Set(val string) error {
    var v uint64
    v, _ = strconv.ParseUint(val, 0, 64)
	(*list) = append(*list, v)
	return nil
}

type flagVar struct {
    short string
    long string
    desc string
    thetype string
}

type columnWidth struct {
    short int
    long int
    desc int
    left int
}

var Align = false
var Banner = fmt.Sprintf("Usage: %s [OPTIONS]", os.Args[0])
var colWidth = columnWidth{0, 0, math.MaxInt64, 0}
var flags []flagVar
var help bool
var Info = ""
var MaxWidth = 80
var TabWidth = 4

func init() {
    Bool(&help, "h", "help", false, "Display this help message")
}

func Arg(i int) string {
    return flag.Arg(i)
}

func Args() []string {
    return flag.Args()
}

func Bool(p *bool, s string, l string, v bool, u string) {
    s, l, u = checkFlags(s, l, u)

    var f = flagVar{s, l, u, ""}
    flags = append(flags, f)

    if (len(f.short) > 0) {flag.BoolVar(p, f.short, v, f.desc)}
    if (len(f.long) > 0) {flag.BoolVar(p, f.long, v, f.desc)}

    updateMaxWidth(f)
}

func checkFlags(
    s string,
    l string,
    u string,
) (string, string, string) {
    var err = false

    if ((len(s) == 0) && (len(l) == 0)) {
        fmt.Fprint(os.Stderr, "No flag provided\n")
        err = true
    }

    if (len(s) > 1) {
        fmt.Fprint(os.Stderr, "Invalid short flag:\n")
        fmt.Fprintf(os.Stderr, "  %s has length greater than 1\n", s)
        err = true
    }

    if (len(l) == 1) {
        fmt.Fprint(os.Stderr, "Invalid long flag:\n")
        fmt.Fprintf(os.Stderr, "  %s has length equal to 1\n", l)
        err = true
    }

    if (len(u) == 0) {
        fmt.Fprint(os.Stderr, "No description provided\n")
        err = true
    }

    if (err) {os.Exit(128)}

    return strings.TrimSpace(s),
           strings.TrimSpace(l),
           strings.TrimSpace(u)
}

func flagToString(f flagVar) string {
    var enoughRoom = ((TabWidth + colWidth.left) <= (MaxWidth / 2))
    var str strings.Builder

    // Leading space
    for i := 0; i < TabWidth; i++ {
        str.WriteString(" ")
    }

    // Flags
    str.WriteString(getFlagColumn(f, Align && enoughRoom))

    // Description
    if (Align && enoughRoom) {
        // Filler
        var fillto = TabWidth + colWidth.left - str.Len()
        for i := 0; i < fillto; i++ {
            str.WriteString(" ")
        }

        var lines = wrap(f.desc, colWidth.desc)
        for i := range(lines) {
            if (i > 0) {
                // Leading space plus filler
                for j := 0; j < (TabWidth + colWidth.left); j++ {
                    str.WriteString(" ")
                }
            }

            // Alignment
            for j := 0; j < TabWidth; j++ {
                str.WriteString(" ")
            }

            // Actual description
            str.WriteString(lines[i])
            str.WriteString("\n")
        }
    } else {
        // New line b/c colWidth.left is too big
        str.WriteString("\n")

        var lines = wrap(f.desc, MaxWidth - (2 * TabWidth))
        for i := range(lines) {
            // Leading space plus filler
            for j := 0; j < (2 * TabWidth); j++ {
                str.WriteString(" ")
            }

            // Actual description
            str.WriteString(lines[i])
            str.WriteString("\n")
        }
        str.WriteString("\n")
    }

    return str.String()
}

func Float64(p *float64, s string, l string, v float64, u string) {
    s, l, u = checkFlags(s, l, u)

    var f = flagVar{s, l, u, "FLOAT"}
    flags = append(flags, f)

    if (len(f.short) > 0) {flag.Float64Var(p, f.short, v, f.desc)}
    if (len(f.long) > 0) {flag.Float64Var(p, f.long, v, f.desc)}

    updateMaxWidth(f)
}

func Float64List(p *Float64ListVar, s string, l string, u string) {
    s, l, u = checkFlags(s, l, u)

    var f = flagVar{s, l, u, "FLOAT"}
    flags = append(flags, f)

    if (len(f.short) > 0) {flag.Var(p, f.short, f.desc)}
    if (len(f.long) > 0) {flag.Var(p, f.long, f.desc)}

    updateMaxWidth(f)
}

func getFlagColumn(f flagVar, align bool) string {
    var str strings.Builder

    // Short flag
    if (len(f.short) > 0) {
        str.WriteString("-")
        str.WriteString(f.short)
        if (len(f.thetype) > 0) {
            str.WriteString(" ")
            str.WriteString(f.thetype)
        }
    }

    // Separator
    var sep = "  "
    if ((len(f.short) > 0) && (len(f.long) > 0)) {sep = ", "}
    if (len(f.short) > 0) {str.WriteString(sep)}

    // Alignment
    if (align) {
        var fillto = colWidth.short + len(sep) - str.Len()
        for i := 0; i < fillto; i++ {
            str.WriteString(" ")
        }
    }

    // Long flag
    if (len(f.long) > 0) {
        str.WriteString("--")
        str.WriteString(f.long)
        if (len(f.thetype) > 0) {
            str.WriteString("=")
            str.WriteString(f.thetype)
        }
    }

    return str.String()
}

func Int(p *int, s string, l string, v int, u string) {
    s, l, u = checkFlags(s, l, u)

    var f = flagVar{s, l, u, "INT"}
    flags = append(flags, f)

    if (len(f.short) > 0) {flag.IntVar(p, f.short, v, f.desc)}
    if (len(f.long) > 0) {flag.IntVar(p, f.long, v, f.desc)}

    updateMaxWidth(f)
}

func Int64(p *int64, s string, l string, v int64, u string) {
    s, l, u = checkFlags(s, l, u)

    var f = flagVar{s, l, u, "INT"}
    flags = append(flags, f)

    if (len(f.short) > 0) {flag.Int64Var(p, f.short, v, f.desc)}
    if (len(f.long) > 0) {flag.Int64Var(p, f.long, v, f.desc)}

    updateMaxWidth(f)
}

func Int64List(p *Int64ListVar, s string, l string, u string) {
    s, l, u = checkFlags(s, l, u)

    var f = flagVar{s, l, u, "INT"}
    flags = append(flags, f)

    if (len(f.short) > 0) {flag.Var(p, f.short, f.desc)}
    if (len(f.long) > 0) {flag.Var(p, f.long, f.desc)}

    updateMaxWidth(f)
}

func IntList(p *IntListVar, s string, l string, u string) {
    s, l, u = checkFlags(s, l, u)

    var f = flagVar{s, l, u, "INT"}
    flags = append(flags, f)

    if (len(f.short) > 0) {flag.Var(p, f.short, f.desc)}
    if (len(f.long) > 0) {flag.Var(p, f.long, f.desc)}

    updateMaxWidth(f)
}

func NArg() int {
    return flag.NArg()
}

func NFlag() int {
    return flag.NFlag()
}

func Parse() {
    flag.Parse()
    if (help) {Usage(0)}
}

func Parsed() bool {
    return flag.Parsed()
}

func PrintDefaults() {
    var less = func (i, j int) bool {
        var left = flags[i].short
        if (len(left) == 0) {left = flags[i].long}

        var right = flags[j].short
        if (len(right) == 0) {right = flags[j].long}

        return (strings.ToLower(left) < strings.ToLower(right))
    }

    if (!sort.SliceIsSorted(flags, less)) {
        sort.SliceStable(flags, less)
    }

    for i := range(flags) {
        fmt.Fprint(os.Stderr, flagToString(flags[i]))
    }
}

func PrintHeader() {
    var header strings.Builder

    var banner = wrap(Banner, MaxWidth)
    for i := range(banner) {
        header.WriteString(banner[i])
        header.WriteString("\n")
    }

    header.WriteString("\nDESCRIPTION\n")

    var info = wrap(Info, MaxWidth - TabWidth)
    for i := range(info) {
        for j := 0; j < TabWidth; j++ {
            header.WriteString(" ")
        }
        header.WriteString(info[i])
        header.WriteString("\n")
    }

    header.WriteString("\nOPTIONS\n")

    fmt.Fprint(os.Stderr, header.String())
}

func String(p *string, s string, l string, v string, u string) {
    s, l, u = checkFlags(s, l, u)

    var f = flagVar{s, l, u, "STRING"}
    flags = append(flags, f)

    if (len(f.short) > 0) {flag.StringVar(p, f.short, v, f.desc)}
    if (len(f.long) > 0) {flag.StringVar(p, f.long, v, f.desc)}

    updateMaxWidth(f)
}

func StringList(p *StringListVar, s string, l string, u string) {
    s, l, u = checkFlags(s, l, u)

    var f = flagVar{s, l, u, "STRING"}
    flags = append(flags, f)

    if (len(f.short) > 0) {flag.Var(p, f.short, f.desc)}
    if (len(f.long) > 0) {flag.Var(p, f.long, f.desc)}

    updateMaxWidth(f)
}

func Uint(p *uint, s string, l string, v uint, u string) {
    s, l, u = checkFlags(s, l, u)

    var f = flagVar{s, l, u, "UINT"}
    flags = append(flags, f)

    if (len(f.short) > 0) {flag.UintVar(p, f.short, v, f.desc)}
    if (len(f.long) > 0) {flag.UintVar(p, f.long, v, f.desc)}

    updateMaxWidth(f)
}

func Uint64(p *uint64, s string, l string, v uint64, u string) {
    s, l, u = checkFlags(s, l, u)

    var f = flagVar{s, l, u, "UINT"}
    flags = append(flags, f)

    if (len(f.short) > 0) {flag.Uint64Var(p, f.short, v, f.desc)}
    if (len(f.long) > 0) {flag.Uint64Var(p, f.long, v, f.desc)}

    updateMaxWidth(f)
}

func Uint64List(p *Uint64ListVar, s string, l string, u string) {
    s, l, u = checkFlags(s, l, u)

    var f = flagVar{s, l, u, "UINT"}
    flags = append(flags, f)

    if (len(f.short) > 0) {flag.Var(p, f.short, f.desc)}
    if (len(f.long) > 0) {flag.Var(p, f.long, f.desc)}

    updateMaxWidth(f)
}

func UintList(p *UintListVar, s string, l string, u string) {
    s, l, u = checkFlags(s, l, u)

    var f = flagVar{s, l, u, "UINT"}
    flags = append(flags, f)

    if (len(f.short) > 0) {flag.Var(p, f.short, f.desc)}
    if (len(f.long) > 0) {flag.Var(p, f.long, f.desc)}

    updateMaxWidth(f)
}

func updateMaxWidth(f flagVar) {
    var sw = 0
    var lw = 0
    var dw = MaxWidth

    if (len(f.short) > 0) {
        sw = 2 + len(f.thetype)
        if (len(f.thetype) > 0) {sw += 1}
    }

    if (len(f.long) > 0) {
        lw = 2 + len(f.long) + len(f.thetype)
        if (len(f.thetype) > 0) {lw += 1}
    }

    dw = MaxWidth - TabWidth - sw - 2 - lw - TabWidth

    if (sw > colWidth.short) {colWidth.short = sw}
    if (lw > colWidth.long) {colWidth.long = lw}
    if (dw < colWidth.desc) {colWidth.desc = dw}

    colWidth.left = colWidth.short + 2 + colWidth.long
}

func Usage(status int) {
    PrintHeader()
    PrintDefaults()
    os.Exit(status)
}

func wrap(str string, cols int) []string {
    var line strings.Builder
    var lines []string
    var words = strings.Fields(str)

    for i := range(words) {
        if (line.Len() + len(words[i]) > cols) {
            lines = append(lines, line.String())
            line.Reset()
            line.WriteString(words[i])
        } else if (line.Len() == 0) {
            line.Reset()
            line.WriteString(words[i])
        } else {
            line.WriteString(" ")
            line.WriteString(words[i])
        }
    }
    if (line.Len() > 0) {lines = append(lines, line.String())}

    return lines
}
