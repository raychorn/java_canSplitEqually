package rlencode

import (
	"testing"
)

func repeatValues(val int, n int) []int {
	var values []int
	for i := 0; i < n; i++ {
		values = append(values, val)
	}
	return values
}

func equal(a, b []int) bool {
    if len(a) != len(b) {
		//fmt.Printf("equal.1 -> false\n")
        return false
    }
    for i, v := range a {
        if v != b[i] {
			//fmt.Printf("equal.2 -> false\n")
            return false
        }
    }
	//fmt.Printf("equal.3 -> true\n")
    return true
}

//TestOne ...
func TestOne(t *testing.T) {
	tables := []struct {
		testCase []int 
		expected []int
	}{
		{[]int{3, 1, 1, 2, 1}, []int{3, 1, -1, 2, 1}},
		{[]int{4, 1, 1, 2, 1},[]int{4, 1, -1, 2, 1}},
		{[]int{8, 8},[]int{8, -1}},
		{[]int{1},[]int{1}},
		{[]int{5, 1, 1, 1, 1, 1},[]int{5, 1, -4}},
		{[]int{5, 1, 1, 1, 1, 1, 1},[]int{5, 1, -5}},
		{[]int{1, 1, 1, 1, 4},[]int{1, -3, 4}},
		{[]int{1, 1, 1, 1, 1, 1, 5},[]int{1, -5, 5}},
		{[]int{1, 0},[]int{1, 0}},
		{[]int{0, 1, 0},[]int{0, 1, 0}},
		{[]int{1, 1, 1},[]int{1, -2}},
		{[]int{0},[]int{0}},
		{[]int{0, 0},[]int{0, -1}},
		{[]int{0, 0, 0},[]int{0, -2}},
		{[]int{},[]int{}},
		{[]int{8, 8, 8},[]int{8, -2}},
		{repeatValues(0, 100),[]int{0, -99}},
		{repeatValues(0, 1000),[]int{0, -999}},
		{repeatValues(0, 10000),[]int{0, -9999}},
		{repeatValues(1, 101),[]int{1, -100}},
		/*
		*/
	}

	count := 0
	successCount := 0
	for _, table := range tables {
		value := RunLengthEncode(table.testCase)
		if !equal(value, table.expected) {
			//t.Errorf("RLE %s -> %s was incorrect, expected: %s.", ArrayAsString(table.testCase), ArrayAsString(value), ArrayAsString(table.expected))
			t.Errorf("RLE %s -> %s was incorrect, expected: %s.", ArrayAsString(table.testCase), ArrayAsString(value), ArrayAsString(table.expected))
		} else {
			t.Logf("RLE %s was correct, expected: %s.", ArrayAsString(value), ArrayAsString(table.expected))
			successCount++
		}
		count++
	}
	t.Logf("Ran %d tests with %d successes.\n", count, successCount)
}
