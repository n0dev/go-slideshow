package exif

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Convert a 2-bytes array into an uint16
func toUint16(bo binary.ByteOrder, buf []byte) uint16 {
	var i uint16
	b := bytes.NewReader(buf)
	err := binary.Read(b, bo, &i)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	return i
}

// Convert a 2-bytes array into an int16
func toInt16(bo binary.ByteOrder, buf []byte) int16 {
	var i int16
	b := bytes.NewReader(buf)
	err := binary.Read(b, bo, &i)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	return i
}

// Convert a 4-bytes array into an uint32
func toUint32(bo binary.ByteOrder, buf []byte) uint32 {
	var i uint32
	b := bytes.NewReader(buf)
	err := binary.Read(b, bo, &i)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	return i
}

// Convert a 4-bytes array into an int32
func toInt32(bo binary.ByteOrder, buf []byte) int32 {
	var i int32
	b := bytes.NewReader(buf)
	err := binary.Read(b, bo, &i)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	return i
}
