package firstPart

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
	Output             string
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
	if len(c.Output) != 0 {
		c.Output += ","
	}
	c.Output += strconv.Itoa(c.GetComboOperand(op) & 7)
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

func Execute(instructions []int) {
	for computer.InstructionPointer < len(instructions) {
		if computer.InstructionPointer >= len(instructions)-1 {
			return
		}
		code := instructions[computer.InstructionPointer]
		fmt.Println("---------------")
		fmt.Println("Code", code)
		regs := append([]int{}, computer.Registers...)
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
		case 6:
			computer.Bdv(operand)
		case 7:
			computer.Cdv(operand)
		}

		fmt.Println("Operand", operand)
		computer.Debug(regs)

		computer.InstructionPointer += 2
	}
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
func main() {
	start := time.Now()
	input, err := os.ReadFile("./input/02.txt")
	if err != nil {
		panic(err)
	}
	instructions := Initialize(input)

	Execute(instructions)
	fmt.Println("Result", computer.Output)
	fmt.Println("Execution time: ", time.Since(start))
}
