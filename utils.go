package main

import (
	"encoding/binary"
	"bytes"
	"fmt"
	"os"
)

func IntToByte(num int64) []byte {
	var buffer bytes.Buffer
	error := binary.Write(&buffer, binary.BigEndian, num)
	checkError("IntToByte", error)
	return buffer.Bytes()
}
func checkError(pos string, e error) {
	if e != nil {
		fmt.Println("error,pos:", pos, e)
		os.Exit(1)
	}
}
