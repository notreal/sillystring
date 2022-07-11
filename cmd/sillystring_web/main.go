package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	ss "github.com/notreal/sillystring"
)

var (
	port = flag.Int("p", 8080, "port of server")
)

func main() {
	// Usage overrides default help message
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.VisitAll(func(f *flag.Flag) {
			dashFlag := "-" + f.Name
			fmt.Fprintf(os.Stderr, "  %3s  %v\n", dashFlag, f.Usage) // f.Value
		})
	}
	flag.Parse()
	http.HandleFunc("/", handler)
	url := "localhost:" + strconv.Itoa(*port)
	fmt.Println("Serving at " + url)
	log.Fatal(http.ListenAndServe(url, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "sillystring.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		what := r.FormValue("what")
		var translated string
		for _, v := range ss.GetAllTranslations() {
			translated += fmt.Sprintf("<div>%s</div>", ss.Translate(what, v))
		}
		fmt.Fprintf(w, "%s", translated)
	default:
		fmt.Fprintf(w, "Only GET and POST methods are supported.")
	}
}
