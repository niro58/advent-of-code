package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)
func abs(n int) int{
	mask := n >> 31
	return (n ^ mask) - mask
}
func firstPart(rows []string) int{
	var l []int
	var r []int
	for i := range rows{
		ps := strings.Split(rows[i], "   ")
		tmp, _ := strconv.Atoi(ps[0])
		l = append(l, tmp)
		tmp, _ = strconv.Atoi(ps[1])
		r = append(r, tmp)
	}
	
	sort.Slice(l, func(i, j int) bool {
		return l[i] < l[j]
	})

	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	var sum int
	for i := range l {
		sum += abs(l[i] - r[i])
	}

	return sum
}
func secondPart(rows []string) int{
	var l []int
	r := make(map[int]int)
	for i := range rows{
		ps := strings.Split(rows[i], "   ")
		tmp, _ := strconv.Atoi(ps[0])
		l = append(l, tmp)
		tmp, _ = strconv.Atoi(ps[1])
		r[tmp] += 1
	}
	var sum int
	
	for i := range l {
		sum += l[i] * r[l[i]]
	}

	return sum
}
func main() {
	start := time.Now()
	input, _ := os.ReadFile(".\\input\\01.txt")
	rows := strings.Split(string(input),"\r\n")
	
	res := secondPart(rows)
	fmt.Println(res)
	fmt.Println("Execution time: ", time.Since(start))
}