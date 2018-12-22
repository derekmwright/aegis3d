package renderer

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

// GLProgram compiles shaders and creates a Program
func GLProgram(vertexShaderSrc, fragmentShaderSrc string) (uint32, error) {
	// Compile Vertex Shader
	vs, err := GLShader(gl.VERTEX_SHADER, vertexShaderSrc)
	if err != nil {
		panic(err)
	}
	// Compile Fragment Shader
	fs, err := GLShader(gl.FRAGMENT_SHADER, fragmentShaderSrc)
	if err != nil {
		panic(err)
	}
	program := gl.CreateProgram()
	// Compile and Link Shaders
	gl.AttachShader(program, vs)
	gl.AttachShader(program, fs)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to link program: %v", log)
	}
	// Cleanup compiled shaders
	gl.DeleteShader(vs)
	gl.DeleteShader(fs)

	return program, nil
}
