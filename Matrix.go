package main

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

func (m *Matrix) Animate() {
	// Updates the matrix
	for _, v := range m.Columns {
		v.Update()
	}
}

func (m *Matrix) ProjectVisualization() *ScreenMap {
	// Projects the resulting matrix
	result := NewScreenMap(m.width, m.height)
	for _, v := range m.Columns {
		v.ProjectVisualization(result)
	}
	return result
}

func (m *Matrix) Produce(output chan *ScreenMap) {
	for {
		m.Animate()
		output <- m.ProjectVisualization()
	}
}
