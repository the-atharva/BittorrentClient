package main

import (
	"fmt"
	"runtime/debug"
	"strings"

	bencode "github.com/jackpal/bencode-go"
)

func (app *application) decodeBencode(bencodedString string) (interface{}, error) {
	reader := strings.NewReader(bencodedString)
	return bencode.Decode(reader) 
}

func (app *application) errorTrace(err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())	
	app.errLog.Output(2, trace)
}






















