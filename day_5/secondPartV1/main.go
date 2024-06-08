package secondPartV1

import (
	"fmt"
	"strconv"
	"strings"
)

// brute force clown fiesta
/*
A better version
insert a classic way of inputs and then create slices of it when needed
for example
49 10 53 12
which is correctly transformed to 49, 50, 51 . . . 59 and 53, 54, 55 . . . 65

insert it as a range struct ?? value and a range ->
if there is a range 49 to 55 that needs to be changed than we split (49, 10) to (49,6) and (56, 4) and do it until the end -> choose the lowest value ,ez clap
*/
type SoilMap struct {
	Key int
	Value int
	MaxOffset int
}
func getSoils(input string) []int {
	soilsLine := strings.Split(input, "\r\n")[0]
	soilsLine = strings.Replace(soilsLine, "seeds: ","",1)
	soils := []int{}
	soilsRaw := strings.Split(soilsLine, " ")
	for i := 0; i < len(soilsRaw); i += 2{
		fmt.Println(i)
		soilRange, err := strconv.Atoi(soilsRaw[i + 1])
		if err != nil {
			panic(err)
		}
		for j := range soilRange{
			soil, err := strconv.Atoi(soilsRaw[i])
			if err != nil{
				panic(err)
			}
			soils = append(soils, soil + j)
		}
	}

	return soils
}
func getSoilMaps(input string) [][]SoilMap{
	maps := [][]SoilMap{}
	temp := []SoilMap{}
	fmt.Println("starting soil maps")
	for index, line := range strings.Split(input, "\r\n"){
		if index == 0 || len(line) == 0 || line[0] < '0' || line[0] > '9'{
			if len(temp) > 0{
				maps = append(maps, temp)
				temp = []SoilMap{}
			}
			continue
		}
		lineSplit := strings.Split(line, " ")

		if (len(lineSplit) != 3){
			panic("not 3 length line split")
		}
		lineSplitInt := []int{}
		for _,i := range lineSplit{
			j, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			lineSplitInt = append(lineSplitInt, j)
		}
		m := SoilMap{
			Key: lineSplitInt[1],
			Value: lineSplitInt[0],
			MaxOffset: lineSplitInt[2],

		}
		temp = append(temp, m)
	}
	if len(temp) > 0{
		maps = append(maps, temp)
	}
	return maps
}
func calculateLowestSoil(soils []int, soilMaps [][]SoilMap) int{
	fmt.Println("starting calculating lowest soil")
	for i := range soils{
		for _, soilMap := range soilMaps{
			for _, soilMapValue := range soilMap{
				if soils[i] < soilMapValue.Key || soils[i] > soilMapValue.Key + soilMapValue.MaxOffset{
					continue
				}
				soils[i] += (soilMapValue.Value - soilMapValue.Key)
				break
			}
			fmt.Println(soils)
		}
	}
	min := soils[0]
	for _, i := range soils{
		if i < min{
			min = i
		}
	}

	return min
}
func Main(input string) int {
	soils := getSoils(input)
	soilMaps := getSoilMaps(input)

	res := calculateLowestSoil(soils, soilMaps)
	return res
}