package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"text/template"
)

var tplTestTables = map[string]string{
	"byte": `
		{0x01, []byte{0x01, 0x02, 0x03}, true},
		{0x01, []byte{0x08, 0x09, 0x00}, false},
        {0x00, []byte{}, false},
        {0x00, []byte{0x00}, true},
`,
	"float32": `
		{.1, []float32{.1, .2, .3}, true},
		{.7, []float32{.8, .9}, false},
        {.01, []float32{}, false},
        {.01, []float32{.3, .2}, false},
`,
	"float64": `
		{.1, []float64{.1, .2, .3}, true},
		{.7, []float64{.8, .9}, false},
        {.01, []float64{}, false},
        {.01, []float64{.3, .2}, false},
`,
	"int": `
		{1, []int{1, 2, 3}, true},
		{7, []int{8, 9, 0}, false},
		{0, []int{}, false},
		{0, []int{-1, -2}, false},
`,
	"int32": `
		{1, []int32{1, 2, 3}, true},
		{7, []int32{8, 9, 0}, false},
		{0, []int32{}, false},
		{0, []int32{-1, -2}, false},
`,
	"int64": `
		{1, []int64{1, 2, 3}, true},
		{7, []int64{8, 9, 0}, false},
		{0, []int64{}, false},
		{0, []int64{-1, -2}, false},
`,
	"rune": `
		{'a', []rune{'a', 'b', 'c'}, true},
		{'w', []rune{'x', 'y', 'z'}, false},
        {0x00, []rune{}, false},
        {0x00, []rune{0x00}, true},
`,
	"string": `
		{"a", []string{"a", "b", "c"}, true},
		{"w", []string{"x", "y", "z"}, false},
        {"empty", []string{}, false},
        {"", []string{"empty", "blank"}, false},
`,
	"uint8": `
		{1, []uint8{1, 2, 3}, true},
		{7, []uint8{8, 9, 0}, false},
		{0, []uint8{}, false},
		{0, []uint8{1, 2}, false},
`,
	"uint16": `
		{1, []uint16{1, 2, 3}, true},
		{7, []uint16{8, 9, 0}, false},
		{0, []uint16{}, false},
		{0, []uint16{1, 2}, false},
`,
	"uint32": `
		{1, []uint32{1, 2, 3}, true},
		{7, []uint32{8, 9, 0}, false},
		{0, []uint32{}, false},
		{0, []uint32{1, 2}, false},
`,
	"uint64": `
		{1, []uint64{1, 2, 3}, true},
		{7, []uint64{8, 9, 0}, false},
		{0, []uint64{}, false},
		{0, []uint64{1, 2}, false},
`,
}

var tplTest = `package {{.Package}}

import "testing"

func Test{{.Type | Title}}In(t *testing.T) {
	tables := []struct{
		a      {{.Type}}
		slice  []{{.Type}}
		result bool
	}{{"{"}}{{.TestTables}}{{"}"}}

	for _, table := range tables {
		result := {{.Type | Title}}In(table.a, table.slice)
		if result != table.result {
			t.Errorf("a %#v in slice %#v incorrect, except %v actual %v", table.a, table.slice, table.result, result)
		}
	}
}
`

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
	Package    string
	Type       string
	TestTables string
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

	v.TestTables = tplTestTables[v.Type]

	outWriter := os.Stdout
	outTestWriter := os.Stdout
	if out != "stdout" {
		outWriter, err = os.Create(out)
		if err != nil {
			log.Fatalf("error open out file: %v", err)
		}

		// Test writter
		outTest := strings.ReplaceAll(out, ".go", "_test.go")
		outTestWriter, err = os.Create(outTest)
		if err != nil {
			log.Fatalf("error open out test file: %v", err)
		}
	}

	t := template.Must(template.New("tpl").Funcs(funcMap).Parse(tpl))
	if err := t.Execute(outWriter, v); err != nil {
		log.Fatalf("error generate code: %v", err)
	}

	// Test generate
	t = template.Must(template.New("tplTest").Funcs(funcMap).Parse(tplTest))
	if err := t.Execute(outTestWriter, v); err != nil {
		log.Fatalf("error generate test code: %v", err)
	}
}
