package main

import (
	"flag"
	"html/template"
	"log"
	"os"
	"strings"
)

var tpl = `package {{.Package}}

// {{.Type | Title}}In godoc
//
// Check some {{.Type}} tpl slice of {{.Type}} types exists.
func {{.Type | Title}}In(a {{.Type}}, slice []{{.Type}}) bool {
	for _, b := range slice {
		if b == a {
			return true
		}
	}
	return false
}
`

type vars struct {
	Package string
	Type    string
}

func main() {
	var err error
	var v vars
	var out string
	flag.StringVar(&v.Package, "package", "coreslices", "The package name")
	flag.StringVar(&v.Type, "type", "", "The type used for the coreslices package functions")
	flag.StringVar(&out, "out", "stdout", "Out result")
	flag.Parse()

	funcMap := template.FuncMap{
		"Title": strings.Title,
	}

	outWriter := os.Stdout
	if out != "stdout" {
		outWriter, err = os.Create(out)
		if err != nil {
			log.Fatalf("error open out file: %v", err)
		}
	}

	t := template.Must(template.New("tpl").Funcs(funcMap).Parse(tpl))
	if err := t.Execute(outWriter, v); err != nil {
		log.Fatalf("error generate code: %v", err)
	}
}
