package main

import (
	"fmt"
)

type torrentInfo struct {
	length int64
	name string
	pieceLength int64 `bencode:"piece length"`
	pieces string 
}

func (ti *torrentInfo) parseInfo(infoMap map[string]any) {
	if l, exists := infoMap["length"]; exists {
		ti.length= l.(int64)
	} else {
		fmt.Println("\nERROR: Length field doesn't exist")
	}
	if n, exists := infoMap["name"]; exists {
		ti.name = n.(string)
	} else {
		fmt.Println("\nERROR: Name field doesn't exist")
	}
	if p, exists := infoMap["piece length"]; exists {
		ti.pieceLength = p.(int64)
	} else {
		fmt.Println("\nERROR: Piece length field doesn't exist")
	}
	if p, exists := infoMap["pieces"]; exists {
		ti.pieces = p.(string)
	} else {
		fmt.Println("\nERROR: Piece length field doesn't exist")
	}
}  

func (ti *torrentInfo) printParsedFile() {
	fmt.Printf("\nLength: %d", ti.length)	
	fmt.Printf("\nName: %s", ti.name)	
	fmt.Printf("\nPiece length: %d", ti.pieceLength)	
	fmt.Printf("\nPieces: % x", ti.pieces)	
}










