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
	// Debugging purposes
	// fmt.Println("Current Location")
	// fmt.Println(string(point.AsciiChar), point.Distance, point.Position)
	// fmt.Println("Returning")
	// for _, r := range res{
	// 	fmt.Println(string(r.AsciiChar), r.Distance, r.Position)
	// }

	return res
}

func isAvailableCell (checkedCells *[]mapPoint, point *mapPoint) bool{
	for _, cell := range (*checkedCells){
		if cell.Position == point.Position{
			return false
		}
	}
	return true
}
func isEnclosedCell(matrix *[][]mapPoint, point mapPoint, alreadyCheckedCells *[]mapPoint) (bool, []mapPoint){
	
	var checkedCells []mapPoint
	var availableCells []mapPoint
	availableCells = append(availableCells, point)
	for len(availableCells) > 0{
		cell := availableCells[0]
		checkedCells = append(checkedCells, cell)
		availableCells = availableCells[1:]

		for x := -1; x <= 1; x++{
			for y := -1; y <= 1; y++{
				if x == y{
					continue
				}
				xNew := cell.Position.x + x
				yNew := cell.Position.y + y
				if yNew < 0 || xNew < 0 ||  len((*matrix)) <= yNew || len((*matrix)[0]) <= xNew {

					return false,checkedCells
				}

				cell := (*matrix)[yNew][xNew]
				if !isAvailableCell(alreadyCheckedCells, &cell) || isAvailableCell(&checkedCells, &cell){
					checkedCells = append(checkedCells, cell)
					continue
				}
				availableCells = append(availableCells, cell)
			}
		}
	}
	return true,checkedCells
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

	//printMatrix(&matrix)

	totalEnclosedChars := 0


	var checkedCells []mapPoint

	fmt.Println("Total", len(matrix) * len(matrix[0]))
	for _, row := range matrix{
		// fmt.Println(y, "/",len(matrix))
		for _, val := range row{
			if val.Distance != 0{
				fmt.Print(val.Distance, " ")
				if val.Distance < 10 {
					fmt.Print(" ")
				}
				continue
			}
			if !isAvailableCell(&checkedCells, &val){
				continue
			}

			ok, funcCheckedCells := isEnclosedCell(&matrix, val, &checkedCells)
			if ok {
				fmt.Print("I", " ", " ")
				totalEnclosedChars += 1
			}else{
				fmt.Print("0", " ", " ")
			}
			checkedCells = append(checkedCells, funcCheckedCells...)

		}
		fmt.Println()
	}


	return totalEnclosedChars
}


func main() {
	input,err := os.ReadFile("..\\inputs\\01_001.txt")
    if err != nil {
        fmt.Println(err)
    }

	res := calculateMaxDistance(string(input))
	fmt.Println(res)

}