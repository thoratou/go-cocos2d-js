package cc

/////////////////
// V3F_C4B_T2F //
/////////////////

type V3F_C4B_T2F interface{}

func NewV3F_C4B_T2F(vertices Vertex3F,
	colors Color,
	texCoords Tex2F,
	arrayBuffer interface{},
	offset int) V3F_C4B_T2F {
	return pcc.Call("V3F_C4B_T2F", vertices, colors, texCoords, arrayBuffer, offset)
}

const (
	V3F_C4B_T2F_BYTES_PER_ELEMENT = 24
)

//////////////////////
// V3F_C4B_T2F_Quad //
//////////////////////

type V3F_C4B_T2F_Quad interface{}

func NewV3F_C4B_T2F_Quad(tl V3F_C4B_T2F,
	bl V3F_C4B_T2F,
	tr V3F_C4B_T2F,
	br V3F_C4B_T2F,
	arrayBuffer interface{},
	offset int) V3F_C4B_T2F_Quad {
	return pcc.Call("V3F_C4B_T2F_Quad", tl, bl, tr, br, arrayBuffer, offset)
}

const (
	V3F_C4B_T2F_Quad_BYTES_PER_ELEMENT = 96
)
