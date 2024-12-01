package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var nonAlphaRegex = regexp.MustCompile(`[^0-9]+`)
func clearString(str string) string {
	return nonAlphaRegex.ReplaceAllString(str, "")
}
func calculateFirstPart(input string) int {
	lines := strings.Split(input, "\r\n")
	sum := 0
	for _, line := range lines{
		cleanLine := clearString(line)
		sum += int(cleanLine[0] - '0') * 10 + int(cleanLine[len(cleanLine) - 1] - '0')
	}
	return sum
}

func calculateSecondPart(input string) int {
	lines := strings.Split(input, "\r\n")
	sum := 0
	for _, line := range lines{
		cleanLine := clearString(line)
		sum += int(cleanLine[0] - '0') * 10 + int(cleanLine[len(cleanLine) - 1] - '0')
	}
	return sum
}
func main() {
	input,err := os.ReadFile("input.txt")
    if err != nil {
        fmt.Println(err)
    }

	res := calculateFirstPart(string(input))
	fmt.Println(res)

}