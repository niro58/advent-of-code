package firstPart

import (
	"fmt"
	"strconv"
	"strings"
)

type SoilMap struct {
	Key int
	Value int
	MaxOffset int
}
func getSoils(input string) []int {
	soilsLine := strings.Split(input, "\r\n")[0]
	soilsLine = strings.Replace(soilsLine, "seeds: ","",1)
	soils := []int{}
	for _, soil := range strings.Split(soilsLine, " "){

		j, err := strconv.Atoi(soil)
		if err != nil{
			panic(err)
		}
		soils = append(soils, j)
	}

	return soils
}
func getSoilMaps(input string) [][]SoilMap{
	maps := [][]SoilMap{}
	temp := []SoilMap{}
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
	fmt.Println(soilMaps)
	for i := range soils{
		for _, soilMap := range soilMaps{
			for _, soilMapValue := range soilMap{
				if soils[i] < soilMapValue.Key || soils[i] > soilMapValue.Key + soilMapValue.MaxOffset{
					continue
				}
				soils[i] += (soilMapValue.Value - soilMapValue.Key)
				break
			}
			fmt.Println(soils, i)
		}
	}
	fmt.Println(soils)
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