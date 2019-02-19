package main

import "testing"

func TestAcute(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"B", "B"},
		{"Nick Rains", "Ńíćk Ŕáíńś"},
	}
	for _, test := range tests {
		if got := Acute(test.input); got != test.want {
			t.Errorf("Acute(%v) -> %v", test.input, got)
		}
	}
}
