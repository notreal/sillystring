package main

// this file for cli & http
import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	// HTTP
	serve = flag.Bool("s", false, "run http server")
	port  = flag.Int("p", 8080, "port of server")
	// CLI // translations
	acute       = flag.Bool("a", false, "acute")
	caron       = flag.Bool("c", false, "caron")
	diaeresis   = flag.Bool("d", false, "diaeresis")
	doubleGrave = flag.Bool("dg", false, "double_grave")
	grave       = flag.Bool("g", false, "grave")
	tilde       = flag.Bool("t", false, "tilde")
)

func main() {
	flag.Parse()
	if *serve {
		http.HandleFunc("/", handler)
		url := "localhost:" + strconv.Itoa(*port)
		fmt.Println("Serving at " + url)
		log.Fatal(http.ListenAndServe(url, nil))
	}
	var which string
	text := strings.Join(flag.Args(), " ")
	switch {
	case *acute:
		which = "acute"
	case *caron:
		which = "caron"
	case *diaeresis:
		which = "diaeresis"
	case *doubleGrave:
		which = "double_grave"
	case *grave:
		which = "grave"
	case *tilde:
		which = "tilde"
	default:
		fmt.Println(text)
		return
	}
	translated := Translate(text, getAllTranslations()[which])
	fmt.Println(translated)
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
		which := r.FormValue("which")
		what := r.FormValue("what")
		translated := Translate(what, getAllTranslations()[which])
		fmt.Fprintf(w, "%s", translated)
	default:
		fmt.Fprintf(w, "Only GET and POST methods are supported.")
	}
}
