package math

import "math"

type Vector3 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

func NewVector3(x, y, z float64) *Vector3 {
	return &Vector3{
		x,
		y,
		z,
	}
}

func (this *Vector3) ApplyQuaternion(q *Quaternion) *Vector3 {

	var x = this.X
	var y = this.Y
	var z = this.Z
	var qx = q.X
	var qy = q.Y
	var qz = q.Z
	var qw = q.W

	// calculate quat * vector

	var ix = qw*x + qy*z - qz*y
	var iy = qw*y + qz*x - qx*z
	var iz = qw*z + qx*y - qy*x
	var iw = - qx*x - qy*y - qz*z

	// calculate result * inverse quat

	this.X = ix*qw + iw * - qx + iy * - qz - iz * - qy
	this.Y = iy*qw + iw * - qy + iz * - qx - ix * - qz
	this.Z = iz*qw + iw * - qz + ix * - qy - iy * - qx

	return this

}

func (this *Vector3) Clone() *Vector3 {
	return NewVector3(this.X, this.Y, this.Z)
}

func (this *Vector3) MultiplyScalar(scalar float64) *Vector3 {

	this.X *= scalar
	this.Y *= scalar
	this.Z *= scalar

	return this

}

func (this *Vector3) Length() float64 {
	return math.Sqrt(this.X*this.X + this.Y*this.Y + this.Z*this.Z)
}

func (this *Vector3) Normalize() *Vector3 {
	if this.Length() == 0 {
		return this.DivideScalar(1)
	} else {
		return this.DivideScalar(this.Length())
	}
}

func (this *Vector3) DivideScalar(scalar float64) *Vector3 {

	return this.MultiplyScalar(1 / scalar)

}

func (this *Vector3) Dot(v *Vector3) float64 {
	return this.X*v.X + this.Y*v.Y + this.Z*v.Z
}

func (this *Vector3) Mul(v *Vector3) *Vector3 {

	this.X *= v.X
	this.Y *= v.Y
	this.Z *= v.Z

	return this

}

func (this *Vector3) Sub(v *Vector3) *Vector3 {

	this.X -= v.X
	this.Y -= v.Y
	this.Z -= v.Z

	return this

}

func (this *Vector3) Add(v *Vector3) *Vector3 {

	this.X += v.X
	this.Y += v.Y
	this.Z += v.Z

	return this

}

func (this *Vector3) CrossVectors(a, b *Vector3) *Vector3 {

	var ax = a.X
	var ay = a.Y
	var az = a.Z
	var bx = b.X
	var by = b.Y
	var bz = b.Z

	this.X = ay*bz - az*by
	this.Y = az*bx - ax*bz
	this.Z = ax*by - ay*bx
	return this
}

func (this *Vector3) ApplyMatrix4(m *Matrix4) *Vector3 {
	var x = this.X
	var y = this.Y
	var z = this.Z
	var e = m.Elements

	var w = 1 / ( e[ 3 ]*x + e[ 7 ]*y + e[ 11 ]*z + e[ 15 ] )

	this.X = ( e[ 0 ]*x + e[ 4 ]*y + e[ 8 ]*z + e[ 12 ] ) * w
	this.Y = ( e[ 1 ]*x + e[ 5 ]*y + e[ 9 ]*z + e[ 13 ] ) * w
	this.Z = ( e[ 2 ]*x + e[ 6 ]*y + e[ 10 ]*z + e[ 14 ] ) * w

	return this
}

func (this *Vector3) DistanceTo(v *Vector3) float64 {
	return math.Sqrt(this.DistanceToSquared(v))
}

func (this *Vector3) DistanceToSquared(v *Vector3) float64 {
	var dx = this.X - v.X
	var dy = this.Y - v.Y
	var dz = this.Z - v.Z
	return dx*dx + dy*dy + dz*dz
}
