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
func (self *position) increment(add *position){
	self.x += add.x
	self.y += add.y
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
				x: -1,
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
				x: 1,
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
				x: 1,
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
				x: -1,
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
	var xSize, ySize int
	lines := strings.Split(input, "\r\n")
	xSize = len(lines)
	ySize = len(lines[0])
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
// todo: start handling , maybe in the main function before anything else, change grow function because it wont work much
func grow(matrix *[][]mapPoint, point *mapPoint) *mapPoint{
	i, ok := commandList[point.AsciiChar]
	if !ok{
		return point
	}
	aPos := position{
		point.Position.x,
		point.Position.y,
	}
	aPos.increment(&i.DirectionA)

	bPos:= position{
		point.Position.x,
		point.Position.y,
	}
	bPos.increment(&i.DirecitonB)

	aPoint := (*matrix)[aPos.y][aPos.x]
	bPoint := (*matrix)[bPos.y][bPos.x]

	if aPoint.Distance != 0 {
		aPoint.Distance = point.Distance + 1
		return &aPoint
	}else if bPoint.Distance != 0{
		bPoint.Distance = point.Distance + 1
		return &bPoint
	}else{
		panic("why no distances around me  ?!?!")
	}

}
func calculateMaxDistance(input string) int {
	matrix, startPoint := createMatrix(input)

	fmt.Println(startPoint)

	var activePoints []*mapPoint
	activePoints = append(activePoints,
		&matrix[startPoint.Position.x - 1][startPoint.Position.y],
		&matrix[startPoint.Position.x + 1][startPoint.Position.y],
		&matrix[startPoint.Position.x][startPoint.Position.y + 1],
		&matrix[startPoint.Position.x][startPoint.Position.y - 1],
	)
	for _, point := range activePoints{
		point.Distance = 1
	}

	for len(activePoints) > 0{
		n := len(activePoints) - 1 // top element
		point := activePoints[n]
		activePoints = activePoints[:n]


		fmt.Println("y",point)
	}



	return 0
}


func main() {
	input,err := os.ReadFile("..\\inputs\\01_001.txt")
    if err != nil {
        fmt.Println(err)
    }

	res := calculateMaxDistance(string(input))
	fmt.Println(res)

}