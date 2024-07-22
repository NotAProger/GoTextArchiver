package vlc

import (
	"archiver/lib/compression/vlc/table"
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"log"
	"strings"
)

type EncoderDecoder struct {
	tblGenerator table.Generator
}

func New(tblGenerator table.Generator) EncoderDecoder {
	return EncoderDecoder{tblGenerator: tblGenerator}
}

// Encode encoding data based chosen algorithm
func (ed EncoderDecoder) Encode(str string) []byte {
	tbl := ed.tblGenerator.NewTable(str)

	encoded := encodeBin(str, tbl)

	return buildEncodedFile(tbl, encoded)
}

// Decode decoding data based on binary file
func (ed EncoderDecoder) Decode(encodedData []byte) string {
	tbl, data := parseFile(encodedData)

	return tbl.Decode(data)
}

// parseFile disassembles binary file to its parts, to pull out
// encoding table and encoded data
func parseFile(data []byte) (table.EncodingTable, string) {
	const (
		tableSizeBytesCount = 4
		dataSizeBytesCount  = 4
	)
	tableSizeBinary, data := data[:tableSizeBytesCount], data[tableSizeBytesCount:]
	dataSizeBinary, data := data[:dataSizeBytesCount], data[dataSizeBytesCount:]

	tableSize := binary.BigEndian.Uint32(tableSizeBinary)
	dataSize := binary.BigEndian.Uint32(dataSizeBinary)

	tblBinary, data := data[:tableSize], data[tableSize:]

	tbl := decodeTable(tblBinary)

	body := NewBinChunks(data).ToString()

	return tbl, body[:dataSize]
}

// buildEncodedFile makes binary file, which contains len of table (uint32),
// len of containig data (uint32), then binary encoded table and binary data
func buildEncodedFile(tbl table.EncodingTable, data string) []byte {
	encodedTbl := encodeTable(tbl)

	var buf bytes.Buffer

	buf.Write((encodeInt((len(encodedTbl)))))
	buf.Write(encodeInt(len(data)))
	buf.Write(encodedTbl)
	buf.Write(splitByChuncks(data, chunksSize).Bytes())

	return buf.Bytes()
}

// encodeInt encode int to binary
func encodeInt(num int) []byte {
	res := make([]byte, 4)
	binary.BigEndian.PutUint32(res, uint32(num))

	return res
}

// decodeTable decode binary table using default go library
func decodeTable(tblBinary []byte) table.EncodingTable {
	var tbl table.EncodingTable

	r := bytes.NewReader(tblBinary)
	if err := gob.NewDecoder(r).Decode(&tbl); err != nil {
		log.Fatal("cant't decode table: ", err)
	}
	return tbl
}

// encodeTable encode to biary table using default go library
func encodeTable(tbl table.EncodingTable) []byte {
	var tableBuf bytes.Buffer

	err := gob.NewEncoder(&tableBuf).Encode((tbl))
	if err != nil {
		log.Fatal("cant't serialize table: ", err)
	}

	return tableBuf.Bytes()
}

// encodeBin encodes str into binary codes sting without spaces.
func encodeBin(str string, table table.EncodingTable) string {
	var buf strings.Builder

	for _, ch := range str {
		buf.WriteString(bin(ch, table))
	}

	return buf.String()
}

// bin takes rune and returns encoded binary code
func bin(ch rune, table table.EncodingTable) string {

	res, ok := table[ch]
	if !ok {
		panic("unknown character: \"" + string(ch) + "\"")
	}
	return res
}
