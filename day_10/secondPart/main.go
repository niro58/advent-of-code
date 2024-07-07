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
type secondMatrixStruct struct {
	Position position
	IsEnclosed bool
	AsciiChar rune
	Distance int
}
var commandList = map[rune]struct {
	Directions []position;
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
func isEnclosed (matrix [][]secondMatrixStruct, startPoint position) bool{
	var checkedPoints []position
	var activePoints []position
	activePoints = append(activePoints, startPoint)

	for len(activePoints) > 0{
		point := activePoints[0]
		activePoints = activePoints[1:]
		checkedPoints = append(checkedPoints,point)

		for y := -1 ; y <= 1; y++{
			for x := -1 ; x <= 1 ; x++{
				if x != 0 && y != 0 && x != y{
					continue
				}
				newPoint := position{
					x: point.x + x,
					y: point.y + y,
				}

				if newPoint.y < 0 || newPoint.x < 0 ||  len((matrix)) <= newPoint.y || len((matrix)[0]) <= newPoint.x {
					return false
				}

				isDuplicate := false
				for _, p := range checkedPoints{
					if p.x == newPoint.x && p.y == newPoint.y{
						isDuplicate = true
					}
				}
				if isDuplicate{
					continue
				}
				for _, p := range activePoints{
					if p.x == newPoint.x && p.y == newPoint.y{
						isDuplicate = true
					}
				}
				if isDuplicate{
					continue
				}

				mp :=  matrix[newPoint.y][newPoint.x]
				if mp.IsEnclosed || mp.AsciiChar == 'S'{
					continue
				}

				activePoints = append(activePoints, newPoint)
			}
		}
	}
	return true
}

func calculateMaxDistance(input string) int {
	matrix, startPoint := createMatrix(input)


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

	var nonLoopPoints []position
	for _, row := range matrix{
		for _, val := range row{
			if val.Distance == 0 && val.AsciiChar != 'S'{
				nonLoopPoints = append(nonLoopPoints, val.Position)
			}
		}
	}


	ySize := len(matrix)
	xSize := len(matrix[0])
	fmt.Println("Matrix size Y:", ySize, "X:", xSize )

	secondMatrix := make([][]secondMatrixStruct, ySize * 2)
	for y := range secondMatrix{
		secondMatrix[y] = make([]secondMatrixStruct, xSize * 2)
	}

	fmt.Println("Matrix size Y:", len(secondMatrix), "X:", len(secondMatrix[0]) )

	for y, row := range matrix{
		for x, cell := range row{

			if cell.Distance == 0 && cell.AsciiChar != 'S'{
				continue
			}
			secondMatrixPos := position{
				x: x * 2,
				y: y * 2,
			}
			secondMatrix[secondMatrixPos.y][secondMatrixPos.x] = secondMatrixStruct{
				Position: position{
					x: x * 2,
					y: y * 2,
				},
				IsEnclosed: cell.Distance > 0,
				AsciiChar: cell.AsciiChar,
				Distance: cell.Distance,
			}

			commandValue, ok := commandList[cell.AsciiChar]
			if !ok{
				panic(ok)
			}

			for _, dir := range commandValue.Directions{
				mainMatrixDirPos := position{
					x: x + dir.x,
					y: y + dir.y,
				}
				secondMatrixDirPos := position{
					x: secondMatrixPos.x + dir.x,
					y: secondMatrixPos.y + dir.y,
				}
				if mainMatrixDirPos.y < 0 || mainMatrixDirPos.x < 0 ||  len(matrix) <= mainMatrixDirPos.y || len(matrix[0]) <= mainMatrixDirPos.x {
					continue
				}
				distDiff := cell.Distance - matrix[mainMatrixDirPos.y][mainMatrixDirPos.x].Distance

				if distDiff != -1 && distDiff != 1 {
					continue
				}


				secondMatrix[secondMatrixDirPos.y][secondMatrixDirPos.x] = secondMatrixStruct{
					Position: position{
						x: secondMatrixPos.x,
						y: secondMatrixPos.y,
					},
					IsEnclosed: true,
					AsciiChar: '-',
					Distance: cell.Distance,
				}


			}


		}
	}
	res := 0
	for i, p := range nonLoopPoints{
		fmt.Println(i, "/", len(nonLoopPoints))
		secondMatrixPos := position{
			x: p.x * 2,
			y: p.y * 2,
		}
		secondMatrix[secondMatrixPos.y][secondMatrixPos.x].Distance = -1
		if isEnclosed(secondMatrix, secondMatrixPos){
			res += 1
			fmt.Println("Is enclosed", secondMatrixPos)
		}
	}

	fmt.Println("Total len is", res)
	return res
}


func main() {
	input,err := os.ReadFile("..\\inputs\\01_001.txt")
    if err != nil {
        fmt.Print(err)
    }

	res := calculateMaxDistance(string(input))
	fmt.Println(res)

}