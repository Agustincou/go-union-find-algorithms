package algorithms

type unionFind struct {
	elementsIDs []int
}

type unionFindInterface interface {
	Initialize(capacity int)
	AreConnected(elementIndex1, elementIndex2 int) bool
	MakeUnion(elementIndex1, elementIndex2 int) error
}

func GetInitializedSlice(cap int) []int {
	slice := make ([]int, cap)

	for index, _ := range slice {
		slice[index] = index
	}

	return slice
}
