package main

import (
	"fmt"
	"strconv"
	
	"unicode"
)

func decodeBencode(bencodedString string) (interface{}, error) {
	length := len(bencodedString)
	if unicode.IsDigit(rune(bencodedString[0])) {
		return decodeBencodeString(bencodedString, length)
	} else if rune(bencodedString[0]) == 'i' {
		return decodeBencodeInt(bencodedString, length)
	} else {
		return "", fmt.Errorf("This format is not supported supported at the moment")
	}
}

func decodeBencodeString(bencodedString string, n int) (interface{}, error) {
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
		return "", err
	}
	return bencodedString[firstColonIndex + 1 : firstColonIndex + 1 + length], nil
}

func decodeBencodeInt(bencodedString string, n int) (interface{}, error) {
	numStr :=bencodedString[1 : n - 1]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return "", err
	}
	return num, nil
}





















