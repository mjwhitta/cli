package cli

import (
	"flag"
	"strings"

	"github.com/mjwhitta/errors"
)

type cliFlag struct {
	desc    string
	gotVal  bool
	hidden  bool
	isList  bool
	long    string
	short   string
	thetype string
	ptr     any
	val     any
}

func newFlag(args ...any) (*cliFlag, error) {
	var f *cliFlag = &cliFlag{}

	for _, arg := range args {
		switch arg := arg.(type) {
		case *bool, *float64, *int, *int64, *string, *uint, *uint64:
			f.ptr = arg
		case *Counter, *FloatList, *IntList, *StringList, *UintList:
			f.isList = true
			f.ptr = arg
		case bool:
			// First time thru, set val
			if !f.gotVal {
				f.gotVal = true
				f.val = arg
			} else { // Otherwise, set hidden
				f.hidden = arg
			}
		case float64, int, int64, uint, uint64:
			f.gotVal = true
			f.val = arg
		case string:
			f.processString(arg)
		default:
			return nil, errors.Newf("unsupported flag type")
		}
	}

	f.setType()

	return f, nil
}

func (f *cliFlag) column(align bool) string {
	var fillto int
	var sb strings.Builder
	var sep string

	// Short flag
	if f.short != "" {
		sb.WriteString("-" + f.short)

		if (f.thetype != "") && (f.long == "") {
			sb.WriteString(" " + f.thetype)
		}
	}

	// Separator
	sep = "  "

	if f.short != "" {
		if f.long != "" {
			sep = ", "
		}

		sb.WriteString(sep)
	}

	// Alignment
	if align {
		fillto = colWidth.short + len(sep) - sb.Len()
		for range fillto {
			sb.WriteString(" ")
		}
	}

	// Long flag
	if f.long != "" {
		sb.WriteString("--" + f.long)

		if f.thetype != "" {
			sb.WriteString("=" + f.thetype)
		}
	}

	return sb.String()
}

//nolint:cyclop,gocyclo // sometimes ugly is necessary
func (f *cliFlag) enable(s string) error {
	if s == "" {
		return nil
	}

	switch ptr := f.ptr.(type) {
	case *bool:
		if val, ok := f.val.(bool); ok {
			flag.BoolVar(ptr, s, val, f.desc)
			break
		}

		return errors.Newf("invalid bool %v", f.val)
	case *Counter:
		flag.Var(ptr, s, f.desc)
	case *float64:
		if val, ok := f.val.(float64); ok {
			flag.Float64Var(ptr, s, val, f.desc)
			break
		}

		return errors.Newf("invalid float64 %v", f.val)
	case *FloatList:
		flag.Var(ptr, s, f.desc)
	case *int:
		if val, ok := f.val.(int); ok {
			flag.IntVar(ptr, s, val, f.desc)
			break
		}

		return errors.Newf("invalid int %v", f.val)
	case *int64:
		if val, ok := f.val.(int64); ok {
			flag.Int64Var(ptr, s, val, f.desc)
			break
		}

		return errors.Newf("invalid int64 %v", f.val)
	case *IntList:
		flag.Var(ptr, s, f.desc)
	case *string:
		if val, ok := f.val.(string); ok {
			flag.StringVar(ptr, s, val, f.desc)
			break
		}

		return errors.Newf("invalid string %v", f.val)
	case *StringList:
		flag.Var(ptr, s, f.desc)
	case *uint:
		if val, ok := f.val.(int); ok {
			if val < 0 {
				return errors.Newf("invalid uint %v", val)
			}

			flag.UintVar(ptr, s, uint(val), f.desc)

			break
		}

		return errors.Newf("invalid uint %v", f.val)
	case *uint64:
		if val, ok := f.val.(int64); ok {
			if val < 0 {
				return errors.Newf("invalid uint64 %v", val)
			}

			flag.Uint64Var(ptr, s, uint64(val), f.desc)

			break
		}

		return errors.Newf("invalid uint64 %v", f.val)
	case *UintList:
		flag.Var(ptr, s, f.desc)
	}

	return nil
}

func (f *cliFlag) processString(arg string) {
	var validLong bool = !f.gotVal && !strings.Contains(arg, " ")

	switch {
	case f.desc != "":
		f.desc += " " + strings.TrimSpace(arg)
	case (f.short == "") && (len(arg) == 1):
		f.short = strings.TrimSpace(arg)
	case (f.long == "") && (len(arg) > 1) && validLong:
		f.long = strings.TrimSpace(arg)
	case !f.isList && (f.val == nil):
		f.gotVal = true
		f.val = arg
	default:
		f.desc = strings.TrimSpace(arg)
	}
}

func (f *cliFlag) setType() {
	switch f.ptr.(type) {
	case *float64, *FloatList:
		f.thetype = "FLOAT"
	case *int, *int64, *IntList:
		f.thetype = "INT"
	case *string, *StringList:
		f.thetype = "STRING"
	case *uint, *uint64, *UintList:
		f.thetype = "UINT"
	}
}

// String will return a string representation of the cliFlag.
func (f *cliFlag) String() string {
	//nolint:mnd // 2 is not a magic number
	var enoughRoom bool = (TabWidth + colWidth.left) <= (MaxWidth / 2)
	var fillto int
	var lines []string
	var sb strings.Builder

	// Leading space
	for range TabWidth {
		sb.WriteString(" ")
	}

	// Flags
	sb.WriteString(f.column(Align && enoughRoom))

	// Description
	if Align && enoughRoom {
		// Filler
		fillto = TabWidth + colWidth.left - sb.Len()
		for range fillto {
			sb.WriteString(" ")
		}

		lines = wrap(f.desc, colWidth.desc)
		for i, line := range lines {
			if i > 0 {
				// Leading space plus filler
				for range TabWidth + colWidth.left {
					sb.WriteString(" ")
				}
			}

			// Alignment
			for range TabWidth {
				sb.WriteString(" ")
			}

			// Actual description
			sb.WriteString(line + "\n")
		}
	} else {
		// New line b/c colWidth.left is too big
		sb.WriteString("\n")

		//nolint:mnd // 2 is not a magic number
		lines = wrap(f.desc, MaxWidth-(2*TabWidth))
		for _, line := range lines {
			// Leading space plus filler
			for range 2 * TabWidth {
				sb.WriteString(" ")
			}

			// Actual description
			sb.WriteString(line + "\n")
		}

		sb.WriteString("\n")
	}

	return sb.String()
}

func (f *cliFlag) table() string {
	var sb strings.Builder

	// Option
	if f.short != "" {
		sb.WriteString("`-" + f.short + "`")

		if f.long != "" {
			sb.WriteString(", ")
		}
	}

	if f.long != "" {
		sb.WriteString("`--" + f.long + "`")
	}

	// Separator
	sb.WriteString(" | ")

	// Args
	if f.thetype != "" {
		sb.WriteString("`" + f.thetype + "`")
	}

	// Separator
	sb.WriteString(" | ")

	// Description
	sb.WriteString(f.desc + "\n")

	return sb.String()
}

func (f *cliFlag) updateMaxWidth() {
	var dw int
	var lw int = 0
	var sep int = 2
	var sw int = 2

	if f.long != "" {
		//nolint:mnd // 2 is not a magic number
		lw = 2 + len(f.long) + len(f.thetype)

		if f.thetype != "" {
			lw++
		}
	}

	dw = MaxWidth - TabWidth - sw - sep - lw - TabWidth

	if dw < colWidth.desc {
		colWidth.desc = dw
	}

	if lw > colWidth.long {
		colWidth.long = lw
	}

	colWidth.short = 2
	colWidth.left = colWidth.short + sep + colWidth.long
}

func (f *cliFlag) validate() error {
	if (f.short == "") && (f.long == "") {
		return errors.New("no flag provided")
	}

	if len(f.short) > 1 {
		return errors.Newf("invalid short flag \"-%s\"", f.short)
	}

	if len(f.long) == 1 {
		return errors.Newf("invalid long flag \"--%s\"", f.long)
	}

	if f.desc == "" {
		if f.long != "" {
			return errors.Newf("no description for \"--%s\"", f.long)
		}

		return errors.Newf("no description for \"-%s\"", f.short)
	}

	return nil
}
