package algorithms

import "errors"

// Association by connecting roots to each other

type QuickUnion struct {
	unionFind
	unionFindInterface
}

func (q QuickUnion) Initialize (capacity int) {
	q.elementsIDs = GetInitializedSlice(capacity)
}

func (q QuickUnion) AreConnected(elementIndex1, elementIndex2 int) bool {
	if elementIndex1 < len(q.elementsIDs) && elementIndex2 < len(q.elementsIDs) {
		return q.getRootIndex(elementIndex1) == q.getRootIndex(elementIndex2)
	} else {
		return false
	}
}

func (q QuickUnion) MakeUnion(elementIndex1, elementIndex2 int) error {
	if elementIndex1 < len(q.elementsIDs) && elementIndex2 < len(q.elementsIDs) {
		element1rootIndex := q.getRootIndex(elementIndex1)
		element2rootIndex := q.getRootIndex(elementIndex2)

		//element2rootIndex used like "ID" because it´s a ROOT index, so element2rootIndex == element2rootID
		q.elementsIDs[element1rootIndex] = element2rootIndex

		return nil
	} else {
		return errors.New("element index outbound")
	}
}

func (q QuickUnion) getRootIndex(elementIndex int) int {
	for tryNumber := 0; tryNumber < len(q.elementsIDs); tryNumber++ {
		if q.elementsIDs[elementIndex] != elementIndex {
			elementIndex = q.elementsIDs[elementIndex]
		} else {
			break
		}
	}

	return elementIndex
}
