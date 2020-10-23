package main

import (
	"fmt"
)

// TestCase ...
type TestCase struct {
	test []int
	result bool
}

func arrayAsString(items []int) string {
	sOut := "{"
	m := len(items) - 1
	lastValue := -1
	lastValueRun := 0
	if (m > -1) {
		lastValue = items[0]
		lastValueRun = 0
	}
	for n, i := range items {
		if (i == lastValue) {
			lastValueRun++
			continue
		}
		if (lastValueRun > 1) {
			sOut += fmt.Sprintf("(%d * %d)", i, lastValueRun)
			lastValueRun = 0
		} else {
			sOut += fmt.Sprintf("%d", i)
		}
		if (n < m) {
			sOut += ", "
		}
		lastValue = i
	}
	if (lastValueRun > 0) {
		sOut += fmt.Sprintf("(%d * %d)", lastValue, lastValueRun)
	}
	sOut += "}"
	return sOut
}

func (t *TestCase) asString() string {
	sOut := fmt.Sprintf("%s -> %t", arrayAsString(t.test), t.result)
	return sOut
}

func repeatValues(val int, n int) []int {
	var values []int
	for i := 0; i < n; i++ {
		values = append(values, val)
	}
	return values
}

func main() {
	testCases := []TestCase { 
        {
            test: []int{3, 1, 1, 2, 1}, 
            result: true,
        },
        {
            test: []int{4, 1, 1, 2, 1}, 
            result: false, 
        },
        {
            test: []int{8, 8}, 
            result: true, 
        },
        {
            test: []int{1}, 
            result: false, 
        },
        {
            test: []int{5, 1, 1, 1, 1, 1}, 
            result: true, 
        },
        {
            test: []int{5, 1, 1, 1, 1, 1, 1}, 
            result: false, 
        },
        {
            test: []int{1, 1, 1, 1, 4}, 
            result: true, 
        },
        {
            test: []int{1, 1, 1, 1, 1, 1, 5}, 
            result: false, 
        },
        {
            test: []int{1, 0}, 
            result: false, 
        },
        {
            test: []int{0, 1, 0}, 
            result: false, 
        },
        {
            test: []int{1, 1, 1}, 
            result: false, 
        },
        {
            test: []int{0}, 
            result: false, 
        },
        {
            test: []int{0, 0}, 
            result: true, 
        },
        {
            test: []int{0, 0, 0}, 
            result: true, 
        },
        {
            test: []int{}, 
            result: false, 
        },
        {
            test: []int{8, 8, 8}, 
            result: false, 
        },
        {
            test: repeatValues(0, 100),
            result: true, 
        },
        {
            test: repeatValues(0, 1000),
            result: true, 
        },
        {
            test: repeatValues(0, 10000),
            result: true, 
        },
        {
            test: repeatValues(1, 101),
            result: false, 
        },
    }

	for _, t := range testCases {
		fmt.Println(t.asString())
	}
	fmt.Println("Done.")
}
