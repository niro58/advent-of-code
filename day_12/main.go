package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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
type HashWithIndex struct{
	Hash string
	Index int
}
func processRow(hash string, validOccs []int) int{
	var closedHashes []string
	openHashes := []HashWithIndex {{
		Hash: hash,
		Index: 0,
	}}
	for len(openHashes) > 0{
		
		h := openHashes[0]
		openHashes = openHashes[1:]

		
		occ := calculateOccurences(h.Hash)
		position := matchingOccs(occ, validOccs)
		
		if position == len(validOccs) && !contains(closedHashes,h.Hash) {
			closedHashes = append(closedHashes, h.Hash)
			continue
		}
		
		if h.Index >= len(hash){
			continue
		}
		openHashes = append(openHashes, HashWithIndex{
			Hash:h.Hash,
			Index: h.Index + 1,
		})

		if h.Hash[h.Index] != '?'{
			continue
		}

		hArr := []rune(h.Hash)
		hArr[h.Index] = '#'

		openHashes = append(openHashes, HashWithIndex{
			Hash:string(hArr),
			Index: h.Index + 1,
		})
	}
	fmt.Println("==========")
	fmt.Println(hash,validOccs, len(closedHashes))
	fmt.Println("Results")
	for _, r := range closedHashes{
		fmt.Println(r)
	}

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
		resInt := processRow(hash, occs)

		total += resInt
		// fmt.Println("Total for", hash, "Results", resInt)
	}
	return total
}
func main() {
	input, _ := os.ReadFile(".\\input\\1.txt")

	res:= firstPart(string(input))
	// output, _ := os.ReadFile(".\\output\\1.txt")
	// outputInt, _ := strconv.Atoi(string(output))
	fmt.Println("Result is ", res)
	// if (res != outputInt){
	// 	 panic("Test failed")
	// }
	fmt.Println("POG")
}