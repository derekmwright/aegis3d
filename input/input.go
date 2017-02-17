package input

import (
	"fmt"

	"github.com/go-gl/glfw/v3.2/glfw"
)

// Process is the callback for any keypresses in glfw
func Process(
	w *glfw.Window,
	key glfw.Key,
	scancode int,
	action glfw.Action,
	mods glfw.ModifierKey,
) {
	var event string
	switch action {
	case glfw.Press:
		event = "Pressed"
	case glfw.Release:
		event = "Released"
	}
	fmt.Println("Action: " + event)
	return
}
