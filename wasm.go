package main

import (
	"unsafe"
)

type Vec struct {
	X, Y, R, G, B, A float32
}

type Triangle struct {
	A, B, C Vec
}

var triangles []Triangle

var triangle *Triangle

var tptr uintptr

func update(ptr uintptr, size int)

func whynot(...int) int

func main() {
	Exec(`console.log("HAI");`)
	println("hai")
	requestAnimationFrame()
}

var outbuf []byte

var ptr uintptr

func init() {
	triangles = make([]Triangle, 1)
	sh := (*SliceHeader)(unsafe.Pointer(&triangles))
	tptr = sh.Data
	triangle = &triangles[0]
	triangle.A = Vec{X: -0.5, Y: 0.5, R: 1.0, G: 1.0, B: 1.0, A: 1.0}
	triangle.B = Vec{X: -0.5, Y: -0.5, R: 1.0, G: 1.0, B: 1.0, A: 1.0}
	triangle.C = Vec{X: 0.5, Y: -0.5, R: 1.0, G: 1.0, B: 1.0, A: 1.0}
	outbuf = make([]byte, 1024*10)
	sh = (*SliceHeader)(unsafe.Pointer(&outbuf))
	ptr = sh.Data
}

var frames uint64

func Exec(s string) {
	n := copy(outbuf, s)
	exec(ptr, n)
}

func exec(ptr uintptr, len int)

func Set(k string, s string) {
	kh := (*StringHeader)(unsafe.Pointer(&k))
	vh := (*StringHeader)(unsafe.Pointer(&s))
	setstring(kh.Data, kh.Len, vh.Data, vh.Len)
}

func setstring(uintptr, int, uintptr, int)

func Buffer(size int) uintptr {
	buf := make([]byte, size)
	return uintptr(unsafe.Pointer(&buf))
}

func requestAnimationFrame()

var r, g, b float32

func clamp(f float32) float32 {
	i := int(f)
	return f - float32(i)
}

type StringHeader struct {
	Data uintptr
	Len  int
}

func ToString(s *string) {
	sh := (*StringHeader)(unsafe.Pointer(s))
	toString(sh.Data, sh.Len)
}

func toString(ptr uintptr, size int)

//go:export frame
func frame(f int) {
	frames++
	if frames == 1 {
		Exec(
			"var fragshader = `" + fshader + "`\n" +
				"var vshader = `" + vshader + "`\n" +
				`  
		var canvas = document.getElementById("derp");
		var rect = canvas.getBoundingClientRect();
		canvas.addEventListener("click",function(e) {
			e.preventDefault();
			wasm.exports.mouseclick(e.clientX-rect.x,e.clientY-rect.y)
		});
		window.gl = canvas.getContext("webgl");
		gl.enable(gl.BLEND);
		gl.blendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA);
		var vs = gl.createShader(gl.VERTEX_SHADER);
		gl.shaderSource(vs, vshader);
		gl.compileShader(vs);
		console.log(gl.getShaderInfoLog(vs));
		var fs = gl.createShader(gl.FRAGMENT_SHADER);
		gl.shaderSource(fs, fragshader);
		gl.compileShader(fs)
		console.log(gl.getShaderInfoLog(fs));
		var prog = gl.createProgram();
		gl.attachShader(prog, vs);
		gl.attachShader(prog, fs);
		gl.linkProgram(prog);
		console.log(gl.getProgramInfoLog(prog));
		gl.useProgram(prog);
		var vbo = gl.createBuffer();
		gl.bindBuffer(gl.ARRAY_BUFFER, vbo);
		var vertattrib = gl.getAttribLocation(prog, "pos");
		console.log(vertattrib);
		gl.enableVertexAttribArray(vertattrib);
		gl.vertexAttribPointer(vertattrib, 2, gl.FLOAT, false, 4 * 6, 0)
		var colattrib = gl.getAttribLocation(prog, "uv")
		console.log(colattrib);
		gl.enableVertexAttribArray(colattrib);
		gl.vertexAttribPointer(colattrib, 4, gl.FLOAT, false, 4 * 6, 4 * 2)
		gl.clearColor(0, 0, 0, 1);
		gl.clear(gl.COLOR_BUFFER_BIT);
		`)
		requestAnimationFrame()
		return
	}
	r += 0.1
	g += 0.01
	b += 0.05
	ClearColor(clamp(r), clamp(g), clamp(b), 1.0)
	// Clear()
	update(tptr, 6*4*3)
	requestAnimationFrame()
}

func Clear()

func ClearColor(r, g, b, a float32)

//go:export mousemove
func mousemove(x, y int) {
	// println("mousemove")
	// println(x)
	// println(y)
}

//go:export mouseclick
func mouseclick(x, y, btn int) {
	println("mouseclick")
	x -= 100
	y -= 100
	dx := float32(x) / 100.0
	triangle.A.X = -0.5 + dx
	triangle.B.X = -0.5 + dx
	triangle.C.X = 0.5 + dx
	triangle.A.Y = 0.5
	triangle.B.Y = -0.5
	triangle.C.Y = -0.5
}

//go:export add
func add(a, b int) int {
	return a + b
}

type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

//go:export derp
func derp(size int) uintptr {
	s := make([]byte, size)
	sh := (*SliceHeader)(unsafe.Pointer(&s))
	return sh.Data
}

const vshader = `
#ifdef GL_ES
precision mediump float;
#endif
attribute vec2 pos;
attribute vec4 uv;

varying vec4 v_uv;

void main() {
v_uv = uv;
gl_Position = vec4(pos,1.0, 1.0);
// gl_Position = projection * vec4(pos,1.0, 1.0);
}
`

const fshader = `
#ifdef GL_ES
#ifdef GL_FRAGMENT_PRECISION_HIGH
	precision highp float;
#else
	precision mediump float;
#endif
#endif
varying vec4 v_uv;
void main() {
gl_FragColor = v_uv;
}
`
