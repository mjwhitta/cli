package cli

import (
    "flag"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

const Version = "1.1.1"

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
}

var Banner = fmt.Sprintf("Usage: %s [OPTIONS]", os.Args[0])
var flags []flagVar
var help bool
var Info = ""
var maxWidth = 0

func init() {
    flags = append(
        flags,
        flagVar{"h", "help", "Display this help message"},
    )
    flag.BoolVar(&help, "h", false, "Display this help message")
    flag.BoolVar(&help, "help", false, "Display this help message")
}

func Arg(i int) string {
    return flag.Arg(i)
}

func Args() []string {
    return flag.Args()
}

func Bool(p *bool, s string, l string, v bool, u string) {
    s, l, u = checkFlags(s, l, u)
    flags = append(flags, flagVar{s, l, u})
    if (len(s) > 0) {flag.BoolVar(p, s, v, u)}
    if (len(l) > 0) {flag.BoolVar(p, l, v, u)}
    updateMaxWidth(s, l)
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

func flagToString(flag flagVar) string {
    var str strings.Builder

    // Leading space
    str.WriteString("    ")

    // Flags
    var f = getFlagColumn(flag.short, flag.long)
    str.WriteString(f)

    // Filler
    for j := 0; j < (maxWidth - len(f)); j++ {
        str.WriteString(" ")
    }

    // Description
    if (maxWidth <= 32) {
        var lines = wrap(flag.desc, 72 - maxWidth)
        for j := range(lines) {
            if (j > 0) {
                // Leading space
                str.WriteString("    ")

                // Filler
                for k := 0; k < maxWidth; k++ {
                    str.WriteString(" ")
                }
            }

            // Filler
            str.WriteString("    ")

            // Actual description
            str.WriteString(lines[j])
            str.WriteString("\n")
        }
    } else {
        // New line b/c maxWidth is too big
        str.WriteString("\n")

        var lines = wrap(flag.desc, 72)
        for j := range(lines) {
            // Leading space
            str.WriteString("        ")

            // Actual description
            str.WriteString(lines[j])
            str.WriteString("\n")
        }
    }

    return str.String()
}

func Float64(p *float64, s string, l string, v float64, u string) {
    s, l, u = checkFlags(s, l, u)
    var long = fmt.Sprintf("%s=FLOAT", l)
    flags = append(flags, flagVar{s, long, u})
    if (len(s) > 0) {flag.Float64Var(p, s, v, u)}
    if (len(l) > 0) {flag.Float64Var(p, l, v, u)}
    updateMaxWidth(s, long)
}

func Float64List(p *Float64ListVar, s string, l string, u string) {
    s, l, u = checkFlags(s, l, u)
    var long = fmt.Sprintf("%s=FLOAT", l)
    flags = append(flags, flagVar{s, long, u})
    if (len(s) > 0) {flag.Var(p, s, u)}
    if (len(l) > 0) {flag.Var(p, l, u)}
    updateMaxWidth(s, long)
}

func getFlagColumn(s string, l string) string {
    var f strings.Builder

    if (len(s) > 0) {
        f.WriteString("-")
        f.WriteString(s)
    }

    if ((len(s) > 0) && (len(l) > 0)) {f.WriteString(", ")}

    if (len(l) > 0) {
        f.WriteString("--")
        f.WriteString(l)
    }

    return f.String()
}

func Int(p *int, s string, l string, v int, u string) {
    s, l, u = checkFlags(s, l, u)
    var long = fmt.Sprintf("%s=INT", l)
    flags = append(flags, flagVar{s, long, u})
    if (len(s) > 0) {flag.IntVar(p, s, v, u)}
    if (len(l) > 0) {flag.IntVar(p, l, v, u)}
    updateMaxWidth(s, long)
}

func Int64(p *int64, s string, l string, v int64, u string) {
    s, l, u = checkFlags(s, l, u)
    var long = fmt.Sprintf("%s=INT", l)
    flags = append(flags, flagVar{s, long, u})
    if (len(s) > 0) {flag.Int64Var(p, s, v, u)}
    if (len(l) > 0) {flag.Int64Var(p, l, v, u)}
    updateMaxWidth(s, long)
}

func Int64List(p *Int64ListVar, s string, l string, u string) {
    s, l, u = checkFlags(s, l, u)
    var long = fmt.Sprintf("%s=INT", l)
    flags = append(flags, flagVar{s, long, u})
    if (len(s) > 0) {flag.Var(p, s, u)}
    if (len(l) > 0) {flag.Var(p, l, u)}
    updateMaxWidth(s, long)
}

func IntList(p *IntListVar, s string, l string, u string) {
    s, l, u = checkFlags(s, l, u)
    var long = fmt.Sprintf("%s=INT", l)
    flags = append(flags, flagVar{s, long, u})
    if (len(s) > 0) {flag.Var(p, s, u)}
    if (len(l) > 0) {flag.Var(p, l, u)}
    updateMaxWidth(s, long)
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
        var left = flags[i].long
        if (len(left) == 0) {left = flags[i].short}

        var right = flags[j].long
        if (len(right) == 0) {right = flags[j].short}

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

    var banner = wrap(Banner, 80)
    for i := range(banner) {
        header.WriteString(banner[i])
        header.WriteString("\n")
    }

    header.WriteString("\nDESCRIPTION\n")

    var info = wrap(Info, 76)
    for i := range(info) {
        header.WriteString("    ")
        header.WriteString(info[i])
        header.WriteString("\n")
    }

    header.WriteString("\nOPTIONS\n")

    fmt.Fprint(os.Stderr, header.String())
}

func String(p *string, s string, l string, v string, u string) {
    s, l, u = checkFlags(s, l, u)
    var long = fmt.Sprintf("%s=STRING", l)
    flags = append(flags, flagVar{s, long, u})
    if (len(s) > 0) {flag.StringVar(p, s, v, u)}
    if (len(l) > 0) {flag.StringVar(p, l, v, u)}
    updateMaxWidth(s, long)
}

func StringList(p *StringListVar, s string, l string, u string) {
    s, l, u = checkFlags(s, l, u)
    var long = fmt.Sprintf("%s=STRING", l)
    flags = append(flags, flagVar{s, long, u})
    if (len(s) > 0) {flag.Var(p, s, u)}
    if (len(l) > 0) {flag.Var(p, l, u)}
    updateMaxWidth(s, long)
}

func Uint(p *uint, s string, l string, v uint, u string) {
    s, l, u = checkFlags(s, l, u)
    var long = fmt.Sprintf("%s=UINT", l)
    flags = append(flags, flagVar{s, long, u})
    if (len(s) > 0) {flag.UintVar(p, s, v, u)}
    if (len(l) > 0) {flag.UintVar(p, l, v, u)}
    updateMaxWidth(s, long)
}

func Uint64(p *uint64, s string, l string, v uint64, u string) {
    s, l, u = checkFlags(s, l, u)
    var long = fmt.Sprintf("%s=UINT", l)
    flags = append(flags, flagVar{s, long, u})
    if (len(s) > 0) {flag.Uint64Var(p, s, v, u)}
    if (len(l) > 0) {flag.Uint64Var(p, l, v, u)}
    updateMaxWidth(s, long)
}

func Uint64List(p *Uint64ListVar, s string, l string, u string) {
    s, l, u = checkFlags(s, l, u)
    var long = fmt.Sprintf("%s=UINT", l)
    flags = append(flags, flagVar{s, long, u})
    if (len(s) > 0) {flag.Var(p, s, u)}
    if (len(l) > 0) {flag.Var(p, l, u)}
    updateMaxWidth(s, long)
}

func UintList(p *UintListVar, s string, l string, u string) {
    s, l, u = checkFlags(s, l, u)
    var long = fmt.Sprintf("%s=UINT", l)
    flags = append(flags, flagVar{s, long, u})
    if (len(s) > 0) {flag.Var(p, s, u)}
    if (len(l) > 0) {flag.Var(p, l, u)}
    updateMaxWidth(s, long)
}

func updateMaxWidth(s string, l string) {
    var width = len(getFlagColumn(s, l))
    if (width > maxWidth) {maxWidth = width}
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
