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

func main() {
	input,err := os.ReadFile("inputs\\01_003.txt")
    if err != nil {
        fmt.Println(err)
    }

	res := calculateFirstPart(string(input))
	fmt.Println(res)

}