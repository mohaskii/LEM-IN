package objects

import (
	"errors"
	"fmt"
)

type Room struct {
	Name string
	LInkedRooms []string
	//Cords
	X int
	Y int
}


type Ant struct{
	Name string
	CurentPlace string
	VisitedRoom []string
}

var Start Room
var End Room
var OtherRooms []*Room
var RoomRegister = map[string]*Room{} 
var Ants []*Ant

func (r *Room) LinkWith(thelikedRoomName string){
	r.LInkedRooms = append(r.LInkedRooms,  thelikedRoomName)
}

func (a *Ant) MoveOn(theRoomName string ) error{
	if _, OnIt  := RoomRegister[theRoomName]; !OnIt {
		return errors.New(fmt.Sprintf("the room '%v' do not exist ", theRoomName))
	}
	a.CurentPlace = theRoomName
	a.VisitedRoom = append(a.VisitedRoom, theRoomName)
	return nil
}
