package main

import (
	"fmt"
	"lemIn/objects"
	"lemIn/tools"
)

func main() {
	tools.PutAntsOnPaths()
	TheFirstPath := tools.ValidePaths[0]
	for _, v:= range objects.Start.AntsInTheRoom{
		objects.IsOut[v.Name]=true
	}
	fmt.Println(TheFirstPath.RommsOfThePath)
	fmt.Println(TheFirstPath.MoveAllTheAntsInThePath())
	fmt.Println(TheFirstPath.MoveAllTheAntsInThePath())
	fmt.Println(TheFirstPath.MoveAllTheAntsInThePath())
	fmt.Println(TheFirstPath.MoveAllTheAntsInThePath())
	fmt.Println(TheFirstPath.MoveAllTheAntsInThePath())
	fmt.Println(TheFirstPath.MoveAllTheAntsInThePath())
	fmt.Println(TheFirstPath.MoveAllTheAntsInThePath())
	fmt.Println(TheFirstPath.MoveAllTheAntsInThePath())
	fmt.Println(TheFirstPath.MoveAllTheAntsInThePath())
	// for _, v := range TheFirstPath.PlacedAnts {
	// 	fmt.Println(v.Name)
	// 	fmt.Println(v.CurentPlace)
	// 	fmt.Printf("this is the start room name : %v\n", objects.Start.Name)
	// 	fmt.Println("==========================================================================================================")
	// }
}

/* exemple that return empty Board
exemple 02
*/
