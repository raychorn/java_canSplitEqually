package rlencode

import "fmt"

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
            var result []int
            n := len(input)
            for i := 0; i < n; i++ {
                val := input[i]
                //fmt.Printf("(1) val=%d\n", val)
                count := 1
                j := i+1
                jj := j
                for ; j < n; j++ {
                    val2 := input[j]
                    //fmt.Printf("(2) val2=%d\n", val2)
                    if (val == val2) {
                        count++
                        //fmt.Printf("(3) count=%d\n", count)
                        continue
                    } else {
                        jj = j
                        //fmt.Printf("(4) jj=%d\n", jj)
                        break
                    }
                }
                if (count > 2) {
                    result = append(result, val)
                    result = append(result, -count)
                    i = jj-1
                    //fmt.Printf("(5) i=%d, %s\n", i, ArrayAsString(result))
                } else {
                    result = append(result, val)
                    //fmt.Printf("(6) i=%d, %s\n", i, ArrayAsString(result))
                }
                //fmt.Println("")
            }
            //fmt.Println("======================================")
            return result
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