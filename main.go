package main

import "C"
import (
	"assimp/assimp/c"
	"assimp/assimp/math"
	"assimp/assimp/model"
	"fmt"
)

func main()  {
/*
	transformation := math.SQ{}
	transformation.Set(4,4,[]float64{11.1, 11.2, 11.3, 11.4,
		12.1, 12.2, 12.3, 12.4,
		13.1, 13.2, 13.3, 13.4,
		14.1, 14.2, 14.3, 14.4})

	assimp_transformation := math.SQ{}
	assimp_transformation.Set(4,4,[]float64{1.1, 1.2, 1.3, 1.4,
		2.1, 2.2, 2.3, 2.4,
		3.1, 3.2, 3.3, 3.4,
		4.1, 4.2, 4.3, 4.4})

	v := math.Mul(transformation,assimp_transformation)
*/
	transformation := math.NewMatrix4().Set(
		11.1,12.1,13.1,14.1,
		11.2,12.2,13.2,14.2,
		11.3,12.3,13.3,14.3,
		11.4,12.4,13.4,14.4)
	assimp_transformation := math.NewMatrix4().Set(
		1.1,2.1,3.1,4.1,
		1.2,2.2,3.2,4.2,
		1.3,2.3,3.3,4.3,
		1.4,2.4,3.4,4.4)

	fmt.Println(assimp_transformation.MultiplyMatrices(transformation))
	//fmt.Println(v,math.Det(v,3),math.Det(v,1),math.Det(v,2),math.Det(v,0))

	//fmt.Println(assimp_transformation.Determinant(),assimp_transformation.Determinant2() ,assimp_transformation.Determinant2()/2)
	//scene := assimp.ImportFile("C:\\Users\\WXH\\Documents\\untitled.fbx",1)
	scene := c.ImportFileExWithProperties("C:\\Users\\WXH\\Documents\\untitled.fbx")
	defer func() {
		scene.ReleaseImport()
	}()

	fmt.Println("mesh num:",scene.NumMeshes())
	/*for _, mesh := range scene.Meshes() {
		fmt.Println("mesh:",mesh)
	}*/
	fmt.Println("Material num:",scene.NumMaterials())
	fmt.Println("Animation num:",scene.NumAnimations())
	fmt.Println("Texture num:",scene.NumTextures())
	fmt.Println("Light num:",scene.NumLights())
	fmt.Println("Camera num:",scene.NumCameras())

//	fmt.Println("metadata:",scene.MetaData())
//	fmt.Println("metadata:",scene.MetaData().Keys() )
//	fmt.Println("metadata:",scene.MetaData().Values() )

	s := model.Scene{}
	s.Convert(scene)
	p := s.MetaData["UnitScaleFactor"]
	fmt.Println(p)
	//fmt.Println(s)
}
