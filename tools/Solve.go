package tools

import (
	"fmt"
	"lemIn/objects"
)

func Solve() {
	GetMaxEndedPath()
	GetAllChunk()
	GetTheChuckWithTHeMaxPath()
	InitValidPAth()
	PutAntsOnPaths()
	//set The map
	for _, v := range objects.Start.AntsInTheRoom {

		objects.IsOut[v.Name] = true
	}
	for _, v := range MoveAllAnts() {
		fmt.Println(v[1:])
	}
}
