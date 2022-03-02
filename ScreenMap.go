package main

type SymbolMap struct {
	symbol rune
	shade  byte
}
type ScreenMap struct {
	width   int
	height  int
	content [][]SymbolMap
}

func NewScreenMap(w, h int) *ScreenMap {
	result := make([][]SymbolMap, h)
	for i := range result {
		result[i] = make([]SymbolMap, w)
	}
	return &ScreenMap{
		width:   w,
		height:  h,
		content: result,
	}
}
