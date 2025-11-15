package cli

import "strings"

type section struct {
	alignOn string
	md      bool
	text    string
	title   string
}

// String will return a string representation of the section.
func (s section) String() string {
	var key string
	var keyMaxLen int
	var sb strings.Builder
	var val string
	var wrapped []string

	if s.md {
		sb.WriteString("\n## " + s.title + "\n\n")

		for _, line := range wrap(s.text, MaxWidth) {
			sb.WriteString(line + "\n")
		}
	} else {
		sb.WriteString(s.title + "\n")

		if s.alignOn == "" {
			for _, line := range wrap(s.text, MaxWidth-TabWidth) {
				for range TabWidth {
					sb.WriteString(" ")
				}

				sb.WriteString(line + "\n")
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
				//nolint:mnd // 4 spaces for indent
				wrapped = wrap(val, MaxWidth-TabWidth-keyMaxLen-4)

				for i, line := range wrapped {
					for range TabWidth {
						sb.WriteString(" ")
					}

					if i == 0 {
						sb.WriteString(key)

						for range keyMaxLen - len(key) {
							sb.WriteString(" ")
						}
					} else {
						for range keyMaxLen {
							sb.WriteString(" ")
						}
					}

					for range TabWidth {
						sb.WriteString(" ")
					}

					sb.WriteString(line + "\n")
				}
			}
		}

		sb.WriteString("\n")
	}

	return sb.String()
}
