package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)
type Node struct {
	Left *Node
	Right *Node
	Value int
}
func NewNode(value int) *Node{
	return &Node{
		Left: nil,
		Right:nil,
		Value:value,
	}
}
func getSequences(input string) [][]int{
	lines := strings.Split(input, "\r\n")
	arr := make([][]int, len(lines))
	for index, line := range lines{
		lineArr := strings.Split(line, " ")
		lineArrInt := []int{}
		for _, i := range lineArr{
			j, err := strconv.Atoi(i)
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
	seqSteps := [][]int{}
	seqSteps = append(seqSteps, seq)

	for !allZero(seqSteps[len(seqSteps) - 1]){
		latestSeq := seqSteps[len(seqSteps) - 1]
		currSeq := []int{}
		for index, i := range latestSeq{
			if index == 0{
				continue
			}
			currSeq = append(currSeq, i - latestSeq[index - 1])
		}
		seqSteps = append(seqSteps, currSeq)
	}
	increment := 0
	for i := len(seqSteps)-1; i>=0; i--{
		val := seqSteps[i][len(seqSteps[i]) - 1]
		increment = val + increment

	}
	fmt.Println(increment)
	fmt.Println(seqSteps)
	fmt.Println("------------")
	return increment
}
func calculateFirstPart(input string) int {
	sequences := getSequences(input)
	sum := 0
	for _, sequence := range sequences{
		sum += calculateSequence(sequence)
	}
	return sum
}

func main() {
	input,err := os.ReadFile("inputs\\01_002.txt")
    if err != nil {
        fmt.Println(err)
    }

	res := calculateFirstPart(string(input))
	fmt.Println(res)

}