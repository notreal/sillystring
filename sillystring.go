// aesthetic string conversions

package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	// TODO make this a webserver
	input := strings.Join(os.Args[1:], " ")
	fmt.Println(Acute(input))
}

// Acute returns string with available characters replaced with acute version
func Acute(in string) string {
	matches := "ACEGILNORSUWYZacegilnorsuwyz"
	accents := RuneSlicer("ÁĆÉǴÍĹŃÓŔŚÚẂÝŹáćéǵíĺńóŕśúẃýź")

	var out strings.Builder
	for _, r := range in {
		match := strings.IndexRune(matches, r)
		if match >= 0 {
			out.WriteRune(accents[match])
		} else {
			out.WriteRune(r)
		}
	}
	return out.String()
}

// RuneSlicer returns rune slice given a string
func RuneSlicer(in string) []rune {
	var out []rune
	for len(in) > 0 {
		r, size := utf8.DecodeRuneInString(in)
		out = append(out, r)
		in = in[size:]
	}
	return out
}
