package c

//#include <assimp/scene.h>
import "C"

import (
	"reflect"
	"unsafe"
)

type MetadataType C.enum_aiMetadataType

const (
	AI_BOOL       MetadataType = C.AI_BOOL
	AI_INT32      MetadataType = C.AI_INT32
	AI_UINT64     MetadataType = C.AI_UINT64
	AI_FLOAT      MetadataType = C.AI_FLOAT
	AI_DOUBLE     MetadataType = C.AI_DOUBLE
	AI_AISTRING   MetadataType = C.AI_AISTRING
	AI_AIVECTOR3D MetadataType = C.AI_AIVECTOR3D
	AI_META_MAX   MetadataType = C.AI_META_MAX
)

type MetadataEntry C.struct_aiMetadataEntry

func (this *MetadataEntry) Value() interface{} {
	switch MetadataType(this.mType) {
	case AI_BOOL:
		var p1 *C.char = (*C.char)(this.mData)
		a := *p1
		if a == 1 {
			return true
		}
		return false
	case AI_INT32:
		var p1 *C.int32_t = (*C.int32_t)(this.mData)
		return int32(*p1)
	case AI_UINT64:
		var p1 *C.int64_t = (*C.int64_t)(this.mData)
		return int64(*p1)
	case AI_FLOAT:
		var p1 *C.float = (*C.float)(this.mData)
		return float32(*p1)
	case AI_DOUBLE:
		var p1 *C.double = (*C.double)(this.mData)
		return float64(*p1)
	case AI_AISTRING:
		var p1 *C.struct_aiString = (*C.struct_aiString)(unsafe.Pointer(this.mData))
		return C.GoString(&p1.data[0])
	case AI_AIVECTOR3D:
		var p1 *C.struct_aiVector3D = (*C.struct_aiVector3D)(unsafe.Pointer(this.mData))
		x := float32(p1.x)
		y := float32(p1.y)
		z := float32(p1.z)
		return [3]float32{x,y,z}
	}

	return nil
}

type Metadata C.struct_aiMetadata

func (this *Metadata) NumProperties() int {
	return int(this.mNumProperties)
}

func (this *Metadata) Keys() []string {
	if this.mKeys != nil {
		var ret []string = make([]string, 0)

		var result []C.struct_aiString

		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(this.mNumProperties)
		header.Len = int(this.mNumProperties)
		header.Data = uintptr(unsafe.Pointer(this.mKeys))

		for _, i2 := range result {
			ret = append(ret, C.GoString(&i2.data[0]))
		}
		return ret
	} else {
		return nil
	}
}

func (this *Metadata) Values() []MetadataEntry {
	if this.mValues != nil {
		var ret []MetadataEntry = make([]MetadataEntry, 0)

		var result []C.struct_aiMetadataEntry

		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(this.mNumProperties)
		header.Len = int(this.mNumProperties)
		header.Data = uintptr(unsafe.Pointer(this.mValues))
		//fmt.Println(result[0].mType, MetadataType(result[0].mType) == AI_BOOL, AI_INT32 == MetadataType(result[0].mType))

		for _, i2 := range result {
			ret = append(ret, MetadataEntry(i2))

			//t := MetadataEntry(i2)
			//fmt.Println(t.Value())
		}

		return ret
	} else {
		return nil
	}
}

type Node C.struct_aiNode

func (this *Node) Name() string {
	return C.GoString(&this.mName.data[0])
}

func (this *Node) Transformation() Matrix4x4 {
	return Matrix4x4(this.mTransformation)
}

func (this *Node) Parent() *Node {
	return (*Node)(this.mParent)
}

func (this *Node) NumChildren() int {
	return int(this.mNumChildren)
}

func (this *Node) Children() []*Node {
	if this.mNumChildren > 0 && this.mChildren != nil {
		var result []*Node
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(this.mNumChildren)
		header.Len = int(this.mNumChildren)
		header.Data = uintptr(unsafe.Pointer(this.mChildren))
		return result
	} else {
		return nil
	}
}

func (this *Node) NumMeshes() int {
	return int(this.mNumMeshes)
}

func (this *Node) Meshes() []int32 {
	if this.mNumMeshes > 0 && this.mMeshes != nil {
		var result []int32
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(this.mNumMeshes)
		header.Len = int(this.mNumMeshes)
		header.Data = uintptr(unsafe.Pointer(this.mMeshes))
		return result
	} else {
		return nil
	}
}


func (this *Node) MetaData() *Metadata {
	return (*Metadata)(this.mMetaData)
}

const (
	SceneFlags_Incomplete        = C.AI_SCENE_FLAGS_INCOMPLETE
	SceneFlags_Validated         = C.AI_SCENE_FLAGS_VALIDATED
	SceneFlags_ValidationWarning = C.AI_SCENE_FLAGS_VALIDATION_WARNING
	SceneFlags_NonVerboseFormat  = C.AI_SCENE_FLAGS_NON_VERBOSE_FORMAT
	SceneFlags_Terrain           = C.AI_SCENE_FLAGS_TERRAIN
)

type Scene C.struct_aiScene

func (this *Scene) Flags() uint {
	return uint(this.mFlags)
}

func (this *Scene) RootNode() *Node {
	return (*Node)(this.mRootNode)
}

func (this *Scene) NumMeshes() int {
	return int(this.mNumMeshes)
}

func (this *Scene) Meshes() []*Mesh {
	if this.mNumMeshes > 0 && this.mMeshes != nil {
		var result []*Mesh
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(this.mNumMeshes)
		header.Len = int(this.mNumMeshes)
		header.Data = uintptr(unsafe.Pointer(this.mMeshes))
		return result
	} else {
		return nil
	}
}

func (this *Scene) NumMaterials() int {
	return int(this.mNumMaterials)
}

func (this *Scene) Materials() []*Material {
	if this.mNumMaterials > 0 && this.mMaterials != nil {
		var result []*Material
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(this.mNumMaterials)
		header.Len = int(this.mNumMaterials)
		header.Data = uintptr(unsafe.Pointer(this.mMaterials))
		return result
	} else {
		return nil
	}
}

func (this *Scene) NumAnimations() int {
	return int(this.mNumAnimations)
}

func (this *Scene) Animations() []*Animation {
	if this.mNumAnimations > 0 && this.mAnimations != nil {
		var result []*Animation
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(this.mNumAnimations)
		header.Len = int(this.mNumAnimations)
		header.Data = uintptr(unsafe.Pointer(this.mAnimations))
		return result
	} else {
		return nil
	}
}

func (this *Scene) NumTextures() int {
	return int(this.mNumTextures)
}

func (this *Scene) Textures() []*Texture {
	if this.mNumTextures > 0 && this.mTextures != nil {
		var result []*Texture
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(this.mNumTextures)
		header.Len = int(this.mNumTextures)
		header.Data = uintptr(unsafe.Pointer(this.mTextures))
		return result
	} else {
		return nil
	}
}

func (this *Scene) NumLights() int {
	return int(this.mNumLights)
}

func (this *Scene) Lights() []*Light {
	if this.mNumLights > 0 && this.mLights != nil {
		var result []*Light
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(this.mNumLights)
		header.Len = int(this.mNumLights)
		header.Data = uintptr(unsafe.Pointer(this.mLights))
		return result
	} else {
		return nil
	}
}

func (this *Scene) NumCameras() int {
	return int(this.mNumCameras)
}

func (this *Scene) Cameras() []*Camera {
	if this.mNumCameras > 0 && this.mCameras != nil {
		var result []*Camera
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(this.mNumCameras)
		header.Len = int(this.mNumCameras)
		header.Data = uintptr(unsafe.Pointer(this.mCameras))
		return result
	} else {
		return nil
	}
}

func (this *Scene) MetaData() *Metadata {
	return (*Metadata)(this.mMetaData)
}
