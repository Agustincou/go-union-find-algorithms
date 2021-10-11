package algorithms

import "errors"

// Association by connecting roots to each other

type quickUnion struct {
	unionFind
	UnionFindInterface
	weightSlice []int
	withCompression bool
	withWeightSlice bool
}

func NewQuickUnion(capacity int, withCompression bool, withWeightedSlice bool) UnionFindInterface {
	quickUnion := quickUnion{}

	quickUnion.elementsIDs = GetElementsIDsSlice(capacity)
	quickUnion.weightSlice = GetElementsWeightSlice(capacity)
	quickUnion.withCompression = withCompression
	quickUnion.withWeightSlice = withWeightedSlice

	return &quickUnion
}

func (q *quickUnion) AreConnected(elementIndex1, elementIndex2 int) bool {
	if elementIndex1 < len(q.elementsIDs) && elementIndex2 < len(q.elementsIDs) {
		return q.getRootIndex(elementIndex1) == q.getRootIndex(elementIndex2)
	} else {
		return false
	}
}

func (q *quickUnion) MakeUnion(elementIndex1, elementIndex2 int) error {
	if elementIndex1 < len(q.elementsIDs) && elementIndex2 < len(q.elementsIDs) {
		element1rootIndex := q.getRootIndex(elementIndex1)
		element2rootIndex := q.getRootIndex(elementIndex2)

		if element1rootIndex == element2rootIndex {
			return nil
		}

		//element2rootIndex used like "ID" because it´s a ROOT index, so element2rootIndex == element2rootID
		if q.weightSlice[element1rootIndex] > q.weightSlice[element2rootIndex] {
			q.elementsIDs[element2rootIndex] = element1rootIndex
			if q.withWeightSlice {
				q.weightSlice[element1rootIndex] += q.weightSlice[element2rootIndex]
			}
		} else {
			q.elementsIDs[element1rootIndex] = element2rootIndex
		}

		return nil
	} else {
		return errors.New("element index outbound")
	}
}

func (q *quickUnion) getRootIndex(elementIndex int) int {
	for tryNumber := 0; tryNumber < len(q.elementsIDs); tryNumber++ {
		if q.elementsIDs[elementIndex] != elementIndex {
			if q.withCompression {
				q.elementsIDs[elementIndex] = q.elementsIDs[q.elementsIDs[elementIndex]]
			}
			elementIndex = q.elementsIDs[elementIndex]
		} else {
			break
		}
	}

	return elementIndex
}
