package ascii

import "strings"

const Height = 8

func Generate(input string, bannerPath string) (string, error) {
	banner, err := LoadBanner(bannerPath)
	if err != nil {
		return "", err
	}

	lines := strings.Split(input, "\\n")
	var result strings.Builder

	for _, line := range lines {
		if line == "" {
			result.WriteString("\n")
			continue
		}

		for i := 0; i < Height; i++ {
			for _, ch := range line {
				if art, ok := banner[ch]; ok {
					result.WriteString(art[i])
				}
			}
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}
