package impl

import "github.com/epnushtaeva/sorts/internal"

type BubbleSort struct {
}

func (sort BubbleSort) Sort(sortingArray []int) []int {
	result := sortingArray
	sortingArrayLength := len(sortingArray)

	for mainArrayIndex := 0; mainArrayIndex < sortingArrayLength-1; mainArrayIndex++ {
		for subArrayIndex := 0; subArrayIndex < sortingArrayLength-mainArrayIndex-1; subArrayIndex++ {
			if result[subArrayIndex] > result[subArrayIndex+1] {
				internal.Swap(result[:], subArrayIndex, subArrayIndex+1)
			}
		}
	}

	return result
}

