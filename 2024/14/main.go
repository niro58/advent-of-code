// 427200000 too high
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var cols, rows int

type Vector2 struct {
	X int
	Y int
}

func (v Vector2) Add(v2 Vector2) Vector2 {
	v.X = (v.X + v2.X)
	v.Y = (v.Y + v2.Y)
	v.X %= cols
	if v.X < 0 {
		v.X += cols
	}

	v.Y %= rows
	if v.Y < 0 {
		v.Y += rows
	}
	return v
}
func (v Vector2) AddBasic(v2 Vector2) Vector2 {
	return Vector2{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}

func RepeatMove(position Vector2, vector Vector2, times int) Vector2 {
	return position.Add(
		Vector2{
			X: vector.X * times,
			Y: vector.Y * times,
		},
	)
}
func SetColsRows(input string) {
	rs := strings.Split(input, "\r\n")
	c, err := strconv.Atoi(rs[0])
	if err != nil {
		panic(err)
	}
	r, err := strconv.Atoi(rs[1])
	if err != nil {
		panic(err)
	}

	cols = c
	rows = r
}
func GetRobots(input string) (positions []Vector2, vectors []Vector2) {
	for i, row := range strings.Split(input, "\r\n") {
		if i < 2 {
			continue
		}
		dirArr := strings.Split(strings.Split(row, "v=")[1], ",")
		positionsArr := strings.Split(strings.Split(strings.Split(row, "p=")[1], " ")[0], ",")

		var vec = make([]int, 2)
		var pos = make([]int, 2)

		for i := range dirArr {
			d, _ := strconv.Atoi(dirArr[i])
			vec[i] = d

			p, _ := strconv.Atoi(positionsArr[i])
			pos[i] = p
		}
		positions = append(positions, Vector2{
			X: pos[0],
			Y: pos[1],
		})
		vectors = append(vectors, Vector2{
			X: vec[0],
			Y: vec[1],
		})

	}
	return positions, vectors
}
func Includes(arr []Vector2, el Vector2) bool {
	for _, e := range arr {
		if e == el {
			return true
		}
	}
	return false
}
func IncludeCount(arr []Vector2, el Vector2) int {
	c := 0
	for _, e := range arr {
		if e == el {
			c += 1
		}
	}
	return c
}
func PreviewGrid(positions []Vector2, highlight []Vector2) {
	for y := range rows {
		for x := range cols {
			v := Vector2{
				X: x,
				Y: y,
			}
			c := IncludeCount(positions, v)
			if c > 0 {
				if Includes(highlight, v) {
					fmt.Print("X", " ")
				} else {
					fmt.Print(strconv.Itoa(c), " ")
				}
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}
func (v Vector2) Distance(v2 Vector2) int {
	return int(math.Sqrt(math.Pow(float64(v.X)-float64(v2.X), 2) + math.Pow(float64(v.Y)-float64(v2.Y), 2)))
}
func HasAround(arr [][]bool, el Vector2) bool {
	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			if x == 0 && y == 0 {
				continue
			}
			np := el.AddBasic(Vector2{X: x, Y: y})
			if np.Y < 0 || np.Y > rows || np.X < 0 || np.X > cols {
				continue
			}
			if arr[np.Y][np.X] {
				return true
			}
		}
	}
	return false
}

func Calculate(input string) int {
	SetColsRows(input)
	positions, vectors := GetRobots(input)
	needFloods := 50
	i := 1
	for {
		for i, pos := range positions {
			positions[i] = RepeatMove(pos, vectors[i], 1)
		}

		var open []Vector2
		open = append(open, positions...)
		var walked []Vector2

		for len(open) > needFloods {
			curr := open[0]
			open = open[1:]
			if Includes(walked, curr) {
				continue
			}

			var flood []Vector2
			flood = append(flood, curr)
			floodSize := 0
			var currFlood []Vector2
			for len(flood) > 0 {
				floodSize += 1

				point := flood[0]
				flood = flood[1:]
				walked = append(walked, point)
				currFlood = append(currFlood, point)
				for y := -1; y <= 1; y++ {
					for x := -1; x <= 1; x++ {
						if x == 0 || y == 0 && (x != 0 && y != 0) {
							continue
						}
						newPoint := point.AddBasic(Vector2{
							X: x,
							Y: y,
						})
						if Includes(walked, newPoint) || Includes(flood, newPoint) || !Includes(open, newPoint) {
							continue
						}

						flood = append(flood, newPoint)
					}
				}
			}
			if floodSize > needFloods {
				PreviewGrid(positions, currFlood)
				fmt.Println("--------", i, "--------")
				fmt.Println()
			}
		}

		if i%100000 == 0 {
			fmt.Println(i)
		}
		i += 1

	}
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
