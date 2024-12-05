package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)
func isCorrectSequence(a string) bool{
	for _, c := range a {
		if c < '0' || c > '9'{
			return false
		}
	}
	return true
}
func reverse(str string) (result string) { 
    for _, v := range str { 
        result = string(v) + result 
    } 
    return
} 
func secondPart(input string) int{
	var res int
	parts := strings.Split(input, "mul(")

	toIgnore := false
	for _, mul := range parts{
		fmt.Println("------")
		fmt.Println(mul)
		revMul := reverse(mul)
		dontIndex := strings.Index(revMul, ")(t'nod")
		doIndex := strings.Index(revMul, ")(od")
		if dontIndex == -1{
			dontIndex = 99999999999999999
		}
		if doIndex == -1 {
			doIndex = 99999999999999999
		}

		lastIgnore := toIgnore
		
		if dontIndex != doIndex {
			if doIndex < dontIndex{
				toIgnore = false
			}else{
				toIgnore = true
			}
			fmt.Println("Set new ignore", toIgnore)
		}
		if lastIgnore {
			fmt.Println("Skipping")
			continue
		}	
		
		inside := strings.Split(mul, ")")
		if len(inside) < 1 {
			continue
		}

		lrSplitted := strings.Split(inside[0], ",")
		
		if len(lrSplitted) != 2 {
			continue
		}
		l,r := lrSplitted[0], lrSplitted[1]
		if !isCorrectSequence(l + r) {
			continue
		}

		lInt, err := strconv.Atoi(l)
		if err != nil {
			continue
		}
		rInt, err := strconv.Atoi(r)
		if err != nil {
			continue
		}
		fmt.Println("Summing", inside[0], lInt,rInt)
		res += lInt * rInt
	}
	return res
}
func main() {
	start := time.Now()
	input, _ := os.ReadFile(".\\input\\01.txt")
	res := secondPart(string(input))

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