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
func contains(slice []string, item string) bool {
    for _, v := range slice {
        if v == item {
            return true
        }
    }
    return false
}
func isEqual(a []int, b []int) bool{
	if len(a) != len(b){
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func matchingOccs(a []int, toMatch []int) int{
	for i := range toMatch{
		
		if len(a) - 1 < i || a[i] != toMatch[i] {
			return i
		}else if a[i] < toMatch[i]{
			return -1
		}
	}
	if len(a) > len(toMatch){
		return -1
	}
	return len(toMatch)
}
func startsWith(s string, with string) bool{
	for i := range with{
		if s[i] != with[i]{
			return false
		}
	}
	return true
}

func processRow(hash string, validOccs []int) int{
	var closedHashes []string
	var currHashes []string
	currHashes = append(currHashes, hash)
	
	var nextHashes []string
	i := 0


	for{
		if len(currHashes) == 0{
			i += 1
			fmt.Println(i, "/",len(hash), "Len",len(nextHashes))
			currHashes = append([]string{},nextHashes...)
			if len(currHashes) == 0{
				break
			}
			nextHashes = []string{}
		} 

		h := currHashes[0]
		currHashes = currHashes[1:]

		
		occ := calculateOccurences(h)
		position := matchingOccs(occ, validOccs)
		if position == -1{
			continue
		}
		if position == len(validOccs) && !contains(closedHashes,h) {
			closedHashes = append(closedHashes, h)
			continue
		}
		
		if i >= len(hash){
			continue
		}
		nextHashes = append(nextHashes, h)

		if h[i] != '?'{
			continue
		}

		hArr := []rune(h)
		hArr[i] = '#'

		nextHashes = append(nextHashes, string(hArr))
	}
	fmt.Println("==========")
	fmt.Println(hash,validOccs, len(closedHashes))
	// fmt.Println("Results")
	// for _, r := range closedHashes{
	// 	fmt.Println(r)
	// }

	return len(closedHashes)
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
		// fmt.Println("Total for", hash, "Results", resInt)
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