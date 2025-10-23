package cli

import "strings"

type section struct {
	alignOn string
	md      bool
	text    string
	title   string
}

// String will return a string representation of the section.
func (s section) String() (ret string) {
	var key string
	var keyMaxLen int
	var val string
	var wrapped []string

	if s.md {
		ret += "\n## " + s.title + "\n\n"

		for _, line := range wrap(s.text, MaxWidth) {
			ret += line + "\n"
		}
	} else {
		ret += s.title + "\n"

		if s.alignOn == "" {
			for _, line := range wrap(s.text, MaxWidth-TabWidth) {
				for range TabWidth {
					ret += " "
				}

				ret += line + "\n"
			}
		} else {
			for _, line := range strings.Split(s.text, "\n") {
				line = strings.TrimSpace(line)
				key, _, _ = strings.Cut(line, s.alignOn)

				if len(key) > keyMaxLen {
					keyMaxLen = len(key)
				}
			}

			for _, line := range strings.Split(s.text, "\n") {
				line = strings.TrimSpace(line)

				key, val, _ = strings.Cut(line, s.alignOn)
				//nolint:mnd // 4 is not a magic number
				wrapped = wrap(val, MaxWidth-TabWidth-keyMaxLen-4)

				for i, line := range wrapped {
					for range TabWidth {
						ret += " "
					}

					if i == 0 {
						ret += key
						for range keyMaxLen - len(key) {
							ret += " "
						}
					} else {
						for range keyMaxLen {
							ret += " "
						}
					}

					for range TabWidth {
						ret += " "
					}

					ret += line + "\n"
				}
			}
		}

		ret += "\n"
	}

	return ret
}
