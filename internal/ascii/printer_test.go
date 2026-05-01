package ascii

import (
	"strings"
	"testing"
)

func TestGenerateSimpleAuto(t *testing.T) {
	input := "Hi"
	bannerPath := "../../banners/standard.txt"

	out, err := Generate(input, bannerPath)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	banner, err := LoadBanner(bannerPath)
	if err != nil {
		t.Fatalf("failed to load banner: %v", err)
	}

	expected := buildExpected(input, banner)

	if out != expected {
		t.Errorf("\nEXPECTED:\n%q\n\nGOT:\n%q\n", expected, out)
	}
}

func TestGenerateNewlineAuto(t *testing.T) {
	input := "Hi\\nThere"
	bannerPath := "../../banners/standard.txt"

	out, err := Generate(input, bannerPath)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	banner, err := LoadBanner(bannerPath)
	if err != nil {
		t.Fatalf("failed to load banner: %v", err)
	}

	expected := buildExpected(input, banner)

	if out != expected {
		t.Errorf("\nEXPECTED:\n%q\n\nGOT:\n%q\n", expected, out)
	}
}

func buildExpected(input string, banner map[rune][]string) string {
	lines := strings.Split(input, "\\n")
	var result strings.Builder

	for _, line := range lines {
		if line == "" {
			result.WriteString("\n")
			continue
		}

		for i := 0; i < 8; i++ {
			for _, ch := range line {
				if art, ok := banner[ch]; ok {
					result.WriteString(art[i])
				}
			}
			result.WriteString("\n")
		}
	}

	return result.String()
}
