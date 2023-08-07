package tools

import (
	"fmt"
	"io/ioutil"
	"lemIn/objects"
	"strconv"
	"strings"
)

func Reader(s string) string {
	Data, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Println(err)
		return ""
	} else {
		return string(Data)
	}
}

func ExploitTheExpemle(TheExempleFileName string) bool {
	ExempleContain := Reader(TheExempleFileName)
	ExempleContainSplited := strings.Split(ExempleContain, "\n")
	//remove the epmty at the end of the list
	ExempleContainSplited = ExempleContainSplited[:len(ExempleContainSplited)-1]
	//***
	NumbersOfAnts, _ := strconv.Atoi(ExempleContainSplited[0])
	//remove the firt line after use it
	ExempleContainSplited = ExempleContainSplited[1:]
	gotStart := false
	gotEnd := false
	for _, v := range ExempleContainSplited {
		if v == "##start" {
			gotStart = true
			continue
		}
		if v == "##end" {
			gotEnd = true
			continue
		}
		// recover Start and end Rooms
		if gotStart && IsRoom(v) {
			RoomElements := strings.Split(v, " ")
			RoomInit("Start", RoomElements)
			gotStart = false
			continue
		}
		if gotEnd && IsRoom(v) {
			RoomElements := strings.Split(v, " ")
			RoomInit("End", RoomElements)
			gotEnd = false
			continue
		}
		//recover Other rooms
		if !gotEnd && !gotStart && IsRoom(v) {
			RoomElements := strings.Split(v, " ")
			RoomInit("other", RoomElements)
			continue
		}
		//Now We'll exploit the Links
		if IsLink(v) {
			theLiks := strings.Split(v, "-")
			firtsLinkedRoom := objects.RoomRegister[theLiks[0]]
			SecondLinkedRoom := objects.RoomRegister[theLiks[1]]
			firtsLinkedRoom.LinkWith(SecondLinkedRoom.Name)
			SecondLinkedRoom.LinkWith(firtsLinkedRoom.Name)
		}

	}
	AntInint(NumbersOfAnts)
	LOLO()

	return true
}

func AntInint(Numbers int) {
	for i := 0; i < Numbers; i++ {
		NewAnt := objects.Ant{}
		NewAnt.Name = "L" + strconv.Itoa(i+1)
		NewAnt.VisitedRoom = append(NewAnt.VisitedRoom, objects.RoomRegister[objects.Start.Name].Name)
		objects.Start.AntsInTheRoom = append(objects.Start.AntsInTheRoom, &NewAnt)
	}

}

func IsRoom(element string) bool {
	return len(strings.Split(element, " ")) == 3
}

func IsLink(element string) bool {
	return len(strings.Split(element, "-")) == 2
}

func RoomInit(typeOfRoom string, RoomElements []string) {
	if typeOfRoom == "Start" {
		objects.Start.Name = RoomElements[0]
		x, _ := strconv.Atoi(RoomElements[1])
		y, _ := strconv.Atoi(RoomElements[2])
		objects.Start.X, objects.Start.Y = x, y
		objects.RoomRegister[objects.Start.Name] = &objects.Start
		return
	}
	if typeOfRoom == "End" {
		objects.End.Name = RoomElements[0]
		x, _ := strconv.Atoi(RoomElements[1])
		y, _ := strconv.Atoi(RoomElements[2])
		objects.End.X, objects.End.Y = x, y
		objects.RoomRegister[objects.End.Name] = &objects.End
		return
	}
	NewRoom := objects.Room{}
	NewRoom.Name = RoomElements[0]
	x, _ := strconv.Atoi(RoomElements[1])
	y, _ := strconv.Atoi(RoomElements[2])
	NewRoom.X, NewRoom.Y = x, y
	objects.RoomRegister[NewRoom.Name] = &NewRoom

}
func removeDuplicates(arr []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for _, value := range arr {
		if !encountered[value] {
			encountered[value] = true
			result = append(result, value)
		}
	}

	return result
}
func LOLO() {
	for _, v := range objects.RoomRegister {
		v.LInkedRooms = removeDuplicates(v.LInkedRooms)
	}
}
func MoveAllAnts() []string {
	tabOfMove := []string{}
	AmontOfAllAnts := len(objects.Start.AntsInTheRoom)
	for len(objects.End.AntsInTheRoom) != AmontOfAllAnts {
		NewMove := ""
		for i := range ValidePaths {
			NewMove += ValidePaths[i].MoveAllTheAntsInTheCurentPath()
		}
		tabOfMove = append(tabOfMove, NewMove)
	}
	return tabOfMove
}
