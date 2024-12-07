// 6417120397561
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)
func isEqualSteps(steps string, target int64) bool{
	var res int64
	stps := strings.Split(steps, " ")
	operation := '-'
	for i := range stps{
		if stps[i] == "+"{
			operation = '+'
		}else if stps[i] == "||" {
			operation = '|'
		}else if stps[i] == "*"{
			operation = '*'
		}else{
			num, err := strconv.ParseInt(stps[i], 10, 64)
			if err != nil {
				panic(err)
			}
			if operation == '*'{
				res *= num
			}else if operation == '|'{
				concat := strconv.FormatInt(num,10)
				res = (res * 10 * int64(len(concat))) + num
			}else{
				res += num
			}
		}
	}
	return res == target
}
func isValidSequenece(sum int64, values []int64, target int64, steps string, isValid*bool){	
	if len(values) == 0 && sum == target {
		// areEq:= isEqualSteps(steps, target)
		// if !areEq{
		// 	panic(steps)
		// }
		fmt.Println("Steps", steps, " = ", target, isEqualSteps(steps, target))
		*isValid= true
		return
	}
	if sum > target || len(values) == 0{
		return
	}

	isValidSequenece(sum + values[0], values[1:],target, steps + " + " + strconv.FormatInt(values[0],10),isValid)
	isValidSequenece(sum * values[0], values[1:],target, steps + " * " + strconv.FormatInt(values[0],10),isValid)

	concat := strconv.FormatInt(sum,10) + strconv.FormatInt(values[0],10)
	n,_ := strconv.ParseInt(concat, 10,64) 
	isValidSequenece(n, values[1:],target, steps + " || " + strconv.FormatInt(values[0],10),isValid)
	
}
func firstPart(inp string) int64{
	var res int64
	rows := strings.Split(inp, "\r\n")
	
	for i := range rows {
		parts := strings.Split(rows[i], ":")
		target,_ := strconv.ParseInt(parts[0], 10, 64)
		numsStr := strings.Split(parts[1], " ")
		var nums []int64

		for j := range numsStr{
			n, err := strconv.ParseInt(numsStr[j],10,64)
			if err == nil {
				nums = append(nums, n)
			}
		}
		isValid := false
		if len(nums) == 1 {
			isValid = nums[0] == target
		}else{
			isValidSequenece(nums[0], nums[1:], target,  strconv.FormatInt(nums[0],10),&isValid)
		}

		if isValid{
			res += target
		}
	}

	return res
}
func main() {
	start := time.Now()
	input, _ := os.ReadFile(".\\input\\01.txt")

	
	res := firstPart(string(input))
	fmt.Println("Result" , res)

	fmt.Println("Execution time: ", time.Since(start))
}
