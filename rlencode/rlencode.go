package rlencode

import "fmt"

type rle struct {
    current int
    num int
    encoded []int
    debug bool
}

func (r *rle) init(args ...interface{}) {
    r.current = -1
    r.num = 0
    r.encoded = []int{}
    r.debug = false
}

func (r *rle) asString() string {
    return fmt.Sprintf("current=%d, num=%d, encoded=%s\n", r.current, r.num, ArrayAsString(r.encoded))
}

func (r *rle) add(val int) {
    if (r.debug) {
        fmt.Printf("(1) %d -> %s", val, r.asString())
    }
    if (r.current == -1) {
        r.current = val
        r.num = 1
        r.encoded = append(r.encoded, val)
        if (r.debug) {
            fmt.Printf("\t(2) -> %s", r.asString())
        }
    } else {
        if (r.current == val) {
            r.num++
            if (r.debug) {
                fmt.Printf("\t(3) -> %s", r.asString())
            }
        } else {
            if (r.num > 1) {
                r.encoded = append(r.encoded, -(r.num-1))
                r.current = val
                r.num = 1
                r.encoded = append(r.encoded, val)
                if (r.debug) {
                    fmt.Printf("\t(4) -> %s", r.asString())
                }
            } else {
                r.current = val
                r.num = 1
                r.encoded = append(r.encoded, val)
                if (r.debug) {
                    fmt.Printf("\t(5) -> %s", r.asString())
                }
            }
        }
    }
}

func (r *rle) flush() {
    if (r.num > 1) {
        r.encoded = append(r.encoded, -(r.num-1))
        if (r.debug) {
            fmt.Printf("\t(4) -> %s", r.asString())
        }
    }
}

type runLength struct {
    enc, dec func([]int) []int
}

//ArrayAsString ...
func ArrayAsString(items []int) string {
	sOut := "{"
	m := len(items) - 1
	for n, i := range items {
		sOut += fmt.Sprintf("%d", i)
		if (n < m) {
			sOut += ", "
		}
	}
	sOut += "}"
	return sOut
}

func newRunLength() *runLength {
    return &runLength{
        enc: func(input []int) []int {
            processor := new(rle)
            processor.init()
            for _, item := range input {
                processor.add(item)
            }
            processor.flush()
            return processor.encoded
        },
        dec: func(input []int) []int {
            var result []int
            for _, i := range input {
                if (i < 0) {
                    lastVal := result[len(result)-1]
                    for j := 0; j < -i; j++ {
                        result = append(result, lastVal)
                    }
                } else {
                    result = append(result, i)
                }
            }
            return result
        }}
}

func (rl runLength) encode(input []int) []int {
    return rl.enc(input)
}

func (rl runLength) decode(input []int) []int {
    return rl.dec(input)
}

//RunLengthEncode ...
func RunLengthEncode(input []int) []int {
    return newRunLength().encode(input)
}
//RunLengthDecode ...
func RunLengthDecode(input []int) []int {
    return newRunLength().decode(input)
}