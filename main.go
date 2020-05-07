package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// TODO
// move todos to gitlab webui
var (
	// HTTP
	serve = flag.Bool("s", false, "run http server")
	port  = flag.Int("p", 8080, "port of server, default 8080")
	// translations
	acute = flag.Bool("a", false, "acute")
	// TODO add more translations
)

func main() {
	flag.Parse()
	if *serve {
		http.HandleFunc("/", handler)
		url := "localhost:" + strconv.Itoa(*port)
		fmt.Println(url)
		log.Fatal(http.ListenAndServe(url, nil))
	}
	var which string
	// TODO remove flags from text
	text := strings.Join(os.Args[1:], " ")
	switch {
	case *acute:
		which = "acute"
	default:
		fmt.Println(text)
		return
	}
	translated := translate(text, getTranslations()[which])
	fmt.Println(translated)
}

func handler(w http.ResponseWriter, r *http.Request) {
	which := strings.Split(r.URL.Path, "/")[1]
	text := strings.Split(r.URL.Path, "/")[2]
	translated := translate(text, getTranslations()[which])
	fmt.Fprintf(w, "%s\n", translated)
}
