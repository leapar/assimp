package math


type Matrix4 struct {
	Elements [16]float64
}

func NewMatrix4() *Matrix4 {
	elements := [16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}
	return &Matrix4{
		Elements: elements,
	}
}

func (this *Matrix4) Compose(position *Vector3, quaternion *Quaternion, scale *Vector3) {
	if scale == nil {
		scale = NewVector3(1, 1, 1)
	}
	this.MakeRotationFromQuaternion(quaternion)
	this.Scale(scale)
	this.SetPosition(position)
}

func (this *Matrix4) Scale(v *Vector3) {
	var te = this.Elements
	var x = v.X
	var y = v.Y
	var z = v.Z
	te[ 0 ] *= x
	te[ 4 ] *= y
	te[ 8 ] *= z
	te[ 1 ] *= x
	te[ 5 ] *= y
	te[ 9 ] *= z
	te[ 2 ] *= x
	te[ 6 ] *= y
	te[ 10 ] *= z
	te[ 3 ] *= x
	te[ 7 ] *= y
	te[ 11 ] *= z

	this.Elements = te
}

func (this *Matrix4) SetPosition(v *Vector3) {
	var te = this.Elements
	te[ 12 ] = v.X
	te[ 13 ] = v.Y
	te[ 14 ] = v.Z

	this.Elements = te
}

func (this *Matrix4) MakeRotationFromQuaternion(q *Quaternion) *Matrix4{
	var te = this.Elements

	var x = q.X
	var y = q.Y
	var z = q.Z
	var w = q.W
	var x2 = x + x
	var y2 = y + y
	var z2 = z + z
	var xx = x * x2
	var xy = x * y2
	var xz = x * z2
	var yy = y * y2
	var yz = y * z2
	var zz = z * z2
	var wx = w * x2
	var wy = w * y2
	var wz = w * z2

	te[ 0 ] = 1 - ( yy + zz )
	te[ 4 ] = xy - wz
	te[ 8 ] = xz + wy
	te[ 1 ] = xy + wz
	te[ 5 ] = 1 - ( xx + zz )
	te[ 9 ] = yz - wx
	te[ 2 ] = xz - wy
	te[ 6 ] = yz + wx
	te[ 10 ] = 1 - ( xx + yy )

	// last column
	te[ 3 ] = 0
	te[ 7 ] = 0
	te[ 11 ] = 0

	// bottom row
	te[ 12 ] = 0
	te[ 13 ] = 0
	te[ 14 ] = 0
	te[ 15 ] = 1

	this.Elements = te

	return this
}

func (this *Matrix4) MultiplyMatrices(a *Matrix4) *Matrix4 {

	var ae = this.Elements
	var be = a.Elements
	var te = this.Elements

	var a11 = ae[ 0 ]
	var a12 = ae[ 4 ]
	var a13 = ae[ 8 ]
	var a14 = ae[ 12 ]
	var a21 = ae[ 1 ]
	var a22 = ae[ 5 ]
	var a23 = ae[ 9 ]
	var a24 = ae[ 13 ]
	var a31 = ae[ 2 ]
	var a32 = ae[ 6 ]
	var a33 = ae[ 10 ]
	var a34 = ae[ 14 ]
	var a41 = ae[ 3 ]
	var a42 = ae[ 7 ]
	var a43 = ae[ 11 ]
	var a44 = ae[ 15 ]

	var b11 = be[ 0 ]
	var b12 = be[ 4 ]
	var b13 = be[ 8 ]
	var b14 = be[ 12 ]
	var b21 = be[ 1 ]
	var b22 = be[ 5 ]
	var b23 = be[ 9 ]
	var b24 = be[ 13 ]
	var b31 = be[ 2 ]
	var b32 = be[ 6 ]
	var b33 = be[ 10 ]
	var b34 = be[ 14 ]
	var b41 = be[ 3 ]
	var b42 = be[ 7 ]
	var b43 = be[ 11 ]
	var b44 = be[ 15 ]

	te[ 0 ] = a11*b11 + a12*b21 + a13*b31 + a14*b41
	te[ 4 ] = a11*b12 + a12*b22 + a13*b32 + a14*b42
	te[ 8 ] = a11*b13 + a12*b23 + a13*b33 + a14*b43
	te[ 12 ] = a11*b14 + a12*b24 + a13*b34 + a14*b44

	te[ 1 ] = a21*b11 + a22*b21 + a23*b31 + a24*b41
	te[ 5 ] = a21*b12 + a22*b22 + a23*b32 + a24*b42
	te[ 9 ] = a21*b13 + a22*b23 + a23*b33 + a24*b43
	te[ 13 ] = a21*b14 + a22*b24 + a23*b34 + a24*b44

	te[ 2 ] = a31*b11 + a32*b21 + a33*b31 + a34*b41
	te[ 6 ] = a31*b12 + a32*b22 + a33*b32 + a34*b42
	te[ 10 ] = a31*b13 + a32*b23 + a33*b33 + a34*b43
	te[ 14 ] = a31*b14 + a32*b24 + a33*b34 + a34*b44

	te[ 3 ] = a41*b11 + a42*b21 + a43*b31 + a44*b41
	te[ 7 ] = a41*b12 + a42*b22 + a43*b32 + a44*b42
	te[ 11 ] = a41*b13 + a42*b23 + a43*b33 + a44*b43
	te[ 15 ] = a41*b14 + a42*b24 + a43*b34 + a44*b44

	this.Elements = te
	return this

}

func (this *Matrix4) Determinant() float64 {

	var te = this.Elements

	var n11 = te[ 0 ]
	var n12 = te[ 4 ]
	var n13 = te[ 8 ]
	var n14 = te[ 12 ]
	var n21 = te[ 1 ]
	var n22 = te[ 5 ]
	var n23 = te[ 9 ]
	var n24 = te[ 13 ]
	var n31 = te[ 2 ]
	var n32 = te[ 6 ]
	var n33 = te[ 10 ]
	var n34 = te[ 14 ]
	var n41 = te[ 3 ]
	var n42 = te[ 7 ]
	var n43 = te[ 11 ]
	var n44 = te[ 15 ]

	//TODO: make this more efficient
	//( based on http://www.euclideanspace.com/maths/algebra/matrix/functions/inverse/fourD/index.htm )

	return n41*(n14*n23*n32-n13*n24*n32-n14*n22*n33+n12*n24*n33+n13*n22*n34-n12*n23*n34) +
		n42*(+ n11*n23*n34-n11*n24*n33+n14*n21*n33-n13*n21*n34+n13*n24*n31-n14*n23*n31) +
		n43*(+ n11*n24*n32-n11*n22*n34-n14*n21*n32+n12*n21*n34+n14*n22*n31-n12*n24*n31) +
		n44*(- n13*n22*n31-n11*n23*n32+n11*n22*n33+n13*n21*n32-n12*n21*n33+n12*n23*n31)

}

func (this *Matrix4) Copy(m *Matrix4) *Matrix4 {

	var te = this.Elements
	var me = m.Elements
	te[ 0 ] = me[ 0 ]
	te[ 1 ] = me[ 1 ]
	te[ 2 ] = me[ 2 ]
	te[ 3 ] = me[ 3 ]
	te[ 4 ] = me[ 4 ]
	te[ 5 ] = me[ 5 ]
	te[ 6 ] = me[ 6 ]
	te[ 7 ] = me[ 7 ]
	te[ 8 ] = me[ 8 ]
	te[ 9 ] = me[ 9 ]
	te[ 10 ] = me[ 10 ]
	te[ 11 ] = me[ 11 ]
	te[ 12 ] = me[ 12 ]
	te[ 13 ] = me[ 13 ]
	te[ 14 ] = me[ 14 ]
	te[ 15 ] = me[ 15 ]
	this.Elements = te
	return this
}

func (this *Matrix4) Decompose(position *Vector3, quaternion *Quaternion, scale *Vector3) *Quaternion{

	var te = this.Elements
	var v1 = Vector3{te[ 0 ], te[ 1 ], te[ 2 ]}
	var sx = v1.Length()
	var v2 = Vector3{te[ 4 ], te[ 5 ], te[ 6 ]}
	var sy = v2.Length()
	var v3 = Vector3{te[ 8 ], te[ 9 ], te[ 10 ]}
	var sz = v3.Length()

	// if determine is negative, we need to invert one scale
	var det = this.Determinant()
	if det < 0 {
		sx = - sx
	}

	if position != nil {
		position.X = te[ 12 ]
		position.Y = te[ 13 ]
		position.Z = te[ 14 ]
	}


	// scale the rotation part
	matrix := NewMatrix4()

	matrix.Copy(this)

	var invSX = 1 / sx
	var invSY = 1 / sy
	var invSZ = 1 / sz

	matrix.Elements[ 0 ] *= invSX
	matrix.Elements[ 1 ] *= invSX
	matrix.Elements[ 2 ] *= invSX

	matrix.Elements[ 4 ] *= invSY
	matrix.Elements[ 5 ] *= invSY
	matrix.Elements[ 6 ] *= invSY

	matrix.Elements[ 8 ] *= invSZ
	matrix.Elements[ 9 ] *= invSZ
	matrix.Elements[ 10 ] *= invSZ

	if quaternion == nil {
		quaternion = NewQuaternion3()
	}
	quaternion.SetFromRotationMatrix(matrix)

	if scale != nil {
		scale.X = sx
		scale.Y = sy
		scale.Z = sz
	}

	return quaternion
}
