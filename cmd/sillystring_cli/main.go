package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	ss "github.com/notreal/sillystring"
)

var (
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
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.VisitAll(func(f *flag.Flag) {
			dashFlag := "-" + f.Name
			fmt.Fprintf(os.Stderr, "  %3s  %v\n", dashFlag, f.Usage) // f.Value
		})
	}
	flag.Parse()

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
	translated := ss.Translate(text, ss.GetAllTranslations()[which])
	fmt.Println(translated)
}
