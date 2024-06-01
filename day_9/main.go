package main

import (
	"advent-of-code-day-9/firstPart"
	"fmt"
	"os"
)


func main() {
	input,err := os.ReadFile("inputs\\02_001.txt")
    if err != nil {
        fmt.Println(err)
    }

	res := firstPart.Main(string(input))
	fmt.Println(res)

}