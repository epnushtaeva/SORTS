package impl

import "github.com/epnushtaeva/sorts/internal"

type FastSort struct {
}

func (sort FastSort) Sort(sortingArray []int) []int {
	result := sortingArray

	sort.sort(result, 0, len(result)-1)

	return result
}

func (sort FastSort) sort(sortingArray []int, minIndex int, maxIndex int) {
	if maxIndex-minIndex <= 0 {
		return
	}

	arrayCenterIndex := internal.GetCenterIndex(maxIndex, minIndex)

	leftElementIndex := 0
	rightElementIndex := maxIndex

	for ; ; {
		leftElementIndex = internal.GetLeftElementIndex(sortingArray, arrayCenterIndex, leftElementIndex)
		rightElementIndex = internal.GetRightElementIndex(sortingArray, rightElementIndex, arrayCenterIndex)

		if leftElementIndex == arrayCenterIndex && rightElementIndex == arrayCenterIndex {
			break
		}

		internal.Swap(sortingArray[:], leftElementIndex, rightElementIndex)

		leftElementIndex++
		rightElementIndex--
	}

	sort.sort(sortingArray, minIndex, arrayCenterIndex)
	sort.sort(sortingArray, arrayCenterIndex+1, maxIndex)
}
