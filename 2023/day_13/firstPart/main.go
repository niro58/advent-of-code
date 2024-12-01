package firstPart

import (
	"fmt"
	"os"
	"strings"
	"time"
)
func isMirrored(left int, rows []string) bool{
	right := left + 1
	for{
		
		if left < 0 || right >= len(rows) {
			return true
		}
		
		if rows[left] != rows[right] {
			return false
		}

		left -= 1
		right += 1 
	}
}
func findMatch(pattern string) int{
	fmt.Println(pattern)
	rows := strings.Split(pattern, "\r\n")
	
	for i := range rows{
		if i == len(rows) - 1 {
			continue
		}

		if rows[i] != rows[i + 1] {
			continue
		}

		if isMirrored(i,rows){
			return i + 1
		}
	}
	return 0
}

func transposePattern(pattern string) string {
	rows := strings.Split(pattern, "\r\n")
	i := 0
	transposed := make([]string,len(rows[0]))
	for i < len(rows[0]) {
		newRow := ""
		for j := range rows{
			newRow += string(rows[j][i])
		}
		transposed[i] = newRow
		i += 1
	}
	var res string
	for i := range transposed{
		res += transposed[i]
		if i == len(transposed) - 1{
			break
		}
		res += "\r\n"
	}
	return res
}
func firstPart(input string) int{
	patterns := strings.Split(input, "\r\n\r\n")
	var res int
	for _,pattern := range patterns {
		fmt.Println("------")
		vRes := findMatch(pattern)
		
		if vRes != 0{
			res += vRes * 100
			continue
		}
		
		hRes := findMatch(transposePattern(pattern))
		res += hRes
		
	}
	return res
}
func main() {
	start := time.Now()
	

	input, _ := os.ReadFile(".\\input\\1.txt")

	res:= firstPart(string(input))
	// output, _ := os.ReadFile(".\\output\\1.txt")
	// outputInt, _ := strconv.Atoi(string(output))
	fmt.Println("Result is ", res)
	// if (res != outputInt){
	// 	 panic("Test failed")
	// }
	fmt.Println("POG")
	fmt.Println("Execution time: ", time.Since(start))
}