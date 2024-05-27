package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Direction struct{
	left string
	right string
}
func writeDirection(str string) {
	f, err := os.OpenFile("file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString(str + "\n")
	if err != nil {
		log.Fatal(err)
	}
}
func calculateFirstPart(input string) int {
	lines := strings.Split(input, "\r\n")

	directions := make(map[string]Direction, len(lines) - 2)
	currentPos := "AAA"
	commands := lines[0]

	for _, line := range lines[2:]{
		line = strings.Replace(line, " ", "",-1)
		commaDelimiter := strings.Split(line, ",")
		key := commaDelimiter[0][0:3]
		leftValue := commaDelimiter[0][len(commaDelimiter[0]) - 3:len(commaDelimiter[0])]
		rightValue := commaDelimiter[1][0:3]
		fmt.Println(key, "= (",leftValue,rightValue, ")")
		writeDirection(key+ " = ( " + leftValue+" "+rightValue+ " )")
		directions[key] = Direction{
			leftValue,
			rightValue,
		}
		if currentPos == ""{
			currentPos = key
		}
	}
	fmt.Println("Commands", commands, len(commands))

	index := 0
	for currentPos != "ZZZ"{
		for _,cmd := range commands{
			dir,ok := directions[currentPos]
			if (!ok){
				fmt.Println(currentPos,ok)
			}
			if(cmd == 'L'){
				currentPos = dir.left
			}else{
				currentPos = dir.right
			}
			index += 1

			writeDirection(currentPos)
		}
	}
	fmt.Println("Found the cuck in", index,"steps")
	return 0
}
func GCD(a, b int) int {
	for b != 0 {
			t := b
			b = a % b
			a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
			result = LCM(result, integers[i])
	}

	return result
}
func allElementsNotEqualToZero(arr []int) bool {
    for _, value := range arr {
		if value == 0 {
			return false
		}
    }
    return true
}
func calculateSecondPart(input string) int {
	lines := strings.Split(input, "\r\n")

	rawDirections := make(map[string]Direction, len(lines) - 2)

	commands := lines[0]

	for _, line := range lines[2:]{
		line = strings.Replace(line, " ", "",-1)
		commaDelimiter := strings.Split(line, ",")
		key := commaDelimiter[0][0:3]
		leftValue := commaDelimiter[0][len(commaDelimiter[0]) - 3:len(commaDelimiter[0])]
		rightValue := commaDelimiter[1][0:3]

		rawDirections[key] = Direction{
			leftValue,
			rightValue,
		}
	}

	var selectedDirections []string

	for key := range rawDirections {
		if strings.HasSuffix(key, "A"){
			selectedDirections = append(selectedDirections, key)
		}
	}
	selectedDirectionsSteps := make([]int,len(selectedDirections))

	step := 0
	for !allElementsNotEqualToZero(selectedDirectionsSteps){
		fmt.Println("----")
		fmt.Println(selectedDirections)
		fmt.Println(selectedDirectionsSteps)
		for _, cmd := range commands{
			for index, dir := range selectedDirections{
				if strings.HasSuffix(dir, "Z"){
					continue
				}

				nextDirection := rawDirections[dir]
				var nextDestination string
				if(cmd == 'L'){
					nextDestination = nextDirection.left
				}else{
					nextDestination = nextDirection.right
				}
				selectedDirections[index] = nextDestination
				if strings.HasSuffix(nextDestination, "Z"){
					selectedDirectionsSteps[index] = step + 1
				}
			}
			step += 1;
		}
	}
	return LCM(selectedDirectionsSteps[0],selectedDirectionsSteps[1],selectedDirectionsSteps[2:]...)
}

func main() {
	input,err := os.ReadFile("inputs\\02_002.txt")
    if err != nil {
        fmt.Println(err)
    }

	res := calculateSecondPart(string(input))
	fmt.Println(res)

}