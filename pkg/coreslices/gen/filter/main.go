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
		{
			slice:  []byte{0x00, 0x01, 0x02},
			filter:  func(i int, e byte) bool { return e >= 0x01 },
			result: []byte{0x01, 0x02},
		},
`,
	"float32": `
		{
			slice:  []float32{0.00, 0.01, 0.02},
			filter:  func(i int, e float32) bool { return e >= 0.01 },
			result: []float32{0.01, 0.02},
		},
`,
	"float64": `
		{
			slice:  []float64{0.00, 0.01, 0.02},
			filter:  func(i int, e float64) bool { return e >= 0.01 },
			result: []float64{0.01, 0.02},
		},
`,
	"int": `
		{
			slice:  []int{0, 1, 2},
			filter:  func(i int, e int) bool { return e >= 1 },
			result: []int{1, 2},
		},
`,
	"int32": `
		{
			slice:  []int32{0, 1, 2},
			filter:  func(i int, e int32) bool { return e >= 1 },
			result: []int32{1, 2},
		},
`,
	"int64": `
		{
			slice:  []int64{0, 1, 2},
			filter:  func(i int, e int64) bool { return e >= 1 },
			result: []int64{1, 2},
		},
`,
	"rune": `
		{
			slice:  []rune{0x31, 0x32, 0x33},
			filter:  func(i int, e rune) bool { return e >= 0x32 },
			result: []rune{0x32, 0x33},
		},
`,
	"string": `
		{
			slice:  []string{"a", "b", "c"},
			filter:  func(i int, e string) bool { return e != "b" },
			result: []string{"a", "c"},
		},
`,
	"uint8": `
		{
			slice:  []uint8{0, 1, 2},
			filter:  func(i int, e uint8) bool { return e >= 1 },
			result: []uint8{1, 2},
		},
`,
	"uint16": `
		{
			slice:  []uint16{0, 1, 2},
			filter:  func(i int, e uint16) bool { return e >= 1 },
			result: []uint16{1, 2},
		},
`,
	"uint32": `
		{
			slice:  []uint32{0, 1, 2},
			filter:  func(i int, e uint32) bool { return e >= 1 },
			result: []uint32{1, 2},
		},
`,
	"uint64": `
		{
			slice:  []uint64{0, 1, 2},
			filter:  func(i int, e uint64) bool { return e >= 1 },
			result: []uint64{1, 2},
		},
`,
}

var tplTest = `package {{.Package}}

import "testing"

func Test{{.Type | Title}}Filter(t *testing.T) {
	tables := []struct {
		slice  []{{.Type}}
		filter func(int, {{.Type}}) bool
		result []{{.Type}}
	}{{"{"}}{{.TestTables}}{{"}"}}

	for _, table := range tables {
		result := {{.Type | Title}}Filter(table.slice, table.filter)
		if len(table.result) != len(result) {
			t.Fatalf("excepted %d elements in result, but %d elements actual", len(table.result), len(result))
		}
		for i, e := range result {
			ee := table.result[i]
			if e != ee {
				t.Errorf("excepted %d element %v, but %v element actual", i, ee, e)
			}
		}
	}
}
`

var tpl = `package {{.Package}}

// {{.Type | Title}}Filter godoc
//
// Filter slice of types {{.Type}}.
func {{.Type | Title}}Filter(slice []{{.Type}}, filter func(int, {{.Type}}) bool) []{{.Type}} {
	result := make([]{{.Type}}, 0)
	for i, e := range slice {
		if filter(i, e) {
			result = append(result, e)
		}
	}
	return result
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
