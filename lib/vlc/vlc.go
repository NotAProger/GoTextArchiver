package vlc

import (
	"strings"
	"unicode"
)

func Encode(str string) []byte {
	str = prepareText(str)

	bStr := encodeBin(str)

	chunks := splitByChuncks(bStr, chunksSize)

	return chunks.Bytes()
}

func Decode(encodedData []byte) string {
	bString := NewBinChunks(encodedData).ToString()

	// build decoding tree
	// bString (dTree) -> text
	dTree := getEncodingTable().DecodingTree()
	// bString (dTree) -> text

	// return decoded text
	return exportText(dTree.Decode(bString))
}

// encodeBin encodes str into binary codes sting without spaces.
//
//	i.g. hello -> 0100100001000101010011000100110001001111
func encodeBin(str string) string {
	var buf strings.Builder

	for _, ch := range str {
		buf.WriteString(bin(ch))
	}

	return buf.String()
}

func bin(ch rune) string {
	table := getEncodingTable()

	res, ok := table[ch]
	if !ok {
		panic("unknown character: \"" + string(ch) + "\"")
	}
	return res
}

func getEncodingTable() encodingTable {
	return encodingTable{
		' ': "11",
		't': "1001",
		'n': "10000",
		's': "0101",
		'r': "01000",
		'd': "00101",
		'!': "001000",
		'c': "000101",
		'm': "000011",
		'g': "0000100",
		'b': "0000010",
		'v': "00000001",
		'k': "0000000001",
		'q': "000000000001",
		'e': "101",
		'o': "10001",
		'a': "011",
		'i': "01001",
		'h': "0011",
		'l': "001001",
		'u': "00011",
		'f': "000100",
		'p': "0000101",
		'w': "0000011",
		'y': "0000001",
		'j': "000000001",
		'x': "00000000001",
		'z': "000000000000",
	}
}

// prepareText prepares text to be fit for encode:
// canges uppe case latters to: !+lower case letter
//
//	i.g.: My name is Vlad -> !my name is !vlad
func prepareText(str string) string {
	var buf strings.Builder // used because strings immutable and it will be very expensive in memory and processor time

	for _, ch := range str {
		if unicode.IsUpper(ch) {
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(ch))
		} else {
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}

// exportText prepares text to export to be writen in file
//
// i.g. !my name is !vlad -> My name is Vlad
func exportText(str string) string {
	var buf strings.Builder

	var isCapital bool
	for _, ch := range str {
		if isCapital {
			buf.WriteRune(unicode.ToUpper(ch))
			isCapital = false

			continue
		}
		if ch == '!' {
			isCapital = true

			continue
		} else {
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}