// 1514522 too low
// 1521453 sheesh
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var grid [][]rune
var robot Robot

var WALL_RUNE = '#'
var OBSTACLE_RUNE = 'O'
var PLAYER_RUNE = '@'
var RUNE_TO_GRID = map[rune]string{
	WALL_RUNE:     "##",
	OBSTACLE_RUNE: "[]",
	PLAYER_RUNE:   "..",
	'.':           "..",
}

type Vector2 struct {
	X int
	Y int
}

func (v Vector2) AddBasic(v2 Vector2) Vector2 {
	res := Vector2{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
	if res.X < 0 || res.Y < 0 {
		fmt.Println("yo")
	}
	return Vector2{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}

type Robot struct {
	Position Vector2
}

func IsObstacle(c rune) bool {
	return c == rune(RUNE_TO_GRID[OBSTACLE_RUNE][0]) || c == rune(RUNE_TO_GRID[OBSTACLE_RUNE][1])
}
func Includes(arr []Vector2, el Vector2) bool {
	for _, v := range arr {
		if v == el {
			return true
		}
	}
	return false
}
func (r *Robot) MoveY(v Vector2) {
	oppositeSideMap := map[rune]Vector2{
		'[': {
			X: 1,
			Y: 0,
		},
		']': {
			X: -1,
			Y: 0,
		},
	}

	var open []Vector2
	var walked []Vector2

	open = append(open, r.Position.AddBasic(v))
	open = append(open, open[0].AddBasic(oppositeSideMap[grid[open[0].Y][open[0].X]]))
	//rethink
	for len(open) > 0 {
		curr := open[0]
		open = open[1:]
		if !Includes(walked, curr) {
			walked = append(walked, curr)
		}
		nextPos := curr.AddBasic(v)
		c := grid[nextPos.Y][nextPos.X]
		if c != '[' && c != ']' {
			continue
		}
		open = append(open, nextPos, nextPos.AddBasic(oppositeSideMap[c]))
	}
	canMove := true
	for i := len(walked) - 1; i >= 0; i-- {
		if grid[walked[i].Y+v.Y][walked[i].X] == WALL_RUNE {
			canMove = false
		}
	}
	if canMove {
		for i := len(walked) - 1; i >= 0; i-- {
			c := grid[walked[i].Y][walked[i].X]

			grid[walked[i].Y][walked[i].X] = '.'
			grid[walked[i].Y+v.Y][walked[i].X] = c
		}
		r.Position = r.Position.AddBasic(v)
	}

}
func (r *Robot) MoveX(v Vector2) {
	newPos := r.Position.AddBasic(v)
	posAfter := newPos.AddBasic(v)
	for IsObstacle(grid[posAfter.Y][posAfter.X]) {
		posAfter = posAfter.AddBasic(v)
	}

	if grid[posAfter.Y][posAfter.X] == '.' {
		grid[newPos.Y][newPos.X] = '.'

		diff := newPos.X - posAfter.X
		if diff < 0 {
			diff *= -1
		}
		for i := 0; i < diff; i++ {
			if i%2 == 0 {
				if v.X == 1 {
					grid[newPos.Y][posAfter.X+(i*v.X*-1)] = ']'
				} else {
					grid[newPos.Y][posAfter.X+(i*v.X*-1)] = '['
				}
			} else {
				if v.X == -1 {
					grid[newPos.Y][posAfter.X+(i*v.X*-1)] = ']'
				} else {
					grid[newPos.Y][posAfter.X+(i*v.X*-1)] = '['
				}
			}
		}
	} else {
		return
	}

	r.Position = newPos
}
func (r *Robot) Move(v Vector2) {
	newPos := r.Position.AddBasic(v)
	curr := grid[newPos.Y][newPos.X]
	if curr == WALL_RUNE || newPos.X < 0 {
		return
	}
	if !IsObstacle(curr) {
		r.Position = newPos
		return
	}

	if v.X != 0 {
		r.MoveX(v)
	} else {
		r.MoveY(v)
	}
}
func GetGrid(input string) (grid [][]rune, robotPosition Vector2, lastRow int) {
	for y, row := range strings.Split(input, "\r\n") {
		grid = append(grid, []rune{})
		if len(row) == 0 {
			lastRow = y
			break
		}
		var rowRunes string
		for _, cell := range row {
			if cell == PLAYER_RUNE {
				robotPosition = Vector2{
					X: len(rowRunes),
					Y: y,
				}
			}
			rowRunes += RUNE_TO_GRID[cell]

		}
		grid[y] = []rune(rowRunes)
	}
	return grid, robotPosition, lastRow
}
func GetMoves(input string, lastRow int) []rune {
	var moves []rune
	for _, row := range strings.Split(input, "\r\n")[lastRow:] {
		moves = append(moves, []rune(row)...)
	}
	return moves
}
func PreviewGrid(robot Robot) {
	for y, row := range grid {
		for x, cell := range row {
			if robot.Position.Y == y && robot.Position.X == x {
				fmt.Print(string(PLAYER_RUNE))
			} else {
				fmt.Print(string(cell))
			}
		}
		fmt.Println()
	}
}
func Calculate(input string) int {
	g, robotPosition, lastRow := GetGrid(input)
	moves := GetMoves(input, lastRow)

	grid = g
	robot = Robot{
		Position: robotPosition,
	}
	for len(moves) > 0 {

		move := moves[0]
		moves = moves[1:]

		var v Vector2
		if move == '>' {
			v.X = 1
		} else if move == '^' {
			v.Y = -1
		} else if move == '<' {
			v.X = -1
		} else if move == 'v' {
			v.Y = 1
		} else {
			panic(move)
		}

		// fmt.Println(string(move))
		robot.Move(v)
		// PreviewGrid(robot)
		// fmt.Println()
	}

	var res int
	for y, row := range grid {
		for x, cell := range row {
			if cell == '[' {
				res += y*100 + x
			}
		}
	}
	return res
}
func main() {
	start := time.Now()
	input, err := os.ReadFile("./input/02.txt")
	if err != nil {
		panic(err)
	}
	res := Calculate(string(input))
	fmt.Println("Result", res)
	fmt.Println("Execution time: ", time.Since(start))
}
