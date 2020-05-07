// aesthetic string translations
// helpful https://unicode-table.com
package main

import "strings"

// Translation maps input & output characters
type Translation struct {
	source string
	dest   []rune
}

func getAllTranslations() map[string]Translation {
	// global maps are tricky so get translations here
	return map[string]Translation{
		// keep alphabetized
		"acute": Translation{
			"ACEGILNORSUWYZacegilnorsuwyz",
			[]rune("ÁĆÉǴÍĹŃÓŔŚÚẂÝŹáćéǵíĺńóŕśúẃýź"),
		},
		"caron": Translation{
			"ACDEGHIKNORSTUZaceghijknorsuz",
			[]rune("ǍČĎĚǦȞǏǨŇǑŘŠŤǓŽǎčěǧȟǐǰǩňǒřšǔž"),
		},
		"diaeresis": Translation{
			"AEHIOUWXYaehiotwxy",
			[]rune("ӒËḦÏÖÜẄẌŸӓëḧïöẗẅẍÿ"),
		},
		"double_grave": Translation{
			"AEIORUaeioru",
			[]rune("ȀȄȈȌȐȔȁȅȉȍȑȕ"),
		},
		"grave": Translation{
			"AEINOUWYaeinouwy",
			[]rune("ÀÈÌǸÒÙẀỲàèìǹòùẁỳ"),
		},
		"tilde": Translation{
			"AEINOUVYaeinouvy",
			[]rune("ÃẼĨÑÕŨṼỸãẽĩñõũṽỹ"),
		},
	}
}

// GetTranslation gives you a named Translation
func GetTranslation(which string) Translation {
	// looks cleaner in calling code
	return getAllTranslations()[which]
}

// AvailableTranslations
// func AvailableTranslations() []string {

// Translate a string
func Translate(s string, t Translation) string {
	var out strings.Builder
	for _, r := range s {
		match := strings.IndexRune(t.source, r)
		if match >= 0 {
			out.WriteRune(t.dest[match])
		} else {
			out.WriteRune(r)
		}
	}
	return out.String()
}

// MaxTranslate returns translation with most characters changed
// func MaxTranslate(s string) string {
