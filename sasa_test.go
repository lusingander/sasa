package sasa

import "testing"

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
