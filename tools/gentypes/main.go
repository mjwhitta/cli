package main

import (
	"os"
	"strings"

	"github.com/mjwhitta/errors"
)

func generateFuncs(f *os.File, t string) {
	var T string = strings.ToUpper(t[0:1]) + t[1:]
	var tn string = T + "List"

	// Type declaration
	f.WriteString("\n// " + tn + " allows setting a value multiple")
	f.WriteString(" times, as in:\n")
	f.WriteString("// --flag=" + t + "1 --flag=" + t + "2\n")
	f.WriteString("type " + tn + " []" + t)
	switch t {
	case "float", "int", "uint":
		f.WriteString("64")
	}
	f.WriteString("\n\n")

	// String() func
	f.WriteString("// String will convert " + tn + " to a string.\n")
	f.WriteString("func (list *" + tn + ") String() string {\n")
	f.WriteString("    if len(*list) == 0 {\n")
	f.WriteString("        return \"[]\"\n")
	f.WriteString("    }\n\n")
	f.WriteString("    return fmt.Sprint(*list)\n")
	f.WriteString("}\n\n")

	// Set() func
	f.WriteString("// Set will append a " + t + " to a " + tn + ".\n")
	f.WriteString("func (list *" + tn + ") Set(val string) error {\n")
	switch t {
	case "float":
		f.WriteString("    var e error\n")
		f.WriteString("    var v " + t + "64\n\n")
		f.WriteString("    if v, e = strconv.Parse" + T + "(val, 64)")
		f.WriteString("; e != nil {\n")
		f.WriteString("  return errors.Newf(\"failed to parse %s as ")
		f.WriteString(t + ": %w\", val, e)\n")
		f.WriteString("    }\n\n")
		f.WriteString("    (*list) = append(*list, v)\n")
	case "int", "uint":
		f.WriteString("    var e error\n")
		f.WriteString("    var v " + t + "64\n\n")
		f.WriteString("    if v, e = strconv.Parse" + T + "(val, 0,")
		f.WriteString(" 64); e != nil {\n")
		f.WriteString("  return errors.Newf(\"failed to parse %s as ")
		f.WriteString(t + ": %w\", val, e)\n")
		f.WriteString("    }\n\n")
		f.WriteString("    (*list) = append(*list, v)\n")
	case "string":
		f.WriteString("    (*list) = append(*list, val)\n")
	}
	f.WriteString("    return nil\n")
	f.WriteString("}\n")
}

func header(f *os.File) {
	f.WriteString("// Code generated by gentypes; DO NOT EDIT.\n")
	f.WriteString("package cli\n\n")
	f.WriteString("import (\n")
	f.WriteString("    \"fmt\"\n")
	f.WriteString("    \"strconv\"\n\n")
	f.WriteString("    \"github.com/mjwhitta/errors\"\n")
	f.WriteString(")\n")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			panic(r.(error).Error())
		}
	}()

	var e error
	var f *os.File
	var fn string = "generated.go"
	var types = []string{"float", "int", "string", "uint"}

	if f, e = os.Create(fn); e != nil {
		panic(errors.Newf("failed to create %s: %w", fn, e))
	}
	defer f.Close()

	header(f)

	for _, thetype := range types {
		generateFuncs(f, thetype)
	}
}
