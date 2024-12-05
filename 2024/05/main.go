package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
) 

func createRules(rows []string) (map[string][]string, int){
	res := make(map[string][]string)
	var lastRow int
	for i := range rows{
		if len(rows[i]) == 0 {
			lastRow = i
			break
		}
		ps := strings.Split(rows[i], "|")
		
		res[ps[0]] = append(res[ps[0]], ps[1])
	}
	return res, lastRow
}
func includesIndex (arr []string, arr2 []string) int{
	for i := range arr {
		for j := range arr2 {
			if arr[i] == arr2[j]{
				return i
			}
		}
	}
	return -1
}
func getPages(pages []string, rules map[string][]string) []string {
	prevPages := []string{}
	hasSwapped := false
	var i int 
	for i < len(pages) {
		p := pages[i]
		rule := rules[p]

		matchIndex := includesIndex(prevPages,rule)
		if matchIndex != -1{
			hasSwapped = true
			tmp := pages[i]
			pages[i] = pages[matchIndex]
			pages[matchIndex] = tmp
			prevPages = []string{}
			i = 0
			continue
		}

		prevPages = append(prevPages, p)
		i += 1
	}
	if !hasSwapped{
		return []string{}
	}

	return pages
}
func main() {
	start := time.Now()
	input, _ := os.ReadFile(".\\input\\01.txt")

	rows := strings.Split(string(input), "\r\n")
	rules, lastRow := createRules(rows)
	var res int
	
	for _,row := range rows[lastRow+1:]{
		pages := strings.Split(row, ",")
		// fmt.Println("------")
		// fmt.Println(pages)
		
		pagesRes := getPages(pages, rules)
		if len(pagesRes) == 0 {
			continue
		}

		// fmt.Println(pages)

		r, err := strconv.Atoi(pages[len(pages) / 2])
		if err != nil {
			panic(err)
		}
		res += r
	}

	fmt.Println(res)
	fmt.Println("Execution time: ", time.Since(start))
}
