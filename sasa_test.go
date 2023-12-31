package sasa

import (
	"fmt"
	"testing"
)

func TestTrimMargin(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "simple",
			s: `foo
			|bar
			|baz`,
			want: `foo
bar
baz`,
		},
		{
			name: "first line is empty",
			s: `
			|foo
			|bar
			|baz`,
			want: `foo
bar
baz`,
		},
		{
			name: "last line is empty",
			s: `foo
			|bar
			|baz
`,
			want: `foo
bar
baz`,
		},
		{
			name: "first and last lines are blank",
			s: `    
			|foo
			|bar
			|baz
			`,
			want: `foo
bar
baz`,
		},
		{
			name: "different marginPrefix indentation",
			s: `
			|foo
	|bar
		|baz`,
			want: `foo
bar
baz`,
		},
		{
			name: "spaces after marginPrefix",
			s: `
			| foo
			|   bar
			|  baz`,
			want: ` foo
   bar
  baz`,
		},
		{
			name: "single line",
			s:    `  |  foo`,
			want: `  foo`,
		},
		{
			name: "line contains multiple marginPrefix",
			s: `foo
			| | bar
			||baz`,
			want: `foo
 | bar
|baz`,
		},
		{
			name: "empty",
			s:    ``,
			want: ``,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := TrimMargin(tt.s)
			if got != tt.want {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestTrimMargin_MarginPrefix(t *testing.T) {
	tests := []struct {
		name         string
		marginPrefix string
		s            string
		want         string
	}{
		{
			name:         "simple",
			marginPrefix: "*",
			s: `foo
			*bar
			*baz`,
			want: `foo
bar
baz`,
		},
		{
			name:         "multiple characters",
			marginPrefix: "///",
			s: `    
			///foo
			///bar
			///baz
			`,
			want: `foo
bar
baz`,
		},
		{
			name:         "not ascii characters",
			marginPrefix: "☆",
			s: `
			☆ foo
			☆   bar
			☆  baz`,
			want: ` foo
   bar
  baz`,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := TrimMargin(tt.s, MarginPrefix(tt.marginPrefix))
			if got != tt.want {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestTrimIndent(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "simple",
			s: `
			foo
			bar
			baz`,
			want: `foo
bar
baz`,
		},
		{
			name: "contains spaces after indentation",
			s: `
			foo
				bar
			  baz`,
			want: `foo
	bar
  baz`,
		},
		{
			name: "empty",
			s:    ``,
			want: ``,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := TrimIndent(tt.s)
			if got != tt.want {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func ExampleTrimMargin() {
	s := TrimMargin(`
	|foo
	|bar
	`)
	fmt.Println(s)
	// Output:
	// foo
	// bar
}

func ExampleTrimIndent() {
	s := TrimIndent(`
	foo
	bar
	`)
	fmt.Println(s)
	// Output:
	// foo
	// bar
}
