package internal

import (
	"fmt"
	"strconv"
)

func GetArrayForSort(args []string) []int{
	inputArray := make([]int, len(args))
	var err error

	for index, value := range args {
		inputArray[index], err = strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
	}

	return inputArray
}

func PrintBeforeSortMessage(){
	fmt.Println("Массив до сортировки")
}

func PrintAfterSortMessage(){
	fmt.Println("Массив после сортировки")
}

func PrintArray(array []int){
	for _ , value := range array {
		fmt.Printf("%d ", value)
	}

	fmt.Println()
}

func Swap(array []int, index1 int, index2 int) {
	tempNumber := array[index1]
	array[index1] = array[index2]
	array[index2] = tempNumber
}

func GetCenterIndex(maxIndex int, minIndex int) int {
	arrayLength := maxIndex - minIndex + 1
	centerIndex := 0

	if arrayLength%2 == 0 {
		centerIndex = int(arrayLength/2) - 1
	} else {
		centerIndex = int((arrayLength - 1) / 2)
	}

	centerIndex += minIndex

	return centerIndex
}

func GetLeftElementIndex(sortingArray []int, arrayCenterIndex int, minIndex int) int{
	for index := minIndex; index <= arrayCenterIndex; index++ {
		if sortingArray[index] > sortingArray[arrayCenterIndex] {
			return index
		}
	}

	return arrayCenterIndex
}

func GetRightElementIndex(sortingArray []int, maxIndex int, arrayCenterIndex int) int{
	for index := maxIndex; index > arrayCenterIndex; index-- {
		if sortingArray[index] < sortingArray[arrayCenterIndex] {
			return index
		}
	}

	return arrayCenterIndex
}

