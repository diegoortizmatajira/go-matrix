package main

import "github.com/nsf/termbox-go"

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

func (screenMap *ScreenMap) Paint() {
	for i := 0; i < screenMap.height; i++ {
		for j := 0; j < screenMap.width; j++ {
			fg := termbox.RGBToAttribute(byte(0), screenMap.content[i][j].shade, byte(0))
			if screenMap.content[i][j].shade == 255 {
				// Is a head, then add bold to the color
				fg = termbox.RGBToAttribute(237, 255, 242) | termbox.AttrBold
			}
			termbox.SetCell(j*spacing, i, screenMap.content[i][j].symbol, fg, background)
		}
	}
	termbox.Flush()
}
