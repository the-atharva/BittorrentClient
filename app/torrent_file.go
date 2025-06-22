package main

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
	"os"

	bencode "github.com/jackpal/bencode-go"
)

type torrentFile struct {
	announce string
	info map[string]any
	infoHash []byte
}

func (app *application) calculateInfoHash() {
	var buf bytes.Buffer
	err := bencode.Marshal(&buf, app.torrentFile.info)
	if err != nil {
		app.errorTrace(err)
	}
	bencodedInfo := buf.Bytes()
	h := sha1.New()
	h.Write(bencodedInfo)
	app.torrentFile.infoHash = h.Sum(nil)
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
	app.torrentFile.info["length"] = app.torrentFile.info["length"].(int64) 
	app.torrentFile.info["name"] = app.torrentFile.info["name"].(string) 
	app.torrentFile.info["piece length"] = app.torrentFile.info["piece length"].(int64) 
	app.torrentFile.info["pieces"] = []byte(app.torrentFile.info["pieces"].(string)) 
	app.calculateInfoHash()
}
























