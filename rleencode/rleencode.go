package rleencode

import (
	"strconv"
	"strings"
	"unicode"
)

type runLength struct {
    enc, dec func(string) string
}

func newRunLength() *runLength {
    return &runLength{
        enc: func(input string) string {
            var result strings.Builder
            for len(input) > 0 {
                firstLetter := input[0]
                inputLength := len(input)
                input = strings.TrimLeft(input, string(firstLetter))
                if counter := inputLength - len(input); counter > 1 {
                    result.WriteString(strconv.Itoa(counter))
                }
                result.WriteString(string(firstLetter))
            }
            return result.String()
        },
        dec: func(input string) string {
            var result strings.Builder
            for len(input) > 0 {
                letterIndex := strings.IndexFunc(input, func(r rune) bool {return !unicode.IsDigit(r)})
                multiply := 1
                if letterIndex != 0 {
                    multiply, _ = strconv.Atoi(input[:letterIndex])
                }
                result.WriteString(strings.Repeat(string(input[letterIndex]), multiply))
                input = input[letterIndex+1:]
            }
            return result.String()
        }}
}

func (rl runLength) encode(input string) string {
    return rl.enc(input)
}

func (rl runLength) decode(input string) string {
    return rl.dec(input)
}

//RunLengthEncode ...
func RunLengthEncode(input string) string {
    return newRunLength().encode(input)
}
//RunLengthDecode ...
func RunLengthDecode(input string) string {
    return newRunLength().decode(input)
}