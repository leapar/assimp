package math

type Vector2 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func NewVector2()*Vector2  {
	return &Vector2{}
}

func (this *Vector2)ToBytes()[]byte  {
	f := make([]byte,0)
	f = append(f,Float32ToByte(float32(this.X))...)
	f = append(f,Float32ToByte(float32(this.Y))...)
	return f
}