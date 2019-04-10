package main

import (
    "fmt"
    "os"
    "strings"

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
    cli.Info = strings.Join(
        []string{
            "Lorem ipsum dolor sit amet, consectetur adipiscing",
            "elit. Mauris ut augue euismod, cursus nulla ut, semper",
            "eros. Integer pulvinar a lectus sed pretium. Cras a",
            "luctus odio, eget sagittis leo. Interdum et malesuada",
            "fames ac ante ipsum primis in faucibus. Nunc lectus",
            "metus, consectetur et tellus tempus, accumsan rhoncus",
            "arcu. Ut velit dui, aliquet eget tellus eget, commodo",
            "auctor enim. Duis blandit, metus vitae sagittis",
            "vestibulum, diam nisl fermentum lorem, non convallis",
            "enim urna rutrum mauris. Duis in scelerisque mauris.",
        },
        "",
    )
    cli.Bool(
        &abool,
        "a",
        "praesent",
        false,
        strings.Join(
            []string{
                "Fusce blandit nisi eu sem maximus, ut eleifend",
                "lectus semper.",
            },
            " ",
        ),
    )
    cli.Float64(
        &afloat,
        "b",
        "facilisis",
        0.0,
        strings.Join(
            []string{
                "Cras ac dolor ante. Nulla posuere non purus ac",
                "vehicula.",
            },
            " ",
        ),
    )
    cli.Int64(
        &aint,
        "c",
        "lorem",
        0,
        strings.Join(
            []string{
                "Nullam efficitur elit vel venenatis mattis. Aenean",
                "viverra tincidunt ligula.",
            },
            " ",
        ),
    )
    cli.String(
        &astring,
        "d",
        "ut",
        "",
        strings.Join(
            []string{
                "Nulla a sollicitudin ex. Sed faucibus lacus congue",
                "dapibus vestibulum.",
            },
            " ",
        ),
    )
    cli.StringList(
        &astringlist,
        "e",
        "placerat",
        strings.Join(
            []string{
                "Etiam efficitur interdum leo nec luctus. Etiam est",
                "erat, bibendum.",
            },
            " ",
        ),
    )
    cli.Uint64(
        &auint,
        "f",
        "sagittis",
        0,
        strings.Join(
            []string{
                "Donec posuere efficitur massa, ut imperdiet mauris",
                "sodales sit amet.",
            },
            " ",
        ),
    )
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
