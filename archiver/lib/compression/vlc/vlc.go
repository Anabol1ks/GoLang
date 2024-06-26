package vlc

import (
	"archiver/lib/compression/vlc/table"
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"log"
	"strings"
	"unicode"
)

type EncoderDecoder struct {
	tblGenerator table.Generator
}

func New(tblGenerator table.Generator) EncoderDecoder {
	return EncoderDecoder{tblGenerator: tblGenerator}
}

func (ed EncoderDecoder) Encode(str string) []byte {
	tbl := ed.tblGenerator.NewTable(str)
	encoded := encodeBin(str, tbl)
	//encode to binary: some text -> 10010101
	//split binary by chunks (8): bits to bytes -> '10010101 10010101 10001010'
	//bytes to hex -> '20 30 3C'
	//retur hexChunksStr
	return buildEncodedFile(tbl, encoded)
}

func buildEncodedFile(tbl table.EncodingTable, data string) []byte {
	encodedTbl := encodeTable(tbl)

	var buf bytes.Buffer

	buf.Write(encodeInt(len(encodedTbl)))
	buf.Write(encodeInt(len(data)))
	buf.Write(encodedTbl)
	buf.Write(splitByChunks(data, chunkSize).Bytes())

	return buf.Bytes()
}

func (ed EncoderDecoder) Decode(encodedData []byte) string {
	tbl, data := parseFile(encodedData)

	return tbl.Decode(data)
}

func parseFile(data []byte) (table.EncodingTable, string) {
	const (
		tableSizeBinaryCount = 4
		dataSizeBinaryCount  = 4
	)
	tableSizeBinary, data := data[:tableSizeBinaryCount], data[tableSizeBinaryCount:]
	dataSizeBinary, data := data[:dataSizeBinaryCount], data[dataSizeBinaryCount:]

	tableSize := binary.BigEndian.Uint32(tableSizeBinary)
	dataSize := binary.BigEndian.Uint32(dataSizeBinary)

	tblBinary, data := data[:tableSize], data[tableSize:]

	tbl := decodeTable(tblBinary)

	body := NewBinChunks(data).Join()

	return tbl, body[:dataSize]

}

func encodeInt(num int) []byte {
	res := make([]byte, 4)
	binary.BigEndian.PutUint32(res, uint32(num))

	return res
}

func decodeTable(tblBinary []byte) table.EncodingTable {
	var tbl table.EncodingTable

	r := bytes.NewReader(tblBinary)
	if err := gob.NewDecoder(r).Decode(&tbl); err != nil {
		log.Fatal("can't dacode table: ", err)
	}

	return tbl
}

func encodeTable(tbl table.EncodingTable) []byte {
	var tableBuf bytes.Buffer

	if err := gob.NewEncoder(&tableBuf).Encode(tbl); err != nil {
		log.Fatal("can't serialize table: ", err)
	}

	return tableBuf.Bytes()
}

// splitByChunks splits binary string by chunks with given,
// i.e.: '100101011001010110001010' -> 10010101 10010101 10001010'

// encodeBin encodes str into binary codes string without spaces.
func encodeBin(str string, table table.EncodingTable) string {
	var buf strings.Builder

	for _, ch := range str {
		buf.WriteString(bin(ch, table))
	}

	return buf.String()
}

func bin(ch rune, table table.EncodingTable) string {
	res, ok := table[ch]
	if !ok {
		panic("unknown character: " + string(ch))
	}

	return res
}

// func getEncodingTable() table.EncodingTable {
// 	return table.EncodingTable{
// 		' ': "11",
// 		'e': "101",
// 		't': "1001",
// 		'o': "10001",
// 		'n': "10000",
// 		'a': "011",
// 		's': "0101",
// 		'i': "01001",
// 		'r': "01000",
// 		'h': "0011",
// 		'd': "00101",
// 		'l': "001001",
// 		'!': "001000",
// 		'u': "00011",
// 		'c': "000101",
// 		'f': "000100",
// 		'm': "000011",
// 		'p': "0000101",
// 		'g': "0000100",
// 		'w': "0000011",
// 		'b': "0000010",
// 		'y': "0000001",
// 		'v': "00000001",
// 		'j': "000000001",
// 		'k': "0000000001",
// 		'x': "00000000001",
// 		'q': "000000000001",
// 		'z': "000000000000",
// 	}
// }

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
