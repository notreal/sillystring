package sillystring

import (
	"reflect"
	"testing"
)

// Not worth testing

// getAllTranslations
// getTranslation

func Test_translate(t *testing.T) {
	type args struct {
		s string
		t Translation
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Test Acute",
			args{"Test Acute", GetTranslation("acute")},
			"Téśt Áćúté",
		},
		{
			"Test Caron",
			args{"Test Caron", GetTranslation("caron")},
			"Ťěšt Čǎřǒň",
		},
		{
			"Test Diaeresis",
			args{"Test Diaeresis", GetTranslation("diaeresis")},
			"Tësẗ Dïӓërësïs",
		},
		{
			"Test Double Grave",
			args{"Test Double Grave", GetTranslation("double_grave")},
			"Tȅst Dȍȕblȅ Gȑȁvȅ",
		},
		{
			"Test Grave",
			args{"Test Grave", GetTranslation("grave")},
			"Tèst Gràvè",
		},
		{
			"Test Tilde",
			args{"Test Tilde", GetTranslation("tilde")},
			"Tẽst Tĩldẽ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Translate(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("translate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAvailableTranslations(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{"Test Available Translations",
			[]string{
				"acute",
				"caron",
				"diaeresis",
				"double_grave",
				"grave",
				"tilde",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AvailableTranslations(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AvailableTranslations() = %v, want %v", got, tt.want)
			}
		})
	}
}
