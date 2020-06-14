package main

import "C"
import (
	"assimp/assimp/c"
	"assimp/assimp/model"
	"fmt"
)

func main()  {
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

	//fmt.Println(s)
}
