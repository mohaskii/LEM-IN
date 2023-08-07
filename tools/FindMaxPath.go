package tools

import (
	"fmt"
	"lemIn/objects"
)

var AlreadyExploredRoom map[string]bool
var AllEndedPath [][]string

func GetValidPath() []objects.Path {
	ValidPaths := []objects.Path{}

	//resset the neework
	SetAllRoomsFalse()
	//check If one Start Childs reach the End Room
	for _, v := range objects.Start.LInkedRooms {
		if v == objects.End.Name {
			NewPath := objects.Path{}
			NewPath.RommsOfThePath = []string{objects.Start.Name, objects.End.Name}
			break
		}
	}
	//get the firts path that camm from The Start Room
	rootPaths := GetUnEploredStartLinkedRoom()
	
	for i := range rootPaths {
		UpDatePaths(rootPaths[i])
	}


	fmt.Println(AllEndedPath)
	return ValidPaths
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

func UpDatePaths(Path []string) {
	SetAllRoomsFalse()
	PathBuffer := [][]string{}
	PathBuffer = append(PathBuffer, Path)
	for !AllPathAreStuck(PathBuffer) {
		for i, v := range PathBuffer {
			// if c == 5 {
			// 	fmt.Println(PathBuffer)
			// 	time.Sleep(5 * time.Second)
			// }
			LastRoom := v[len(v)-1]
			if LastRoom == objects.End.Name{
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

func ValidePathsFounded(Paths [][]string) [][]string {
	ValidePath := [][]string{}
	for _, v := range Paths {
		if v[len(v)-1] == objects.End.Name {
			ValidePath = append(ValidePath, v)
		}
	}
	return ValidePath
}

//make an function to remove path
// 	To do

func TabCompare(tab1 []string, tab2 []string) bool {
	if len(tab1) != len(tab2) {
		return false
	}
	for i, v := range tab1 {
		if v != tab2[i] {
			return false
		}
	}
	return true
}

func AllPathAreStuck(PathBuffer [][]string) bool {
	for i := range PathBuffer {
		// recover the last room of the path

		LastRoom := objects.RoomRegister[PathBuffer[i][len(PathBuffer[i])-1]]
		DontHaveChilds := true
		for _, v := range LastRoom.LInkedRooms {
			if !AlreadyExploredRoom[v] {
				DontHaveChilds = false
				break
			}
		}
		// fmt.Printf("ddddd :  %v \n", DontHaveChilds)
		if LastRoom.Name != objects.End.Name && !DontHaveChilds {
			return false
		}
	}
	return true
}
