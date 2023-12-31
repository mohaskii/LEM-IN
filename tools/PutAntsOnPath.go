package tools

import (
	"lemIn/objects"
)

var _ = ExploitTheExpemle("ExempleToTest/exemple03")

var ValidePaths = []objects.Path{}

var AlreadyPlaced map[string]bool

var LenValidfePath = len(ValidePaths)

func PutAntsOnPaths() {
	for _, Ant := range objects.Start.AntsInTheRoom {
		TheIndexOfGoodPath := FindTheGoodPath()
		ValidePaths[TheIndexOfGoodPath].PlacedAnts = append(ValidePaths[TheIndexOfGoodPath].PlacedAnts, Ant)
	}
}

func FindTheGoodPath() int {
	//we'll get the sum of the number of ant and the number of room on the of the first path
	s1 := len(ValidePaths[0].RommsOfThePath) + len(ValidePaths[0].PlacedAnts)
	for i := range ValidePaths {

		s2 := len(ValidePaths[i].RommsOfThePath) + len(ValidePaths[i].PlacedAnts)
		if s2 < s1 {
			return i
		}
	}
	return 0
}

func InitValidPAth() {
	ValidePathss := []objects.Path{}
	for _, v := range TheGoodChunk {
		New := objects.Path{}
		New.RommsOfThePath = v
		ValidePathss = append(ValidePathss, New)
	}
	ValidePaths = ValidePathss
}
