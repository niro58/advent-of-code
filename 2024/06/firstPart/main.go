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
func visualize(inp string ,walkedPos map[int]bool, cols int){
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
func firstPart(inp string) int{

	cols := strings.Index(inp,"\r")
	rows := len(strings.Split(inp, "\r\n"))
	inp = strings.Replace(inp, "\r\n", "", -1)
	
	pos := strings.Index(inp, "^")
	direction := 0
	
	walkedPos := make(map[int]bool)
	for{
		nextPos := move(pos,movementOrder[direction].X,movementOrder[direction].Y,cols,rows)
		if nextPos < 0 {
			break			
		}

		if inp[nextPos] == '#'{
			direction = (direction + 1) % len(movementOrder)
			continue
		}
		walkedPos[pos] = true
		pos = nextPos
	}
	walkedPos[pos] = true
	visualize(inp,walkedPos, cols)
	return len(walkedPos)
}
func main() {
	start := time.Now()
	input, _ := os.ReadFile(".\\input\\01.txt")

	
	res := firstPart(string(input))
	fmt.Println(res)

	fmt.Println("Execution time: ", time.Since(start))
}
