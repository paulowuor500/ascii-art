package ascii

import (
	"bufio"
	"fmt"
	"os"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open banner file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	bannerMap := make(map[rune][]string)

	currentChar := rune(32)
	var currentLines []string

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" && len(currentLines) == 0 {
			continue
		}

		currentLines = append(currentLines, line)

		if len(currentLines) == 8 {
			bannerMap[currentChar] = currentLines
			currentLines = nil
			currentChar++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading banner file: %v", err)
	}

	return bannerMap, nil
}
