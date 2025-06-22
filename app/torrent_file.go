package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	bencode "github.com/jackpal/bencode-go"
)

type torrentFile struct {
	announce string
	info map[string]any
}

func (app *application) calcInfoHash(parsedFile any) ([]byte, error) {
	return nil, nil
}

func (app *application) decodeBencode(reader io.Reader) (any, error) {
	decoded, err := bencode.Decode(reader)
	if err != nil {
		app.errorTrace(err)
		return nil, err
	}
	return decoded, nil
}

func (app *application) parseTorrentFile(fileName string) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		app.errorTrace(err)
	}
	parsedFile, err := app.decodeBencode(bytes.NewReader(file))
	if fileMap, ok := parsedFile.(map[string]any); ok {
		if trackerURL, exists := fileMap["announce"]; exists {
			app.torrentFile.announce = trackerURL.(string)
		} else {
			fmt.Println("ERROR: Announce key doesn't exist")		
			os.Exit(1)
		}
		if info, exists := fileMap["info"]; exists {
			infoMap := info.(map[string]any)
			keysToCheck := [4]string{"length", "name", "piece length", "pieces"}
			for _, key := range keysToCheck {
				if value, exists := infoMap[key]; exists {
					app.torrentFile.info[key] = value
				} else {
					fmt.Printf("\nERROR: Key %s doesn't exist", key)
					os.Exit(1)
				}
			}
		} else {
			fmt.Println("ERROR: Info doesn't exist")
			os.Exit(1)
		}
	} else {
		fmt.Println("Can't convert decoded torrent file to map")
		os.Exit(1)
	}
}
























