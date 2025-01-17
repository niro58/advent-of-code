package firstPart

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var cols, rows int
var toElapse = 100

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

type Robot struct {
	Position  Vector2
	Direction Vector2
}

func (r Robot) RepeatMove(times int) Vector2 {
	return r.Position.Add(
		Vector2{
			X: r.Direction.X * times,
			Y: r.Direction.Y * times,
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
func GetRobots(input string) []Robot {
	var robots []Robot
	for i, row := range strings.Split(input, "\r\n") {
		if i < 2 {
			continue
		}
		dirArr := strings.Split(strings.Split(row, "v=")[1], ",")
		positionsArr := strings.Split(strings.Split(strings.Split(row, "p=")[1], " ")[0], ",")

		var dirs = make([]int, 2)
		var positions = make([]int, 2)

		for i := range dirArr {
			d, _ := strconv.Atoi(dirArr[i])
			dirs[i] = d

			p, _ := strconv.Atoi(positionsArr[i])
			positions[i] = p
		}
		robots = append(robots, Robot{
			Position: Vector2{
				X: positions[0],
				Y: positions[1],
			},
			Direction: Vector2{
				X: dirs[0],
				Y: dirs[1],
			},
		})

	}
	return robots
}
func IncludesVector(robots []Robot, v Vector2) bool {
	for _, robot := range robots {
		if robot.Position == v {
			return true
		}
	}
	return false
}
func IncludesVectorCount(robots []Robot, v Vector2) int {
	var result int
	for _, robot := range robots {
		if robot.Position == v {
			result += 1
		}
	}
	return result
}
func PreviewGrid(robots []Robot) {
	for y := range rows {
		for x := range cols {
			v := Vector2{
				X: x,
				Y: y,
			}
			c := IncludesVectorCount(robots, v)
			if c > 0 {
				fmt.Print(strconv.Itoa(c), " ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}
func SplitToQuadrants(robots []Robot) []int {
	quadrants := make([]int, 4)
	xMiddle := cols / 2
	yMiddle := rows / 2
	var xOffset, yOffset int
	if cols%2 == 1 {
		xOffset = 1
	}
	if rows%2 == 1 {
		yOffset = 1
	}
	positions := []Vector2{
		{
			X: 0,
			Y: 0,
		},
		{
			X: xMiddle + xOffset,
			Y: 0,
		},
		{
			X: 0,
			Y: yMiddle + yOffset,
		},
		{
			X: xMiddle + xOffset,
			Y: yMiddle + yOffset,
		},
	}
	fmt.Println(positions)
	for y := range yMiddle {
		for x := range xMiddle {
			for i, p := range positions {
				v := p.AddBasic(Vector2{
					X: x,
					Y: y,
				})
				count := IncludesVectorCount(robots, v)
				quadrants[i] += count
			}
		}
	}
	return quadrants
}
func Calculate(input string) int {
	SetColsRows(input)
	robots := GetRobots(input)

	for i, robot := range robots {
		robots[i].Position = robot.RepeatMove(100)
		PreviewGrid(robots)
	}
	quadrantCounts := SplitToQuadrants(robots)

	result := 1
	for _, quadrant := range quadrantCounts {
		result *= quadrant
	}
	return result
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
