package cli

import "strings"

type section struct {
	alignOn string
	md      bool
	text    string
	title   string
}

func (s section) String() (ret string) {
	var key string
	var max int
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
				for i := 0; i < TabWidth; i++ {
					ret += " "
				}

				ret += line + "\n"
			}
		} else {
			for _, line := range strings.Split(s.text, "\n") {
				line = strings.TrimSpace(line)
				key, _, _ = strings.Cut(line, s.alignOn)

				if len(key) > max {
					max = len(key)
				}
			}

			for _, line := range strings.Split(s.text, "\n") {
				line = strings.TrimSpace(line)

				key, val, _ = strings.Cut(line, s.alignOn)
				wrapped = wrap(val, MaxWidth-TabWidth-len(key)-4)

				for i, line := range wrapped {
					for j := 0; j < TabWidth; j++ {
						ret += " "
					}

					if i == 0 {
						ret += key
						for j := 0; j < max-len(key); j++ {
							ret += " "
						}
					} else {
						for j := 0; j < max; j++ {
							ret += " "
						}
					}

					for j := 0; j < TabWidth; j++ {
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
