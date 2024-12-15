package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)
func firstPart(stones []int) int {
	for i := range 25 {
		var newStones []int
		for len(stones) > 0{
			stone := stones[0]
			stones = stones[1:]
			stoneStr := strconv.Itoa(stone)
			
			if len(stoneStr) % 2 == 0  {
				left, _ := strconv.Atoi(stoneStr[:len(stoneStr)/2])
				right, _ := strconv.Atoi(stoneStr[len(stoneStr)/2:])
				newStones = append(newStones,left)
				newStones = append(newStones,right)
		
			}else if stone == 0 {
				newStones = append(newStones, 1)
			}else{
				newStones = append(newStones, stone * 2024)
			}
		}
		// fmt.Println("------------")
		fmt.Println(i, len(newStones))
		stones = newStones
	}
	return len(stones)
}
func main() {
	start := time.Now()
	input, _ := os.ReadFile(".\\input\\01.txt")
	stonesStr := strings.Split(string(input), " ")
	var stonesInt []int
	for i := range stonesStr{
		sInt, _ := strconv.Atoi(stonesStr[i])
		stonesInt = append(stonesInt, sInt)
	}
	res := firstPart(stonesInt)
	fmt.Println("Result" , res)

	fmt.Println("Execution time: ", time.Since(start))
}
