package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Garden struct {
	grid []string
	cols int
	rows int
	res  int
}

var garden Garden

func (g Garden) hasVector(v Vector2) bool {
	return (v.X >= 0 && v.X <= g.cols) && (v.Y >= 0 && v.Y <= g.rows)
}
func (g Garden) getGridChar(v Vector2) rune {
	if g.hasVector(v) {
		return rune(g.grid[v.Y][v.X])
	}
	return ' '

}
func (g Garden) getCorners(v Vector2) int {
	vs := []Vector2{
		{
			X: 0,
			Y: -1,
		},
		{
			X: 1,
			Y: 0,
		},
		{
			X: 0,
			Y: 1,
		},
		{
			X: -1,
			Y: 0,
		},
		{
			X: 0,
			Y: -1,
		},
	}
	var corners int

	for i := range vs {
		if i == len(vs)-1 {
			break
		}

		curr := g.getGridChar(v)
		side1 := g.getGridChar(v.Add(vs[i]))
		side2 := g.getGridChar(v.Add(vs[i+1]))
		diag := g.getGridChar(v.Add(vs[i]).Add(vs[i+1]))

		if (curr != side1 && curr != side2) || (curr == side1 && curr == side2 && curr != diag) {
			corners += 1
		}
	}
	return corners
}

type Vector2 struct {
	X int
	Y int
}

func (v Vector2) Add(v2 Vector2) Vector2 {
	return Vector2{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}
func (v Vector2) GetNeighbors() []Vector2 {
	var res []Vector2
	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			if (x != 0 && y != 0) || (x == 0 && y == 0) {
				continue
			}

			res = append(res, v.Add(Vector2{
				X: x,
				Y: y,
			}))
		}
	}
	return res
}
func (v Vector2) InArr(arr []Vector2) bool {
	for i := range arr {
		if arr[i] == v {
			return true
		}
	}
	return false
}

func FindRegion(start Vector2, seen *[]Vector2) {
	q := []Vector2{
		start,
	}
	var area int
	var corners int
	c := garden.getGridChar(start)
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		*seen = append(*seen, p)

		area += 1
		corners += garden.getCorners(p)

		neighbors := p.GetNeighbors()
		for _, n := range neighbors {
			if n.InArr(*seen) || n.InArr(q) || !garden.hasVector(n) || garden.getGridChar(n) != c {
				continue
			}

			q = append(q, n)
		}
	}
	fmt.Println("Grid", string(garden.getGridChar(start)), "area", area, "corners", corners, "res", area*corners)
	garden.res += area * corners
}

func secondPart() {
	var seen []Vector2

	for y := range garden.grid {
		for x := range garden.grid[y] {
			place := Vector2{
				X: x,
				Y: y,
			}
			// if garden.getGridChar(place) != 'B' {
			// 	continue
			// }
			if !place.InArr(seen) {
				FindRegion(place, &seen)
			}
		}
	}
}

func main() {
	start := time.Now()
	input, _ := os.ReadFile(".\\input\\01.txt")
	rowsStr := strings.Split(string(input), "\r\n")
	garden = Garden{
		grid: rowsStr,
		cols: len(rowsStr) - 1,
		rows: len(rowsStr[0]) - 1,
	}

	secondPart()

	fmt.Println("Result", garden.res)

	fmt.Println("Execution time: ", time.Since(start))
}
