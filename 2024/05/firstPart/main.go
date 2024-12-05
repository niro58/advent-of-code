package firstPart

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
func includes (arr []string, arr2 []string) bool{
	for i := range arr {
		for j := range arr2 {
			if arr[i] == arr2[j]{
				return true
			}
		}
	}
	return false
}
func isValidRow(pages []string, rules map[string][]string) bool {
	prevPages := []string{}
	for i,p := range pages {
		rule := rules[p]
		if includes(prevPages, rule){
			fmt.Println(i,prevPages,p,rule)
			return false
		}
		prevPages = append(prevPages, p)
	}
	return true
}
func main() {
	start := time.Now()
	input, _ := os.ReadFile(".\\input\\01.txt")

	rows := strings.Split(string(input), "\r\n")
	rules, lastRow := createRules(rows)
	fmt.Println(lastRow)
	var res int
	for _,row := range rows[lastRow+1:]{
		pages := strings.Split(row, ",")
		if !isValidRow(pages, rules){
			continue
		}

		r, err := strconv.Atoi(pages[len(pages) / 2])
		if err != nil {
			panic(err)
		}
		res += r
	}

	fmt.Println(res)
	fmt.Println("Execution time: ", time.Since(start))
}
