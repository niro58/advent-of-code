//4,1,2,7,6,6,3,0,3

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Computer struct {
	Registers          []int
	InstructionPointer int
	Output             []int
}

// Division
func (c Computer) Division(r int, op int) int {
	return r / (1 << c.GetComboOperand(op))
}
func (c *Computer) Adv(op int) {
	c.Registers[0] = c.Division(c.Registers[0], op)
}

// Bitwise Xor
func (c *Computer) Bxl(op int) {
	c.Registers[1] = c.Registers[1] ^ op
}

// Modulo
func (c *Computer) Bst(op int) {
	c.Registers[1] = c.GetComboOperand(op) & 7
}

func (c *Computer) Jnz(op int) {
	if c.Registers[0] == 0 || c.InstructionPointer == op {
		return
	}
	c.InstructionPointer = op
}
func (c *Computer) Bxc() {
	c.Registers[1] = c.Registers[1] ^ c.Registers[2]
}
func (c *Computer) Out(op int) {
	c.Output = append(c.Output, c.GetComboOperand(op)&7)
}
func (c *Computer) Bdv(op int) {
	c.Registers[1] = c.Division(c.Registers[0], op)

}
func (c *Computer) Cdv(op int) {
	c.Registers[2] = c.Division(c.Registers[0], op)

}
func (c *Computer) Debug(oldRegs []int) {
	fmt.Println("Instruction Pointer", c.InstructionPointer)
	for i := range c.Registers {
		fmt.Println("Register", i, oldRegs[i], "->", c.Registers[i])
	}
}

var computer Computer

func Execute(instructions []int) bool {
	outpI := 0
	for computer.InstructionPointer < len(instructions) {
		if computer.InstructionPointer >= len(instructions)-1 {
			return isArrEqual(computer.Output, instructions)
		}
		code := instructions[computer.InstructionPointer]
		operand := instructions[computer.InstructionPointer+1]

		switch code {
		case 0:
			computer.Adv(operand)
		case 1:
			computer.Bxl(operand)
		case 2:
			computer.Bst(operand)
		case 3:
			tmp := computer.InstructionPointer
			computer.Jnz(operand)
			if tmp != computer.InstructionPointer {
				continue
			}
		case 4:
			computer.Bxc()
		case 5:
			computer.Out(operand)
			// if computer.Output[outpI] != instructions[outpI] {
			// 	return false
			// }
			outpI += 1
		case 6:
			computer.Bdv(operand)
		case 7:
			computer.Cdv(operand)
		}

		computer.InstructionPointer += 2
	}

	return outpI == len(instructions)
}
func (c *Computer) Clear() {
	for i := range c.Registers {
		c.Registers[i] = 0
	}
	c.InstructionPointer = 0
	c.Output = []int{}
}
func (c Computer) GetComboOperand(operand int) int {
	if operand <= 3 {
		return operand
	} else if operand >= 7 {
		panic(operand)
	} else {
		return c.Registers[operand-4]
	}
}
func Initialize(input []byte) []int {
	rows := strings.Split(string(input), "\r\n")
	for _, row := range rows[0:3] {
		v := strings.Split(row, ":")[1]
		vInt, err := strconv.Atoi(v[1:])
		if err != nil {
			panic(err)
		}
		computer.Registers = append(computer.Registers, vInt)
	}
	var instructions []int
	for _, c := range strings.Split(strings.Split(rows[4], ":")[1][1:], ",") {
		cInt, err := strconv.Atoi(c)
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, cInt)
	}
	return instructions
}
func isArrEqual(arr []int, comp []int) bool {
	if len(arr) != len(comp) {
		return false
	}
	for i := range arr {
		if arr[i] != comp[i] {
			return false
		}
	}
	return true
}
func bruteForce(instructions []int) int {
	i := 0

	for !Execute(instructions) {
		//create file if doesnt exist
		if _, err := os.Stat("./result-testing"); os.IsNotExist(err) {
			os.Mkdir("./result-testing", 0755)
		}
		file, _ := os.OpenFile(fmt.Sprintf("./result-testing/%d.txt", len(computer.Output)), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		defer file.Close()
		file.WriteString(fmt.Sprintf("%d | %v\n", i, computer.Output))
		computer.Clear()
		computer.Registers[0] = i
		i += 1
	}
	return i - 1
}
func main() {
	start := time.Now()
	input, err := os.ReadFile("./input/02.txt")
	if err != nil {
		panic(err)
	}
	instructions := Initialize(input)
	res := bruteForce(instructions)
	fmt.Println("Result", res)
	fmt.Println("Execution time: ", time.Since(start))
}
