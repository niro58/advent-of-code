package firstPart

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
type PlantSize struct{
	Area int
	Perimeter int
}
func includes(arr []int, el int) bool{
	for i := range arr {
		if arr[i] == el{
			return true
		}
	}
	return false
}
func firstPart(gridFlattened string, cols,rows int) int {
	var plants []PlantSize
	var res int

	var closedPoints []int
	
	for index, c := range gridFlattened{
		if includes(closedPoints, index){
			continue
		}
		var plant PlantSize
		var openPoints []int
		openPoints = append(openPoints, index)
		for len(openPoints) > 0 {
			point := openPoints[0]
			closedPoints = append(closedPoints, point)

			openPoints = openPoints[1:]

			moves := []int{
				move(point,1,0,cols,rows),
				move(point,-1,0,cols,rows),
				move(point,0,1,cols,rows),
				move(point,0,-1,cols,rows),
			}
			plant.Area += 1

			for _, m := range moves{
				if m == -1 || rune(gridFlattened[m]) != c {
					plant.Perimeter += 1
					continue
				}
				if !includes(closedPoints, m) && !includes(openPoints, m){
					openPoints = append(openPoints, m)
				}
			}
		}
		plants = append(plants, plant)
	}
	for i := range plants{
		res += plants[i].Area * plants[i].Perimeter
	}

	return res
}

func main() {
	start := time.Now()
	input, _ := os.ReadFile(".\\input\\01.txt")
	inputFlat := strings.Replace(string(input), "\r\n", "", -1)
	rows := strings.Split(string(input), "\r\n")
	colsC := len(rows[0])

	res := firstPart(inputFlat, len(rows),colsC)


	fmt.Println("Result" , res)

	fmt.Println("Execution time: ", time.Since(start))
}
