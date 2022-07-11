package main

import (
	"flag"
	"fmt"
	"strings"

	ss "github.com/notreal/sillystring"
)

func main() {
	flag.Parse()
	text := strings.Join(flag.Args(), " ")
	var translated string
	for _, t := range ss.GetAllTranslations() {
		translated += ss.Translate(text, t) + "\n"
	}
	fmt.Println(translated)
}
