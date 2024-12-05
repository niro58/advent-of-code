package firstPart

import (
	"fmt"
	"os"
	"strings"
	"time"
) 
func indexToGrid(i int, cols int) (x int, y int){
	return i % cols, i / cols
}
func move(start int, x int, y int, cols int, rows int) int {
	newX := start % cols + x
	newY := start / cols + y

	if newX >= cols || newY < 0 || newY >= rows {
		return -1
	}

	return newX + newY * cols
}

type Movable struct {
	Pos int
	DirectionX int
	DirectionY int
	LetterIndex int
}
func isValidWord(i int, cols int, rows int, inp string) int{
	var matches int
	moves := []Movable{
		{
			Pos: move(i, 3,0,cols,rows),
			DirectionX: -1,
			DirectionY: 0,
			LetterIndex: 0,
		},
		{
			Pos: move(i, -3,0,cols,rows),
			DirectionX: 1,
			DirectionY: 0,
			LetterIndex: 0,
		},
		{
			Pos: move(i, 0,-3,cols,rows),
			DirectionX: 0,
			DirectionY: 1,
			LetterIndex: 0,
		},
		{
			Pos: move(i, 0,3,cols,rows),
			DirectionX: 0,
			DirectionY: -1,
			LetterIndex: 0,
		},
		{
			Pos: move(i, 3,3,cols,rows),
			DirectionX: -1,
			DirectionY: -1,
			LetterIndex: 0,
		},
		{
			Pos: move(i, -3,3,cols,rows),
			DirectionX: 1,
			DirectionY: -1,
			LetterIndex: 0,
		},
		{
			Pos: move(i, 3,-3,cols,rows),
			DirectionX: -1,
			DirectionY: 1,
			LetterIndex: 0,
		},
		{
			Pos: move(i, -3,-3,cols,rows),
			DirectionX: 1,
			DirectionY: 1,
			LetterIndex: 0,
		},
	};
	letters := []rune{'S','A','M','X'}
	for len(moves) > 0 {
		mv := moves[len(moves) - 1]
		moves = moves[:len(moves)- 1]
		if mv.Pos < 0{
			continue
		}
		fmt.Println("Index",mv.Pos, "Cell", inp[mv.Pos], "Letter", string(letters[mv.LetterIndex]))

		if rune(inp[mv.Pos]) != letters[mv.LetterIndex] {
			continue
		}
		if mv.LetterIndex == len(letters) - 1{
			matches += 1
			continue
		}
		newPos := move(mv.Pos, mv.DirectionX, mv.DirectionY,cols,rows)
		mv.Pos = newPos
		mv.LetterIndex += 1
		moves = append(moves, mv)
	}
	return matches
}
func firstPart(inp string) int {
	var res int
	cols := strings.Index(inp,"\r")
	rows := len(strings.Split(inp, "\r\n"))
	fmt.Println(inp, cols)

	inp = strings.Replace(inp, "\r\n","",-1)
	for y, row := range strings.Split(inp, "\r\n") {
		for x := range row {
			cell := x + y*cols
			if inp[cell] != 'X'{
				continue
			}
			res += isValidWord(cell, cols,rows, inp)
	
		}
	}
	return res
}
func main() {
	start := time.Now()
	input, _ := os.ReadFile(".\\input\\01.txt")
	res := firstPart(string(input))

	fmt.Println(res)

	fmt.Println("Execution time: ", time.Since(start))
}
