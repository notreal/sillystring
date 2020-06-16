package main

// this file for cli & http
import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	// HTTP
	serve = flag.Bool("s", false, "run http server")
	port  = flag.Int("p", 8080, "port of server")
	// CLI // translations
	acute      = flag.Bool("a", false, "acute")
	caron      = flag.Bool("ca", false, "caron")
	circumflex = flag.Bool("ci", false, "circumflex")
	diaeresis  = flag.Bool("d", false, "diaeresis")
	dotAbove   = flag.Bool("da", false, "dot_above")
	dotBelow   = flag.Bool("db", false, "dot_below")
	fraktur    = flag.Bool("f", false, "fraktur")
	grave      = flag.Bool("g", false, "grave")
	hook       = flag.Bool("ho", false, "hook")
	tilde      = flag.Bool("t", false, "tilde")
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
	case *circumflex:
		which = "circumflex"
	case *diaeresis:
		which = "diaeresis"
	case *dotAbove:
		which = "dot_above"
	case *dotBelow:
		which = "dot_below"
	case *fraktur:
		which = "fraktur"
	case *grave:
		which = "grave"
	case *hook:
		which = "hook"
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
