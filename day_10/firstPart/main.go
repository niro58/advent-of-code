package main

import (
	"fmt"
	"os"
	"strings"
)
type position struct{
	x int
	y int
}
type mapPoint struct{
	AsciiChar rune
	Position position
	Distance int
}
var commandList = map[rune]struct {
	Directions []position
}{
	'.': {
		Directions: nil,
	},
	'S': {
		Directions: []position{
			{
				x: 1,
				y: 0,
			},
			{
				x: -1,
				y: 0,
			},
			{
				x: 0,
				y: 1,
			},
			{
				x: 0,
				y: -1,
			},
		},
	},
	'|': {
		Directions: []position{
			{
				x: 0,
				y: 1,
			},
			{
				x: 0,
				y: -1,
			},
		},
	},
	'-': {
		Directions: []position{
			{
				x: 1,
				y: 0,
			},
			{
				x: -1,
				y: 0,
			},
		},
	},
	'L': {
		Directions: []position{
			{
				x: 1,
				y: 0,
			},
			{
				x: 0,
				y: -1,
			},
		},
	},
	'J': {
		Directions: []position{
			{
				x: -1,
				y: 0,
			},
			{
				x: 0,
				y: -1,
			},
		},
	},
	'7': {
		Directions: []position{
			{
				x: -1,
				y: 0,
			},
			{
				x: 0,
				y: 1,
			},
		},
	},
	'F': {
		Directions: []position{
			{
				x: 1,
				y: 0,
			},
			{
				x: 0,
				y: 1,
			},
		},
	},
}
func createMatrix(input string) ([][]mapPoint, *mapPoint){
	var ySize int
	var xSize int
	lines := strings.Split(input, "\r\n")
	xSize = len(lines[0])
	ySize = len(lines)
	matrix := make([][]mapPoint, ySize)
	var startPosition *mapPoint

	for index := range matrix{
		matrix[index] = make([]mapPoint, xSize)
	}

	for y, str := range lines{
		for x, r := range str{
			matrix[y][x] = mapPoint{
				AsciiChar: r,
				Position: position{
					x: x,
					y: y,
				},
			}
			if r == 'S'{
				startPosition = &matrix[y][x]
			}
		}
	}
	return matrix, startPosition
}

func grow(matrix *[][]mapPoint, point *mapPoint) []*mapPoint{
	commandValue, ok := commandList[point.AsciiChar]
	if !ok{
		return nil
	}
	var res []*mapPoint

	for _, direction := range commandValue.Directions{
		y := point.Position.y + direction.y
		x := point.Position.x + direction.x
		if y < 0 || x < 0 ||  len((*matrix)) <= y || len((*matrix)[0]) <= x {
			continue
		}
		matrixPoint := &((*matrix)[y][x])
		possibleDir, ok := commandList[matrixPoint.AsciiChar]
		if !ok {
			panic("errr")
		}
		if matrixPoint.Distance == 0 && matrixPoint.AsciiChar != '.' && matrixPoint.AsciiChar != 'S'{
			for _, dir := range possibleDir.Directions{
				if dir.x * -1 == direction.x && dir.y * -1 == direction.y{
					res = append(res, matrixPoint)
					matrixPoint.Distance = point.Distance + 1
				}
			}
		}
	}
	fmt.Println("Current Location")
	fmt.Println(string(point.AsciiChar), point.Distance, point.Position)
	fmt.Println("Returning")
	for _, r := range res{
		fmt.Println(string(r.AsciiChar), r.Distance, r.Position)
	}

	return res
}
func calculateMaxDistance(input string) int {
	matrix, startPoint := createMatrix(input)
	fmt.Println("Start point", startPoint.Position)


	var activePoints []*mapPoint
	activePoints = append(activePoints,
		&matrix[startPoint.Position.y][startPoint.Position.x],
	)
	takeStart := true
	for len(activePoints) > 0{
		var point *mapPoint
		if takeStart{
			point = activePoints[0]
			activePoints = activePoints[1:]
			takeStart = false
		}else{
			n := len(activePoints) - 1
			point = activePoints[n]
			activePoints = activePoints[:n]
			takeStart = true
		}

		newPoints := grow(&matrix, point)
		activePoints = append(activePoints, newPoints...)
	}
	max := 0
	for _, row := range matrix{

		for _, val := range row{
			if val.Distance > max{
				max = val.Distance
			}
			fmt.Print(val.Distance)
			fmt.Print(" ")
			if val.Distance < 10{
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}


	return max
}


func main() {
	input,err := os.ReadFile("..\\inputs\\01_001.txt")
    if err != nil {
        fmt.Println(err)
    }

	res := calculateMaxDistance(string(input))
	fmt.Println(res)

}