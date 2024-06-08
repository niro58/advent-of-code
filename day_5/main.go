package main

import (
	"aoc-day-5/secondPart"
	"fmt"
	"os"
)


func main() {
	input,err := os.ReadFile("inputs\\01_001.txt")
    if err != nil {
        fmt.Println(err)
    }

	res := secondPart.Main(string(input))
	fmt.Println(res)

}