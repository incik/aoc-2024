package shared

import "os"

func check(e error) {
	if e != nil {
		println("Error")
		panic(e)
	}
}

func invalidArguments() {
	println("Invalid arguments")
	println("Usage: aoc04 <input>")
}

func preflight() {
	var args = os.Args[1:]

	if len(args) < 1 {
		invalidArguments()
		return
	}
}

func GetFileContents() []byte {
	preflight()

	var args = os.Args[1:]

	data, err := os.ReadFile(args[0])
	check(err)

	return data
}
