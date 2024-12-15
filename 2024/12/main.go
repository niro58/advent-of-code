// Result 194482
// Execution time:  19.5141453s

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func firstPart(stones []int, memo map[int][]int) []int {
}

func main() {
	start := time.Now()
	input, _ := os.ReadFile(".\\input\\01.txt")
	stonesStr := strings.Split(string(input), " ")
	// var stones []int
	memo := createMemo([]int{0})
	var totalRes int
	var stones []int
	for i := range stonesStr{
		sInt, _ := strconv.Atoi(stonesStr[i])
		stones = append(stones, sInt)
	}
	stones = firstPart(stones, memo)
	fmt.Println("1")
	stones = firstPart(stones, memo)
	fmt.Println("2")
	stones = firstPart(stones, memo)
	totalRes += len(stones)
	fmt.Println("Memo'd")

	fmt.Println("Result" , totalRes)

	fmt.Println("Execution time: ", time.Since(start))
}
