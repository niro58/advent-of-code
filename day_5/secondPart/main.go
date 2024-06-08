package secondPart

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
type Soil struct{
	Value int
	Range int
}
func getSoils(input string) []Soil {
	soilsLine := strings.Split(input, "\r\n")[0]
	soilsLine = strings.Replace(soilsLine, "seeds: ","",1)
	soils := []Soil{}
	soilsRaw := strings.Split(soilsLine, " ")
	for i := 0; i < len(soilsRaw); i += 2{
		fmt.Println(i)
		soilRange, err := strconv.Atoi(soilsRaw[i + 1])
		if err != nil {
			panic(err)
		}
		soilValue, err := strconv.Atoi(soilsRaw[i])
		if err != nil {
			panic(err)
		}
		soil := Soil{
			Value: soilValue,
			Range: soilRange,
		}
		soils = append(soils, soil)

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
func calculateLowestSoil(soils []Soil, soilMaps [][]SoilMap) int{
	fmt.Println(soils)
	fmt.Println("starting calculating lowest soil")
	fmt.Println(soilMaps)
	for _, soilMap := range soilMaps{
		for _, soilMapValue := range soilMap{
			temp := []Soil{}
			for _, soil := range soils{
				if soil.Range == 0 {
					continue
				}
				isOutOfRange := soilMapValue.Key > soil.Value + soil.Range || soilMapValue.Key + soilMapValue.MaxOffset < soil.Value
				if isOutOfRange {
					temp = append(temp, soil)
					continue
					}
				fmt.Println("----")
				fmt.Println(soilMapValue)
				fmt.Println(soil)

				isWholeRange :=  soilMapValue.Key <= soil.Value && soilMapValue.Key + soilMapValue.MaxOffset >= soil.Value + soil.Range
				isInBetween := soilMapValue.Key > soil.Value && soilMapValue.Key + soilMapValue.MaxOffset < soil.Value + soil.Range
				if isWholeRange{
					fmt.Println("Is whole range")
					rangeSoil := Soil{
						Value: soil.Value - soilMapValue.Key + soilMapValue.Value,
						Range: soil.Range,
					}
					fmt.Println(rangeSoil)
					temp = append(temp, rangeSoil)
				}else if isInBetween {
					fmt.Println("is in between")
					startSoil := Soil {
						Value: soil.Value,
						Range: soilMapValue.Key - soil.Value,
					}
					inBetweenSoil := Soil {
						Value: soilMapValue.Value,
						Range: (soilMapValue.Key + soilMapValue.MaxOffset) - soilMapValue.Key,
					}
					endSoil := Soil{
						Value: soilMapValue.Key + soilMapValue.MaxOffset,
						Range: (soil.Value + soil.Range) - (soilMapValue.Key + soilMapValue.MaxOffset),
					}
					fmt.Println(startSoil,inBetweenSoil,endSoil)
					temp = append(temp, startSoil, inBetweenSoil,endSoil)
				}else if (soilMapValue.Key <= soil.Value){
					fmt.Println("is from left")
					startSoil := Soil{
						Value: soilMapValue.Value  + (soil.Value  - soilMapValue.Key),
						Range: (soilMapValue.Key + soilMapValue.MaxOffset) - soil.Value,
					}
					endSoil := Soil{
						Value: soilMapValue.Key + soilMapValue.MaxOffset ,
						Range: (soil.Value + soil.Range) - (soilMapValue.Key + soilMapValue.MaxOffset),
					}
					fmt.Println(startSoil,endSoil)
					temp = append(temp, startSoil,endSoil)
				}else{
					fmt.Println("is from right")
					startSoil:= Soil{
						Value: soil.Value,
						Range: soilMapValue.Key - soil.Value,
					}
					endSoil := Soil{
						Value: soilMapValue.Value,
						Range: soil.Value + soil.Range - soilMapValue.Key,
					}
					fmt.Println(startSoil,endSoil)
					temp = append(temp, startSoil,endSoil)
				}

				}
			soils = temp
		}
	}
	min := soils[0].Value
	for _, s := range soils{
		if s.Value < min && s.Range > 0{
			min = s.Value
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