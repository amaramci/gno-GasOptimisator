package main

import (
	"fmt"
	"os"

	"github.com/amaramci/gnogas/analyzer"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gnogas-analyze <path-to-gno-file>")
		os.Exit(1)
	}

	path := os.Args[1]
	stats, err := analyzer.AnalyzeFile(path)
	if err != nil {
		fmt.Println("Error analyzing file:", err)
		os.Exit(1)
	}

	fmt.Printf("Lines: %d\nAssignments: %d\nFunction Calls: %d\n", stats.Lines, stats.Assignments, stats.FunctionCalls)
}
