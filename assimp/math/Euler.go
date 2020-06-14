package math

import "math"

type Euler struct {
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	Z     float64 `json:"z"`
	Order string //'XYZ', 'YZX', 'ZXY', 'XZY', 'YXZ', 'ZYX'
}

const (
	DEG2RAD = math.Pi / 180
	RAD2DEG = 180 / math.Pi
)

func DegToRad(degrees float64) float64 {

	return degrees * DEG2RAD

}

func RadToDeg(radians float64) float64 {

	return radians * RAD2DEG

}

func clamp(value, min, max float64) float64 {
	return math.Max(min, math.Min(max, value))
}

func NewEuler(x, y, z float64) *Euler {
	return &Euler{
		x,
		y,
		z,
		"XYZ",
	}
}

func (this *Euler) SetFromQuaternion(q *Quaternion) *Euler {
	var matrix = NewMatrix4()
	matrix.MakeRotationFromQuaternion(q)
	return this.setFromRotationMatrix(matrix, "XYZ")
}

func (this *Euler) setFromRotationMatrix(m *Matrix4, order string) *Euler {

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

	if order == "XYZ" {
		this.Y = math.Asin(clamp(m13, - 1, 1))
		if math.Abs(m13) < 0.99999 {
			this.X = math.Atan2(- m23, m33)
			this.Z = math.Atan2(- m12, m11)
		} else {
			this.X = math.Atan2(m32, m22)
			this.Z = 0
		}
	} else if order == "YXZ" {
		this.X = math.Asin(- clamp(m23, - 1, 1))
		if math.Abs(m23) < 0.99999 {
			this.Y = math.Atan2(m13, m33)
			this.Z = math.Atan2(m21, m22)
		} else {
			this.Y = math.Atan2(- m31, m11)
			this.Z = 0
		}
	} else if order == "ZXY" {
		this.X = math.Asin(clamp(m32, - 1, 1))
		if math.Abs(m32) < 0.99999 {
			this.Y = math.Atan2(- m31, m33)
			this.Z = math.Atan2(- m12, m22)
		} else {
			this.Y = 0
			this.Z = math.Atan2(m21, m11)
		}
	} else if order == "ZYX" {
		this.Y = math.Asin(- clamp(m31, - 1, 1))
		if math.Abs(m31) < 0.99999 {
			this.X = math.Atan2(m32, m33)
			this.Z = math.Atan2(m21, m11)
		} else {
			this.X = 0
			this.Z = math.Atan2(- m12, m22)
		}
	} else if order == "YZX" {
		this.Z = math.Asin(clamp(m21, - 1, 1))
		if math.Abs(m21) < 0.99999 {
			this.X = math.Atan2(- m23, m22)
			this.Y = math.Atan2(- m31, m11)
		} else {
			this.X = 0
			this.Y = math.Atan2(m13, m33)
		}
	} else if order == "XZY" {
		this.Z = math.Asin(- clamp(m12, - 1, 1))
		if math.Abs(m12) < 0.99999 {
			this.X = math.Atan2(m32, m22)
			this.Y = math.Atan2(m13, m11)
		} else {
			this.X = math.Atan2(- m23, m33)
			this.Y = 0
		}
	} else {
		//console.warn("THREE.Euler:.setFromRotationMatrix() given unsupported order: " + order)
	}

	this.Order = order

	return this
}
