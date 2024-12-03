package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
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

	data, err := os.ReadFile(args[0])
	check(err)

	re := regexp.MustCompile(`do(n\'t)?\(\)|mul\(\d+,\d+\)`)
	instructions := re.FindAll(data, -1)

	re2 := regexp.MustCompile(`\d+`)

	var products []int

	enabled := true

	for _, instr := range instructions {
		strInst := string(instr)
		if strInst == "don't()" {
			enabled = false
			continue
		}
		if strInst == "do()" {
			enabled = true
			continue
		} else if enabled {
			strNums := re2.FindAll(instr, -1)
			num1, _ := strconv.Atoi(string(strNums[0]))
			num2, _ := strconv.Atoi(string(strNums[1]))
			products = append(products, num1*num2)
		}
	}

	sum := 0

	for _, product := range products {
		sum += product
	}

	fmt.Println(sum)
}
