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
	}

	for name, test := range testcases {
		t.Run(name, func(t *testing.T) {
			fset := token.NewFileSet()
			f, err := parser.ParseFile(fset, "", test.input, parser.AllErrors)

			if err != nil {
				t.Error("\nActual: ", err, "\n Did not expected error")
			}

			r := Run(f, fset, test.lineLimit, test.stmtLimit)
			actual := r[0].Message

			if !strings.Contains(actual, test.expected) {
				t.Error("\nActual: ", actual, "\nExpected: ", test.expected)
			}
		})
	}
}
