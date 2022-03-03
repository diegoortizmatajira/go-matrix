package main

import (
	"math/rand"
	"sync"
)

type MatrixColumn struct {
	Data         []rune
	Colors       []byte
	Streams      []*Stream
	columnNumber int
	height       int
	shadingDelta byte
}

type Stream struct {
	Head   int
	Target int
}

func NewMatrixColumn(height, columnNumber int) *MatrixColumn {
	return &MatrixColumn{
		Data:   make([]rune, height),
		Colors: make([]byte, height),
		Streams: []*Stream{
			{
				Head:   0,
				Target: height,
			},
		},
		columnNumber: columnNumber,
		height:       height,
		shadingDelta: byte(255 / height),
	}
}

func NewStreamFrom(head, screenHeight int) *Stream {
	target := rand.Intn(screenHeight-head) + head
	return &Stream{
		Head:   head,
		Target: target,
	}
}
func NewStream(screenHeight int) *Stream {
	head := rand.Intn(screenHeight)
	return NewStreamFrom(head, screenHeight)
}

func getShadedColor(currentColor byte, shadingDelta byte) byte {
	if currentColor < shadingDelta {
		return 0
	} else {
		return currentColor - shadingDelta
	}
}

func (s *Stream) Advance(screenHeight int) bool {
	if rand.Float32() < 0.3 { // Remains static on 10% of the times
		return true
	}
	s.Head++
	return s.Head < s.Target
}

func (c *MatrixColumn) EnsureStreams() {
	haveFinished := false
	remainingStreams := make([]*Stream, 0)
	for _, stream := range c.Streams {
		c.Data[stream.Head] = generateSymbol()
		c.Colors[stream.Head] = byte(255)
		if stream.Advance(c.height) {
			remainingStreams = append(remainingStreams, stream)
		} else {
			haveFinished = true
		}
	}
	streamCount := len(remainingStreams)
	if streamCount == 0 {
		remainingStreams = append(remainingStreams, NewStream(c.height))
		streamCount++
	}
	if haveFinished && streamCount == 1 && rand.Float32() < 0.3 {
		// If there is a single stream, then there is a 30% chance of a second stream starting from the top
		remainingStreams = append(remainingStreams, NewStreamFrom(0, c.height))
	}

	c.Streams = remainingStreams
}

func (c *MatrixColumn) Update(waitGroup *sync.WaitGroup) {
	for row := 0; row < c.height; row++ {
		c.Colors[row] = getShadedColor(c.Colors[row], c.shadingDelta)
	}
	c.EnsureStreams()
	waitGroup.Done()
}

func (c *MatrixColumn) ProjectVisualization(screenMap *ScreenMap) {
	for i := 0; i < c.height; i++ {
		screenMap.content[i][c.columnNumber] = SymbolMap{
			symbol: c.Data[i],
			shade:  c.Colors[i],
		}
	}
}
