package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)
func calculateOccurences(hash string) []int{
	hash += "."

	var occurences []int
	currCount := 0
	for _,  v := range hash{
		if v == '?' || v == '.'{
			if(currCount > 0){
				occurences = append(occurences, currCount)
			}
			currCount = 0
		}else{
			currCount += 1
		}
	}
	return occurences
}
func canFinish(a []int, position int, left int) bool{
	sum := -1
	for i:= len(a) - 1 ;i > position ; i--{
		sum += a[i] + 1
	}

	return left >= sum
}
func matchingOccs(a []int,toMatch []int, maxToMatch int) int{

	for i := range toMatch{
		if toMatch[i] > maxToMatch {
			return -1
		}else if len(a) - 1 < i || a[i] != toMatch[i]{
			return i
		}
	}
	
	if len(a) > len(toMatch){
		return -1
	}
	return len(toMatch)
}
func processMemo(hash string, i int, validOccs []int, res *int, totalIter *int) {
	*totalIter += 1

	var maxOcc int
	for i := range validOccs{
		maxOcc = max(maxOcc, validOccs[i])
	}

	occ := calculateOccurences(hash)
	position := matchingOccs(occ, validOccs, maxOcc)
	if position == -1 {
		return
	}else if  !canFinish(validOccs,position, len(hash) - i){
		fmt.Println("Hash",hash,"Position",position,"Valid occs",validOccs)
		return
	}
	if position == len(validOccs) {
		*res += 1
		return
	}
	
	if i == len(hash) {
		return
	}

	processMemo(hash, i + 1, validOccs, res, totalIter)

	if hash[i] != '?'{
		return
	}

	hArr := []rune(hash)
	hArr[i] = '#'

	processMemo(string(hArr), i + 1, validOccs, res,totalIter)
}

func processRow(hash string, validOccs []int) int{
	
	var res int
	//iter for testing
	var iter int

	processMemo(hash, 0, validOccs, &res, &iter)

	fmt.Println("!!!!!!!!")
	fmt.Println(hash)
	fmt.Println(res,iter)	
	fmt.Println("---------")
	return res
}
func firstPart(input string) int{
	rows := strings.Split(input, "\r\n")
	total := 0
	for _, row := range rows {
		parts := strings.Split(row, " ")
		
		hash := parts[0]
		
		var occs []int
		for _, occ := range strings.Split(parts[1], ","){
			i,_ := strconv.Atoi(occ)
			occs = append(occs, i)
		}

		var fOccs []int 
		var fHashes string
		
		for i := range 5 {
			fOccs = append(fOccs, occs...)
			if i == 4 {
				fHashes += hash
			}else{
				fHashes += hash + "?"
			}
		}
		resInt := processRow(fHashes, fOccs)

		total += resInt
	}
	return total
}
func main() {
	start := time.Now()
	

	input, _ := os.ReadFile(".\\input\\1.txt")

	res:= firstPart(string(input))
	// output, _ := os.ReadFile(".\\output\\1.txt")
	// outputInt, _ := strconv.Atoi(string(output))
	fmt.Println("Result is ", res)
	// if (res != outputInt){
	// 	 panic("Test failed")
	// }
	fmt.Println("POG")
	fmt.Println("Execution time: ", time.Since(start))
}