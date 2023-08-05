package tools

import (
	"lemIn/objects"
)

var AlreadyExploredRoom map[string]bool

func GetValidPath() []objects.Path {
	ValidPaths := []objects.Path{}
	//resset the neework
	SetAllRoomsFalse()
	//check If on Start Childs reach the End Room
	for _, v := range objects.Start.LInkedRooms {
		if v == objects.End.Name {
			NewPath := objects.Path{}
			NewPath.RommsOfThePath = []string{objects.Start.Name, objects.End.Name}
			break
		}
	}
	//get the firts path that camm from The Start Room
	rootPaths := GetUnEploredStartLinkedRoom()
	//intialize the  previous befor the update
	PreviousPath := []string{}
	for len(rootPaths) != 0 {
		if len(rootPaths) == 1 && TabCompare(PreviousPath, rootPaths[0]) {
			break
		}
		PotentatielValidPath := ValidePathsFounded(rootPaths)
		for _, v := range PotentatielValidPath {
			NewPath := objects.Path{}
			NewPath.RommsOfThePath = v
			ValidPaths = append(ValidPaths, NewPath)
			SetAllRoomsFalse()
			RemovePathFromTheNetwork(ValidPaths)
			//restar the path finder after founded one or many paths at the same time
			rootPaths = GetUnEploredStartLinkedRoom()
		}
		if len(rootPaths) == 1 {
			PreviousPath = rootPaths[0]
		}
		UpDatePaths(&rootPaths)
	}
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
			AlreadyExploredRoom[v] = true
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
		child := GetChildreen(v)
		for _, v2 := range child {
			NewPath := []string{}
			NewPath = append(NewPath, objects.Start.Name, v, v2)
			thePAthToReturn = append(thePAthToReturn, NewPath)
		}
	}
	return thePAthToReturn
}

func UpDatePaths(Paths *[][]string) {
	NewPathsFounded := [][]string{}
	for i, v := range *Paths {
		lastRoom := v[len(v)-1]
		childsOftheLastRoom := GetChildreen(lastRoom)
		if len(childsOftheLastRoom) == 0 {
			continue
		}
		if len(childsOftheLastRoom) == 1 {

			(*Paths)[i] = append((*Paths)[i], childsOftheLastRoom[0])
			continue
		}
		(*Paths)[i] = append((*Paths)[i], childsOftheLastRoom[0])
		childsOftheLastRoom = childsOftheLastRoom[1:]
		for _, v2 := range childsOftheLastRoom {
			NewPath := []string{}
			for _, v3 := range v {
				NewPath = append(NewPath, v3)
			}
			NewPath = append(NewPath, v2)
			NewPathsFounded = append(NewPathsFounded, NewPath)
		}
	}
	*Paths = append(*Paths, NewPathsFounded...)
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
