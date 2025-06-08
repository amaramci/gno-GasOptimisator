package analyzer

import (
	"bufio"
	"os"
	"strings"
)

type Stats struct {
	Lines         int
	Assignments   int
	FunctionCalls int
}

func AnalyzeFile(path string) (*Stats, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats := &Stats{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		stats.Lines++

		line = strings.TrimSpace(line)
		if strings.Contains(line, "<==") || strings.Contains(line, ":=") {
			stats.Assignments++
		}
		if strings.Contains(line, "(") && strings.Contains(line, ")") && !strings.HasPrefix(line, "package") {
			stats.FunctionCalls++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}
