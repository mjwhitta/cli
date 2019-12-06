package cli

import (
	"fmt"
	"strconv"
)

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
