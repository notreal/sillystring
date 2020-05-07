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
		if got := translate(test.input, getTranslations()["acute"]); got != test.want {
			t.Errorf("acute: %v -> %v", test.input, got)
		}
	}
}
