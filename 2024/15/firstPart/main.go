package firstPart

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

func (r *Robot) Move(v Vector2) {
	newPos := r.Position.AddBasic(v)
	if grid[newPos.Y][newPos.X] == WALL_RUNE {
		return
	}
	if grid[newPos.Y][newPos.X] == OBSTACLE_RUNE {
		posAfter := newPos.AddBasic(v)
		for grid[posAfter.Y][posAfter.X] == 'O' {
			posAfter = posAfter.AddBasic(v)
		}
		if grid[posAfter.Y][posAfter.X] == '.' {
			grid[newPos.Y][newPos.X] = '.'
			grid[posAfter.Y][posAfter.X] = 'O'
		} else {
			return
		}
	}
	r.Position = newPos
}
func GetGrid(input string) (grid [][]rune, robotPosition Vector2, lastRow int) {
	for y, row := range strings.Split(input, "\r\n") {
		grid = append(grid, []rune{})
		fmt.Println(len(row))
		if len(row) == 0 {
			lastRow = y
			break
		}
		grid[y] = make([]rune, len(row))
		for x, cell := range row {
			grid[y][x] = cell
			if cell == PLAYER_RUNE {
				grid[y][x] = '.'
			}

			if cell == PLAYER_RUNE {
				robotPosition = Vector2{
					X: x,
					Y: y,
				}
			}
		}
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
	PreviewGrid(robot)
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
		robot.Move(v)
	}

	var res int
	for y, row := range grid {
		for x, cell := range row {
			if cell == OBSTACLE_RUNE {
				res += y*100 + x
			}
		}
	}
	return res
}
func main() {
	start := time.Now()
	input, err := os.ReadFile("./input/03.txt")
	if err != nil {
		panic(err)
	}
	res := Calculate(string(input))
	fmt.Println("Result", res)
	fmt.Println("Execution time: ", time.Since(start))
}
