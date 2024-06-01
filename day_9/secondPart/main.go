package secondPart

import (
	"fmt"
	"strconv"
	"strings"
)


func getSequences(input string) [][]int{
	lines := strings.Split(input, "\r\n")
	arr := make([][]int, len(lines))
	for index, line := range lines{
		lineArr := strings.Split(line, " ")
		lineArrInt := []int{}
		for i := len(lineArr)-1; i >= 0; i--{
			j, err := strconv.Atoi(lineArr[i])
			if err != nil {
				panic(err)
			}
			lineArrInt = append(lineArrInt, j)
		}
		arr[index] = lineArrInt
	}
	return arr
}
func allZero(input []int) bool{
	for _, i := range input{
		if i != 0{
			return false
		}
	}
	return true
}
func calculateSequence(seq []int) int {
	fmt.Println(seq)
	seqSteps := [][]int{}
	seqSteps = append(seqSteps, seq)

	for !allZero(seqSteps[len(seqSteps) - 1]){
		latestSeq := seqSteps[len(seqSteps) - 1]
		currSeq := []int{}
		for index, i := range latestSeq{
			if index == 0{
				continue
			}
			currSeq = append(currSeq, latestSeq[index - 1] - i)
		}
		seqSteps = append(seqSteps, currSeq)
	}
	increment := 0
	for i := len(seqSteps)-1; i>=0; i--{
		val := seqSteps[i][len(seqSteps[i]) - 1]
		increment = val - increment

	}
	fmt.Println(increment)
	fmt.Println(seqSteps)
	fmt.Println("------------")
	return increment
}
func Main(input string) int {
	sequences := getSequences(input)
	sum := 0
	for _, sequence := range sequences{
		sum += calculateSequence(sequence)
	}
	return sum
}