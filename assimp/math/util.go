package math

import (
	"encoding/binary"
	_math "math"
)



func ToSnorm8(x_value float64) byte{

	return byte(_math.Round(clamp(x_value, -1.0, 1.0) * 127))
}



func Float32ToByte(float float32) []byte {
	bits := _math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)

	return bytes
}


func ByteToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	return _math.Float32frombits(bits)
}

func Float64ToByte(float float64) []byte {
	bits := _math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)

	return bytes
}


func ByteToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	return _math.Float64frombits(bits)
}

func ByteToInt(bytes []byte) int {
	return int(binary.LittleEndian.Uint64(bytes))
}

func Byte4ToInt(bytes []byte) int {
	return int(binary.LittleEndian.Uint32(bytes))
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(i))
	return buf
}

func IntToBytes(i int) []byte {
	var buf = make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(i))
	return buf
}

func Uint32ToBytes(i uint32) []byte {
	var buf = make([]byte, 8)
	binary.LittleEndian.PutUint32(buf, uint32(i))
	return buf
}