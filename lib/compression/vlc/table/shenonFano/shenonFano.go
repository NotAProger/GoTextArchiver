package shenonFano

import (
	"archiver/lib/compression/vlc/table"
	"fmt"
	"math"
	"sort"
	"strings"
)

// gust convinient abstaction
type Generator struct{}

// for initialization from outside
func NewGenerator() Generator {
	return Generator{}
}

type encodingTable map[rune]code

type code struct {
	Char     rune   // letter in the code
	Quantity int    // how many times this char has been used in text
	Bits     uint32 // code which will have char, now its in decimal, i will be converted to binary
	Size     int    // support stat, wich will help to understand how long shoud code 'Bits' be
}

// how many times in text letter have been used
type charStat map[rune]int

// NewTable generates new encoding table based on incoming text
func (g Generator) NewTable(text string) table.EncodingTable {
	stat := newCharStat(text)

	return build(stat).Export()
}

// Export exports encodingTable like a map[rune]string
// throug compressing table in needed format
func (et encodingTable) Export() map[rune]string {
	res := make(map[rune]string)

	for k, v := range et {
		byteStr := fmt.Sprintf("%b", v.Bits)

		if lenDiff := v.Size - len(byteStr); lenDiff != 0 {
			byteStr = strings.Repeat("0", lenDiff) + byteStr
		}
		res[k] = byteStr
	}
	return res
}

// build func is taking charStat and makes out of it encodingTable
// in the process filling out information, which is needed to correclty
// build this table
func build(stat charStat) encodingTable {
	codes := make([]code, 0, len(stat))

	for ch, qty := range stat {
		codes = append(codes, code{
			Char:     ch,
			Quantity: qty,
		})
	}

	// sorting slice of codes from max uses of char to less usees of char
	sort.Slice(codes, func(i, j int) bool {
		if codes[i].Quantity != codes[j].Quantity {
			return codes[i].Quantity > codes[j].Quantity
		}

		return codes[i].Char < codes[j].Char
	})

	assignCodes(codes)

	res := make(encodingTable)

	for _, code := range codes {
		res[code.Char] = code
	}
	return res
}

var isFirstLoop bool = true

// assignCodes assigns codes to its chars
func assignCodes(codes []code) {
	if len(codes) < 2 {
		if isFirstLoop {
			codes[0].Size++
			codes[0].Bits = 0
		}
		return
	}
	isFirstLoop = false
	divider := bestDividerPosition(codes)

	for i := 0; i < len(codes); i++ {
		codes[i].Bits <<= 1
		codes[i].Size++

		if i >= divider {
			codes[i].Bits |= 1
		}
	}
	assignCodes(codes[:divider])
	assignCodes(codes[divider:])
	isFirstLoop = true
}

// besDividerPosition choses best devider position
//
//	returns int position, which devides binary tree
//
// depends on evenly used in text both devided sides
func bestDividerPosition(codes []code) int {
	total := 0
	for _, code := range codes {
		total += code.Quantity
	}

	left := 0
	pevDiff := math.MaxInt
	bestPosition := 0

	for i := 0; i < len(codes)-1; i++ {
		left += codes[0].Quantity
		right := total - left

		diff := abs(right - left)
		if diff >= pevDiff {
			break
		}

		pevDiff = diff
		bestPosition = i + 1
	}

	return bestPosition
}

// module of the int
func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

// newCharStat sorts through text rune by rune, filling out charStat
func newCharStat(text string) charStat {
	res := make(charStat)

	for _, ch := range text {
		res[ch]++
	}

	return res
}
