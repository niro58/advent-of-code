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
func duplicateCount(s []int) int {
	m := make(map[int]int)
	for _, v := range s {
		m[v]++
	}
	count := 0
	for _, v := range m {
		if v > 1 {
			count += v - 1
		}
	}
	return count
}
func formatBlinks(blinks map[int]int){
	for k,v := range blinks {
		for range v {
			fmt.Print(k, "  ")
		}
	}
	fmt.Println()
	fmt.Println()
}
const NeededBlinks = 75
func firstPart(stones []int, neededBlinks int) int {
	blinks := make(map[int]int)
	for i := range stones {
		blinks[stones[i]] += 1
	}

	for range neededBlinks{
		nextBlink := make(map[int]int)
		for k, v := range blinks{
			stoneStr := strconv.Itoa(k)

			if len(stoneStr) % 2 == 0{
				left, _ := strconv.Atoi(stoneStr[:len(stoneStr)/2])
				right, _ := strconv.Atoi(stoneStr[len(stoneStr)/2:])
				nextBlink[left] += v
				nextBlink[right] += v
			}else if k == 0 {
				nextBlink[1] += v
			}else{
				nextBlink[k * 2024] += v
			}
			delete(blinks, k)
		}
		blinks = nextBlink
	}
	
	var res int
	for _, v := range blinks{
		res += v
	}
	return res
}

func main() {
	start := time.Now()
	input, _ := os.ReadFile(".\\input\\01.txt")
	stonesStr := strings.Split(string(input), " ")
	// memo := make(map[int][]int)

	// fmt.Println("Memoing")
	// for i := range NeededBlinks{
	// 	memo[i] = firstPart([]int{0}, memo, i)
	// }
	// fmt.Println("Memo'd")
	var totalRes int
	var stones []int
	for i := range stonesStr{
		sInt, _ := strconv.Atoi(stonesStr[i])
		stones = append(stones, sInt)
	}
	totalRes = firstPart(stones, NeededBlinks)

	fmt.Println("Result" , totalRes)

	fmt.Println("Execution time: ", time.Since(start))
}
