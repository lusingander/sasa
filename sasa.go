package sasa

import (
	"strings"
	"unicode"
)

func TrimMargin(s string) string {
	marginPrefix := "|"

	lines := strings.Split(s, "\n")
	ln := len(lines)

	ss := make([]string, 0, ln)

	for i, line := range lines {
		if i == 0 || i == ln-1 {
			if strings.TrimSpace(line) == "" {
				continue
			}
		}

		nonSpaceIdx := strings.IndexFunc(line, isNotSpace)
		if nonSpaceIdx == -1 {
			ss = append(ss, line)
			continue
		}

		if strings.HasPrefix(line[nonSpaceIdx:], marginPrefix) {
			ss = append(ss, line[nonSpaceIdx+len(marginPrefix):])
		} else {
			ss = append(ss, line)
		}
	}

	return strings.Join(ss, "\n")
}

func isNotSpace(r rune) bool {
	return !unicode.IsSpace(r)
}
