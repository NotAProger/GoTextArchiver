package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type BinaryChunks []BinaryChunk

type BinaryChunk string

// if you want to change this viriable
// change NewBinChunk fmt.Sprintf("%08b", code), where 08 - new variable
const chunksSize = 8

// NewBinChunks splits add data by BinaryChunks, with size of variable chunksSize
func NewBinChunks(data []byte) BinaryChunks {

	res := make(BinaryChunks, 0, len(data))

	for _, code := range data {
		res = append(res, NewBinChunk(code))
	}

	return res
}

// NewBinChunk creates unified BinaryChunk with given size
func NewBinChunk(code byte) BinaryChunk {
	return BinaryChunk(fmt.Sprintf("%08b", code))
}

// method BinaryChunks takes them and makes out of it one string
func (bcs BinaryChunks) ToString() string {
	var buf strings.Builder

	for _, bc := range bcs {
		buf.WriteString(string(bc))
	}

	return buf.String()
}

// method Bytes takes BinaryChunks and crates one big slice []byte
func (bcs BinaryChunks) Bytes() []byte {
	res := make([]byte, 0, len(bcs))

	for _, bc := range bcs {
		res = append(res, bc.Byte())
	}

	return res
}

// method Byte takes BinaryChunk and crates byte
func (bc BinaryChunk) Byte() byte {
	bum, err := strconv.ParseUint(string(bc), 2, chunksSize)
	if err != nil {
		panic("can't parse binary chunk: " + err.Error())
	}
	return byte(bum)
}

// splitByChuncks splits binary string by chunks with given size
//
//	i.g.: '101011101001011010101110' -> '10101110 10010110 10101110'
func splitByChuncks(bStr string, chunkSize int) BinaryChunks {
	strLen := utf8.RuneCountInString(bStr)
	chunkCount := strLen / chunkSize // dont using len() bcs with non eng letters apear troubles
	if strLen%chunkSize != 0 {
		chunkCount++
	}

	res := make(BinaryChunks, 0, chunkCount)

	var buf strings.Builder

	for _, ch := range bStr {
		buf.WriteString(string(ch))
		if buf.Len() == chunksSize {
			res = append(res, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}

	if buf.Len() != 0 {
		lastChunk := buf.String()

		lastChunk += strings.Repeat("0", chunkSize-len(lastChunk))

		res = append(res, BinaryChunk(lastChunk))
	}
	return res
}
