package main

import (
	"advent-of-code-day-9/secondPart"
	"fmt"
	"os"
)


func main() {
	input,err := os.ReadFile("inputs\\02_002.txt")
    if err != nil {
        fmt.Println(err)
    }

	res := secondPart.Main(string(input))
	fmt.Println(res)

}