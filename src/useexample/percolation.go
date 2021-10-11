package useexample

import (
	"errors"
	"go-union-find-algorithms/src/algorithms"
)

type Percolation struct {
	matrixPercolated      [][]bool                      // NxN
	unionFindAlgorithm    algorithms.UnionFindInterface // NxN + 2(top and bottom extra nodes to detect leak)
	squareMatrixSideSize  int
	squareMatrixTotalSize int
	openSitesQuantity     int
	topElementIndex       int
	bottomElementIndex    int
}

func NewPercolation(squareMatrixSideSize int, desiredAlgorithm algorithms.AvailableAlgorithms) *Percolation {
	percolation := Percolation{}
	percolation.squareMatrixSideSize = squareMatrixSideSize
	percolation.squareMatrixTotalSize = squareMatrixSideSize * squareMatrixSideSize
	percolation.matrixPercolated = getSquareBooleanMatrix(squareMatrixSideSize)

	percolation.topElementIndex = percolation.squareMatrixTotalSize
	percolation.bottomElementIndex = percolation.topElementIndex + 1

	switch desiredAlgorithm {
	case algorithms.QuickFindAlgorithm:
		percolation.unionFindAlgorithm = algorithms.NewQuickFind(percolation.squareMatrixTotalSize + 2)
	case algorithms.QuickUnionAlgorithm:
		percolation.unionFindAlgorithm = algorithms.NewQuickUnion(percolation.squareMatrixTotalSize+2, false, false)
	case algorithms.WeightedQuickUnionAlgorithm:
		percolation.unionFindAlgorithm = algorithms.NewQuickUnion(percolation.squareMatrixTotalSize+2, false, true)
	case algorithms.CompressedWeightedQuickUnionAlgorithm:
		percolation.unionFindAlgorithm = algorithms.NewQuickUnion(percolation.squareMatrixTotalSize+2, true, true)
	default:
		percolation.unionFindAlgorithm = algorithms.NewQuickUnion(percolation.squareMatrixTotalSize+2, true, true)
	}

	// Top elements of square matrix was connected to topElement
	for index := 0; index < percolation.squareMatrixSideSize; index++ {
		_ = percolation.unionFindAlgorithm.MakeUnion(index, percolation.topElementIndex)
	}

	// Bottom elements of square matrix was connected to bottomElement
	for index := percolation.squareMatrixTotalSize - 1; index >= (percolation.squareMatrixTotalSize - percolation.squareMatrixSideSize); index-- {
		_ = percolation.unionFindAlgorithm.MakeUnion(index, percolation.bottomElementIndex)
	}

	return &percolation
}

func (p *Percolation) Open(row, col int) error {
	outboundCheckError := p.checkOutboundError(row, col)
	if outboundCheckError == nil {
		if p.matrixPercolated[row-1][col-1] == false {
			p.matrixPercolated[row-1][col-1] = true
			p.openSitesQuantity++
			p.connectWithOpenNeighbors(row, col)
		}
		return nil
	} else {
		return outboundCheckError
	}
}

func (p *Percolation) IsOpen(row, col int) (bool, error) {
	outboundCheckError := p.checkOutboundError(row, col)
	if outboundCheckError == nil {
		return p.matrixPercolated[row-1][col-1], nil
	} else {
		return false, outboundCheckError
	}
}

func (p *Percolation) IsFull(row, col int) (bool, error) {
	outboundCheckError := p.checkOutboundError(row, col)
	if outboundCheckError == nil {
		return p.unionFindAlgorithm.AreConnected(p.topElementIndex, p.matrixIndexToArrayIndex(row, col)), nil
	} else {
		return false, outboundCheckError
	}
}

func (p *Percolation) NumberOfOpenSites() int {
	return p.openSitesQuantity
}

func (p *Percolation) Percolates() bool {
	return p.unionFindAlgorithm.AreConnected(p.topElementIndex, p.bottomElementIndex)
}

func (p *Percolation) checkOutboundError(row, col int) error {
	if row > 0 && row <= len(p.matrixPercolated[0]) &&
		col > 0 && col <= len(p.matrixPercolated[0]) {
		return nil
	} else {
		return errors.New("index outbound exception")
	}
}

func (p *Percolation) matrixIndexToArrayIndex(row, col int) int {
	return p.squareMatrixSideSize*(row-1) + p.squareMatrixSideSize*(col-1)
}

func (p *Percolation) connectWithOpenNeighbors(row, col int) {

	neighborsMatrixIndexes := [][]int{
		{row+1,col},
		{row-1,col},
		{row,col+1},
		{row,col-1},
	}

	for _, matrixIndex := range neighborsMatrixIndexes {
		outboundCheckError := p.checkOutboundError(matrixIndex[0], matrixIndex[1])
		if outboundCheckError == nil {
			if p.matrixPercolated[matrixIndex[0]][matrixIndex[1]] {
				_ = p.unionFindAlgorithm.MakeUnion(p.matrixIndexToArrayIndex(row, col), p.matrixIndexToArrayIndex(matrixIndex[0], matrixIndex[1]))
			}
		}
	}
}

func getSquareBooleanMatrix(cap int) [][]bool {
	booleanMatrix := make([][]bool, cap)
	for i := range booleanMatrix {
		booleanMatrix[i] = make([]bool, cap)
	}

	return booleanMatrix
}
