package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		println("Error")
		panic(e)
	}
}

func invalidArguments() {
	println("Invalid arguments")
	println("Usage: aoc02 <input>")
}

func copySlice(s []int) []int {
	newSlice := make([]int, len(s)) // allocate memory for new slice
	copy(newSlice, s)
	return newSlice
}

func removeElement(s []int, i int) []int {
	return append(s[:i], s[i+1:]...) // Go doesn't have the remove, you have to slice the array before and after and glue it together
}

func checkReport(report []int) bool {
	var ascending = false
	var descending = false
	var stagnating = false
	var outOfBounds = false

	// compare neighbors, detect the trend and set the flags
	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]

		if diff < 0 {
			descending = true

			if diff < -3 {
				outOfBounds = true
				break
			}

		} else if diff == 0 { // no change is "unsafe"
			stagnating = true
		} else {
			ascending = true

			if diff > 3 {
				outOfBounds = true
				break
			}

		}

		// if both "ascending" and "descending" are set, there was a flip of directions
		if (ascending && descending) || stagnating || outOfBounds {
			break
		}
	}

	if (ascending && descending) || stagnating || outOfBounds {
		return false
	} else {
		return true
	}
}

func recheck(numbers []int) bool {
	valid := false

	for index := 0; index < len(numbers); index++ {
		var fixed []int

		if index+1 >= len(numbers) {
			fixed = numbers[:index]
		} else {
			// slices work by references, so removing an item from `numbers` in loop would create nonsenses
			// hence the copy of the slice is necessary
			fixed = removeElement(copySlice(numbers), index)
		}

		valid = checkReport(fixed)

		if valid {
			break
		}
	}

	return valid
}

func main() {
	var args = os.Args[1:]

	if len(args) < 1 {
		invalidArguments()
		return
	}

	if len(args) == 1 {
		data, err := os.ReadFile(args[0])
		check(err)

		reports := strings.Split(string(data), "\n")

		var safeCount = 0
		var dampenedSafeCount = 0

		for _, report := range reports {
			if report == "" {
				break
			}

			var numbers []int

			for _, item := range strings.Split(report, " ") {
				num, _ := strconv.Atoi(item)
				numbers = append(numbers, num)
			}

			valid := checkReport(numbers)

			if valid {
				safeCount += 1
			} else if recheck(numbers) {
				dampenedSafeCount += 1
			}
		}

		fmt.Println("Total undampened safes: ", safeCount)
		fmt.Println("Total dampened safes: ", safeCount+dampenedSafeCount)
	}
}
