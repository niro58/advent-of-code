package main

import (
	"aoc-day-5/firstPart"
	"fmt"
	"os"
)


func main() {
	input,err := os.ReadFile("inputs\\01_002.txt")
    if err != nil {
        fmt.Println(err)
    }

	res := firstPart.Main(string(input))
	fmt.Println(res)

}