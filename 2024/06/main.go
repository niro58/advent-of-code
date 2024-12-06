package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)
func move(start int, x int, y int, cols int, rows int) int {
	newX := start % cols + x
	newY := start / cols + y

	if newX >= cols  || newX < 0|| newY < 0 || newY >= rows {
		return -1
	}
	res := newX + newY * cols

	return res
}
type Vector2 struct {
	X int
	Y int
}


var movementOrder = []Vector2{
	{
		X: 0,
		Y: -1,
	},
	{
		X: 1,
		Y: 0,
	},
	{
		X:0,
		Y: 1,
	},
	{
		X: -1,
		Y: 0,
	},
}
func visualizeMovement(inp string ,walkedPos map[int]bool, cols int){
	runes := []rune(inp)

	for i := range walkedPos {
		runes[i] = '@'
	}

	res := ""
	for i := range runes {
		if i % cols == 0{
			res += "\r\n"
		}
		res += string(runes[i])
	}
	fmt.Println(res)
}
func visualize(inp string, cols int){
	runes := []rune(inp)
	res := ""
	for i := range runes {
		if i % cols == 0{
			res += "\r\n"
		}
		res += string(runes[i])
	}
	fmt.Println(res)
}

func isLooped(inp string, initPos,cols,rows int) bool {
	test := []rune(inp)
	pos := initPos
	dir := 0 
	repeatSteps := make(map[int]int)
	for {
		nextPos := move(pos,movementOrder[dir].X,movementOrder[dir].Y,cols,rows)
		if nextPos == -1{
			return false
		}
		if repeatSteps[nextPos] >= 5{
			return true
		}
		
		if nextPos == initPos && repeatSteps[initPos] > 3 {
			visualize(string(test), cols)
			return true
		}
		repeatSteps[nextPos] += 1

		if inp[nextPos] == '#'{
			dir = (dir + 1) % len(movementOrder)
			continue
		}
		pos = nextPos
	}
}
func firstPart(inp string) int{

	cols := strings.Index(inp,"\r")
	rows := len(strings.Split(inp, "\r\n"))
	inp = strings.Replace(inp, "\r\n", "", -1)
	
	initPos := strings.Index(inp, "^")


	
	var loopPositions []int

	for i := range inp {
		if i == initPos{
			continue
		}
		inpRunes := []rune(inp)
		if inpRunes[i] == '#'{
			continue
		}
		inpRunes[i] = '#'	
		if isLooped(string(inpRunes), initPos, cols,rows){
			fmt.Println("-----",i,"-----")

			loopPositions = append(loopPositions, i)
		}
	}
	// visualizeMovement(inp,walkedPos,cols)
	fmt.Println("=======")

	fmt.Println("Before", len(loopPositions))
	return len(loopPositions)
}
func main() {
	start := time.Now()
	input, _ := os.ReadFile(".\\input\\01.txt")

	
	res := firstPart(string(input))
	fmt.Println(res)

	fmt.Println("Execution time: ", time.Since(start))
}
