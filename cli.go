package cli

import (
    "flag"
    "fmt"
    "os"
    "strconv"
    "strings"
)

const Version = "1.0.0"

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
    flags = append(flags, flagVar{s, l, u})
    if (len(s) > 0) {flag.BoolVar(p, s, v, u)}
    if (len(l) > 0) {flag.BoolVar(p, l, v, u)}
}

func Float64(p *float64, s string, l string, v float64, u string) {
    flags = append(flags, flagVar{s, fmt.Sprintf("%s=FLOAT", l), u})
    if (len(s) > 0) {flag.Float64Var(p, s, v, u)}
    if (len(l) > 0) {flag.Float64Var(p, l, v, u)}
}

func Float64List(p *Float64ListVar, s string, l string, u string) {
    flags = append(flags, flagVar{s, fmt.Sprintf("%s=FLOAT", l), u})
    if (len(s) > 0) {flag.Var(p, s, u)}
    if (len(l) > 0) {flag.Var(p, l, u)}
}

func Int(p *int, s string, l string, v int, u string) {
    flags = append(flags, flagVar{s, fmt.Sprintf("%s=INT", l), u})
    if (len(s) > 0) {flag.IntVar(p, s, v, u)}
    if (len(l) > 0) {flag.IntVar(p, l, v, u)}
}

func Int64(p *int64, s string, l string, v int64, u string) {
    flags = append(flags, flagVar{s, fmt.Sprintf("%s=INT", l), u})
    if (len(s) > 0) {flag.Int64Var(p, s, v, u)}
    if (len(l) > 0) {flag.Int64Var(p, l, v, u)}
}

func Int64List(p *Int64ListVar, s string, l string, u string) {
    flags = append(flags, flagVar{s, fmt.Sprintf("%s=INT", l), u})
    if (len(s) > 0) {flag.Var(p, s, u)}
    if (len(l) > 0) {flag.Var(p, l, u)}
}

func IntList(p *IntListVar, s string, l string, u string) {
    flags = append(flags, flagVar{s, fmt.Sprintf("%s=INT", l), u})
    if (len(s) > 0) {flag.Var(p, s, u)}
    if (len(l) > 0) {flag.Var(p, l, u)}
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
    for i := range(flags) {
        fmt.Fprint(os.Stderr, "    ")
        if (len(flags[i].short) > 0) {
            if (len(flags[i].short) > 1) {fmt.Fprint(os.Stderr, "-")}
            fmt.Fprintf(os.Stderr, "-%s, ", flags[i].short)
        }
        fmt.Fprintf(os.Stderr, "--%s\n", flags[i].long)

        var lines = wrap(flags[i].desc, 72)
        for i := range(lines) {
            fmt.Fprintf(os.Stderr, "        %s\n", lines[i])
        }
    }
}

func String(p *string, s string, l string, v string, u string) {
    flags = append(flags, flagVar{s, fmt.Sprintf("%s=STRING", l), u})
    if (len(s) > 0) {flag.StringVar(p, s, v, u)}
    if (len(l) > 0) {flag.StringVar(p, l, v, u)}
}

func StringList(p *StringListVar, s string, l string, u string) {
    flags = append(flags, flagVar{s, fmt.Sprintf("%s=STRING", l), u})
    if (len(s) > 0) {flag.Var(p, s, u)}
    if (len(l) > 0) {flag.Var(p, l, u)}
}

func Uint(p *uint, s string, l string, v uint, u string) {
    flags = append(flags, flagVar{s, fmt.Sprintf("%s=UINT", l), u})
    if (len(s) > 0) {flag.UintVar(p, s, v, u)}
    if (len(l) > 0) {flag.UintVar(p, l, v, u)}
}

func Uint64(p *uint64, s string, l string, v uint64, u string) {
    flags = append(flags, flagVar{s, fmt.Sprintf("%s=UINT", l), u})
    if (len(s) > 0) {flag.Uint64Var(p, s, v, u)}
    if (len(l) > 0) {flag.Uint64Var(p, l, v, u)}
}

func Uint64List(p *Uint64ListVar, s string, l string, u string) {
    flags = append(flags, flagVar{s, fmt.Sprintf("%s=UINT", l), u})
    if (len(s) > 0) {flag.Var(p, s, u)}
    if (len(l) > 0) {flag.Var(p, l, u)}
}

func UintList(p *UintListVar, s string, l string, u string) {
    flags = append(flags, flagVar{s, fmt.Sprintf("%s=UINT", l), u})
    if (len(s) > 0) {flag.Var(p, s, u)}
    if (len(l) > 0) {flag.Var(p, l, u)}
}

func Usage(status int) {
    var banner = wrap(Banner, 80)
    for i := range(banner) {
        fmt.Fprintf(os.Stderr, "%s\n", banner[i])
    }
    fmt.Fprint(os.Stderr, "\n")

    fmt.Fprint(os.Stderr, "DESCRIPTION\n")
    var info = wrap(Info, 76)
    for i := range(info) {
        fmt.Fprintf(os.Stderr, "    %s\n", info[i])
    }
    fmt.Fprint(os.Stderr, "\n")

    fmt.Fprint(os.Stderr, "OPTIONS\n")
    PrintDefaults()

    os.Exit(status)
}

func wrap(str string, cols int) []string {
    var line  = ""
    var lines []string
    var words = strings.Fields(str)

    for i := range(words) {
        if (len(line) + len(words[i]) > cols) {
            lines = append(lines, line)
            line = words[i]
        } else if (len(line) == 0) {
            line = words[i]
        } else {
            line = fmt.Sprintf("%s %s", line, words[i])
        }
    }
    if (len(line) > 0) {lines = append(lines, line)}

    return lines
}
