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

func printParsedFile(decodedFile *torrentFile) {
	fmt.Printf("\nannounce: %s", decodedFile.announce)
	fmt.Printf("\nlength: %d", decodedFile.info["length"])	
	fmt.Printf("\nname: %s", decodedFile.info["name"])	
	fmt.Printf("\npiece length: %d", decodedFile.info["piece length"])	
	fmt.Printf("\npieces: % x", decodedFile.info["pieces"])	

}
