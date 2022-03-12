package main

import (
	"time"

	"github.com/nsf/termbox-go"
)

const spacing = 2
const background = termbox.ColorBlack
const FrameDuration = 30 * time.Millisecond

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
			break
		case screenMap := <-output:
			screenMap.Paint()
			time.Sleep(FrameDuration)
			break
		}
	}
}
