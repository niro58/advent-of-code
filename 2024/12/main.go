// 907858 too high
// 898664 too low
// 896810
package main

import (
	"fmt"
	"os"
	"strconv"
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

func includes(arr []int, el int) bool{
	for i := range arr {
		if arr[i] == el{
			return true
		}
	}
	return false
}

func indexToCoords(index, cols int) (int, int){
	return index % cols, index / cols
}
func previewGrid(points []int, walls []int, corners []int, cols,rows int){
	if len(points) == 0 {
		return  
	}
	r := ""
	for y := range rows{
		row := ""
		for x := range cols{
			index := y * cols + x
		 	if includes(corners, index) {
				row += "X "
			}else if includes(walls, index){
				row += "- "
			}else if includes(points, index){
				row += "Z "
			}else{
				row += "  "
			}
		}

		for _,c:=range row{
			if c != ' '{
				r +=  row + "\r\n"
				break
			}
		}
	}
	fmt.Println(r)
	fmt.Println("---------")
}
func previewGridWithNumbers(points []int, walls map[int]int, cols,rows int){
	if len(points) == 0 {
		return  
	}
	r := ""
	for y := range rows{
		row := ""
		for x := range cols{
			index := y * cols + x
		 	if walls[index] != 0{
				if walls[index] == 9 {

					row += "- "
				}else{
					row += strconv.Itoa(walls[index]) + " "
				}
			}else if includes(points, index){
				row += "Z "
			}else{
				row += "  "
			}
		}

		for _,c:=range row{
			if c != ' '{
				r +=  row + "\r\n"
				break
			}
		}
	}
	fmt.Println(r)
	fmt.Println("---------")
}
type Vector2 struct{
	X int
	Y int
}
func createBorderPoints(points []int, cols,rows int) []int{
    var borderPoints []int
    for _, p := range points {
        moves := []int{
            move(p, 1, 0, cols, rows),
            move(p, 1, 1, cols, rows),
            move(p, 1, -1, cols, rows),
            move(p, 0, 1, cols, rows),
            move(p, 0, -1, cols, rows),
            move(p, -1, 0, cols, rows),
            move(p, -1, 1, cols, rows),
            move(p, -1, -1, cols, rows),
        }
        for _, m := range moves {
            if m != -1 && !includes(points, m) && !includes(borderPoints, m) {
				previewGrid(points,[]int{m},[]int{},cols,rows)
                borderPoints = append(borderPoints, m)
            }
        }
    }
	return borderPoints
}
func createSubGrids(borderPoints []int, cols,rows int) [][]int{
	var grids [][]int
	var open []int
	var closed []int
	for len(borderPoints) > 0{
		point := borderPoints[0]
		borderPoints = borderPoints[1:]
		if includes(closed, point){
			continue
		}
		var grid []int

		open = append(open, point)
		for len(open) > 0 {
			p := open[0]
			borderPoints = remove(borderPoints, p)
		
			closed = append(closed, p)
			open = open[1:] 
			moves := []int{
				move(p, 1,0,cols,rows),
				move(p, -1,0,cols,rows),
				move(p, 0,1,cols,rows),
				move(p, 0,-1,cols,rows),
			}
			for _, m := range moves{
				if includes(open,m) || includes(closed,m){
					continue
				}
				if includes(borderPoints,m){
					open = append(open, m)
				}
			}

			grid = append(grid, p)
		}
		grids = append(grids, grid)
		previewGrid(grid,[]int{},[]int{},cols,rows)
		fmt.Println("---")
	}
	return grids
}
func remove(arr []int, v int) []int{
	p := []int{}
	for i := range arr{
		if arr[i] == v{
			continue
		}
		p = append(p, arr[i])
	}
	return p
}
func calculatePerimeter(points []int, cols,rows int) int{
	cols += 3
	rows += 3
	for i := range points {
		x1, y1 := indexToCoords(points[i], cols - 3)
		points[i] = (y1 + 1) * cols + x1 + 1
		points[i] = move(points[i],1,1,cols,rows)
	}

	borderPoints := createBorderPoints(points, cols,rows)
	grids := createSubGrids(borderPoints,cols,rows)
	
	var total int
	for gridIndex, grid := range grids{
		var edges = make(map[int]int)
		var borderPoint []int

		if gridIndex == 0 {
			borderPoint = grid
		}else{
			borderPoint = createBorderPoints(grid, cols,rows)
		}
		for _, index := range borderPoint{
	
			dirs := []Vector2{
				{
					X: 1,
					Y: 0,
				},
				{
					X: -1,
					Y: 0,
				},
				{
					X: 0,
					Y: 1,
				},
				{
					X: 0,
					Y: -1,
				},
			}
	
			edgeCount := 0
			for i := range dirs{
				m := move(index, dirs[i].X,dirs[i].Y, cols,rows)
				if gridIndex == 0 {
					if includes(points,m){
						edgeCount += 1
					}
				}else{
					if includes(grid,m){
						edgeCount += 1
					}
				}
			} 
	
	
			if edgeCount == 1{
				edges[index] = 9
			}else if edgeCount == 2 {
				var dir Vector2
				for i := range dirs{
					m := move(index, dirs[i].X,dirs[i].Y, cols,rows)
					if includes(points,m){
						if dir.X == 0 && dir.Y == 0{
							dir = dirs[i]
						}else if dir.X * -1 == dirs[i].X && dir.Y * -1 == dirs[i].Y{
							edges[index] = 9
						}else{
							total += 1
							edges[index] = 1
						}
					}
				} 
			}else if edgeCount == 3 {
				edges[index] = 3
				total += 3 
			}else if edgeCount == 4 {
				total += 4
				edges[index] = 4
			}else{
				total += 1
				edges[index] = 1
			}
			// previewGridWithNumbers(borderPoint, edges,cols,rows)
			// fmt.Println("Added", total - lTotal, "Total",total)
			// fmt.Println("------------")
		}
		previewGridWithNumbers(points, edges,cols,rows)
		previewGridWithNumbers(borderPoint, edges,cols,rows)
	}
	if total < 4 {
		return 4
	}
	fmt.Println("total walls", total)
	fmt.Println("---")
	return total
}

func secondPart(gridFlattened string, cols,rows int) int {
	var res int

	var walkedPoints []int
	
	for index, c := range gridFlattened{
		if includes(walkedPoints, index){
			continue
		}
		if c != 'P'{
			continue
		}

		var area,perimeter int
		var openPoints []int		
		var perimeterPoints []int

		openPoints = append(openPoints, index)
		
		for len(openPoints) > 0 {
			point := openPoints[0]
			perimeterPoints = append(perimeterPoints, point)
			
			walkedPoints = append(walkedPoints, point)

			openPoints = openPoints[1:]

			moves := []int{
				move(point,1,0,cols,rows),
				move(point,-1,0,cols,rows),
				move(point,0,1,cols,rows),
				move(point,0,-1,cols,rows),
			}
			area += 1

			for _, m := range moves{
				if m == -1 {
					continue
				}

				if !includes(walkedPoints, m) && !includes(openPoints, m) && rune(gridFlattened[m]) == c{
					openPoints = append(openPoints, m)
				}
			}
		}
		perimeter = calculatePerimeter(perimeterPoints, cols,rows)
		res += area * perimeter
	}

	return res
}

func main() {
	start := time.Now()
	input, _ := os.ReadFile(".\\input\\01.txt")
	inputFlat := strings.Replace(string(input), "\r\n", "", -1)
	rows := strings.Split(string(input), "\r\n")
	colsC := len(rows[0])

	res := secondPart(inputFlat, len(rows),colsC)


	fmt.Println("Result" , res)

	fmt.Println("Execution time: ", time.Since(start))
}
