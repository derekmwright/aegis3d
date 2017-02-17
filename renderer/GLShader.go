package renderer

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.5-core/gl"
)

// GLShader compiles GLShader and returns the shader
func GLShader(shaderType uint32, source string) (uint32, error) {
	// Create shader based on shaderType
	shader := gl.CreateShader(shaderType)
	// Compile the Shader with the source
	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLen int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLen)
		log := strings.Repeat("\x00", int(logLen+1))
		gl.GetShaderInfoLog(shader, logLen, nil, gl.Str(log))

		return 0, fmt.Errorf("Failed to compile %v: %v", source, log)
	}

	return shader, nil
}
