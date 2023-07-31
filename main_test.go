package funlen

import (
	"go/parser"
	"go/token"
	"strings"
	"testing"
)

func TestRunTable(t *testing.T) {
	testcases := map[string]struct {
		input     string
		expected  string
		lineLimit int
		stmtLimit int
	}{
		"too-many-statements": {
			input: `package main
	func main() {
	print("Hello, world!")
	print("Hello, world!")}`,
			expected:  "Function 'main' has too many statements (2 > 1)",
			lineLimit: 1,
			stmtLimit: 1,
		},
		"too-many-lines": {
			input: `package main
	import "fmt"
	func main() {
	print("main!")
	print("is!")
	print("too!")
	print("long")}`,
			expected:  "Function 'main' is too long (3 > 1)",
			lineLimit: 1,
			stmtLimit: 10,
		},
		"too-many-statements-inline-func": {
			input: `package main
	func main() {
	print("Hello, world!")
	if true {
		y := []int{1,2,3,4}
		for k, v := range y {
			f := func() { print("test") }
			f()
		}
	}
	print("Hello, world!")}`,
			expected:  "Function 'main' has too many statements (8 > 1)",
			lineLimit: 1,
			stmtLimit: 1,
		},
	}

	for name, test := range testcases {
		t.Run(name, func(t *testing.T) {
			fset := token.NewFileSet()
			f, err := parser.ParseFile(fset, "", test.input, parser.ParseComments)

			if err != nil {
				t.Error("\nActual: ", err, "\n Did not expected error")
			}

			r := Run(f, fset, test.lineLimit, test.stmtLimit, false)
			actual := r[0].Message

			if !strings.Contains(actual, test.expected) {
				t.Error("\nActual: ", actual, "\nExpected: ", test.expected)
			}
		})
	}
}

func TestRunIgnoresComments(t *testing.T) {

	input := `package main
	func main() {
	// Comment 1
	// Comment 2
	// Comment 3
	print("Hello, world!")}
	// Comment Doc
	func unittest() {
	// Comment 1
	// Comment 2
	print("Hello, world!")}
	// Comment 3`

	lineLimit := 2
	stmtLimit := 2

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", input, parser.ParseComments)

	if err != nil {
		t.Error("\nActual: ", err, "\n Did not expected error")
	}

	r := Run(f, fset, lineLimit, stmtLimit, true)

	if len(r) > 0 {
		t.Error("\nActual: ", r, "\nExpected no lint errors")
	}
}
