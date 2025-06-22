package main

import (
	"encoding/json"
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
	res, _ := process(app, command, argument)
	jsonOutput, _ := json.Marshal(res)
	fmt.Println(string(jsonOutput))
}

func  process(app *application, command, argument string) (interface{}, error) {
	switch command {
		case "decode":
			decodedString, err := app.decodeBencode(strings.NewReader(argument))
			return decodedString, err
		case "info":
			decodedfile, err := app.parseTorrentFile(argument)
			return decodedfile, err	
		default:
			fmt.Println("Unknown command: " + command)
			os.Exit(1)
	}
	return "", nil
}
















