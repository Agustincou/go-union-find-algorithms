package algorithms

import "errors"

// Association using groups IDs

type QuickFind struct {
	unionFind
	unionFindInterface
}

func (q QuickFind) Initialize(capacity int) {
	q.elementsIDs = GetInitializedSlice(capacity)
}

func (q QuickFind) AreConnected(elementIndex1, elementIndex2 int) bool {
	if elementIndex1 < len(q.elementsIDs) && elementIndex2 < len(q.elementsIDs) {
		return q.elementsIDs[elementIndex1] == q.elementsIDs[elementIndex2]
	} else {
		return false
	}
}

func (q QuickFind) MakeUnion(elementIndex1, elementIndex2 int) error {
	if elementIndex1 < len(q.elementsIDs) && elementIndex2 < len(q.elementsIDs) {
		element1ID := q.elementsIDs[elementIndex1]
		element2ID := q.elementsIDs[elementIndex2]

		for index, elementID := range q.elementsIDs {
			if elementID == element1ID {
				q.elementsIDs[index] = element2ID
			}
		}

		return nil
	} else {
		return errors.New("element index outbound")
	}
}
