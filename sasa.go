package sasa

import (
	"math"
	"strings"
	"unicode"
)

type trimMarginOptions struct {
	marginPrefix string
}

type trimMarginOption func(*trimMarginOptions)

// MarginPrefix is an option that specifies the string to use as a delimiter for TrimMargin.
func MarginPrefix(s string) trimMarginOption {
	return func(opt *trimMarginOptions) {
		opt.marginPrefix = s
	}
}

// TrimMargin returns a string with leading whitespace removed up to the delimiter.
// The delimiter can be specified with the MarginPrefix option (default: "|").
// If the first and last blank lines are blank, those lines are also removed.
func TrimMargin(s string, options ...trimMarginOption) string {
	opts := &trimMarginOptions{
		marginPrefix: "|",
	}
	for _, opt := range options {
		opt(opts)
	}

	lines := strings.Split(s, "\n")
	ln := len(lines)

	ss := make([]string, 0, ln)

	for i, line := range lines {
		if i == 0 || i == ln-1 {
			if isBlank(line) {
				continue
			}
		}

		nonSpaceIdx := strings.IndexFunc(line, isNotSpace)
		if nonSpaceIdx == -1 {
			ss = append(ss, line)
			continue
		}

		if strings.HasPrefix(line[nonSpaceIdx:], opts.marginPrefix) {
			ss = append(ss, line[nonSpaceIdx+len(opts.marginPrefix):])
		} else {
			ss = append(ss, line)
		}
	}

	return strings.Join(ss, "\n")
}

// TrimIndent returns a string with leading whitespace characters removed by the minimum number of whitespace characters on each line.
// If the first and last blank lines are blank, those lines are also removed.
func TrimIndent(s string) string {
	lines := strings.Split(s, "\n")
	ln := len(lines)

	ss := make([]string, 0, ln)

	width := math.MaxInt
	for _, line := range lines {
		if isBlank(line) {
			continue
		}
		nonSpaceIdx := strings.IndexFunc(line, isNotSpace)
		if nonSpaceIdx == -1 {
			n := len(line)
			if n < width {
				width = n
			}
		} else {
			n := nonSpaceIdx
			if n < width {
				width = n
			}
		}
	}

	for i, line := range lines {
		if i == 0 || i == ln-1 {
			if isBlank(line) {
				continue
			}
		}

		if len(line) <= width {
			ss = append(ss, "")
		} else {
			ss = append(ss, line[width:])
		}
	}

	return strings.Join(ss, "\n")
}

func isNotSpace(r rune) bool {
	return !unicode.IsSpace(r)
}

func isBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}
