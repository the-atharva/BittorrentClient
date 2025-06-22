package main

import (
	"bytes"
	"io"
	"os"

	bencode "github.com/jackpal/bencode-go"
)

func (app *application) decodeBencode(reader io.Reader) (any, error) {
	decoded, err := bencode.Decode(reader)
	if err != nil {
		app.errorTrace(err)
		return nil, err
	}
	return decoded, nil
}

func (app *application) parseTorrentFile(fileName string) (any, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		app.errorTrace(err)
		return nil, err
	}
	parsedFile, err := app.decodeBencode(bytes.NewReader(file))
	return parsedFile, nil
}
























