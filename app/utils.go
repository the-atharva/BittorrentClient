package main

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime/debug"
)

func (app *application) errorTrace(err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())	
	app.errLog.Output(2, trace)
	os.Exit(1)
}

func printDecodedString(decodedString any) {
	jsonOutput, _ := json.Marshal(decodedString)
	fmt.Println(string(jsonOutput))
}

func printParsedFile(decodedFile any) {
	if fileMap, ok := decodedFile.(map[string]any); ok {
		if trackerURL, exists := fileMap["announce"]; exists {
			fmt.Printf("\nTracker URL: %s", trackerURL)
		} else {
			fmt.Println("ERROR: Announce key doesn't exist")		
		}
		if info, exists := fileMap["info"]; exists {
			if infoMap, ok := info.(map[string]any); ok {
				if length, exists := infoMap["length"]; exists {
					fmt.Printf("\nLength: %v\n", length)
				} else {
					fmt.Println("ERROR: Length doesnmt exist")
				}
			}
		} else {
			fmt.Println("ERROR: Info doesn't exist")
		}
	} else {
		fmt.Println("Can't convert decoded torrent file to map")
	}
}
