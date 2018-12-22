package utils

import "github.com/go-gl/gl/v4.1-core/gl"

// ToggleWireframe toggles between LINE and FILL
func ToggleWireframe() {
	var pgMode int32
	gl.GetIntegerv(gl.POLYGON_MODE, &pgMode)
	if pgMode == gl.LINE {
		gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
	}
	if pgMode == gl.FILL {
		gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)
	}
}
