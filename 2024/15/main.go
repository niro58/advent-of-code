// 1514522 too low
// 1521453 sheesh
// 2.7067ms before cleanup
// 1.6ms before cleanup

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
var EMPTY_RUNE = '.'
var OBSTACLE_LEFT = '['
var OBSTACLE_RIGHT = ']'

var MOVE_TO_VECTOR = map[rune]Vector2{
	'>': {
		X: 1,
		Y: 0,
	},
	'^': {
		X: 0,
		Y: -1,
	}, 'v': {
		X: 0,
		Y: 1,
	}, '<': {
		X: -1,
		Y: 0,
	},
}

type Vector2 struct {
	X int
	Y int
}

func (v Vector2) AddBasic(v2 Vector2) Vector2 {
	return Vector2{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}

type Robot struct {
	Position Vector2
}

func IsObstacle(c rune) bool {
	return c == OBSTACLE_LEFT || c == OBSTACLE_RIGHT
}
func (el Vector2) IncludedIn(arr []Vector2) bool {
	for _, v := range arr {
		if v == el {
			return true
		}
	}
	return false
}
func (r *Robot) MoveY(v Vector2) {
	obstacleToOppositeVector := map[rune]Vector2{
		OBSTACLE_LEFT: {
			X: 1,
			Y: 0,
		},
		OBSTACLE_RIGHT: {
			X: -1,
			Y: 0,
		},
	}

	var open []Vector2
	var walked []Vector2

	open = append(open, r.Position.AddBasic(v))
	open = append(open, open[0].AddBasic(obstacleToOppositeVector[grid[open[0].Y][open[0].X]]))

	for len(open) > 0 {
		curr := open[0]
		open = open[1:]
		if !curr.IncludedIn(walked) {
			walked = append(walked, curr)
		}

		nextPos := curr.AddBasic(v)

		c := grid[nextPos.Y][nextPos.X]
		if !IsObstacle(c) {
			continue
		}

		open = append(open, nextPos, nextPos.AddBasic(obstacleToOppositeVector[c]))
	}

	for _, pos := range walked {
		if grid[pos.Y+v.Y][pos.X] == WALL_RUNE {
			return
		}
	}

	for i := len(walked) - 1; i >= 0; i-- {
		cell := walked[i]
		c := grid[cell.Y][cell.X]

		grid[cell.Y][cell.X] = '.'
		grid[cell.Y+v.Y][cell.X] = c
	}
	r.Position = r.Position.AddBasic(v)

}
func (r *Robot) MoveX(v Vector2) {
	nextPos := r.Position.AddBasic(v)

	posAfterNext := nextPos.AddBasic(v)
	for IsObstacle(grid[posAfterNext.Y][posAfterNext.X]) {
		posAfterNext = posAfterNext.AddBasic(v)
	}

	if grid[posAfterNext.Y][posAfterNext.X] == WALL_RUNE {
		return
	}

	grid[nextPos.Y][nextPos.X] = '.'

	distance := nextPos.X - posAfterNext.X
	if distance < 0 {
		distance *= -1
	}

	var c rune
	if v.X < 0 {
		c = OBSTACLE_LEFT
	} else if v.X > 0 {
		c = OBSTACLE_RIGHT
	}
	for i := 0; i < distance; i++ {
		grid[nextPos.Y][posAfterNext.X+(i*v.X*-1)] = c

		if c == OBSTACLE_LEFT {
			c = OBSTACLE_RIGHT
		} else {
			c = OBSTACLE_LEFT
		}
	}

	r.Position = nextPos
}
func (r *Robot) Move(v Vector2) {
	nextPos := r.Position.AddBasic(v)
	curr := grid[nextPos.Y][nextPos.X]
	if curr == WALL_RUNE {
		return
	}
	if !IsObstacle(curr) {
		r.Position = nextPos
		return
	}

	if v.X != 0 {
		r.MoveX(v)
	}
	if v.Y != 0 {
		r.MoveY(v)
	}
}
func GetGrid(input string) (grid [][]rune, robotPosition Vector2, lastRow int) {
	RUNE_TO_GRID := map[rune]string{
		WALL_RUNE:     string([]rune{WALL_RUNE, WALL_RUNE}),
		OBSTACLE_RUNE: string([]rune{OBSTACLE_LEFT, OBSTACLE_RIGHT}),
		PLAYER_RUNE:   string([]rune{EMPTY_RUNE, EMPTY_RUNE}),
		EMPTY_RUNE:    string([]rune{EMPTY_RUNE, EMPTY_RUNE}),
	}

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
func PreviewGrid() {
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
		robot.Move(MOVE_TO_VECTOR[move])
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
	input, err := os.ReadFile("./input/01.txt")
	if err != nil {
		panic(err)
	}
	res := Calculate(string(input))
	fmt.Println("Result", res)
	fmt.Println("Execution time: ", time.Since(start))
}
