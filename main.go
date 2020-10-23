package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	rle "./rlencode"
)

// TestCase ...
type TestCase struct {
	test []int
	result bool
	actual bool
}

func (t *TestCase) asString() string {
	sOut := ""
	fmt.Printf("%s\n", rle.ArrayAsString(rle.RunLengthEncode(t.test)))
	//sOut += fmt.Sprintf("%s -> %t (%t)\n", rle.ArrayAsString(t.test), t.result, t.actual)
	value := rle.ArrayAsString(t.test)
	runes := []rune(value)
	n := len(value)
	if (n > 30) {
		value = string(runes[0:15]) + "..." + string(runes[n-15:])
	}
	sOut += fmt.Sprintf("%s -> %t (%t)\n", value, t.result, t.actual)
	return sOut
}

func repeatValues(val int, n int) []int {
	var values []int
	for i := 0; i < n; i++ {
		values = append(values, val)
	}
	return values
}

// ===============================================================
func findSplitPoint1(arr []int, n int) int {
    if (len(arr) <= 1) {
        return -1
	}
    leftSum := 0
    for i := 0; i < n; i++ {
        leftSum += arr[i]
		rightSum := 0
		for j := i+1; j < n; j++ {
            rightSum += arr[j]
		}

        if (leftSum == rightSum) {
            return i+1
		} 
	}
    return -1
} 

func sumArray(arr []int, start int, end int) int {
	sum := 0
	for i := start; i < end; i++ {
        sum += arr[i]
	}
	return sum
}

func getNumKeysFromMap(counter map[string]int) []string {
    keys := make([]string, 0, len(counter))
    for k := range counter {
        keys = append(keys, k)
	}
	return keys
}

func countUniquesInArray(arr []int) int {
	counter := make( map[string]int )    
	for _, item := range arr {
		counter[strconv.Itoa(item)]++
		if (len(counter) > 1) {
			return -1
		}
	}
	return 1
} 

func findSplitPoint2(arr []int, n int) int {
	/*
	As you can see, this method that worked much better in Python does not have the same effect here because
	GO does not have an optimized Array summation or Array slicing like Python does.  So there is no optimal
	method other than the original code.  Threading will also not help as demonstrated by the Java version
	because there is a direct dependency between the outer loop and the inner loop.
	*/
	if (n <= 1) {
        return -1
	}
	numUniques := countUniquesInArray(arr)
	if ( (numUniques == 1) && (n > 1) && (arr[0] != 0) && ((n % 2) == 1) )  {
        return -1
	}
    totalSum := sumArray(arr, 0, n)
    if (totalSum == 0) {
		if (n > 1) {
			return 1
		} else {
			return 0
		}
	}
    targetSum := totalSum / 2
	leftSum := 0
	isEQ := false
	for i := 0; i < n; i++ {
        leftSum += arr[i]
        if (leftSum == targetSum) {
			sa := sumArray(arr, i+1, n)
            isEQ = leftSum == sa
            if (isEQ) {
                return i+1
			}
		}
	}
    return -1
}

//================================================================
func canSplitEqually(nums []int) (bool, time.Duration, time.Duration, bool) {
	startTime1 := time.Now()
    splitPoint1 := findSplitPoint1(nums, len(nums))
	elapsedTime1 := time.Since(startTime1)

	startTime2 := time.Now()
    splitPoint2 := findSplitPoint2(nums, len(nums))
	elapsedTime2 := time.Since(startTime2)

	cannotSplit := ( (splitPoint1 == -1) || (splitPoint1 == len(nums)) )
	
	fmt.Printf("(1) splitPoint1 is %d, splitPoint2 is %d\n", splitPoint1, splitPoint2)
	if (splitPoint1 != splitPoint2) {
		fmt.Printf("(2) splitPoint1 is %d, splitPoint2 is %d\n", splitPoint1, splitPoint2)
	}

	return (cannotSplit == false), elapsedTime1, elapsedTime2, (splitPoint1 == splitPoint2)
}
//================================================================

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

	for testNum, t := range testCases {
        bValue, runDuration1, runDuration2, runsMatch := canSplitEqually(t.test)
		//resultOfTest := (bValue == t.result)
		t.actual = (bValue == t.result)
		sResult := "Correct"
		if (!t.actual) {
			sResult = "Incorrect"
		}
		fmt.Printf("%d. %s -> %s (%s vs %s) %t\n", testNum+1, sResult, t.asString(), runDuration1, runDuration2, runsMatch)

		if (!runsMatch) {
			fmt.Printf("FAILURE after %d tests because methods do not agree !!!\n", testNum)
            os.Exit(2)
		}

        if (t.actual == false) {
			fmt.Printf("FAILURE after %d tests because expected result did not match (actual %t vs expected %t) !!!\n", testNum, bValue, t.result)
            os.Exit(1)
		}
	}
	fmt.Println("Done.")
}
