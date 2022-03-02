package main

import (
	"sync"
)

type Matrix struct {
	Columns []*MatrixColumn
	width   int
	height  int
}

func NewMatrix(w, h int) *Matrix {
	columns := make([]*MatrixColumn, w)
	for i := 0; i < w; i++ {
		columns[i] = NewMatrixColumn(h, i)
	}
	return &Matrix{
		Columns: columns,
		width:   w,
		height:  h,
	}
}

func (m *Matrix) Produce(output chan *ScreenMap) {
	for {
		// Updates the matrix
		waitGroup := &sync.WaitGroup{}
		waitGroup.Add(m.width)
		for _, v := range m.Columns {
			go v.Update(waitGroup)
		}
		waitGroup.Wait()
		// Projects the resulting matrix
		result := NewScreenMap(m.width, m.height)
		for _, v := range m.Columns {
			v.ProjectVisualization(result)
		}
		output <- result
	}
}
