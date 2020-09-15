package main

import (
	"github.com/epnushtaeva/sorts/internal"
	"github.com/epnushtaeva/sorts/pkg/interface"
	"github.com/epnushtaeva/sorts/pkg/impl"
	"os"
)

func main() {
	sortType := os.Args[1]
	inputArray := internal.GetArrayForSort(os.Args[2:])

	internal.PrintBeforeSortMessage()
	internal.PrintArray(inputArray)

    var sort _interface.Sort

	switch sortType {
		case internal.BUBBLE_SORT:
			 sort = impl.BubbleSort{}
		case internal.FAST_SORT:
			 sort = impl.FastSort{}
		case internal.SHAKE_SORT:
			 sort = impl.ShakeSort{}
		default:
			break
	}

	resultArray := sort.Sort(inputArray)

	internal.PrintAfterSortMessage()
	internal.PrintArray(resultArray)
}