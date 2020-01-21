package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func rangeExtraction(list []int) string {
	var integerList []string
	var possibleRangeElements []int
	var rangeElements string
	counter := 0
	index := 0

	// Pass in our list and a func to compare values and sort the integers in ascending order
	sort.Slice(list, func(i, j int) bool {
		numA := list[i]
		numB := list[j]
		return numA < numB
	})

start:
	if index < len(list)-1 {
		if math.Abs(float64(list[index])-float64(list[index+1])) == 1 {
			possibleRangeElements = append(possibleRangeElements, list[index])
			counter++
			//last index in the list
			if index+1 == len(list)-1 {
				possibleRangeElements = append(possibleRangeElements, list[index+1])
				if counter >= 2 {
					rangeElements = strconv.Itoa(possibleRangeElements[0]) + "-" + strconv.Itoa(list[index+1])
					integerList = append(integerList, rangeElements)
				} else {
					integerList = append(integerList, strconv.Itoa(possibleRangeElements[0]))
					integerList = append(integerList, strconv.Itoa(possibleRangeElements[1]))
				}
				possibleRangeElements = nil
			}
		} else {
			if counter > 0 {
				possibleRangeElements = append(possibleRangeElements, list[index])
				if counter >= 2 {
					rangeElements = strconv.Itoa(possibleRangeElements[0]) + "-" + strconv.Itoa(list[index])
					integerList = append(integerList, rangeElements)
				} else {
					integerList = append(integerList, strconv.Itoa(possibleRangeElements[0]))
					integerList = append(integerList, strconv.Itoa(possibleRangeElements[1]))
				}
				possibleRangeElements = nil
			} else {
				integerList = append(integerList, strconv.Itoa(list[index]))
			}

			//last index in the list
			if index+1 == len(list)-1 {
				integerList = append(integerList, strconv.Itoa(list[index+1]))
			}
			counter = 0
		}
		index++
		goto start
	}

	var solution string
	index1 := 0

start2:
	if index1 < len(integerList) {
		solution += integerList[index1]
		if index1 < len(integerList)-1 {
			solution += ","
		}
		index1++
		goto start2
	}

	return solution
}
func main() {
	integers := []int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20, 23, 24, 25, 22}
	fmt.Println(rangeExtraction(integers[0:]))
}
