package table

import "strings"

// decodingTree is binary tree, which contains values on its ends
type decodingTree struct {
	Value string
	Zero  *decodingTree
	One   *decodingTree
}

// Generator interface to create new decoding table from outside
type Generator interface {
	NewTable(text string) EncodingTable
}

// EncodingTable - useful and convinient way to create and operate encoding table
type EncodingTable map[rune]string

func (et EncodingTable) Decode(str string) string {
	dt := et.decodingTree()

	return dt.Decode(str)
}

// decodingTree way to create and fill decong tree
func (et EncodingTable) decodingTree() decodingTree {
	res := decodingTree{}

	for ch, code := range et {
		res.add(code, ch)
	}

	return res
}

// decode all string thru passing each time formed decoding tree
func (dt *decodingTree) Decode(str string) string {
	var buf strings.Builder

	currNode := dt

	for _, ch := range str {
		if currNode.Value != "" {
			buf.WriteString(currNode.Value)
			currNode = dt
		}

		switch ch {
		case '0':
			currNode = currNode.Zero
		case '1':
			currNode = currNode.One
		}
	}
	if currNode.Value != "" {
		buf.WriteString(currNode.Value)
	}

	return buf.String()

}

// method add - simple argorithm to fill tree
func (dt *decodingTree) add(code string, vlaue rune) {
	currNode := dt
	for _, ch := range code {
		switch ch {
		case '0':
			if currNode.Zero == nil {
				currNode.Zero = &decodingTree{}
			}

			currNode = currNode.Zero
		case '1':
			if currNode.One == nil {
				currNode.One = &decodingTree{}
			}

			currNode = currNode.One
		}
	}

	currNode.Value = string(vlaue)
}
