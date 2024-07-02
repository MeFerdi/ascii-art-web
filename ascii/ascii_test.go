package web

import (
	"path/filepath"
	"testing"
)

func TestLoadBanner(t *testing.T) {
	// Test loading a valid banner file
	loadBanner(filepath.Join("banners", "shadow.txt"))
	if _, ok := bannerMap["shadow.txt"]; !ok {
		t.Error("Failed to load banner file")
	}
	loadBanner(filepath.Join("banners", "thinkertoy.txt"))
	if _, ok := bannerMap["thinkertoy.txt"]; !ok {
		t.Error("Failed to load banner file")
	}
	loadBanner(filepath.Join("banners", "standard.txt"))
	if _, ok := bannerMap["standard.txt"]; !ok {
		t.Error("Failed to load the banner file")
	}

	// Add more test cases for loading other banner files or error scenarios
}

func TestPrintAscii(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		bannerStyle string
	}{
		{"Standard banner", "Hello", "standard.txt"},
		{"Shadow banner", "World", "shadow.txt"},
		{"Non-ASCII character", "Hello\x00", "standard.txt"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture the output of PrintAscii
			// ...
		})
	}
}

// Helper function to compare string slices
// func equalSlices(a, b []string) bool {
// 	if len(a) != len(b) {
// 		return false
// 	}

// 	for i := range a {
// 		if i >= len(b) || a[i] != b[i] {
// 			return false
// 		}
// 	}

// 	return true
// }

func Test_loadBanner(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loadBanner(tt.args.filename)
		})
	}
}
