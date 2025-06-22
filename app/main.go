package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)
type application struct {
	errLog *log.Logger
}

func main() {

	errorLog := log.New(os.Stderr, "ERROR\t", log.Lshortfile)
	app := &application {
		errLog: errorLog,
	}
	command := os.Args[1]
	argument := os.Args[2]
	process(app, command, argument)
	
}

func  process(app *application, command, argument string) (any, error) {
	switch command {
		case "decode":
			decodedString, err := app.decodeBencode(strings.NewReader(argument))
			printDecodedString(decodedString)
			return decodedString, err
		case "info":
			parsedFile, err := app.parseTorrentFile(argument)
			printParsedFile(parsedFile)
			return parsedFile, err	
		default:
			fmt.Println("Unknown command: " + command)
			os.Exit(1)
	}
	return "", nil
}
















