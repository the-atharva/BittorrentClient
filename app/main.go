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

	if command == "decode" {
		bencodedValue := os.Args[2]

		decoded, err := app.decodeBencode(bencodedValue)
		if err != nil {
			fmt.Println(err)
			return
		}

		jsonOutput, _ := json.Marshal(decoded)
		fmt.Println(string(jsonOutput))
	} else {
		fmt.Println("Unknown command: " + command)
		os.Exit(1)
	}
}
