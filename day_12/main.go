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
type Hash struct{
	Hash string
	Index int
}
func getResult(hash string, validOccs []int ) string {
	var q []Hash
	q = append(q, Hash{
		Hash: hash,
		Index:0,
	})
	
	var res string
	for res == ""{
		last := len(q) - 1
		h := q[last]
		q = q[:last]

		var maxOcc int
		for i := range validOccs{
			maxOcc = max(maxOcc, validOccs[i])
		}
	
		occ := calculateOccurences(h.Hash)
		position := matchingOccs(occ, validOccs, maxOcc)
		if position == -1 || !canFinish(validOccs,position, len(hash) - h.Index) {
			continue
		}

		if position == len(validOccs) {
			res = h.Hash
			continue
		}
		
		if h.Index == len(h.Hash) {
			continue
		}

		q = append(q, Hash {
			Hash:h.Hash,
			Index: h.Index + 1,
		})
	
		if hash[h.Index] != '?'{
			continue
		}
	
		hArr := []rune(h.Hash)
		hArr[h.Index] = '#'
		
		q = append(q, Hash{
			Hash: string(hArr),
			Index: h.Index + 1,
		})
	}
	return res
}
func isLocked(i int, lockedPositions []int) bool{
	for j := range lockedPositions{
		if lockedPositions[j] == i{
			return true
		}
	}
	return false
}
type OccMovement struct{
	Left int
	Right int
	Distance int
	Locked []int

	LeftOffset int
	RightOffset int
}

func CreateMovements(result string, hash string) []OccMovement{
	var i int
	var movements []OccMovement 

	for i < len(result){
		if(result[i] != '#'){
			i += 1
			continue
		}
		var movement OccMovement
		
		//Left
		movement.Left = i

		//Right and locked positions
		movement.Right = i
		for movement.Right < len(result) && result[movement.Right] == '#'{
			if hash[movement.Right] == '#'{
				movement.Locked = append(movement.Locked, movement.Right)
			}
			movement.Right += 1
		}
		movement.Right -= 1

		//Distance
		movement.Distance = movement.Right - movement.Left + 1
		
		
		j := movement.Left - 1
		for j >= 0{
			if result[j] != '?' || (j - 1 >= 0 && result[j - 1] == '#') {
				break;
			}
			movement.LeftOffset += 1
			j -= 1 
		}
		j = movement.Right + 1
		for j < len(result){
			if result[j] != '?' || (j + 1 < len(result) && result[j + 1] == '#') {
				break;
			}
			movement.RightOffset += 1
			j += 1
		}

		movements = append(movements, movement)
		i = movement.Right + 1
	}
	return movements
}
func (m OccMovement) Debug(){
	fmt.Println("--------------")
	fmt.Println("Left",m.Left)
	fmt.Println("Right",m.Right)
	fmt.Println("Distance",m.Distance)
	fmt.Println("Locked",m.Locked)
	fmt.Println("LeftOffset",m.LeftOffset)
	fmt.Println("RightOffset",m.RightOffset)
}
func getTotalSolutions(movements []OccMovement, result int) int{
	if len(movements) == 0 {
		return result
	}

	movement := movements[0]
	var currRes int
	
	return getTotalSolutions(movements, result * currRes)
}
func processRow(hash string, validOccs []int) int{
	res := getResult(hash, validOccs)
	occMovements := CreateMovements(res, hash)
	fmt.Println("First result", res)
	for _,j :=range occMovements{
		j.Debug()
	}
	var distances []int

	
	var totalSolutions int = 1
	//Solution
	for _, move := range distances{
		totalSolutions *= move + 1
	}
	fmt.Println(hash, totalSolutions)
	return totalSolutions
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
		resInt := processRow(hash, occs)

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