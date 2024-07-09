package art

import (
	"strings"
)

func PrintArt(slicedFile []string, strArg string) string {
	var result strings.Builder

	if strArg == "\\n" {
		result.WriteString("\n")
		return result.String()
	} else if strArg == "" {
		return ""
	} else if strArg == "\\t" {
		result.WriteString("    ")
		return result.String()
	}

	// Handle unprintables
	notHandle := []string{"\\a", "\\b", "\\v", "\\f", "\\r"}

	for _, unpri := range notHandle {
		if strings.Contains(strArg, unpri) {
			return "Contains unprintable"
		}
	}

	// Handle tab
	args := strings.Replace(strArg, "\\t", "    ", -1)
	argsN := strings.ReplaceAll(args, "\\n", "\n")

	// Handle newline
	argsSplit := strings.Split(argsN, "\n")

	// Handle foreign inputs
	for _, v := range argsSplit {
		for _, val := range v {
			if val < 32 || val > 126 {
				return "Unprintable Strings"
			}
		}
	}

	for _, i := range argsSplit {
		if i == "" {
			result.WriteString("\n")
			continue
		}
		// Loop through the eight lines
		for j := 0; j < 8; j++ {
			// Loop through fileData
			for _, k := range i {
				start := int(k-32)*9 + 1
				result.WriteString(slicedFile[start+j])
			}
			result.WriteString("\n")
		}
	}

	return result.String()
}
