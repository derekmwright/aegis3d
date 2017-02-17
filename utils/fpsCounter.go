package utils

import (
	"strconv"

	"github.com/go-gl/glfw/v3.2/glfw"
)

type FpsCounter struct {
	previousTime float64
	frameCount   int
}

func (f *FpsCounter) UpdateFpsCounter(w *glfw.Window) {
	currentTime := glfw.GetTime()
	elapsedTime := currentTime - f.previousTime

	if elapsedTime > 0.25 {
		f.previousTime = currentTime
		fps := strconv.FormatFloat(
			float64(f.frameCount)/elapsedTime,
			'f',
			-1,
			64,
		)
		f.frameCount = 0
		title := "Aegis Game Engine @ fps: " + fps
		w.SetTitle(title)
	}
	f.frameCount++
}
