package tools

import (
	"lemIn/objects"
)

var AlreadyExploredRoom map[string]bool
var AllEndedPath [][]string

func GetMaxEndedPath() {
	//resset the neework
	SetAllRoomsFalse()
	//get the firts path that camm from The Start Room
	rootPaths := GetUnEploredStartLinkedRoom()
	//we will find the max ended path for each child of the star roomm
	for i := range rootPaths {
		FindEndedPath(rootPaths[i])
	}

}

func GetChildreen(nameOfRoom string) []string {
	theRoom := objects.RoomRegister[nameOfRoom]
	theChilds := []string{}
	if nameOfRoom != objects.End.Name {
		AlreadyExploredRoom[nameOfRoom] = true
	}
	for _, v := range theRoom.LInkedRooms {

		if !AlreadyExploredRoom[v] {
			theChilds = append(theChilds, v)
			if v != objects.End.Name {

				AlreadyExploredRoom[v] = true
			}
		}
	}
	return theChilds
}

func SetAllRoomsFalse() {
	TemporaryMap := map[string]bool{}
	for key, _ := range objects.RoomRegister {
		TemporaryMap[key] = false
	}
	AlreadyExploredRoom = TemporaryMap
}

func RemovePathFromTheNetwork(paths []objects.Path) {
	for _, path := range paths {
		if len(path.RommsOfThePath) < 2 {
			return
		}
		for _, v := range path.RommsOfThePath[1 : len(path.RommsOfThePath)-1] {
			(*&AlreadyExploredRoom)[v] = true
		}
	}
}

func GetUnEploredStartLinkedRoom() [][]string {
	AlreadyExploredRoom[objects.Start.Name] = true
	thePAthToReturn := [][]string{}
	theChildsOftheStartRoom := objects.Start.LInkedRooms
	NoExploredChild := []string{}
	for _, v := range theChildsOftheStartRoom {
		if !AlreadyExploredRoom[v] {
			NoExploredChild = append(NoExploredChild, v)
		}
	}
	for _, v := range NoExploredChild {
		NewPath := []string{}
		NewPath = append(NewPath, objects.Start.Name, v)
		thePAthToReturn = append(thePAthToReturn, NewPath)
	}
	return thePAthToReturn
}

func FindEndedPath(Path []string) {
	SetAllRoomsFalse()
	AlreadyExploredRoom[objects.Start.Name] = true
	PathBuffer := [][]string{}
	PathBuffer = append(PathBuffer, Path)
	for !AllPathAreStuck(PathBuffer) {
		for i, v := range PathBuffer {
			LastRoom := v[len(v)-1]
			if LastRoom == objects.End.Name {
				continue
			}
			LastRoomChilds := GetChildreen(LastRoom)
			if len(LastRoomChilds) == 0 {
				continue
			}
			if len(LastRoomChilds) == 1 {
				PathBuffer[i] = append(PathBuffer[i], LastRoomChilds[0])
				continue
			}
			PathBuffer[i] = append(PathBuffer[i], LastRoomChilds[0])
			LastRoomChilds = LastRoomChilds[1:]
			for _, v2 := range LastRoomChilds {
				NewPath := []string{}
				//Copy all v element in the NewPath
				for _, v3 := range v {
					NewPath = append(NewPath, v3)
				}
				NewPath = append(NewPath, v2)
				PathBuffer = append(PathBuffer, NewPath)
			}
		}
	}
	// recover  all ended path in buffer
	for _, v := range PathBuffer {
		if v[len(v)-1] == objects.End.Name {
			AllEndedPath = append(AllEndedPath, v)
		}
	}

}

func AllPathAreStuck(PathBuffer [][]string) bool {
	for i := range PathBuffer {
		// recover the last room of the path

		LastRoom := objects.RoomRegister[PathBuffer[i][len(PathBuffer[i])-1]]
		DontHaveChilds := true
		//we will the check if the last  have child
		for _, v := range LastRoom.LInkedRooms {
			if !AlreadyExploredRoom[v] {
				DontHaveChilds = false
				break
			}
		}

		if LastRoom.Name != objects.End.Name && !DontHaveChilds {
			return false
		}
	}
	return true
}

func HaveSameRoom(Path1, Path2 []string) bool {
	for _, v := range Path1[1 : len(Path1)-1] {
		for _, v2 := range Path2[1 : len(Path2)-1] {
			if v == v2 {
				return true
			}
		}
	}
	return false
}
