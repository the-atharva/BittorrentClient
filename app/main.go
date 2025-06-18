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
	decodedString, _ := app.process(command, argument)
	jsonOutput, _ := json.Marshal(decodedString)
	fmt.Println(string(jsonOutput))
}


