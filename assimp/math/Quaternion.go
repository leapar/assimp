package math

import "math"

type Quaternion struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
	W float64 `json:"w"`
}

func NewQuaternion(x, y, z float64) *Quaternion {
	return &Quaternion{
		x,
		y,
		z,
		1,
	}
}

func NewQuaternion3() *Quaternion {
	return &Quaternion{
		0,
		0,
		0,
		1,
	}
}

func NewQuaternion2(x, y, z, w float64) *Quaternion {
	return &Quaternion{
		x,
		y,
		z,
		w,
	}
}

//setFromUnitVectors
func (this *Quaternion) SetFromUnitVectors(vFrom, vTo *Vector3) *Quaternion {
	var v1 = NewVector3(0, 0, 0)
	var r float64
	var EPS = 0.000001

	r = vFrom.Dot(vTo) + 1
	if r < EPS {
		r = 0
		if math.Abs(vFrom.X) > math.Abs(vFrom.Z) {
			v1 = NewVector3(- vFrom.Y, vFrom.X, 0)
		} else {
			v1 = NewVector3(0, - vFrom.Z, vFrom.Y)
		}
	} else {
		v1.CrossVectors(vFrom, vTo)
	}

	this.X = v1.X
	this.Y = v1.Y
	this.Z = v1.Z
	this.W = r

	return this.Normalize()
}

func (this *Quaternion) Length() float64 {
	return math.Sqrt(this.X*this.X + this.Y*this.Y + this.Z*this.Z + this.W*this.W)
}

func (this *Quaternion) Normalize() *Quaternion {
	var l = this.Length()
	if l == 0.0 {
		this.X = 0
		this.Y = 0
		this.Z = 0
		this.W = 1
	} else {
		l = 1 / l
		this.X = this.X * l
		this.Y = this.Y * l
		this.Z = this.Z * l
		this.W = this.W * l
	}

	return this
}

//THREE.Math.radToDeg(v1.angle())
func (this *Quaternion) SetFromEuler(euler *Euler)*Quaternion {

	var x float64 = euler.X
	var y float64 = euler.Y
	var z float64 = euler.Z
	var order string = euler.Order

	// http://www.mathworks.com/matlabcentral/fileexchange/
	// 	20696-function-to-convert-between-dcm-euler-angles-quaternions-and-euler-vectors/
	//	content/SpinCalc.m

	var cos = math.Cos
	var sin = math.Sin
	var c1 float64 = cos(x / 2)
	var c2 float64 = cos(y / 2)
	var c3 float64 = cos(z / 2)

	var s1 float64 = sin(x / 2)
	var s2 float64 = sin(y / 2)
	var s3 float64 = sin(z / 2)
	switch order {
	case "XYZ":
		this.X = s1*c2*c3 + c1*s2*s3
		this.Y = c1*s2*c3 - s1*c2*s3
		this.Z = c1*c2*s3 + s1*s2*c3
		this.W = c1*c2*c3 - s1*s2*s3
	case "YXZ":
		this.X = s1*c2*c3 + c1*s2*s3
		this.Y = c1*s2*c3 - s1*c2*s3
		this.Z = c1*c2*s3 - s1*s2*c3
		this.W = c1*c2*c3 + s1*s2*s3
	case "ZXY":
		this.X = s1*c2*c3 - c1*s2*s3
		this.Y = c1*s2*c3 + s1*c2*s3
		this.Z = c1*c2*s3 + s1*s2*c3
		this.W = c1*c2*c3 - s1*s2*s3
	case "ZYX":
		this.X = s1*c2*c3 - c1*s2*s3
		this.Y = c1*s2*c3 + s1*c2*s3
		this.Z = c1*c2*s3 - s1*s2*c3
		this.W = c1*c2*c3 + s1*s2*s3
	case "YZX":
		this.X = s1*c2*c3 + c1*s2*s3
		this.Y = c1*s2*c3 + s1*c2*s3
		this.Z = c1*c2*s3 - s1*s2*c3
		this.W = c1*c2*c3 - s1*s2*s3
	case "XZY":
		this.X = s1*c2*c3 - c1*s2*s3
		this.Y = c1*s2*c3 - s1*c2*s3
		this.Z = c1*c2*s3 + s1*s2*c3
		this.W = c1*c2*c3 + s1*s2*s3
	}

	return this
}

func (this *Quaternion) SetFromRotationMatrix(m *Matrix4) *Quaternion {

	// http://www.euclideanspace.com/maths/geometry/rotations/conversions/matrixToQuaternion/index.htm

	// assumes the upper 3x3 of m is a pure rotation matrix (i.e, unscaled)

	var te = m.Elements

	var m11 = te[ 0 ]
	var m12 = te[ 4 ]
	var m13 = te[ 8 ]
	var m21 = te[ 1 ]
	var m22 = te[ 5 ]
	var m23 = te[ 9 ]
	var m31 = te[ 2 ]
	var m32 = te[ 6 ]
	var m33 = te[ 10 ]

	var trace = m11 + m22 + m33
	var s float64

	if  trace > 0  {

		s = 0.5 / math.Sqrt(trace+1.0)

		this.W = 0.25 / s
		this.X = ( m32 - m23 ) * s
		this.Y = ( m13 - m31 ) * s
		this.Z = ( m21 - m12 ) * s

	} else if  m11 > m22 && m11 > m33  {

		s = 2.0 * math.Sqrt(1.0+m11-m22-m33)

		this.W = ( m32 - m23 ) / s
		this.X = 0.25 * s
		this.Y = ( m12 + m21 ) / s
		this.Z = ( m13 + m31 ) / s

	} else if  m22 > m33  {

		s = 2.0 * math.Sqrt(1.0+m22-m11-m33)

		this.W = ( m13 - m31 ) / s
		this.X = ( m12 + m21 ) / s
		this.Y = 0.25 * s
		this.Z = ( m23 + m32 ) / s

	} else {

		s = 2.0 * math.Sqrt(1.0+m33-m11-m22)

		this.W = ( m21 - m12 ) / s
		this.X = ( m13 + m31 ) / s
		this.Y = ( m23 + m32 ) / s
		this.Z = 0.25 * s

	}

	return this

}
