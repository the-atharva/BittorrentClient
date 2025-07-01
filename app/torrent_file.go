package main

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"os"
	"strings"

	bencode "github.com/jackpal/bencode-go"
)

type torrentFile struct {
	announce string
	info torrentInfo
	infoHash []byte
}

func (tf *torrentFile) calculateInfoHash() {
	var buf bytes.Buffer
	err := bencode.Marshal(&buf, tf.info)
	if err != nil {
		app.errorTrace(err)
	}
	bencodedInfo := buf.Bytes()
	temp, _ := app.decodeBencode(strings.NewReader(string(bencodedInfo)))
	app.printDecodedString(temp)
	h := sha1.New()
	h.Write(bencodedInfo)
	tf.infoHash = h.Sum(nil)
}

func (tf *torrentFile) parseTorrentFile(fileName string) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		app.errorTrace(err)
	}
	parsedFile, err := app.decodeBencode(bytes.NewReader(file))
	if fileMap, ok := parsedFile.(map[string]any); ok {
		if trackerURL, exists := fileMap["announce"]; exists {
			tf.announce = trackerURL.(string)
		} else {
			fmt.Println("ERROR: Announce key doesn't exist")		
			os.Exit(1)
		}
		if infoMap, ok := fileMap["info"].(map[string]any); ok {
			tf.info.parseInfo(infoMap)
		}	else {
			fmt.Println("Can't convert decoded info file to map")
			os.Exit(1)
		}
	} else {
		fmt.Println("Can't convert decoded torrent file to map")
		os.Exit(1)
	}
}

func (tf torrentFile) printParsedFile() {
	fmt.Printf("\nAnnounce: %s", tf.announce)
	tf.info.printParsedFile()
	fmt.Printf("\nInfo Hash: % x", tf.infoHash)
}
























