package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Vector2 struct {
	X int
	Y int
}

func minTokens(aButton, bButton, target Vector2) int {
	leftSide := target.X*bButton.Y - target.Y*bButton.X
	rightSide := aButton.X*bButton.Y - aButton.Y*bButton.X
	fmt.Printf("(%d-(%d*x))/%d = (%d-(%d*x))/%d\r\n", target.X, aButton.X, bButton.X, target.Y, aButton.Y, bButton.Y)
	if leftSide%rightSide == 0 {
		res := leftSide / rightSide
		target.X -= aButton.X * res
		if  target.X  % bButton.X != 0{
			return 0
		}
		rem := target.X / bButton.X
		return res*3 + rem
	}
	return 0
}

func main() {
	start := time.Now()
	input, _ := os.ReadFile(".\\input\\01.txt")
	rows := strings.Split(string(input), "\r\n")
	var total int
	i := 0
	for (i)*4 < len(rows) {
		aButtonX, _ := strconv.Atoi(strings.Split(strings.Split(rows[i*4], "X")[1], ",")[0])
		aButtonY, _ := strconv.Atoi(strings.Split(rows[i*4], "Y")[1])
		aButton := Vector2{
			X: aButtonX,
			Y: aButtonY,
		}

		bButtonX, _ := strconv.Atoi(strings.Split(strings.Split(rows[i*4+1], "X")[1], ",")[0])
		bButtonY, _ := strconv.Atoi(strings.Split(rows[i*4+1], "Y")[1])
		bButton := Vector2{
			X: bButtonX,
			Y: bButtonY,
		}

		tButtonX, _ := strconv.Atoi(strings.Split(strings.Split(rows[i*4+2], "X=")[1], ",")[0])
		tButtonY, _ := strconv.Atoi(strings.Split(rows[i*4+2], "Y=")[1])
		target := Vector2{
			X: tButtonX + 10000000000000,
			Y: tButtonY + 10000000000000,
		}

		res := minTokens(aButton, bButton, target)

		total += res
		if res != 0 {
			fmt.Println(tButtonX, res)
		}
		i += 1
	}
	fmt.Println("Total", total)
	fmt.Println("Execution time: ", time.Since(start))
}
