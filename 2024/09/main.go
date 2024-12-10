// 90994085674
// 6390180901651
// 2 - 6412470176240
// 2 - 6412470176240
// 2 - 6412390114238
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)
type Block struct{
	Id int
	Start int
	End int
}
func createBlocks(inp string) []Block{
	var blocks []Block
	var id int
	var index int
	for i := range inp {
		var length int
		currId := id
		if i % 2 == 0 {
			dataLength, _ := strconv.Atoi(string(inp[i]))
			length = dataLength
			id += 1
		}else{
			spaces, _ := strconv.Atoi(string(inp[i]))
			length = spaces
			currId = -1
		}
		end := index + length - 1

		blocks = append(blocks, Block{
			Id: currId,
			Start: index,
			End: index + length - 1,
		})

		index = end + 1
	}

	return blocks
}

func connectBlocks(blocks []Block) []Block{

	right := len(blocks) - 1
	for right > 0{
		for blocks[right].Id == -1{
			right -= 1
		}
		block := blocks[right]
		size := block.End - block.Start

		for i,b := range blocks {
			if b.Id != -1 || b.Start == -1 || b.Start > block.Start{
				continue
			}
			end := b.End
			bSize := b.End - b.Start
			if bSize >= size{
				// fmt.Println("--")
				// fmt.Println("ID", block.Id)
				// fmt.Println("End block ", block)
				// fmt.Println("Dots", b)
				blocks = append(blocks, 
					Block{
						Id: -1,
						Start: block.Start,
						End: block.End,
					},
				)
				block.Start = b.Start
				block.End = b.Start + size

				if block.End == end {
					blocks = append(blocks[:i], blocks[i+1:]...)
					blocks[right - 1] = block
				}else{
					b.Start = block.End + 1
					b.End = end
					blocks[i] = b
					blocks[right] = block
				}

				// blocksToString(blocks)
				// fmt.Println("End block ", block)
				// fmt.Println("Dots", blocks[i])
				// fmt.Println("--")
			}

		}
		right -= 1 
	}

	return blocks
}
func checksum(blocksConnected []Block) int {
	var maxI int 
	for i := range blocksConnected{
		maxI = max(maxI, blocksConnected[i].End)
	}

	var sum int
	for i := range maxI{
		for j := range blocksConnected{
			if blocksConnected[j].Id == -1{
				continue
			}
			if blocksConnected[j].Start <= i && blocksConnected[j].End >= i{
				sum += i * blocksConnected[j].Id
			}
			
		}
		i++
	}

	return sum
}

func blocksToString(blocks []Block) string{
	var maxI int 
	for i := range blocks{
		maxI = max(maxI, blocks[i].End)
	}
	var res string
	for i := range maxI + 1 {
		for j := range blocks {
			if blocks[j].Start <= i && blocks[j].End >= i{
				if blocks[j].Id == -1{
					res += "."
				}else{
					res += strconv.Itoa(blocks[j].Id)
				}
			}
		} 
	}
	fmt.Println(res)
	return res
}
func firstPart(inp string) int {
	blocks := createBlocks(inp)
	blocksToString(blocks)
	fmt.Println(blocks)
	cBlocks := connectBlocks(blocks)
	
	// os.WriteFile(".\\res\\01.txt", []byte(blocksToString(cBlocks)), 0644)
	res := checksum(cBlocks)
	return res
}
func main() {
	start := time.Now()
	input, _ := os.ReadFile(".\\input\\01.txt")
		
	res := firstPart(string(input))
	fmt.Println("Result" , res)

	fmt.Println("Execution time: ", time.Since(start))
}
