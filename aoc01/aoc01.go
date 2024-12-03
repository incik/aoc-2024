package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Println("Error")
		panic(e)
	}
}

func sumList(list []int) int {
	var sum = 0
	for i := 0; i < len(list); i++ {
		sum += list[i]
	}
	return sum
}

func getLists(data []byte) ([]int, []int) {
	var firstList []int
	var secondList []int

	for _, line := range strings.Split(string(data), "\n") {
		if line == "" {
			break
		}

		splited := strings.Split(line, "   ")
		prvni, _ := strconv.Atoi(splited[0])
		firstList = append(firstList, prvni)
		druhy, _ := strconv.Atoi(splited[1])
		secondList = append(secondList, druhy)
	}

	return firstList, secondList
}

func distaces(filename string) {
	data, err := os.ReadFile(filename)
	check(err)

	firstList, secondList := getLists(data)

	slices.Sort(firstList)
	slices.Sort(secondList)

	var distances []int

	for i := 0; i < len(firstList); i++ {
		distance := firstList[i] - secondList[i]
		if distance < 0 {
			distance = -distance
		}
		distances = append(distances, distance)
	}

	sum := sumList(distances)

	fmt.Println(sum)
}

func similarity(filename string) {
	data, err := os.ReadFile(filename)
	check(err)

	firstList, secondList := getLists(data)

	var hashmap = map[int]int{}

	for i := 0; i < len(firstList); i++ {
		if hashmap[firstList[i]] != 0 { // did we already count this number?
			continue
		}

		for j := 0; j < len(secondList); j++ {
			if firstList[i] == secondList[j] {
				hashmap[firstList[i]] += 1
			}
		}
	}

	//	fmt.Println(hashmap)

	var similarities []int

	// go through the first list and multiply the number of occurences of the number in the second list
	for i := 0; i < len(firstList); i++ {
		similarities = append(similarities, firstList[i]*hashmap[firstList[i]])
	}

	// 	fmt.Println(similarities)

	sum := sumList(similarities)
	fmt.Println(sum)
}

func invalidArguments() {
	fmt.Println("Invalid arguments")
	fmt.Println("Usage: aoc01 [distaces|similarity] <input_file>")
}

func main() {
	var args = os.Args[1:]

	if len(args) < 2 {
		invalidArguments()
		return
	}

	if len(args) == 2 {
		if args[0] == "distances" {
			distaces(args[1])
		} else if args[0] == "similarity" {
			similarity(args[1])
		} else {
			invalidArguments()
		}
	}
}
