package firstPart

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var grid Grid

type Vector2 struct {
	X int
	Y int
}

func (v Vector2) AddVector(v2 Vector2) Vector2 {
	return Vector2{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}
func (v Vector2) AddInts(x int, y int) Vector2 {
	return Vector2{
		X: v.X + x,
		Y: v.Y + y,
	}
}
func (v Vector2) IsEqualInt(x int, y int) bool {
	return v.X == x && v.Y == y
}
func (v Vector2) Includes(arr []Vector2) bool {
	for _, v2 := range arr {
		if v2 == v {
			return true
		}
	}
	return false
}

type Grid struct {
	Start Vector2
	End   Vector2
	Grid  []string
}
type Reindeer struct {
	Position  Vector2
	Direction Vector2
	Score     int
}

func (r Reindeer) getEmptyNeighbourCells() []Vector2 {
	var res []Vector2
	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			if (x == 0 && y == 0) || (x != 0 && y != 0) {
				continue
			}

			v := r.Position.AddInts(x, y)
			if grid.Grid[v.Y][v.X] == '.' {
				res = append(res, Vector2{
					X: x,
					Y: y,
				})
			}
		}
	}
	return res
}
func PreviewGrid(r Reindeer) {
	for y, row := range grid.Grid {
		for x, cell := range row {
			if r.Position.IsEqualInt(x, y) {
				fmt.Print("R")
			} else if grid.Start.IsEqualInt(x, y) {
				fmt.Print("S")
			} else if grid.End.IsEqualInt(x, y) {
				fmt.Print("E")
			} else {
				fmt.Print(string(cell))
			}
		}
		fmt.Println()
	}
}
func Calculate() int {
	var lowestScore int
	var reindeers []*Reindeer

	reindeers = append(reindeers, &Reindeer{
		Position: grid.Start,
		Direction: Vector2{
			X: 1,
			Y: 0,
		},
	})
	var corners []Vector2

	for len(reindeers) > 0 {
		reindeer := reindeers[0]
		if grid.Grid[reindeer.Position.Y][reindeer.Position.X] != '.' || (lowestScore > 0 && reindeer.Score > lowestScore) {
			reindeers = reindeers[1:]
			continue
		}
		if reindeer.Position == grid.End {
			if lowestScore == 0 || lowestScore > reindeer.Score {
				lowestScore = reindeer.Score
			}
			reindeers = reindeers[1:]
			continue
		}

		neighbours := reindeer.getEmptyNeighbourCells()
		if reindeer.Score > 0 && (len(neighbours) <= 1 || reindeer.Position.Includes(corners)) {
			reindeers = reindeers[1:]
			continue
		}

		corners = append(corners, reindeer.Position)
		oppositeDir := Vector2{X: reindeer.Direction.X * -1, Y: reindeer.Direction.Y * -1}
		for _, n := range neighbours {
			if n == reindeer.Direction || n == oppositeDir {
				continue
			}
			reindeers = append(reindeers, &Reindeer{
				Position:  reindeer.Position.AddVector(n),
				Direction: n,
				Score:     reindeer.Score + 1001,
			})
		}

		reindeer.Position = reindeer.Position.AddVector(reindeer.Direction)
		reindeer.Score += 1
	}
	return lowestScore
}
func Initialize(input []byte) {
	grid.Grid = strings.Split(string(input), "\r\n")

	flattenInput := strings.ReplaceAll(string(input), "\r\n", "")
	startIndex := strings.Index(flattenInput, "S")
	endIndex := strings.Index(flattenInput, "E")
	grid.Start = Vector2{
		X: startIndex % len(grid.Grid[0]),
		Y: startIndex / len(grid.Grid),
	}
	grid.End = Vector2{
		X: endIndex % len(grid.Grid[0]),
		Y: endIndex / len(grid.Grid),
	}

	grid.Grid[grid.Start.Y] = strings.ReplaceAll(grid.Grid[grid.Start.Y], "S", ".")
	grid.Grid[grid.End.Y] = strings.ReplaceAll(grid.Grid[grid.End.Y], "E", ".")
}
func main() {
	start := time.Now()
	input, err := os.ReadFile("./input/03.txt")
	if err != nil {
		panic(err)
	}
	Initialize(input)
	res := Calculate()
	fmt.Println("Result", res)
	fmt.Println("Execution time: ", time.Since(start))
}
