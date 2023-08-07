package tools

import "fmt"

var AllChunk = make(map[string][][]string)
var AllReadyGot = make(map[int]bool)

func FindChunk(Index int) [][]string {
	ChunkToReturn := [][]string{}
	fmt.Println(AllEndedPath)
	fmt.Println("==============================")
	ChunkToReturn = append(ChunkToReturn, AllEndedPath[Index])
	AllReadyGot[Index]= true
	for i := range AllEndedPath {
		if i != Index &&!AllReadyGot[i]&& !HaveSameRoom(AllEndedPath[Index], AllEndedPath[i]) && HaveNoCommonElementWith(ChunkToReturn,AllEndedPath[i] ) {
			ChunkToReturn = append(ChunkToReturn, AllEndedPath[i])
			AllReadyGot[i]= true
			Index = i

		}
	}
	return ChunkToReturn
}

func HaveNoCommonElementWith(Prvious [][]string , New  []string)bool {
	for _, v := range Prvious {
		if HaveSameRoom(v, New){
			return false
		}
	}
	return true
}