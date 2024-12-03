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

		for _, report := range reports {
			if report == "" {
				break
			}

			var ascending = false
			var descending = false
			var outOfBounds = false

			var numbers []int

			for _, item := range strings.Split(report, " ") {
				num, _ := strconv.Atoi(item)
				numbers = append(numbers, num)
			}

			// compare neighbors, detect the trend and set the flags
			for i := 0; i < len(numbers)-1; i++ {
				diff := numbers[i+1] - numbers[i]

				if diff < 0 {
					descending = true

					if diff < -3 {
						outOfBounds = true
					}
				} else if diff == 0 { // no change is "unsafe"
					outOfBounds = true
				} else {
					ascending = true

					if diff > 3 {
						outOfBounds = true
					}
				}
			}

			// if both "ascending" and "descending" are set, there was a flip of directions
			if (ascending && descending) || outOfBounds {
				// fmt.Println("Unsafe!")
			} else {
				// fmt.Println("Safe :)")
				safeCount += 1
			}
		}

		fmt.Println("Total safes: ", safeCount)

	}
}
