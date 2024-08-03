package main

import (
	"fmt"
	"os"
	"strings"
)
var GALAXY_RUNE = '#'
var EMPTY_RUNE ='.'
type Position struct{
	X int
	Y int
}
func printMap(m [][]rune){
	for _, row := range m{
		for _, cell := range row{
			fmt.Print(cell, " ")
		}
		fmt.Println()
	}
}
func getGalaxyMap(inp string) ( []Position){
	rows := strings.Split(inp, "\r\n")
	ySize := len(rows)
	xSize := len(rows[0])

	var galaxyPositions []Position
	var colOffsets []int
	var rowOffsets []int

	for y, row:= range rows{
		isBlank := true
		for x, cell := range row{
			if cell == GALAXY_RUNE{
				galaxyPositions = append(galaxyPositions, Position{
					X: x,
					Y: y,
				})
				isBlank = false
			}
		}
		if isBlank{
			rowOffsets = append(rowOffsets, y)
		}
	}
	for x := 0; x < xSize; x++{
		isBlank := true
		for y := 0 ; y < ySize; y++{
			if rune(rows[y][x]) == GALAXY_RUNE{
				isBlank = false
			}
		}
		if isBlank{
			colOffsets = append(colOffsets, x)
		}
	}

	for i, pos := range galaxyPositions{
		colOffset := 0
		rowOffset := 0

		for _, r := range rowOffsets{
			if r <= pos.Y{
				//to modify the distance
				rowOffset += 1000000 -1
			}
		}

		for _, c := range colOffsets{
			if c <= pos.X{
				colOffset += 1000000 - 1
			}
		}

		yPos := pos.Y + rowOffset
		xPos := pos.X + colOffset

		galaxyPositions[i] = Position{
			Y: yPos,
			X: xPos,
		}
	}

	return galaxyPositions
}
func calculateDistance(a Position, b Position ) int{
	distance := 0
	if a.X > b.X{
		distance += a.X - b.X
	}else{
		distance += b.X - a.X
	}
	if a.Y > b.Y {
		distance += a.Y - b.Y
	}else{
		distance += b.Y - a.Y
	}
	return distance
}
func getPairsSum(inp string) int{
	galaxyPositions := getGalaxyMap(inp)

	res := 0
	for i, aPos := range galaxyPositions{
		for j := i + 1; j < len(galaxyPositions);j++{
			res += calculateDistance(aPos, galaxyPositions[j])
		}
	}

	return res
}
func main() {
	input,err := os.ReadFile(".\\inputs\\01_001.txt")
    if err != nil {
        fmt.Print(err)
    }

	res := getPairsSum(string(input))
	fmt.Println(res)

}