// aesthetic string conversions

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"unicode/utf8"
)

var serve = flag.Bool("s", false, "run webserver")

func main() {
	// TODO make this a webserver
	flag.Parse()
	if *serve {
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
	}
	input := strings.Join(os.Args[1:], " ")
	fmt.Println(Acute(input))
}

func handler(w http.ResponseWriter, r *http.Request) {
	var path = strings.Split(r.URL.Path, "/")[1:] //remove leading /
	for _, word := range path {
		fmt.Fprintf(w, "%s\n", Acute(word))
	}
}

// Acute returns string with available characters replaced with acute version
func Acute(in string) string {
	matches := "ACEGILNORSUWYZacegilnorsuwyz"
	accents := runeSlicer("ÁĆÉǴÍĹŃÓŔŚÚẂÝŹáćéǵíĺńóŕśúẃýź")

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

// runeSlicer returns rune slice given a string
func runeSlicer(in string) []rune {
	var out []rune
	for len(in) > 0 {
		r, size := utf8.DecodeRuneInString(in)
		out = append(out, r)
		in = in[size:]
	}
	return out
}
