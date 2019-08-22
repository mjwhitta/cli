# cli

A simple wrapper for golang's `flag` package.

## How to install

Open a terminal and run the following:

```
$ go get -u gitlab.com/mjwhitta/cli
```

## How to use

See the [examples](examples) for in-depth usage. Below is a sample
usage to create simple main package:

```
package main

import (
    "fmt"
    "os"
    "strings"

    "gitlab.com/mjwhitta/cli"
)

var abool bool
var aint int
var astring string

func init() {
    // Configure cli package
    cli.Align = true
    cli.Authors = []string{"Some Person <someperson@some.domain>"}
    cli.Banner = fmt.Sprintf("Usage: %s [OPTIONS] <arg>", os.Args[0])
    cli.BugEmail = "bugs@some.domain"
    cli.ExitStatus = strings.Join(
        []string{
            "Normally the exit status is 0. In the event of invalid",
            "or missing arguments, the exit status will be non-zero.",
        },
        " ",
    )
    cli.Info = "A sample usage for the cli package"
    cli.MaxWidth = 80 // Default
    cli.SeeAlso = []string{"some", "other", "tools"}
    cli.TabWidth = 4 // Default
    cli.Title = "Some tool"

    // Parse cli flags
    cli.Flag(&abool, "bool", false, "Example for bool flag.")
    cli.Flag(&aint, "i", "int", 0, "Example for int flag.")
    cli.Flag(&astring, "string", "value", "Example for string flag.")
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
    fmt.Printf("%d\n", aint)
    fmt.Printf("%s\n", astring)
    fmt.Printf("%d %s\n", cli.NArg(), cli.Args())
}
```

### Configuring

Export                | Default               | Description
------                | -------               | -----------
`cli.Align`           | false                 | Aligned the columns
`cli.Authors`         | [""]                  | List of authors
`cli.Banner`          | "Usage: $0 [OPTIONS]" | The usage example
`cli.BugEmail`        | ""                    | Email for reporting bugs
`cli.Info`            | ""                    | The description of the tool
`cli.MaxWidth`        | 80                    | Maximum width of usage
`cli.SeeAlso`         | [""]                  | List of other packages for more info
`cli.TabWidth`        | 4                     | The number of spaces between columns
`cli.Title`           | ""                    | Title for generated README.md

### Functions

The following methods can be used to create flags:

- `Flag(ptr *interface{}, long string, val interface{}, desc string)`
- `Flag(ptr *interface{}, short string, val interface{}, desc string)`
- `Flag(ptr *interface{}, short string, long string, val interface{}, desc string)`

Other functions that simply wrap the `flag` package include:

- `NArg()`
- `NFlag()`
- `Parse()`
- `Parsed()`
- `PrintDefaults()`

Additional functions include:

- `PrintHeader()`
- `Readme()`

And finally to print the usage message use `Usage(status int)`
