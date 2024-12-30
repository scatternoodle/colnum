package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
)

const appName = "colnum"
const pattern = `^[a-zA-Z]{1,2}$`

func usageInfo() string {
	return "Usage: " + appName + " COLUMN_STRING\n"
}

func helpInfo() string {
	return usageInfo() +
		"Takes a spreadsheet column indicator and returns its column number. " +
		"Takes Excel, Google Sheets and ODS style indicators, e.g. A=1, Z=26, AA=27. " +
		"The returned column number is 1-indexed." +
		"\n\n" +
		"-h --help           - Display help information.\n" +
		"-z --zero-index     - Return a zero-indexed column number (1-indexing is the default).\n"
}

func argFail(reason string) {
	fmt.Println("invalid argument. " + reason)
	fmt.Println(usageInfo())
}

func alphaPos(c rune) int {
	if c >= 'A' && c <= 'Z' {
		return int(c - 'A' + 1)
	}
	if c >= 'a' && c <= 'z' {
		return int(c - 'a' + 1)
	}
	return -1
}

func colNum(s string) (int, error) {
	reg := regexp.MustCompile(pattern)
	if !reg.MatchString(s) {
		return 0, errors.New("invalid column indicator, must match pattern " + pattern)
	}

	pos1 := alphaPos(rune(s[0]))
	if pos1 == -1 {
		return 0, fmt.Errorf("unable to resolve alpha position for indicator position 1: %b", s[0])
	}
	if len(s) == 1 {
		return pos1, nil
	}

	pos1 *= 26
	pos2 := alphaPos(rune(s[1]))
	if pos2 == -1 {
		return 0, fmt.Errorf("unable to resolve alpha position for indicator position 1: %b", s[0])
	}
	return pos1 + pos2, nil
}

func main() {
	args := os.Args
	if len(args) < 2 {
		argFail("Not enough arguments.")
		os.Exit(1)
	}
	if args[1] == "-h" || args[1] == "--help" {
		fmt.Println(helpInfo())
		os.Exit(0)
	}

	var zeroIndex bool
	argPos := 1
	if args[1] == "-z" || args[1] == "--zero-index" {
		argPos++
		zeroIndex = true
	}
	if zeroIndex && len(args) < 3 {
		argFail("Not enough arguments.")
		os.Exit(1)
	}

	num, err := colNum(args[argPos])
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}
	if zeroIndex {
		num--
	}
	fmt.Println(num)
}
