package firstPart

import (
	"fmt"
	"os"
	"strconv"
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
func includes(a []int, b int)bool{
	for i := range a {
		if a[i] == b{
			return true
		}
	}
	return false
}
func findSolutions(grid[]int, index,height,cols,rows int, results *[]int){
	if height == 9 {
		if !includes(*results, index){
			*results = append(*results, index)
		}
		return
	}
	moves := []int{
		move(index,1,0,cols,rows),
		move(index,-1,0,cols,rows),
		move(index,0,1,cols,rows),
		move(index,0,-1,cols,rows),
	}
	for i := range moves{
		if moves[i] == -1 {
			continue
		}
		if grid[moves[i]] == height + 1{
			findSolutions(grid, moves[i], height+1,cols,rows,results)
		}
	}
}
func firstPart(grid []int, cols,rows int) int {
	var res int
	for i := range grid{
		if grid[i] == 0 {
			var sols []int

			findSolutions(grid,i,0,cols,rows,&sols)

			if len(sols) > 0 {
				res += len(sols)
			}
		}
	}
	return res
}
func main() {
	start := time.Now()
	input, _ := os.ReadFile(".\\input\\01.txt")
	inputFlat := strings.Replace(string(input), "\r\n", "", -1)

	var inputInt []int
	for i := range inputFlat{
		if inputFlat[i] == '.'{
			inputInt = append(inputInt, 99)
			continue
		}
		n, _ := strconv.Atoi(string(inputFlat[i]))
		inputInt = append(inputInt, n)
	}
	rows := len(strings.Split(string(input), "\r\n"))
	cols := len(strings.Split(string(input), "\r\n")[0])
	res := firstPart(inputInt, cols,rows)
	fmt.Println("Result" , res)

	fmt.Println("Execution time: ", time.Since(start))
}
