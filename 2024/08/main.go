// 6417120397561
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)
func includes(antinodes []int, i int) bool{
	for j := range antinodes{
		if antinodes[j] == i{
			return true
		}
	}
	return false
}
// func setAntidote(inp *string, i int, antinodes []int) bool {
// 	if i >= len(*inp) || i < 0{
// 		return false
// 	}

// 	r := []rune(*inp)
// 	if r[i] == '#'{
// 		return false
// 	}

// 	r[i] = '#'
// 	*inp = string(r)
// 	return true
// }
func toGrid(inp string, cols int, antinodes []int) {
	inpRunes := []rune(inp)
	for i := range antinodes{
		inpRunes[antinodes[i]] = '#'
	}
	inp = string(inpRunes)
	for i, c := range inp{
		if i % cols == 0{
			fmt.Println()
		}
		fmt.Print(string(c))
	}
	fmt.Println()
}

func indexToGrid(i int, cols int) (x int, y int){
	return i % cols, i / cols
}
func gridToIndex(x int, y int, cols int, rows int) int{
	if x < 0 || x >= cols || y < 0 || y >= rows{
		return -1
	}

	return x + y * cols
}
func firstPart(inp string, cols int, rows int) int{
	var antinodes []int
	positions := make(map[rune][]int, 0)
	for i, c := range inp{
		if c == '.'{
			continue
		}		
		positions[c] = append(positions[c], i)
		
	}
	for _, pos := range positions{
		for i, v1 := range pos {
			if !includes(antinodes, v1){
				antinodes = append(antinodes, v1)
			}
			for j := i + 1; j < len(pos); j++{
				x1,y1:= indexToGrid(v1,cols)
				x2,y2 := indexToGrid(pos[j],cols)
				if !includes(antinodes, pos[j]){
					antinodes = append(antinodes, pos[j])
				}
				xDiff, yDiff := x2-x1, y2-y1

				for {
					isInvalid:= true
					rgx1,rgy1 := x1 - xDiff, y1 - yDiff
					rgx2,rgy2 := x2 + xDiff, y2 + yDiff
					results := []int {
						gridToIndex(rgx1,rgy1,cols, rows),  
						gridToIndex(rgx2,rgy2,cols, rows),
					}
					for _, res := range results{
						if res < 0 || res >= len(inp){
							continue
						}
						
						isInvalid = false
						if !includes(antinodes, res){
							antinodes = append(antinodes, res)
						}
					}
					x1,y1 = rgx1,rgy1
					x2,y2 = rgx2,rgy2
					if isInvalid {
						break
					}
				}

			}
		}
	}
	toGrid(inp, cols, antinodes)

	return len(antinodes)
}
func main() {
	start := time.Now()
	input, _ := os.ReadFile(".\\input\\01.txt")
	
	inputFlat := strings.Replace(string(input), "\r\n", "", -1)

	
	res := firstPart(inputFlat, len(strings.Split(string(input), "\r\n")[0]), len(strings.Split(string(input), "\r\n")))
	fmt.Println("Result" , res)

	fmt.Println("Execution time: ", time.Since(start))
}
