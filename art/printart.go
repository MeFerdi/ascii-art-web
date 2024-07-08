package art

import (
	"fmt"
	"strings"
)

func PrintArt(slicedFile []string, strArg string) {

	if strArg == "\\n" {
		fmt.Println()
		return
	} else if strArg == "" {
		return
	} else if strArg == "\\t" {
		fmt.Println("	")
		return
	}

	// Handle unprintables
	notHandle := []string{"\\a", "\\b", "\\v", "\\f", "\\r"}

	for _, unpri := range notHandle {
		if strings.Contains(strArg, unpri) {
			fmt.Println("Contains unprintable")
			return
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
				fmt.Println("Unprintable Strings")
				return
			}
		}
	}

	for _, i := range argsSplit {
		if i == "" {
			fmt.Println()
			continue
		}
		// Loop through the eight lines
		for j := 0; j < 8; j++ {
			// Loop through fileData
			for _, k := range i {
				start := int(k-32)*9 + 1
				fmt.Printf(slicedFile[start+j])
			}
			fmt.Println()
		}
	}
}
