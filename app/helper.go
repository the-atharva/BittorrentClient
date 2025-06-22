package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime/debug"

	bencode "github.com/jackpal/bencode-go"
)

func (app *application) decodeBencode(reader io.Reader) (interface{}, error) {
	decoded, err := bencode.Decode(reader)
	if err != nil {
		app.errorTrace(err)
		return nil, err
	}
	return decoded, nil
}

func (app *application) errorTrace(err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())	
	app.errLog.Output(2, trace)
	os.Exit(1)
}

func (app *application) parseTorrentFile(fileName string) (interface{}, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		app.errorTrace(err)
		return nil, err
	}
	return app.decodeBencode(bytes.NewReader(file))
}
























