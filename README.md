# cli

[![Yum](https://img.shields.io/badge/-Buy%20me%20a%20cookie-blue?labelColor=grey&logo=cookiecutter&style=for-the-badge)](https://www.buymeacoffee.com/mjwhitta)

[![Go Report Card](https://goreportcard.com/badge/github.com/mjwhitta/cli?style=for-the-badge)](https://goreportcard.com/report/github.com/mjwhitta/cli)
![License](https://img.shields.io/github/license/mjwhitta/cli?style=for-the-badge)

## What is this?

A simple wrapper for Go's `flag` package.

## How to install

Open a terminal and run the following:

```
$ go get --ldflags "-s -w" --trimpath -u github.com/mjwhitta/cli
```

## Usage

See the [example](cmd/example) for an in-depth usage. Below is a
sample usage to create simple main package:

```
package main

import (
    "fmt"
    "os"
    "strings"

    "github.com/mjwhitta/cli"
)

// Flags
var flags struct {
    aBool   bool
    aString string
}

func init() {
    // Configure cli package
    cli.Align = true // Defaults to false
    cli.Authors = []string{"Your Name <your@email.tld>"}
    cli.Banner = fmt.Sprintf("%s [OPTIONS] <arg>", os.Args[0])
    cli.ExitStatus = strings.Join(
        []string{
            "Normally the exit status is 0. In the event of an",
            "error, the exit status will be 1.",
        },
        " ",
    )
    cli.Info = "Lorem ipsum dolor sit amet, consectetur adipiscing"

    // Parse cli flags
    cli.Flag(&flags.aBool, "b", "bool", false, "Sample boolean flag.")
    cli.Flag(&flags.aString, "s", "", "Sample string flag.")
    cli.Parse()

    // Validate cli args
    if cli.NArg() == 0 {
        cli.Usage(1)
    } else if cli.NArg() > 1 {
        cli.Usage(1)
    } else if flags.aString == "" {
        cli.Usage(1)
    }
}

func main() {
    fmt.Printf("%t\n", flags.aBool)
    fmt.Printf("%s\n", flags.aString)
    fmt.Printf("%d - %s\n", cli.NArg(), cli.Args())
}
```

### Configuring

Export           | Default               | Description
------           | -------               | -----------
`cli.Align`      | false                 | Aligned the columns
`cli.Authors`    | [""]                  | List of authors
`cli.Banner`     | "Usage: $0 [OPTIONS]" | The usage example
`cli.BugEmail`   | ""                    | Email for reporting bugs
`cli.ExitStatus` | ""                    | Description of all possible exit statuses
`cli.Info`       | ""                    | The description of the tool
`cli.MaxWidth`   | 80                    | Maximum width of usage
`cli.SeeAlso`    | [""]                  | List of other packages for more info
`cli.TabWidth`   | 4                     | The number of spaces between columns
`cli.Title`      | ""                    | Title for generated README.md

### Functions

The following methods can be used to create flags:

- `Flag(ptr *interface{}, long string, val interface{}, desc string)`
- `Flag(ptr *interface{}, short string, val interface{}, desc string)`
- `Flag(ptr *interface{}, short string, long string, val interface{}, desc string)`

You can use `Section(title string, text string)` to add new custom
sections. Other functions that simply wrap the `flag` package include:

- `NArg()`
- `NFlag()`
- `Parse()`
- `Parsed()`
- `PrintDefaults()`

Additional functions include:

- `PrintExtra()`
- `PrintHeader()`
- `Readme()`

And finally to print the usage message use `Usage(status int)`

## Links

- [Source](https://github.com/mjwhitta/cli)
