package cli

var Align = false
var Authors []string
var Banner = fmt.Sprintf("%s [OPTIONS]", os.Args[0])
var BugEmail string
var colWidth = columnWidth{0, 0, 1024, 0}
var ExitStatus string
var flags []flagVar
var help bool
var Info = ""
var less = func(i, j int) bool {
	// Sort by short flag
	var left = flags[i].short
	if len(left) == 0 {
		left = flags[i].long
	}

	var right = flags[j].short
	if len(right) == 0 {
		right = flags[j].long
	}

	// Fallback to long flag if comparing same short flag
	if strings.ToLower(left) == strings.ToLower(right) {
		if len(flags[i].long) != 0 {
			left = flags[i].long
		}
		if len(flags[j].long) != 0 {
			right = flags[j].long
		}
	}

	return (strings.ToLower(left) < strings.ToLower(right))
}
var MaxWidth = 80
var readme bool
var SeeAlso []string
var TabWidth = 4
var Title string

const Version = "1.7.2"
