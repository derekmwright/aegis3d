package renderer

import (
	"fmt"

	"github.com/go-gl/gl/v4.5-core/gl"
)

func Init() {
	err := gl.Init()
	if err != nil {
		panic(err)
	}
	// Disable depth rest and face culling
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	// Output some information to console
	fmt.Println("AGE OpenGL Version: ", gl.GoStr(gl.GetString(gl.VERSION)))
	fmt.Println("AGE OpenGL Renderer:", gl.GoStr(gl.GetString(gl.RENDERER)))
	fmt.Println("AGE OpenGL Vendor:  ", gl.GoStr(gl.GetString(gl.VENDOR)))
}
