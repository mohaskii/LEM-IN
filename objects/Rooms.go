package objects

import (
	"fmt"
)

type Room struct {
	Name        string
	LInkedRooms []string
	//Cords
	X             int
	Y             int
	AntsInTheRoom []*Ant
}

type Ant struct {
	Name        string
	VisitedRoom []string
}

type Path struct {
	RommsOfThePath []string
	PlacedAnts     []*Ant
}

var Start Room
var End Room
var RoomRegister = map[string]*Room{}
var IsOut = make(map[string]bool, len(Start.AntsInTheRoom))
var AmontOfAllAnts = len(Start.AntsInTheRoom)

func (r *Room) LinkWith(thelikedRoomName string) {
	r.LInkedRooms = append(r.LInkedRooms, thelikedRoomName)
}

func (a *Ant) MoveOn(theRoomName string) error {
	if _, OnIt := RoomRegister[theRoomName]; !OnIt {
		return fmt.Errorf("the room '%v' do not exist ", theRoomName)
	}
	RoomRegister[theRoomName].AntsInTheRoom = append(RoomRegister[theRoomName].AntsInTheRoom, a)
	a.VisitedRoom = append(a.VisitedRoom, theRoomName)

	return nil
}

func (p *Path) MoveAllTheAntsInThePath() string {
	//set the map

	theMouves := ""

	for i := range p.RommsOfThePath {
		theReversed := (len(p.RommsOfThePath) - 1) - i
		if p.RommsOfThePath[theReversed] != End.Name && p.RommsOfThePath[theReversed] != Start.Name {
			//we check if got an Ant in the Roomm
			TheRoom := RoomRegister[p.RommsOfThePath[theReversed]]

			if len(TheRoom.AntsInTheRoom) == 1 {
				// we move the ant on the next room
				TheAnt := TheRoom.AntsInTheRoom[0]
				TheAnt.MoveOn(p.RommsOfThePath[theReversed+1])
				RoomRegister[p.RommsOfThePath[theReversed]].AntsInTheRoom = []*Ant{}
				theMouves += fmt.Sprintf(" %s-%s", TheAnt.Name, p.RommsOfThePath[theReversed+1])
			}
		}
	}
	for _, v := range p.PlacedAnts {
		if IsOut[v.Name] {
			v.MoveOn(p.RommsOfThePath[1])
			IsOut[v.Name] = false
			theMouves += fmt.Sprintf(" %s-%s", v.Name, p.RommsOfThePath[1])
			break
		}
	}
	return theMouves
}
