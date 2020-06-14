package model

import "C"
import (
	"assimp/assimp/c"
	"assimp/assimp/math"
	"encoding/binary"
	_math "math"
)

type Mesh struct {
	Name string
	Materialindex int
	Vertices []*math.Vector3
	Normals []*math.Vector3
	Tangents []*math.Vector3
	Bitangents []*math.Vector3
	Colors []*math.Color
	TextureCoords []*math.Vector3
	Faces []uint32
}

type Light struct {
	Name string
	Type int
	Position *math.Vector3
	Direction *math.Vector3
	Up *math.Vector3
}

type Material struct {
	Properties map[string]interface{}
}

type Scene struct {
	Meshes []*Mesh
	MetaData map[string]interface{}
	Materials []*Material
	Lights []*Light
}

func (this*Scene)convertMetadata(src *c.Metadata)  {
	num :=  src.NumProperties()
	if num <= 0 {
		return
	}
	this.MetaData = make(map[string]interface{})
	k := src.Keys()
	v := src.Values()

	for i := 0; i < len(k); i++ {
		this.MetaData[k[i]] = v[i].Value()
	}
}

func (this*Scene)convertMeshes(src []*c.Mesh) {
	for _, v := range src {
		mesh := &Mesh{}
		this.Meshes = append(this.Meshes,mesh)
		mesh.Name = v.Name()
		mesh.Materialindex = v.MaterialIndex()
		mesh.Vertices = make([]*math.Vector3,v.NumVertices())
		for _, v2 := range v.Vertices() {
			mesh.Vertices = append(mesh.Vertices,math.NewVector3(float64(v2.X()),float64(v2.Y()),float64(v2.Z()),))
		}
		mesh.Normals = make([]*math.Vector3,0)
		for _, v2 := range v.Normals() {
			mesh.Normals = append(mesh.Normals,math.NewVector3(float64(v2.X()),float64(v2.Y()),float64(v2.Z()),))
		}
		mesh.Tangents = make([]*math.Vector3,0)
		for _, v2 := range v.Tangents() {
			mesh.Tangents = append(mesh.Tangents,math.NewVector3(float64(v2.X()),float64(v2.Y()),float64(v2.Z()),))
		}
		mesh.Bitangents = make([]*math.Vector3,0)
		for _, v2 := range v.Bitangents() {
			mesh.Bitangents = append(mesh.Bitangents,math.NewVector3(float64(v2.X()),float64(v2.Y()),float64(v2.Z()),))
		}

		mesh.Colors = make([]*math.Color,0)
		for i := 0; i < c.MaxNumberOfColorSets; i++ {
			t2 := v.Colors(i)
			if t2 == nil {
				break
			}
			for j := 0; j < v.NumVertices(); j++ {
				t := t2[j]
				mesh.Colors = append(mesh.Colors,&math.Color{float64(t.R()),float64(t.G()),float64(t.B()),float64(t.A())})
			}
		}

		mesh.TextureCoords = make([]*math.Vector3,0)
		for i := 0; i < c.MaxNumberOfTextureCoords; i++ {
			t2 := v.TextureCoords(i)
			if t2 == nil {
				break
			}
			for j := 0; j < v.NumVertices(); j++ {
				t := t2[j]
				mesh.TextureCoords = append(mesh.TextureCoords,&math.Vector3{float64(t.X()),float64(t.Y()),float64(t.Z())})
			}
		}

		mesh.Faces = make([]uint32,v.NumFaces())
		for _, face := range v.Faces() {
			mesh.Faces = append(mesh.Faces,face.CopyIndices()...)
		}
	}
}

func (this*Scene)convertLights(src []*c.Light) {
	for _, v := range src {
		light := &Light{}
		light.Name = v.Name()
		light.Type = int(v.Type())
		p := v.Position()
		light.Position = math.NewVector3(float64(p.X()),float64(p.Y()),float64(p.Z()))
		p = v.Direction()
		light.Direction = math.NewVector3(float64(p.X()),float64(p.Y()),float64(p.Z()))
		this.Lights = append(this.Lights,light)
	}
}


func byteToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	return _math.Float32frombits(bits)
}

func byteToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	return _math.Float64frombits(bits)
}

func byteToInt(bytes []byte) int {
	return int(binary.LittleEndian.Uint64(bytes))
}

func getPropVal(prop *c.MaterialProperty) interface{} {
	switch prop.Type() {
	case c.PTI_Float:

		return byteToFloat32([]byte(prop.Data()))
	case c.PTI_Double:
		return byteToFloat64([]byte(prop.Data()))
	case c.PTI_String:

		return string([]byte(prop.Data()))
	case c.PTI_Integer:
		return byteToInt([]byte(prop.Data()))
	case c.PTI_Buffer:
		return []byte(prop.Data())
	}

	return nil
}

func (this*Scene) convertMaterials(src []*c.Material) {
	for _, v := range src {
		material := &Material{}

		props := v.Properties()
		material.Properties = make(map[string]interface{})

		for i := 0; i < v.NumProperties(); i++ {
			prop := props[i]
			material.Properties[prop.Key()] = getPropVal(prop)
		}

		this.Materials = append(this.Materials,material)
	}
}

func (this *Scene)Convert(src *c.Scene)  {
	this.convertMetadata(src.MetaData())
	this.Meshes = make([]*Mesh,src.NumMeshes())
	this.convertMeshes(src.Meshes())

	this.Lights = make([]*Light,0)
	this.convertLights(src.Lights())

	this.Materials = make([]*Material,0)
	this.convertMaterials(src.Materials())

}
