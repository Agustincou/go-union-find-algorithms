package algorithms

type unionFind struct {
	elementsIDs []int
}

type UnionFindInterface interface {
	AreConnected(elementIndex1, elementIndex2 int) bool
	MakeUnion(elementIndex1, elementIndex2 int) error
}

func GetElementsIDsSlice(cap int) []int {
	slice := make([]int, cap)

	for index, _ := range slice {
		slice[index] = index
	}

	return slice
}

func GetElementsWeightSlice(cap int) []int {
	slice := make([]int, cap)

	for index, _ := range slice {
		slice[index] = 1
	}

	return slice
}
