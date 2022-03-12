package main

import (
	"testing"
)

func BenchmarkColumnUpdate(b *testing.B) {
	column := NewMatrixColumn(50, 1)
	for i := 0; i < b.N; i++ {
		column.Update()
	}
}

func BenchmarkMatrixAnimation(b *testing.B) {
	matrix := NewMatrix(80, 50)
	for i := 0; i < b.N; i++ {
		matrix.Animate()
	}
}

func BenchmarkMatrixProjection(b *testing.B) {
	matrix := NewMatrix(80, 50)
	for i := 0; i < b.N; i++ {
		matrix.ProjectVisualization()
	}
}

func BenchmarkScreenMapPaint(b *testing.B) {
	// Creates an empty screenMap
	screenMap := NewScreenMap(80, 50)
	for i := 0; i < b.N; i++ {
		screenMap.Paint()
	}
}
