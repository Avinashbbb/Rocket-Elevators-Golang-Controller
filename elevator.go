package main

import (
	"sort"
	"time"
)

type Elevator struct {
	ID                    int
	status                string
	amountOfFloors        int
	currentFloor         int
	door                  *Door
	floorRequestsList      []int
	direction             string
	overweightAlarm		string
	overweight            bool
	obstruction           bool
	completedRequestsList []int
}

func NewElevator(_elevatorID int, _status string, _amountOfFloors int, _currentFloor int) *Elevator {
	elevator := Elevator{
		ID:                    _elevatorID,
		status:                _status,
		amountOfFloors:        _amountOfFloors,
		currentFloor:         _currentFloor,
		door:                  NewDoor(_elevatorID, "closed"),
		floorRequestsList:      []int{},
		direction:             "nil",
		overweight:            false,
		obstruction:           false,
		completedRequestsList: []int{},
	}
	return &elevator
}

func (e *Elevator) move() {
	for len(e.floorRequestsList) > 0 {
		
		e.status = "moving"
		e.sortFloorList()
		destination := e.floorRequestsList[0]
		if e.direction == "up" {
			for e.currentFloor < destination {
				e.currentFloor++
			}
		} else if e.direction == "down" {
			for e.currentFloor > destination {
				e.currentFloor--
			}
		}
		e.status = "stopped"
		e.operateDoors()
		e.floorRequestsList = e.floorRequestsList[1:]
		e.completedRequestsList = append(e.completedRequestsList, destination)
	}
	e.status = "idle"
}
func (e *Elevator) sortFloorList() {
	if e.direction == "Up" {
		sort.Ints(e.floorRequestsList)
	}else{
		sort.Sort(sort.Reverse(sort.IntSlice(e.floorRequestsList)))
	}
}
func (e *Elevator) operateDoors(){
	e.door.status = "opened"
	time.Sleep( 5* time.Second)
	if !e.overweight{
		e.door.status = "closing"
		if !e.obstruction{
			e.door.status = "closed"
		}else{
			e.operateDoors()
		}
	}else{
		for e.overweight{
			e.overweightAlarm = "Activated"
		}
		e.operateDoors()
	}
}
func(e *Elevator) addNewRequest(requestedFloor int){
	if !contains(e.floorRequestsList,requestedFloor){
		e.floorRequestsList = append(e.floorRequestsList,requestedFloor)
	}
	if e.currentFloor < requestedFloor{
		e.direction = "up"
	}
	if e.currentFloor > requestedFloor{
		e.direction = "down"
	}
}