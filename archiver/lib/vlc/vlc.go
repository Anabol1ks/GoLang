package vlc

import (
	"strings"
	"unicode"
)

func Encode(str string) []byte {
	str = prepareText(str)
	//encode to binary: some text -> 10010101
	//split binary by chunks (8): bits to bytes -> '10010101 10010101 10001010'
	chunks := splitByChunks(encodeBin(str), chunkSize)
	//bytes to hex -> '20 30 3C'
	//retur hexChunksStr
	return chunks.Bytes()
}

func Decode(encodeData []byte) string {
	// hexchunks -> binary chunks
	//bChunks -> binary string
	bString := NewBinChunks(encodeData).Join()

	//build decoding tree
	dTree := getEncodingTable().DecodingTree()
	//bString(dTree)->text
	return exportText(dTree.Decode(bString))
}

// splitByChunks splits binary string by chunks with given,
// i.e.: '100101011001010110001010' -> 10010101 10010101 10001010'

// encodeBin encodes str into binary codes string without spaces.
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
		panic("unknown character: " + string(ch))
	}

	return res
}

func getEncodingTable() encodingTable {
	return encodingTable{
		' ': "11",
		'e': "101",
		't': "1001",
		'o': "10001",
		'n': "10000",
		'a': "011",
		's': "0101",
		'i': "01001",
		'r': "01000",
		'h': "0011",
		'd': "00101",
		'l': "001001",
		'!': "001000",
		'u': "00011",
		'c': "000101",
		'f': "000100",
		'm': "000011",
		'p': "0000101",
		'g': "0000100",
		'w': "0000011",
		'b': "0000010",
		'y': "0000001",
		'v': "00000001",
		'j': "000000001",
		'k': "0000000001",
		'x': "00000000001",
		'q': "000000000001",
		'z': "000000000000",
	}
}

// prepareText prepares text to be fit for encode:
// changes upper case letters to: !+lower case letter
// i.g.: My name is Grisha -> !my name is !grisha
func prepareText(str string) string {
	var buf strings.Builder

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

// exportText is opposite to prepareText, it prepapres decoded text to export:
// it changes: ! + <lower case letter> -> to upper case letter.
// i.g.: !my name is !grisha -> My name is Grisha
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
