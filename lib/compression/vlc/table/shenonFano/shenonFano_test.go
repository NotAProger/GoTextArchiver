package shenonFano

import (
	"reflect"
	"testing"
)

func Test_besDividerPosition(t *testing.T) {
	tests := []struct {
		name  string
		codes []code
		want  int
	}{
		{
			name: "one eleemnt",
			codes: []code{
				{Quantity: 2},
			},
			want: 0,
		},
		{
			name: "two eleemnt",
			codes: []code{
				{Quantity: 2},
				{Quantity: 2},
			},
			want: 1,
		},
		{
			name: "three eleemnt",
			codes: []code{
				{Quantity: 2},

				{Quantity: 1},
				{Quantity: 1},
			},
			want: 1,
		},
		{
			name: "many eleemnt",
			codes: []code{
				{Quantity: 2},
				{Quantity: 2},

				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: 2,
		},
		{
			name: "uncertainty eleemnt (need rightmost)",
			codes: []code{
				{Quantity: 1},

				{Quantity: 1},
				{Quantity: 1},
			},
			want: 1,
		},
		{
			name: "uncertainty eleemnt (need rightmost)",
			codes: []code{
				{Quantity: 2},

				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bestDividerPosition(tt.codes); got != tt.want {
				t.Errorf("besDividerPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_assignCodes(t *testing.T) {
	tests := []struct {
		name  string
		codes []code
		want  []code
	}{
		{
			name: "one elem",
			codes: []code{
				{Quantity: 2},
			},
			want: []code{
				{Quantity: 2, Bits: 0, Size: 1},
			},
		},
		{
			name: "two elems",
			codes: []code{
				{Quantity: 2},
				{Quantity: 2},
			},
			want: []code{
				{Quantity: 2, Bits: 0, Size: 1},
				{Quantity: 2, Bits: 1, Size: 1},
			},
		},
		{
			name: "three elems, crtain position",
			codes: []code{
				{Quantity: 2}, // 0
				{Quantity: 1}, // 10
				{Quantity: 1}, // 11
			},
			want: []code{
				{Quantity: 2, Bits: 0, Size: 1},
				{Quantity: 1, Bits: 2, Size: 2}, // 10
				{Quantity: 1, Bits: 3, Size: 2},
			},
		},
		{
			name: "three elems, crtain position",
			codes: []code{
				{Quantity: 1}, // 0
				{Quantity: 1}, // 10
				{Quantity: 1}, // 11
			},
			want: []code{
				{Quantity: 1, Bits: 0, Size: 1},
				{Quantity: 1, Bits: 2, Size: 2}, // 10
				{Quantity: 1, Bits: 3, Size: 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assignCodes(tt.codes)

			if !reflect.DeepEqual(tt.codes, tt.want) {
				t.Errorf("got: %+v, \nwant: %+v", tt.codes, tt.want)
			}
		})
	}
}

func Test_build(t *testing.T) {
	tests := []struct {
		name string
		text string
		want encodingTable
	}{
		{
			name: "base test",
			text: "abbbcc",
			want: encodingTable{
				'a': code{
					Char:     'a',
					Quantity: 1,
					Bits:     3,
					Size:     2,
				},
				'b': code{
					Char:     'b',
					Quantity: 3,
					Bits:     0,
					Size:     1,
				},
				'c': code{
					Char:     'c',
					Quantity: 2,
					Bits:     2,
					Size:     2,
				},
			},
		},

		{
			name: "one char test",
			text: "aaaa",
			want: encodingTable{
				'a': code{
					Char:     'a',
					Quantity: 4,
					Bits:     0,
					Size:     1,
				},
			},
		},

		{
			name: "uncertain char test",
			text: "aabbcc",
			want: encodingTable{
				'a': code{
					Char:     'a',
					Quantity: 2,
					Bits:     0,
					Size:     1,
				},
				'b': code{
					Char:     'b',
					Quantity: 2,
					Bits:     2,
					Size:     2,
				},
				'c': code{
					Char:     'c',
					Quantity: 2,
					Bits:     3,
					Size:     2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := build(newCharStat(tt.text)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("build() = %v, want %v", got, tt.want)
			}
		})
	}
}
