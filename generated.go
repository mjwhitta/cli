// Code generated by gentypes; DO NOT EDIT.
package cli

import (
	"fmt"
	"strconv"

	"github.com/mjwhitta/errors"
)

// FloatList allows setting a value multiple times, as in:
// --flag=float1 --flag=float2
type FloatList []float64

// String will convert FloatList to a string.
func (list *FloatList) String() string {
	if len(*list) == 0 {
		return "[]"
	}

	return fmt.Sprint(*list)
}

// Set will append a float to a FloatList.
func (list *FloatList) Set(val string) error {
	var e error
	var v float64

	if v, e = strconv.ParseFloat(val, 64); e != nil {
		return errors.Newf("failed to parse %s as float: %w", val, e)
	}

	(*list) = append(*list, v)
	return nil
}

// IntList allows setting a value multiple times, as in:
// --flag=int1 --flag=int2
type IntList []int64

// String will convert IntList to a string.
func (list *IntList) String() string {
	if len(*list) == 0 {
		return "[]"
	}

	return fmt.Sprint(*list)
}

// Set will append a int to a IntList.
func (list *IntList) Set(val string) error {
	var e error
	var v int64

	if v, e = strconv.ParseInt(val, 0, 64); e != nil {
		return errors.Newf("failed to parse %s as int: %w", val, e)
	}

	(*list) = append(*list, v)
	return nil
}

// StringList allows setting a value multiple times, as in:
// --flag=string1 --flag=string2
type StringList []string

// String will convert StringList to a string.
func (list *StringList) String() string {
	if len(*list) == 0 {
		return "[]"
	}

	return fmt.Sprint(*list)
}

// Set will append a string to a StringList.
func (list *StringList) Set(val string) error {
	(*list) = append(*list, val)
	return nil
}

// UintList allows setting a value multiple times, as in:
// --flag=uint1 --flag=uint2
type UintList []uint64

// String will convert UintList to a string.
func (list *UintList) String() string {
	if len(*list) == 0 {
		return "[]"
	}

	return fmt.Sprint(*list)
}

// Set will append a uint to a UintList.
func (list *UintList) Set(val string) error {
	var e error
	var v uint64

	if v, e = strconv.ParseUint(val, 0, 64); e != nil {
		return errors.Newf("failed to parse %s as uint: %w", val, e)
	}

	(*list) = append(*list, v)
	return nil
}
