package sillystring

import (
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
			"Test circle_neg",
			args{"Test circle_neg", GetTranslation("circle_neg")},
			"🅣🅔🅢🅣 🅒🅘🅡🅒🅛🅔_🅝🅔🅖",
		},
		{
			"Test mathematical",
			args{"Test mathematical", GetTranslation("mathematical")},
			"𝓣𝓮𝓼𝓽 𝓶𝓪𝓽𝓱𝓮𝓶𝓪𝓽𝓲𝓬𝓪𝓵",
		},
		{
			"Test upside_down",
			args{"Test upside_down", GetTranslation("upside_down")},
			"uʍop_ǝpᴉsdn ʇsǝꓕ",
		},
		{
			"Test yi",
			args{"Test yi", GetTranslation("yi")},
			"ꋖꈼꌚꋖ ꐞꂑ",
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
