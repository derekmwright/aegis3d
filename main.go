// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Renders a textured spinning cube using GLFW 3 and OpenGL 4.1 core forward-compatible profile.
package main

import (
	"runtime"

	"github.com/derekmwright/aegis/input"
	r "github.com/derekmwright/aegis/renderer"
	"github.com/derekmwright/aegis/utils"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	windowWidth  = 800
	windowHeight = 600
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 5)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(
		windowWidth,
		windowHeight,
		"Aegis Game Engine",
		nil,
		nil,
	)
	// Check for errors creating window
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	// Setup Input
	window.SetKeyCallback(input.Process)

	r.Init()
	program, err := r.GLProgram(vertexShader, fragmentShader)

	// Check for Shader Compilation Errors
	if err != nil {
		panic(err)
	}

	gl.UseProgram(program)

	var vertices = []float32{
		0.0, 0.5, 0.0,
		0.5, -0.5, 0.0,
		-0.5, -0.5, 0.0,
	}

	var colors = []float32{
		1.0, 0.0, 0.0,
		0.0, 1.0, 0.0,
		0.0, 0.0, 1.0,
	}

	// Setup Vertex Buffer
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(
		gl.ARRAY_BUFFER,
		len(vertices)*9,
		gl.Ptr(vertices),
		gl.STATIC_DRAW,
	)
	var cbo uint32
	gl.GenBuffers(1, &cbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, cbo)
	gl.BufferData(
		gl.ARRAY_BUFFER,
		len(colors)*9,
		gl.Ptr(colors),
		gl.STATIC_DRAW,
	)

	// Setup Vertex Array
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
	gl.BindBuffer(gl.ARRAY_BUFFER, cbo)
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 0, nil)

	gl.EnableVertexAttribArray(0)
	gl.EnableVertexAttribArray(1)

	// Setup FPS Counter
	fps := new(utils.FpsCounter)

	for !window.ShouldClose() {
		// Render
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.UseProgram(program)
		gl.BindVertexArray(vao)
		gl.DrawArrays(gl.TRIANGLES, 0, 3)

		// Update FPS counter
		fps.UpdateFpsCounter(window)

		// Maintenance
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

var vertexShader = `
#version 450
layout(location = 0) in vec3 vertex_pos;
layout(location = 1) in vec3 vertex_rgb;

out vec3 color;

void main() {
	color = vertex_rgb;
	gl_Position = vec4(vertex_pos, 1.0);
}` + "\x00"

var fragmentShader = `
#version 450
in vec3 color;
out vec4 frag_color;

void main() {
	frag_color = vec4(color, 1.0);
}` + "\x00"
