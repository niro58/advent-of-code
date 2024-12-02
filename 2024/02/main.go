package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)
func abs(n int) int{
	mask := n >> 31
	return (n ^ mask) - mask
}

func removeFromArray(numbers []int, i int) []int{
	return append(numbers[:i], numbers[i+1:]...)
}
func secondPart(numbers []int) bool{
	var arrs [][]int
	for i := range numbers{
		c := append([]int{}, numbers...)
		arrs = append(arrs, removeFromArray(c, i))
	}
	arrs = append(arrs, numbers)
	for _,arr := range arrs {
		dir := arr[0] > arr[1]
		isCorrect := true
		for i := range len(arr) - 1 {
			diff := abs(arr[i] - arr[i + 1])
			if arr[i] > arr[i + 1] != dir || diff < 1 || diff > 3{
				isCorrect = false
				break
			}
		}

		if isCorrect{
			return true
		}
	}
	
	
	return false
}

func main() {
	start := time.Now()
	input, _ := os.ReadFile(".\\input\\01.txt")
	rows := strings.Split(string(input),"\r\n")
	
	res := 0
	for _, row := range rows{
		fmt.Println("--------------------")
		var nums []int
		for _, s:= range strings.Split(row, " "){
			num, _ := strconv.Atoi(s)
			nums = append(nums, num)
		}
		r :=  secondPart(nums)
		fmt.Println(nums, r)
		if r{
			res += 1
		}else{

		}

	}
	fmt.Println(res)
	fmt.Println("Execution time: ", time.Since(start))
}

// func firstPart(numbers []int) bool{
// 	initialDirection := numbers[0] > numbers[1]
// 	for i := range len(numbers) - 1{
// 		dir := numbers[i] > numbers[i+ 1]
// 		diff := abs(numbers[i]  - numbers[i + 1])
// 		if diff < 1 || diff > 3 {
// 			return false
// 		}
// 		if dir != initialDirection{
// 			return false
// 		}

// 	}
	
// 	return true
// }