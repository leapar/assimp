package model

import "C"
import (
	"assimp/assimp/c"
	"assimp/assimp/math"
)

type Mesh struct {
	Name          string
	Materialindex int
	Vertices      []*math.Vector3
	Normals       []*math.Vector3
	Tangents      []*math.Vector3
	Bitangents    []*math.Vector3
	Colors        []*math.Color
	TextureCoords []*math.Vector3
	Faces         [][]uint32
}

type Light struct {
	Name      string
	Type      int
	Position  *math.Vector3
	Direction *math.Vector3
	Up        *math.Vector3
}

type Material struct {
	Properties map[string]interface{}
}

type Node struct {
	Name           string
	Transformation *math.Matrix4
	Parent         *Node
	NumChildren    int
	Children       []*Node
	NumMeshes      int
	Meshes         []*Mesh
	MetaData       map[string]interface{}
}

type Scene struct {
	Meshes    []*Mesh
	MetaData  map[string]interface{}
	Materials []*Material
	Lights    []*Light
	RootNode  *Node
}

func (this *Scene) convertMetadata(src *c.Metadata) map[string]interface{} {
	if src == nil {
		return nil
	}
	metaData := make(map[string]interface{})

	num := src.NumProperties()
	if num <= 0 {
		return nil
	}

	k := src.Keys()
	v := src.Values()

	for i := 0; i < len(k); i++ {
		metaData[k[i]] = v[i].Value()
	}

	return metaData
}

func (this *Scene) convertMeshes(src []*c.Mesh) []*Mesh {
	meshes := make([]*Mesh, 0)
	for _, v := range src {
		mesh := &Mesh{}
		meshes = append(meshes, mesh)
		mesh.Name = v.Name()
		mesh.Materialindex = v.MaterialIndex()
		mesh.Vertices = make([]*math.Vector3, 0)
		for _, v2 := range v.Vertices() {
			mesh.Vertices = append(mesh.Vertices, math.NewVector3(float64(v2.X()), float64(v2.Y()), float64(v2.Z())))
		}
		mesh.Normals = make([]*math.Vector3, 0)
		for _, v2 := range v.Normals() {
			mesh.Normals = append(mesh.Normals, math.NewVector3(float64(v2.X()), float64(v2.Y()), float64(v2.Z())))
		}
		mesh.Tangents = make([]*math.Vector3, 0)
		for _, v2 := range v.Tangents() {
			mesh.Tangents = append(mesh.Tangents, math.NewVector3(float64(v2.X()), float64(v2.Y()), float64(v2.Z())))
		}
		mesh.Bitangents = make([]*math.Vector3, 0)
		for _, v2 := range v.Bitangents() {
			mesh.Bitangents = append(mesh.Bitangents, math.NewVector3(float64(v2.X()), float64(v2.Y()), float64(v2.Z())))
		}

		mesh.Colors = make([]*math.Color, 0)
		for i := 0; i < c.MaxNumberOfColorSets; i++ {
			t2 := v.Colors(i)
			if t2 == nil {
				break
			}
			for j := 0; j < v.NumVertices(); j++ {
				t := t2[j]
				mesh.Colors = append(mesh.Colors, &math.Color{float64(t.R()), float64(t.G()), float64(t.B()), float64(t.A())})
			}
		}

		mesh.TextureCoords = make([]*math.Vector3, 0)
		for i := 0; i < c.MaxNumberOfTextureCoords; i++ {
			t2 := v.TextureCoords(i)
			if t2 == nil {
				break
			}
			for j := 0; j < v.NumVertices(); j++ {
				t := t2[j]
				mesh.TextureCoords = append(mesh.TextureCoords, &math.Vector3{float64(t.X()), float64(t.Y()), float64(t.Z())})
			}
		}

		mesh.Faces = make([][]uint32, v.NumFaces())
		for index, face := range v.Faces() {
			mesh.Faces[index] = make([]uint32,0)
			mesh.Faces[index] = append(mesh.Faces[index], face.CopyIndices()...)
		}
	}

	return meshes
}

func (this *Scene) convertLights(src []*c.Light) {
	for _, v := range src {
		light := &Light{}
		light.Name = v.Name()
		light.Type = int(v.Type())
		p := v.Position()
		light.Position = math.NewVector3(float64(p.X()), float64(p.Y()), float64(p.Z()))
		p = v.Direction()
		light.Direction = math.NewVector3(float64(p.X()), float64(p.Y()), float64(p.Z()))
		this.Lights = append(this.Lights, light)
	}
}




func getPropVal(prop *c.MaterialProperty) interface{} {
	switch prop.Type() {
	case c.PTI_Float:
		return prop.Float32Data()
	case c.PTI_Double:
		return prop.Float64Data()
	case c.PTI_String:


		return prop.StringData()
	case c.PTI_Integer:
		return prop.Int32Data()
	case c.PTI_Buffer:
		return []byte(prop.Data())
	}

	return nil
}

func (this *Scene) convertMaterials(src []*c.Material) {
	for _, v := range src {
		material := &Material{}

		props := v.Properties()
		material.Properties = make(map[string]interface{})

		for i := 0; i < v.NumProperties(); i++ {
			prop := props[i]
			material.Properties[prop.Key()] = getPropVal(prop)
		}

		this.Materials = append(this.Materials, material)
	}
}

func (this *Scene) convertNodes(src *c.Node, parent *Node) *Node {

	node := &Node{}

	node.Name = src.Name()
	node.MetaData = this.convertMetadata(src.MetaData())
	node.NumChildren = src.NumChildren()
	node.Transformation = math.NewMatrix4()
	trans := src.Transformation()

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			node.Transformation.Elements[i*4+j] = float64(trans.Values()[i][j])
		}
	}
	node.NumMeshes = src.NumMeshes()
	node.Parent = parent

	if parent != nil {
		if parent.Children == nil {
			parent.Children = make([]*Node, 0)
		}

		parent.Children = append(parent.Children, node)
	}

	node.Meshes = make([]*Mesh, 0)
	meshIndexs := src.Meshes()

	for i := 0; i < node.NumMeshes; i++ {
		node.Meshes = append(node.Meshes, this.Meshes[meshIndexs[i]])
	}

	for _, child := range src.Children() {
		this.convertNodes(child, node)
	}

	return node
}

func (this *Scene) Convert(src *c.Scene) {
	this.MetaData = this.convertMetadata(src.MetaData())
	this.Meshes = this.convertMeshes(src.Meshes())

	this.Lights = make([]*Light, 0)
	this.convertLights(src.Lights())

	this.Materials = make([]*Material, 0)
	this.convertMaterials(src.Materials())

	this.RootNode = this.convertNodes(src.RootNode(), nil)
}
