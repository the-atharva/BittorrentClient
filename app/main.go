package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)
type application struct {
	errLog *log.Logger
	torrentFile *torrentFile
}

func main() {

	errorLog := log.New(os.Stderr, "ERROR\t", log.Lshortfile)
	torrentFile := &torrentFile {
		announce: "",
		info: make(map[string]any),
		infoHash: nil,
	}
	app := &application {
		errLog: errorLog,
		torrentFile: torrentFile,
	}
	command := os.Args[1]
	argument := os.Args[2]
	process(app, command, argument)
	fmt.Println()	
}

func  process(app *application, command, argument string) (any, error) {
	switch command {
		case "decode":
			decodedString, err := app.decodeBencode(strings.NewReader(argument))
			printDecodedString(decodedString)
			return decodedString, err
	case "info":
			app.parseTorrentFile(argument)
			printParsedFile(app.torrentFile)
			fmt.Printf("\nInfo Hash: %x", app.torrentFile.infoHash)
			return app.torrentFile, nil	
		default:
			fmt.Println("Unknown command: " + command)
			os.Exit(1)
	}
	return "", nil
}
















