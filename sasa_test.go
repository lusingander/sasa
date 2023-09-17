package sasa

import "testing"

func TestTrimMargin(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{
			s: `foo
			|bar
			|baz`,
			want: `foo
bar
baz`,
		},
		{
			s: `
			|foo
			|bar
			|baz`,
			want: `foo
bar
baz`,
		},
		{
			s: `foo
			|bar
			|baz
`,
			want: `foo
bar
baz`,
		},
		{
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
			s: `
			|foo
	|bar
		|baz`,
			want: `foo
bar
baz`,
		},
		{
			s: `
			| foo
			|   bar
			|  baz`,
			want: ` foo
   bar
  baz`,
		},
		{
			s:    `  |  foo`,
			want: `  foo`,
		},
		{
			s:    ``,
			want: ``,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run("", func(t *testing.T) {
			got := TrimMargin(tt.s)
			if got != tt.want {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}
