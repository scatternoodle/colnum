package main

import "testing"

func TestAlphaPos(t *testing.T) {
	tests := []struct {
		c rune
		p int
	}{
		{'A', 1},
		{'P', 16},
		{'Z', 26},
		{'a', 1},
		{'w', 23},
		{'z', 26},
		{'?', -1},
		{'4', -1},
	}

	for _, tt := range tests {
		t.Run(string(tt.c), func(t *testing.T) {
			if a := alphaPos(tt.c); a != tt.p {
				t.Errorf("have %d, want %d", a, tt.p)
			}
		})
	}
}

func TestColNum(t *testing.T) {
	tests := []struct {
		str     string
		num     int
		mustErr bool
	}{
		{"a", 1, false},
		{"A", 1, false},
		{"z", 26, false},
		{"Z", 26, false},
		{"aA", 27, false},
		{"Bb", 54, false},
		{"BBB", 0, true},
		{"42", 0, true},
		{"R2", 0, true},
		{"Y?", 0, true},
		{"%T", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			a, err := colNum(tt.str)
			if tt.mustErr && err == nil {
				t.Error("must error")
			}
			if !tt.mustErr && err != nil {
				t.Errorf("unexpected error: %s", err)
			}
			if a != tt.num {
				t.Errorf("have %d, want %d", a, tt.num)
			}
		})
	}
}
