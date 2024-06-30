package web

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// bannerMap is a map that stores the ASCII art for different banner files
var bannerMap map[string]string

// init initializes the bannerMap and loads the ASCII art from the banner files
func init() {
	bannerMap = make(map[string]string)
	loadBanner(filepath.Join("banners", "shadow.txt"))
	loadBanner(filepath.Join("banners", "standard.txt"))
	loadBanner(filepath.Join("banners", "thinkertoy.txt"))
}

// loadBanner reads the contents of a banner file and stores it in the bannerMap
func loadBanner(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filename, err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if len(lines) == 0 {
		return
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		os.Exit(1)
	}

	bannerMap[filepath.Base(filename)] = strings.Join(lines, "\n")
}

// GetLetterArray retrieves the ASCII art representation for a given character from the specified banner file
func GetLetterArray(char rune, bannerStyle string) []string {
	banner, ok := bannerMap[bannerStyle+".txt"]
	if !ok {
		return []string{}
	}
	alphabet := strings.Split(banner, "\n")
	start := (int(char) - 32) * 9
	if start < 0 || start >= len(alphabet) {
		return []string{}
	}
	arr := alphabet[start : start+9]
	return arr
}

// PrintAscii returns the ASCII art representation of a given string
func PrintAscii(str, bannerStyle string) (string, error) {
	lines := strings.Split(str, "\n")
	letters := [][]string{}
	for _, line := range lines {
		for _, letter := range line {
			if letter < 32 || letter > 126 {
				return "", fmt.Errorf("non-ASCII character '%c' encountered", letter)
			}
			arr := GetLetterArray(letter, bannerStyle)
			letters = append(letters, arr)
		}
	}
	var result strings.Builder
	// Print the ASCII art vertically
	for i := 1; i < 9; i++ {
		for _, letter := range letters {
			if len(letter) < i {
				return "", fmt.Errorf("error: file content modified")
			}
			result.WriteString(letter[i])
		}
		if i < 8 {
			result.WriteString("\n")
		}
	}
	return result.String(), nil
}
