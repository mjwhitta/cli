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

func (f *cliFlag) enable() {
	f.enableShort()
	f.enableLong()
}

func (f *cliFlag) enableLong() {
	if f.long == "" {
		return
	}

	switch f.ptr.(type) {
	case *bool:
		flag.BoolVar(f.ptr.(*bool), f.long, f.val.(bool), f.desc)
	case *Counter:
		flag.Var(f.ptr.(*Counter), f.long, f.desc)
	case *float64:
		flag.Float64Var(
			f.ptr.(*float64),
			f.long,
			f.val.(float64),
			f.desc,
		)
	case *FloatList:
		flag.Var(f.ptr.(*FloatList), f.long, f.desc)
	case *int:
		flag.IntVar(f.ptr.(*int), f.long, f.val.(int), f.desc)
	case *int64:
		flag.Int64Var(
			f.ptr.(*int64),
			f.long,
			int64(f.val.(int)),
			f.desc,
		)
	case *IntList:
		flag.Var(f.ptr.(*IntList), f.long, f.desc)
	case *string:
		flag.StringVar(
			f.ptr.(*string),
			f.long,
			f.val.(string),
			f.desc,
		)
	case *StringList:
		flag.Var(f.ptr.(*StringList), f.long, f.desc)
	case *uint:
		flag.UintVar(f.ptr.(*uint), f.long, uint(f.val.(int)), f.desc)
	case *uint64:
		flag.Uint64Var(
			f.ptr.(*uint64),
			f.long,
			uint64(f.val.(int)),
			f.desc,
		)
	case *UintList:
		flag.Var(f.ptr.(*UintList), f.long, f.desc)
	}
}

func (f *cliFlag) enableShort() {
	if f.short == "" {
		return
	}

	switch f.ptr.(type) {
	case *bool:
		flag.BoolVar(f.ptr.(*bool), f.short, f.val.(bool), f.desc)
	case *Counter:
		flag.Var(f.ptr.(*Counter), f.short, f.desc)
	case *float64:
		flag.Float64Var(
			f.ptr.(*float64),
			f.short,
			f.val.(float64),
			f.desc,
		)
	case *FloatList:
		flag.Var(f.ptr.(*FloatList), f.short, f.desc)
	case *int:
		flag.IntVar(f.ptr.(*int), f.short, f.val.(int), f.desc)
	case *int64:
		flag.Int64Var(
			f.ptr.(*int64),
			f.short,
			int64(f.val.(int)),
			f.desc,
		)
	case *IntList:
		flag.Var(f.ptr.(*IntList), f.short, f.desc)
	case *string:
		flag.StringVar(
			f.ptr.(*string),
			f.short,
			f.val.(string),
			f.desc,
		)
	case *StringList:
		flag.Var(f.ptr.(*StringList), f.short, f.desc)
	case *uint:
		flag.UintVar(
			f.ptr.(*uint),
			f.short,
			uint(f.val.(int)),
			f.desc,
		)
	case *uint64:
		flag.Uint64Var(
			f.ptr.(*uint64),
			f.short,
			uint64(f.val.(int)),
			f.desc,
		)
	case *UintList:
		flag.Var(f.ptr.(*UintList), f.short, f.desc)
	}
}

func (f *cliFlag) processString(arg string) {
	var validLong bool = !f.gotVal && !strings.Contains(arg, " ")

	if f.desc != "" {
		f.desc += " " + strings.TrimSpace(arg)
	} else if (f.short == "") && (len(arg) == 1) {
		f.short = strings.TrimSpace(arg)
	} else if (f.long == "") && (len(arg) > 1) && validLong {
		f.long = strings.TrimSpace(arg)
	} else if !f.isList && (f.val == nil) {
		f.gotVal = true
		f.val = arg
	} else {
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
	var enoughRoom bool = (TabWidth + colWidth.left) <= (MaxWidth / 2)
	var fillto int
	var lines []string
	var str string

	// Leading space
	for i := 0; i < TabWidth; i++ {
		str += " "
	}

	// Flags
	str += f.column(Align && enoughRoom)

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

func (f *cliFlag) table() string {
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

func (f *cliFlag) updateMaxWidth() {
	var dw int
	var lw int = 0
	var sep int = 2
	var sw int = 2

	if f.long != "" {
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
