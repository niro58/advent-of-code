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
}
type command struct{
	AsciiChar rune
	DirectionA position
	DirecitonB position
}
/*
| is a vertical pipe connecting north and south.
- is a horizontal pipe connecting east and west.
L is a 90-degree bend connecting north and east.
J is a 90-degree bend connecting north and west.
7 is a 90-degree bend connecting south and west.
F is a 90-degree bend connecting south and east.
*/
var commandList = []command{
	{
		AsciiChar: '|',
		DirectionA: position{
			x: 0,
			y: 1,
		},
		DirecitonB: position{
			x: 0,
			y: -1,
		},
	},
	{
		AsciiChar: '-',
		DirectionA: position{
			x: 1,
			y: 0,
		},
		DirecitonB: position{
			x: -1,
			y: 0,
		},
	},
	{
		AsciiChar: 'L',
		DirectionA: position{
			x: -1,
			y: 0,
		},
		DirecitonB: position{
			x: 0,
			y: -1,
		},
	},
	{
		AsciiChar: 'J',
		DirectionA: position{
			x: 1,
			y: 0,
		},
		DirecitonB: position{
			x: 0,
			y: -1,
		},
	},
	{
		AsciiChar: '7',
		DirectionA: position{
			x: 1,
			y: 0,
		},
		DirecitonB: position{
			x: 0,
			y: 1,
		},
	},
	{
		AsciiChar: 'F',
		DirectionA: position{
			x: -1,
			y: 0,
		},
		DirecitonB: position{
			x: 0,
			y: 1,
		},
	},
}
func createMatrix(input string) ([][]mapPoint, position){
	var xSize, ySize int
	lines := strings.Split(input, "\r\n")
	xSize = len(lines)
	ySize = len(lines[0])
	matrix := make([][]mapPoint, ySize)
	var startPosition position

	for index := range matrix{
		matrix[index] = make([]mapPoint, xSize)
	}

	for y, str := range lines{
		for x, r := range str{
			point := mapPoint{
				AsciiChar: r,
				Position: position{
					x: x,
					y: y,
				},
			}
			if r == 'S'{
				startPosition = point.Position
			}
			matrix[y][x] = mapPoint{
				AsciiChar: r,
				Position: position{
					x: x,
					y: y,
				},
			}
		}
	}
	fmt.Println(matrix)
	return matrix, startPosition
}

func calculateMaxDistance(input string) int {
	matrix, startPosition := createMatrix(input)

	var activePositions []position
	activePositions = append(activePositions, startPosition)




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