# cli

A simple wrapper for golang's `flag` package.

## How to install

Open a terminal and run the following:

```
$ go get -u gitlab.com/mjwhitta/cli
```

## How to use

Below is a sample usage to create simple main package:

```
package main

import (
    "fmt"
    "os"

    "gitlab.com/mjwhitta/cli"
)

var abool bool
var afloat float64
var aint int64
var astring string
var astringlist cli.StringListVar
var auint uint64

func init() {
    // Parse cli args
    cli.Banner = fmt.Sprintf("Usage: %s [OPTIONS] <arg>", os.Args[0])
    cli.Info = "A sample usage for the cli package"
    cli.Bool(&abool, "b", "bool", false, "Example for bool flag")
    cli.Float64(
        &afloat,
        "f",
        "float",
        0.0,
        "Example for float64 flag",
    )
    cli.Int64(&aint, "i", "int", 0, "Example for int64 flag")
    cli.String(
        &astring,
        "s",
        "string",
        "value",
        "Example for string flag",
    )
    cli.StringList(
        &astringlist,
        "sl",
        "stringlist",
        "Example for []string flag",
    )
    cli.Uint64(&auint, "u", "uint", 0, "Example for uint64 arg")
    cli.Parse()

    // Validate cli args
    if (cli.NArg() == 0) {
        cli.Usage(1)
    } else if (len(astring) == 0) {
        cli.Usage(2)
    }
}

func main() {
    fmt.Printf("%t\n", abool)
    fmt.Printf("%f\n", afloat)
    fmt.Printf("%d\n", aint)
    fmt.Printf("%s\n", astring)
    fmt.Printf("%d - %s\n", len(astringlist), astringlist)
    fmt.Printf("%d\n", auint)
    fmt.Printf("%d %s\n", cli.NArg(), cli.Args())
}
```

The following methods can be used to create flags:

- `Bool(ptr *bool, short string, long string, init bool, desc string)`
- `Float64(ptr *float64, short string, long string, init float64, desc string)`
- `Float64List(ptr *cli.Float64ListVar, short string, long string, desc string)`
- `Int(ptr *int, short string, long string, init int, desc string)`
- `Int64(ptr *int64, short string, long string, init int64, desc string)`
- `Int64List(ptr *cli.Int64ListVar, short string, long string, desc string)`
- `IntList(ptr *cli.IntListVar, short string, long string, desc string)`
- `String(ptr *string, short string, long string, init string, desc string)`
- `StringList(ptr *cli.StringListVar, short string, long string, desc string)`
- `Uint(ptr *uint, short string, long string, init uint, desc string)`
- `Uint64(ptr *uint64, short string, long string, init uint64, desc string)`
- `Uint64List(ptr *cli.Uint64ListVar, short string, long string, desc string)`
- `UintList(ptr *cli.UintListVar, short string, long string, desc string)`

Other functions that simply wrap the `flag` package include:

- `NArg()`
- `NFlag()`
- `Parse()`
- `Parsed()`
- `PrintDefaults()`

And finally to print the usage message use `Usage(status int)`
