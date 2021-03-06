package cli

import (
	"flag"
	"fmt"
	"strings"
)

type cliFlag struct {
	desc    string
	gotVal  bool
	hidden  bool
	isList  bool
	long    string
	short   string
	thetype string
	ptr     interface{}
	val     interface{}
}

func newFlag(args ...interface{}) (*cliFlag, error) {
	var f *cliFlag = &cliFlag{}

	for _, arg := range args {
		switch arg.(type) {
		case *bool, *float64, *int, *int64, *string, *uint, *uint64:
			f.ptr = arg
		case *FloatList, *IntList, *StringList, *UintList:
			f.isList = true
			f.ptr = arg
		case bool:
			if f.val == nil {
				f.gotVal = true
				f.val = arg
			}
			f.hidden = arg.(bool)
		case float64, int, int64, uint, uint64:
			f.gotVal = true
			f.val = arg
		case string:
			f.processString(arg.(string))
		default:
			return nil, fmt.Errorf("cli: Unsupported flag type")
		}
	}

	f.setType()

	return f, nil
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

func (f *cliFlag) validate() error {
	if (f.short == "") && (f.long == "") {
		return fmt.Errorf("cli: No flag provided")
	}

	if len(f.short) > 1 {
		return fmt.Errorf("cli: Invalid short flag \"%s\"", f.short)
	}

	if len(f.long) == 1 {
		return fmt.Errorf("cli: Invalid long flag \"%s\"", f.long)
	}

	if f.desc == "" {
		if f.long != "" {
			return fmt.Errorf(
				"cli: No description provided for \"%s\" flag",
				f.long,
			)
		}

		return fmt.Errorf(
			"cli: No description provided for \"%s\" flag",
			f.short,
		)
	}

	return nil
}
