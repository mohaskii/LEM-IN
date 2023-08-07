package tools

var AllChunk = make(map[int][][]string)
var AllReadyGot = make(map[int]bool)
var TheGoodChunk = [][]string{}

func FindChunk(Index int) [][]string {
	ChunkToReturn:= [][]string{}
	ChunkToReturn = append(ChunkToReturn, AllEndedPath[Index])
	AllReadyGot[Index] = true
	for i := range AllEndedPath {
		if i != Index && !AllReadyGot[i] && !HaveSameRoom(AllEndedPath[Index], AllEndedPath[i]) && HaveNoCommonElementWith(ChunkToReturn, AllEndedPath[i]) {
			ChunkToReturn = append(ChunkToReturn, AllEndedPath[i])
			AllReadyGot[i] = true
			Index = i

		}
	}
	return ChunkToReturn
}

func HaveNoCommonElementWith(Prvious [][]string, New []string) bool {
	for _, v := range Prvious {
		if HaveSameRoom(v, New) {
			return false
		}
	}
	return true
}

func GetAllChunk() {
	for i := range AllEndedPath {
		AllChunk[i] = FindChunk(i)
	}
}

func GetTheChuckWithTHeMaxPath () {
	theMaxChunk := AllChunk[0]
	for _, v := range AllChunk{
		if len(v) > len(theMaxChunk){
			theMaxChunk =v
		}
	}
	TheGoodChunk = theMaxChunk
}
