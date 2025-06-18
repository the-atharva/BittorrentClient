package main

import (
	"fmt"
	"runtime/debug"
	"strconv"
	"unicode"

	// bencode "github.com/jackpal/bencode-go"
)

func (app *application) decodeBencode(bencodedString string) (interface{}, error) {
	length := len(bencodedString)
	if unicode.IsDigit(rune(bencodedString[0])) {
		return app.decodeBencodeString(bencodedString, length)
	} else if rune(bencodedString[0]) == 'i' {
		return app.decodeBencodeInt(bencodedString, length)
	} else {
		return "", fmt.Errorf("This format is not supported supported at the moment")
	}
}

func (app *application) decodeBencodeInt(bencodedString string, n int) (interface{}, error) {
	numStr :=bencodedString[1 : n - 1]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		app.errorTrace(err)
		return "", err
	}
	return num, nil
}



func (app *application) decodeBencodeString(bencodedString string, n int) (interface{}, error) {
	var firstColonIndex int
	for i, character := range bencodedString {
		if character == ':' {
			firstColonIndex = i
			break
		}
	}
	lengthStr := bencodedString[:firstColonIndex]
	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		app.errorTrace(err)
		return "", err
	}
	return bencodedString[firstColonIndex + 1 : firstColonIndex + 1 + length], nil
}

func (app *application) errorTrace(err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())	
	app.errLog.Output(2, trace)
}

func (app *application) process(command, argument string) (interface{}, error) {
	switch command {
		case "decode":
			decodedString, err := app.decodeBencode(argument)
			return decodedString, err
		default:
			fmt.Println("Unknown command: " + command)
			os.Exit(1)
	}
	return "", nil
}




















