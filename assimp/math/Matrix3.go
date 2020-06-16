package math

type Matrix3 struct {
	Elements [9]float64
}

func NewMatrix3() *Matrix3 {
	elements := [9]float64{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1}
	return &Matrix3{
		Elements: elements,
	}
}
func (this *Matrix3) Set(n11, n12, n13, n21, n22, n23, n31, n32, n33 float64) *Matrix3 {
	 te := &this.Elements

	te[ 0 ] = n11
	te[ 1 ] = n21
	te[ 2 ] = n31

	te[ 3 ] = n12
	te[ 4 ] = n22
	te[ 5 ] = n32

	te[ 6 ] = n13
	te[ 7 ] = n23
	te[ 8 ] = n33

	return this

}

func (this *Matrix3) SetFromMatrix4(m *Matrix4) *Matrix3 {
	me := m.Elements
	this.Set(
		me[0], me[4], me[8],
		me[1], me[5], me[9],
		me[2], me[6], me[10],
	)

	return this

}

func (this *Matrix3) Scale(sx, sy float64) *Matrix3 {
	te := &this.Elements

	te[0] *= sx
	te[3] *= sx
	te[6] *= sx
	te[1] *= sy
	te[4] *= sy
	te[7] *= sy

	return this
}

func (this *Matrix3) Determinant() float64 {
	te := this.Elements

	a := te[0]
	b := te[1]
	c := te[2]
	d := te[3]
	e := te[4]
	f := te[5]
	g := te[6]
	h := te[7]
	i := te[8]

	return a*e*i - a*f*h - b*d*i + b*f*g + c*d*h - c*e*g

}

func (this *Matrix3) MultiplyScalar(s float64) *Matrix3 {

	te := &this.Elements

	te[0] *= s
	te[3] *= s
	te[6] *= s
	te[1] *= s
	te[4] *= s
	te[7] *= s
	te[2] *= s
	te[5] *= s
	te[8] *= s

	return this

}

func (this *Matrix3) multiplyMatrices(a, b *Matrix3) *Matrix3 {

	ae := a.Elements
	be := b.Elements
	te := &this.Elements

	a11 := ae[0]
	a12 := ae[3]
	a13 := ae[6]

	a21 := ae[1]
	a22 := ae[4]
	a23 := ae[7]

	a31 := ae[2]
	a32 := ae[5]
	a33 := ae[8]

	b11 := be[0]
	b12 := be[3]
	b13 := be[6]

	b21 := be[1]
	b22 := be[4]
	b23 := be[7]
	b31 := be[2]
	b32 := be[5]
	b33 := be[8]

	te[0] = a11*b11 + a12*b21 + a13*b31
	te[3] = a11*b12 + a12*b22 + a13*b32
	te[6] = a11*b13 + a12*b23 + a13*b33

	te[1] = a21*b11 + a22*b21 + a23*b31
	te[4] = a21*b12 + a22*b22 + a23*b32
	te[7] = a21*b13 + a22*b23 + a23*b33

	te[2] = a31*b11 + a32*b21 + a33*b31
	te[5] = a31*b12 + a32*b22 + a33*b32
	te[8] = a31*b13 + a32*b23 + a33*b33

	return this

}


// DotPruduct 矩阵点乘
func (this *Matrix3)DotVector3(a *Vector3)*Vector3 {
	result := NewVector3(0,0,0)

	result.X = a.X * this.Elements[0] + a.Y * this.Elements[1] + a.Z * this.Elements[2]
	result.Y = a.X * this.Elements[3] + a.Y * this.Elements[4] + a.Z * this.Elements[5]
	result.Z = a.X * this.Elements[6] + a.Y * this.Elements[7] + a.Z * this.Elements[8]

	return result
}

