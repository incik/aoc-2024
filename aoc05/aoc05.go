package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"adventofcode/shared"
)

func parseData(data []byte) ([]string, []string) {
	parts := strings.Split(string(data), "\n\n")
	rulesRaw, pagesRaw := parts[0], parts[1]
	rules := strings.Split(rulesRaw, "\n")
	pages := strings.Split(pagesRaw, "\n")
	return rules, pages
}

func parseRules(rules []string) map[string][]string {
	// create a map of all the numbers each number should be followed with
	rulesMap := make(map[string][]string)
	for _, rule := range rules {
		parts := strings.Split(rule, "|")
		rulesMap[parts[0]] = append(rulesMap[parts[0]], parts[1])
	}
	return rulesMap
}

func testOrder(order string, rules map[string][]string) bool {
	pages := strings.Split(order, ",")
	var goods []string

	for i, page := range pages {
		for _, good := range goods {
			if slices.Contains(rules[page], good) {
				// fmt.Println(good, "was already here!")
				return false
			}
		}

		if len(rules[page]) == 0 && i != len(pages)-1 {
			// fmt.Println(page, " has no followers and isn't last!")
			return false
		}

		//fmt.Println(page, " => ", rules[page])
		goods = append(goods, page)
	}

	return true
}

func fixOrder(order string, rules map[string][]string) string {
	pages := strings.Split(order, ",")

	followers := make(map[string]int)

	for _, page := range pages {
		// fmt.Println("rules for ", page, " => ", rules[page])

		followers[page] = 0
		for i := 0; i < len(pages); i++ {
			if slices.Contains(rules[page], pages[i]) {
				followers[page] += 1
			}
		}
	}

	var fixed []string

	// screw custom sorting, let's use brute force
	// fmt.Println(followers)
	for i := len(pages) - 1; i >= 0; i-- {
		for _, page := range pages {
			if followers[page] == i {
				fixed = append(fixed, page)
			}
		}
	}

	return strings.Join(fixed, ",")
}

func getMiddleNumber(order string) int {
	pages := strings.Split(order, ",")
	num, _ := strconv.Atoi(pages[len(pages)/2])
	return num
}

func main() {
	data := shared.GetFileContents()
	rules, orders := parseData(data)

	ruleMap := parseRules(rules)

	sum := 0
	failSum := 0
	for _, order := range orders {
		if testOrder(order, ruleMap) {
			sum += getMiddleNumber(order)
		} else {
			fixed := fixOrder(order, ruleMap)
			failSum += getMiddleNumber(fixed)
			continue
		}
	}

	fmt.Println("sum: ", sum)
	fmt.Println("failSum: ", failSum)
}
