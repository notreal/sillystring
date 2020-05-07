// aesthetic string translations

package main

import "strings"

// global maps are tricky so get translations here
func getTranslations() map[string]translation {
	return map[string]translation{
		"acute": translation{
			"ACEGILNORSUWYZacegilnorsuwyz",
			[]rune("ÁĆÉǴÍĹŃÓŔŚÚẂÝŹáćéǵíĺńóŕśúẃýź"),
		},
	}
}

type translation struct {
	source string
	dest   []rune
}

func translate(s string, t translation) string {
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
