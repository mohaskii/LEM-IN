package main

import (
	"fmt"
	"lemIn/objects"
	"lemIn/tools"
)

func main() {
	tools.GetMaxEndedPath()
	tools.GetAllChunk()
	tools.GetTheChuckWithTHeMaxPath()
	tools.InitValidPAth()
	tools.PutAntsOnPaths()
	//set The map
	for _, v:= range objects.Start.AntsInTheRoom{
	
		objects.IsOut[v.Name]= true
	}
	fmt.Println(len(tools.MoveAllAnts()))
	fmt.Println()
	

}
