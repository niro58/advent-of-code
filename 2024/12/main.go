package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)
func move(start int, x int, y int, cols int, rows int) int {
	newX := start % cols + x
	newY := start / cols + y

	if newX >= cols  || newX < 0|| newY < 0 || newY >= rows {
		return -1
	}
	res := newX + newY * cols

	return res
}
type PlantSize struct{
	Area int
	Sides int
}

func includes(arr []int, el int) bool{
	for i := range arr {
		if arr[i] == el{
			return true
		}
	}
	return false
}
func includesIndexDir(arr []IndexDir, index int, dir Vector2) bool{
	for i := range arr {
		if arr[i].Dir == dir && arr[i].Index == index{
			return true
		}
	}
	return false
}
func remove(arr[]int, el int) []int{
	index := -1
	for i := range arr {
		if arr[i] == el{
			index = i
			break 
		}
	}
	if index == -1{
		panic("element not in array")
	}
	return append(arr[:index], arr[index+1:]...)
}
func indexToCoords(index, cols int) (int, int){
	return index % cols, index / cols
}
func previewGrid(points []int, secondPoints []int, cols,rows int){
	for y := range rows{
		for x := range cols{
			index := y * cols + x
			
			if includes(secondPoints, index){
				fmt.Print("L ")
			}else if includes(points, index){
				fmt.Print("Z ")
			}else{
				fmt.Print("X ")
			}
		}
		fmt.Println()
	}
	fmt.Println("---------")
}
type Vector2 struct{
	X int
	Y int
}
type IndexDir struct{
	Index int
	Dir Vector2
}
func calculateVertical(points []int, cols,rows int) int{
	cols += 2
	rows += 2
	for i := range points{
		x1,y1 := indexToCoords(points[i], cols - 2)
		points[i]  = (y1 + 1) * cols + x1 + 1
	}
	previewGrid(points, []int{},cols,rows)
	var total int
	var walls []IndexDir
	for y := range rows {
		for x := range cols{
			index := y * cols + x
			if includes(points,index){
				continue
			}
			moves := []Vector2{	
				{
					X: 1,
					Y: 0,

				},
				{
					X: 0,
					Y: 1,
				},
				{
					X: -1,
					Y: 0,
				} ,
				{
					X: 0,
					Y: -1,
				} ,
			}
			for _,v := range moves{
				pos := move(index, v.X, v.Y, cols,rows)
				if pos == -1  || !includes(points,pos){
					continue
				}
				
				var possibleMoves []Vector2
				if v.X != 0 {
					possibleMoves = append(possibleMoves, Vector2{X: 0,Y: 1}, Vector2{X: 0,Y: -1})
				}else if v.Y != 0 {
					possibleMoves = append(possibleMoves, Vector2{X: -1,Y: 0}, Vector2{X: 1,Y: 0})
				}
				if includesIndexDir(walls, index, possibleMoves[0]) || includesIndexDir(walls, index, possibleMoves[1]) {
					continue
				}
				// fmt.Println(index)
				for _, pm := range possibleMoves{
					p := index
					fmt.Println("Results", total)
					previewGrid(points, []int{p}, cols,rows)
					possiblePos := move(p, v.X, v.Y,cols,rows)
					for p != -1 && includes(points,possiblePos){
						
						fmt.Println("Results", total)
						previewGrid(points, []int{p}, cols,rows)
						walls = append(walls, IndexDir{
							Index: p,
							Dir: pm,
						})
						p = move(p, pm.X, pm.Y, cols,rows)
						if includes(points, p){
							break
						}
						possiblePos = move(p, v.X, v.Y,cols,rows)
					}

				}
				total += 1
			} 

		}
	}
	fmt.Println("total walls", total)
	return total
}

func firstPart(gridFlattened string, cols,rows int) int {
	var plants []PlantSize
	var res int

	var closedPoints []int
	
	for index, c := range gridFlattened{
		if includes(closedPoints, index){
			continue
		}

		var plant PlantSize
		var plantPoints []int		
		var perimeterPoints []int
		plantPoints = append(plantPoints, index)
		
		for len(plantPoints) > 0 {
			point := plantPoints[0]
			
			closedPoints = append(closedPoints, point)

			plantPoints = plantPoints[1:]

			moves := []int{
				move(point,1,0,cols,rows),
				move(point,-1,0,cols,rows),
				move(point,0,1,cols,rows),
				move(point,0,-1,cols,rows),
			}
			plant.Area += 1

			for _, m := range moves{
				if m == -1 {
					continue
				}
				if c == rune(gridFlattened[m]) && !includes(perimeterPoints, m) {
					perimeterPoints = append(perimeterPoints, m)
				}

				if !includes(closedPoints, m) && !includes(plantPoints, m) && rune(gridFlattened[m]) == c{
					plantPoints = append(plantPoints, m)
				}
			}
		}
		
		plant.Sides = calculateVertical(perimeterPoints, cols,rows)
		
		plants = append(plants, plant)
	}
	for i := range plants{
		fmt.Println(plants[i].Area , "|", plants[i].Sides)
		res += plants[i].Area * plants[i].Sides
	}

	return res
}

func main() {
	start := time.Now()
	input, _ := os.ReadFile(".\\input\\01.txt")
	inputFlat := strings.Replace(string(input), "\r\n", "", -1)
	rows := strings.Split(string(input), "\r\n")
	colsC := len(rows[0])

	res := firstPart(inputFlat, len(rows),colsC)


	fmt.Println("Result" , res)

	fmt.Println("Execution time: ", time.Since(start))
}
