package main

import (
	"github.com/chnsz/tf-linter/passes/h001"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(h001.Analyzer)
}
