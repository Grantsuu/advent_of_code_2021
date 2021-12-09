package util

import (
	"bufio"
	"os"
	"path/filepath"
)

func ParseInputFile(path string) ([]string, error) {
	absPath, _ := filepath.Abs(path)
	file, err := os.Open(absPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
