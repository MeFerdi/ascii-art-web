package art

import (
	"fmt"
	"os"
	"strings"
)

func Reading(fileName string) ([]string, error) {

	// Read file
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("%v", err)
		fmt.Println()
		return nil, err
	}

	var splitFileData []string
	if fileName == "thinkertoy.txt" {
		splitFileData = strings.Split(string(fileData), "\r\n")
	} else {
		splitFileData = strings.Split(string(fileData), "\n")
	}

	// File Corrupted or Altered
	if len(splitFileData) != 856 {
		return nil, fmt.Errorf("file corrupted")
	}

	return splitFileData, err
}
