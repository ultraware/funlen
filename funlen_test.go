package funlen

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testCases := []struct {
		dir            string
		lineLimit      int
		stmtLimit      int
		ignoreComments bool
	}{
		{
			dir:       "too_many_statements",
			lineLimit: 1,
			stmtLimit: 1,
		},
		{
			dir:       "too_many_lines",
			lineLimit: 1,
			stmtLimit: 10,
		},
		{
			dir:       "too_many_statements_inline_func",
			lineLimit: 1,
			stmtLimit: 1,
		},
		{
			dir:            "ignores_comments",
			lineLimit:      2,
			stmtLimit:      2,
			ignoreComments: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.dir, func(t *testing.T) {
			t.Parallel()

			a := NewAnalyzer(test.lineLimit, test.stmtLimit, test.ignoreComments)

			analysistest.Run(t, analysistest.TestData(), a, test.dir)
		})
	}
}
