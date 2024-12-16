package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/ultraware/funlen"
)

func main() {
	singlechecker.Main(funlen.NewAnalyzer(0, 0, false))
}
