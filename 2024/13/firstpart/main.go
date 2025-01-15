package firstpart

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
func minTokens(aButton, bButton, target, currPos Vector2, tokens int, lowestRes *int, tokenMap map[int]int) {
    p := currPos.Y*currPos.Y + currPos.X
    if tokenMap[p] != 0 && tokenMap[p] <= tokens {
        return
    }
    tokenMap[p] = tokens

	if currPos.X > target.X || currPos.Y > target.Y || (*lowestRes != 0 && tokens >= *lowestRes){
		// *lol -= 1
		// fmt.Println(*lol) 
		return
	}
	if currPos.X == target.X && currPos.Y == target.Y{
		if *lowestRes == 0 {
			*lowestRes = tokens
		}
		*lowestRes = min(*lowestRes, tokens)
		// fmt.Println(*lowestRes)
		// *lol -= 1
		return
	}
	aVector := Vector2{
		X: currPos.X + aButton.X,
		Y: currPos.Y + aButton.Y,
	}
	bVector := Vector2{
		X: currPos.X + bButton.X,
		Y: currPos.Y + bButton.Y,
	}
	
	minTokens(aButton,bButton, target ,aVector, tokens + 3, lowestRes, tokenMap)
	minTokens(aButton,bButton, target ,bVector, tokens + 1, lowestRes, tokenMap)
	// *lol += 2
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
			X: tButtonX,
			Y: tButtonY,
		}
		tokenMap := make(map[int]int)
		var res int
		minTokens(aButton,bButton, target, Vector2{X: 0,Y: 0},0,&res, tokenMap)
		total += res
		fmt.Println("Res",res)
		i += 1
	}
	fmt.Println("Total",total)
	fmt.Println("Execution time: ", time.Since(start))
}
