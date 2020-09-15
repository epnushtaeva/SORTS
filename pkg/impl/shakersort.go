package impl

import "github.com/epnushtaeva/sorts/internal"

type ShakeSort struct {
}

func (sort ShakeSort) Sort(sortingArray []int) []int {
	result := sortingArray
	leftIndex := 0
	rightIndex := len(result) - 1

	for ; ; {
		if leftIndex > rightIndex {
			break
		}

		for index := rightIndex; index > leftIndex; index-- {
            if result[index - 1] > result[index] {
            	internal.Swap(result[:], index - 1, index);
			}
		}

		leftIndex++

		for index := leftIndex; index < rightIndex; index++ {
			if result[index] > result[index + 1] {
				internal.Swap(result[:], index + 1, index);
			}
		}

		rightIndex--
	}

	return result
}
