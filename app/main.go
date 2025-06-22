package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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
	decodedString, _ := process(app, command, argument)
	jsonOutput, _ := json.Marshal(decodedString)
	fmt.Println(string(jsonOutput))
}

func  process(app *application, command, argument string) (interface{}, error) {
	switch command {
		case "decode":
			decodedString, err := app.decodeBencode(argument)
			return decodedString, err
		default:
			fmt.Println("Unknown command: " + command)
			os.Exit(1)
	}
	return "", nil
}
