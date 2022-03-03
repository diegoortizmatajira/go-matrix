package main

import (
	"time"

	"github.com/nsf/termbox-go"
)

const spacing = 2
const background = termbox.ColorBlack

func Paint(screenMap *ScreenMap) {
	for i := 0; i < screenMap.height; i++ {
		for j := 0; j < screenMap.width; j++ {
			fg := termbox.RGBToAttribute(byte(0), screenMap.content[i][j].shade, byte(0))
			if screenMap.content[i][j].shade == 255 {
				// Is a head, then add bold to the color
				fg = termbox.RGBToAttribute(237, 255, 242)
				fg |= termbox.AttrBold
			}
			termbox.SetCell(j*spacing, i, screenMap.content[i][j].symbol, fg, background)
		}
	}
	termbox.Flush()
	time.Sleep(50 * time.Millisecond)
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetOutputMode(termbox.OutputRGB)
	termbox.SetInputMode(termbox.InputEsc)
	termbox.Clear(termbox.ColorGreen, background)

	output := make(chan *ScreenMap, 3)
	w, h := termbox.Size()
	matrix := NewMatrix(w/spacing, h)
	go matrix.Produce(output)
	event_queue := make(chan termbox.Event)
	go func() {
		for {
			event_queue <- termbox.PollEvent()
		}
	}()
loop:
	for {
		select {
		case ev := <-event_queue:
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
				break loop
			}
		default:
			Paint(<-output)
		}
	}
}
