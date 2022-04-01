package main

import (
	"github.com/chnsz/tf-linter/passes/argsAccCheck"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(argsAccCheck.Analyzer)
}
