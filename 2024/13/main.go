package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
) 
type Vector2 struct{
	X int
	Y int
}
func divmod(numerator, denominator int64) (quotient, remainder int64) {
    quotient = numerator / denominator
    remainder = numerator % denominator
    return
}
func minTokens2(aButton, bButton, target Vector2)  int{
	var curr Vector2
	var tokens int
	for curr.X < target.X && curr.Y < target.Y{
		xDiv, xRem := divmod(int64(target.X - curr.X), int64(bButton.X))
		yDiv, yRem := divmod(int64(target.Y - curr.Y), int64(bButton.Y))
		
		if xDiv == yDiv && xDiv != 0 && yDiv != 0 && xRem == 0 && yRem == 0{
			return tokens + int(xDiv)
		}

		curr = Vector2{
			curr.X + aButton.X,
			curr.Y + aButton.Y,
		}
		tokens += 3

	}
	if curr == target{
		return tokens
	}
	return 0
}

func main() {
	start := time.Now()
	input, _ := os.ReadFile(".\\input\\01.txt")
	rows := strings.Split(string(input), "\r\n")
	var total int
	i := 0
	for (i) * 4 < len(rows){
		aButtonX,_ := strconv.Atoi(strings.Split(strings.Split(rows[i*4], "X")[1], ",")[0])
		aButtonY,_ := strconv.Atoi(strings.Split(rows[i*4], "Y")[1])
		aButton := Vector2{
			X: aButtonX,
			Y: aButtonY,
		}

		bButtonX,_ := strconv.Atoi(strings.Split(strings.Split(rows[i*4+1], "X")[1], ",")[0])
		bButtonY,_ := strconv.Atoi(strings.Split(rows[i*4+1], "Y")[1])
		bButton := Vector2{
			X: bButtonX,
			Y: bButtonY,
		}

		tButtonX,_ := strconv.Atoi(strings.Split(strings.Split(rows[i*4+2], "X=")[1], ",")[0])
		tButtonY,_ := strconv.Atoi(strings.Split(rows[i*4+2], "Y=")[1])
		target := Vector2{
			X: tButtonX  ,
			Y: tButtonY  ,
		}
		res := minTokens(aButton,bButton, target)
		total += res
		fmt.Println("Res",res)
		i += 1
	}
	fmt.Println("Total",total)
	fmt.Println("Execution time: ", time.Since(start))
}
