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
	port  = flag.Int("p", 8080, "port of server, default 8080")
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
		fmt.Println(url)
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
	which := strings.Split(r.URL.Path, "/")[1]
	text := strings.Split(r.URL.Path, "/")[2]
	translated := Translate(text, getAllTranslations()[which])
	fmt.Fprintf(w, "%s\n", translated)
}
