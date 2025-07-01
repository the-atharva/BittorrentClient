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

var app *application

func main() {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Lshortfile)
	torrentInfo := torrentInfo {
		length: 0,
		name: "",
		pieceLength: 0,
		pieces: "", 

	}
	torrentFile := &torrentFile {
		announce: "",
		info: torrentInfo,
		infoHash: nil,
		pieceHashLength: 20,
	}
	app = &application {
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
			app.printDecodedString(decodedString)
			return decodedString, err
	case "info":
			app.torrentFile.parseTorrentFile(argument)
			app.torrentFile.calculateInfoHash()
			app.torrentFile.printParsedFile()
			return app.torrentFile, nil	
		default:
			fmt.Println("Unknown command: " + command)
			os.Exit(1)
	}
	return "", nil
}
















