package cc

import (
	"github.com/gopherjs/gopherjs/js"
)

type WebGLUniformLocation interface{}
type WebGLProgram interface{}
type webGLUniformLocation struct{ *js.Object }
type webGLProgram struct{ *js.Object }

// The Component class
type GLProgram interface {
	DestroyProgram()
	InitWithVertexShaderByteArray(string, string) bool
	InitWithString(string, string) bool
	InitWithVertexShaderFilename(string, string) bool
	Init(string, string) bool
	AddAttribute(string, int)
	Link() bool
	Use()
	UpdateUniforms()
	GetUniformLocationForName(string) int
	GetUniformMVPMatrix() WebGLUniformLocation
	GetUniformSampler() WebGLUniformLocation
	SetUniformLocationWith1i(WebGLUniformLocation, int)
	SetUniformLocationWith2i(WebGLUniformLocation, int, int)
	SetUniformLocationWith3i(WebGLUniformLocation, int, int, int)
	SetUniformLocationWith4i(WebGLUniformLocation, int, int, int, int)
	SetUniformLocationWith2iv(WebGLUniformLocation, []int, int)
	SetUniformLocationWith3iv(WebGLUniformLocation, []int, int)
	SetUniformLocationWith4iv(WebGLUniformLocation, []int, int)
	SetUniformLocationI32(...interface{})
	SetUniformLocationWith1f(WebGLUniformLocation, float64)
	SetUniformLocationWith2f(WebGLUniformLocation, float64, float64)
	SetUniformLocationWith3f(WebGLUniformLocation, float64, float64, float64)
	SetUniformLocationWith4f(WebGLUniformLocation, float64, float64, float64, float64)
	SetUniformLocationWith2fv(WebGLUniformLocation, []float64, int)
	SetUniformLocationWith3fv(WebGLUniformLocation, []float64, int)
	SetUniformLocationWith4fv(WebGLUniformLocation, []float64, int)
	SetUniformLocationWithMatrix4fv(WebGLUniformLocation, []float64, int)
	SetUniformLocationF32(...interface{})
	SetUniformsForBuiltins()
	SetUniformForModelViewProjectionMatrix()
	SetUniformForModelViewProjectionMatrixWithMat4([]float64)
	SetUniformForModelViewAndProjectionMatrixWithMat4()
	VertexShaderLog() string
	GetVertexShaderLog() string
	GetFragmentShaderLog() string
	FragmentShaderLog() string
	ProgramLog() string
	GetProgramLog() string
	Reset()
	GetProgram() WebGLProgram
	Retain()
	Release()
}

type glProgram struct{ *js.Object }

// NewComponent is the constructor for Component.
// glContext could be nil ; in this case, the global context will be used
func NewGlProgram(vShaderFileName string, fShaderFileName string, glContext *js.Object) GLProgram {
	return &glProgram{pcc.Get("GLProgram").New()}
}

// DestroyProgram destroys the program
func (gp *glProgram) DestroyProgram() {
	gp.Call("destroyProgram")
}

// InitWithVertexShaderByteArray initializes the GLProgram with a vertex and fragment with string
func (gp *glProgram) InitWithVertexShaderByteArray(vertShaderStr string, fragShaderStr string) bool {
	return gp.Call("initWithVertexShaderByteArray", vertShaderStr, fragShaderStr).Bool()
}

// InitWithString initializes the GLProgram with a vertex and fragment with string
func (gp *glProgram) InitWithString(vertShaderStr string, fragShaderStr string) bool {
	return gp.Call("initWithString", vertShaderStr, fragShaderStr).Bool()
}

// InitWithVertexShaderFilename initializes the GLProgram with a vertex and fragment with contents of filenames
func (gp *glProgram) InitWithVertexShaderFilename(vShaderFilename string, fShaderFileName string) bool {
	return gp.Call("initWithVertexShaderFilename", vShaderFilename, fShaderFileName).Bool()
}

// Init initializes the GLProgram with a vertex and fragment with contents of filenames
func (gp *glProgram) Init(vShaderFilename string, fShaderFileName string) bool {
	return gp.Call("init", vShaderFilename, fShaderFileName).Bool()
}

// AddAttribute adds a new attribute to the shader
func (gp *glProgram) AddAttribute(attributeName string, index int) {
	gp.Call("addAttribute", attributeName, index)
}

// Link links the glProgram
func (gp *glProgram) Link() bool {
	return gp.Call("link").Bool()
}

// Use calls glUseProgram()
func (gp *glProgram) Use() {
	gp.Call("use")
}

// UpdateUniforms create 4 uniforms:
//  cc.UNIFORM_PMATRIX
//  cc.UNIFORM_MVMATRIX
//  cc.UNIFORM_MVPMATRIX
//  cc.UNIFORM_SAMPLER
func (gp *glProgram) UpdateUniforms() {
	gp.Call("updateUniforms")
}

// GetUniformLocationForName retrieves the named uniform location for this shader program.
func (gp *glProgram) GetUniformLocationForName(name string) int {
	return gp.Call("getUniformLocationForName", name).Int()
}

// GetUniformMVPMatrix retrieves the uniform MVP matrix
func (gp *glProgram) GetUniformMVPMatrix() WebGLUniformLocation {
	return &webGLUniformLocation{gp.Call("getUniformMVPMatrix")}
}

// GetUniformSampler retrieves the uniform sampler
func (gp *glProgram) GetUniformSampler() WebGLUniformLocation {
	return &webGLUniformLocation{gp.Call("getUniformSampler")}
}

// SetUniformLocationWith1i calls glUniform1i only if the values are different than the previous call for this same shader program.
func (gp *glProgram) SetUniformLocationWith1i(location WebGLUniformLocation, i1 int) {
	gp.Call("setUniformLocationWith1i", location, i1)
}

// SetUniformLocationWith2i calls glUniform2i only if the values are different than the previous call for this same shader program.
func (gp *glProgram) SetUniformLocationWith2i(location WebGLUniformLocation, i1 int, i2 int) {
	gp.Call("setUniformLocationWith2i", location, i1, i2)
}

// SetUniformLocationWith3i calls glUniform3i only if the values are different than the previous call for this same shader program.
func (gp *glProgram) SetUniformLocationWith3i(location WebGLUniformLocation, i1 int, i2 int, i3 int) {
	gp.Call("setUniformLocationWith3i", location, i1, i2, i3)
}

// SetUniformLocationWith4i calls glUniform4i only if the values are different than the previous call for this same shader program.
func (gp *glProgram) SetUniformLocationWith4i(location WebGLUniformLocation, i1 int, i2 int, i3 int, i4 int) {
	gp.Call("setUniformLocationWith4i", location, i1, i2, i3, i4)
}

// SetUniformLocationWith2iv calls glUniform2iv only if the values are different than the previous call for this same shader program.
func (gp *glProgram) SetUniformLocationWith2iv(location WebGLUniformLocation, intArray []int, numberOfArrays int) {
	gp.Call("setUniformLocationWith2iv", location, intArray, numberOfArrays)
}

// SetUniformLocationWith3iv calls glUniform3iv only if the values are different than the previous call for this same shader program.
func (gp *glProgram) SetUniformLocationWith3iv(location WebGLUniformLocation, intArray []int, numberOfArrays int) {
	gp.Call("setUniformLocationWith3iv", location, intArray, numberOfArrays)
}

// SetUniformLocationWith4iv calls glUniform4iv only if the values are different than the previous call for this same shader program.
func (gp *glProgram) SetUniformLocationWith4iv(location WebGLUniformLocation, intArray []int, numberOfArrays int) {
	gp.Call("setUniformLocationWith4iv", location, intArray, numberOfArrays)
}

// SetUniformLocationI32 calls glUniform1i only if the values are different than the previous call for this same shader program.
func (gp *glProgram) SetUniformLocationI32(args ...interface{}) {
	gp.Call("setUniformLocationI32", args)
}

// SetUniformLocationWith1f calls glUniform1f only if the values are different than the previous call for this same shader program.
func (gp *glProgram) SetUniformLocationWith1f(location WebGLUniformLocation, f1 float64) {
	gp.Call("setUniformLocationWith1f", location, f1)
}

// SetUniformLocationWith2f calls glUniform2f only if the values are different than the previous call for this same shader program.
func (gp *glProgram) SetUniformLocationWith2f(location WebGLUniformLocation, f1 float64, f2 float64) {
	gp.Call("setUniformLocationWith2f", location, f1, f2)
}

// SetUniformLocationWith3f calls glUniform3f only if the values are different than the previous call for this same shader program.
func (gp *glProgram) SetUniformLocationWith3f(location WebGLUniformLocation, f1 float64, f2 float64, f3 float64) {
	gp.Call("setUniformLocationWith3f", location, f1, f2, f3)
}

// SetUniformLocationWith4f calls glUniform4f only if the values are different than the previous call for this same shader program.
func (gp *glProgram) SetUniformLocationWith4f(location WebGLUniformLocation, f1 float64, f2 float64, f3 float64, f4 float64) {
	gp.Call("setUniformLocationWith4f", location, f1, f2, f3, f4)
}

// SetUniformLocationWith2fv calls glUniform2fv only if the values are different than the previous call for this same shader program.
func (gp *glProgram) SetUniformLocationWith2fv(location WebGLUniformLocation, floatArray []float64, numberOfArrays int) {
	gp.Call("setUniformLocationWith2fv", location, floatArray, numberOfArrays)
}

// SetUniformLocationWith3fv calls glUniform3fv only if the values are different than the previous call for this same shader program.
func (gp *glProgram) SetUniformLocationWith3fv(location WebGLUniformLocation, floatArray []float64, numberOfArrays int) {
	gp.Call("setUniformLocationWith3fv", location, floatArray, numberOfArrays)
}

// SetUniformLocationWith4fv calls glUniform4fv only if the values are different than the previous call for this same shader program.
func (gp *glProgram) SetUniformLocationWith4fv(location WebGLUniformLocation, floatArray []float64, numberOfArrays int) {
	gp.Call("setUniformLocationWith4fv", location, floatArray, numberOfArrays)
}

// SetUniformLocationWithMatrix4fv calls glUniformMatrix4fv only if the values are different than the previous call for this same shader program.
func (gp *glProgram) SetUniformLocationWithMatrix4fv(location WebGLUniformLocation, matrixArray []float64, numberOfMatrices int) {
	gp.Call("setUniformLocationWith2fv", location, matrixArray, numberOfMatrices)
}

func (gp *glProgram) SetUniformLocationF32(args ...interface{}) {
	gp.Call("setUniformLocationF32", args)
}

// SetUniformsForBuiltins updates the builtin uniforms if they are different than the previous call for this same shader program.
func (gp *glProgram) SetUniformsForBuiltins() {
	gp.Call("setUniformsForBuiltins")
}

// SetUniformForModelViewProjectionMatrix updates the MVP matrix on the MVP uniform if it is different than the previous call for this same shader program.
func (gp *glProgram) SetUniformForModelViewProjectionMatrix() {
	gp.Call("setUniformForModelViewProjectionMatrix")
}

func (gp *glProgram) SetUniformForModelViewProjectionMatrixWithMat4(swapMat4 []float64) {
	gp.Call("setUniformForModelViewProjectionMatrixWithMat4", swapMat4)
}

func (gp *glProgram) SetUniformForModelViewAndProjectionMatrixWithMat4() {
	gp.Call("setUniformForModelViewAndProjectionMatrixWithMat4")
}

// VertexShaderLog returns the vertexShader error log
func (gp *glProgram) VertexShaderLog() string {
	return gp.Call("vertexShaderLog").String()
}

// GetVertexShaderLog returns the vertexShader error log
func (gp *glProgram) GetVertexShaderLog() string {
	return gp.Call("getVertexShaderLog").String()
}

// GetFragmentShaderLog returns the fragmentShader error log
func (gp *glProgram) GetFragmentShaderLog() string {
	return gp.Call("getFragmentShaderLog").String()
}

// FragmentShaderLog returns the fragmentShader error log
func (gp *glProgram) FragmentShaderLog() string {
	return gp.Call("fragmentShaderLog").String()
}

// ProgramLog returns the program error log
func (gp *glProgram) ProgramLog() string {
	return gp.Call("programLog").String()
}

// GetProgramLog returns the program error log
func (gp *glProgram) GetProgramLog() string {
	return gp.Call("getProgramLog").String()
}

// Reset reloads all shaders.
// This function is designed for android when opengl context lost, so don't call it
func (gp *glProgram) Reset() {
	gp.Call("reset")
}

// GetProgram retrieves the WebGLProgram object
func (gp *glProgram) GetProgram() WebGLProgram {
	return &webGLProgram{gp.Call("getProgramLog")}
}

func (gp *glProgram) Retain() {
	gp.Call("retain")
}

func (gp *glProgram) Release() {
	gp.Call("release")
}
