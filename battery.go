package main

//var ColumnID = 1
//var ElevatorID = 1
var FloorRequestButtonID = 1
var CallButtonID = 1
//var Floor int

type Battery struct {
	ID                       int
	status                   string
	columnsList              []Column
	floorRequestsButtonsList []*FloorRequestButton //might have to chang the data type
	columnID                 int
}

func NewBattery(_id int,_amountOfColumns int,_amountOfFloors int,_amountOfBasements int, _amountOfElevatorPerColumn int) *Battery{
	battery:= Battery{
		ID: _id,
		status: "online",
		columnsList: [] Column{},
		floorRequestsButtonsList: []*FloorRequestButton{},
		columnID: 1,
	}
		if _amountOfBasements > 0 {
			battery.createBasementFloorRequestButtons(_amountOfBasements)
			battery.createBasementColumn(_amountOfBasements,_amountOfElevatorPerColumn)
			_amountOfColumns --
		}
		battery.createFloorRequestButtons(_amountOfFloors)
		battery.createColumns(_amountOfColumns,_amountOfFloors,_amountOfElevatorPerColumn)
		

		
	
	return &battery
}

func (b *Battery) createBasementColumn(_amountOfBasements int, _amountOfElevatorPerColumn int) {
	servedFloors := []int{}
	floor := -1
	for i := 0; i < _amountOfBasements; i++ {
		servedFloors = append(servedFloors, floor)
		floor--
	}
	var column *Column = NewColumn(b.columnID,_amountOfBasements,_amountOfElevatorPerColumn, servedFloors, true)
	b.columnsList = append(b.columnsList, *column)
	b.columnID++
}
func(b *Battery)createColumns(_amountOfColumns int,_amountOfFloors int,_amountOfElevatorPerColumn int){
	amountOfFloorsPerColumn := (_amountOfFloors -1)/_amountOfColumns +1
	floor := 1

	for i := 0; i < _amountOfColumns; i++ {
		servedFloors := []int{}
		for n := 0; n < amountOfFloorsPerColumn; n++ {
			if floor <= _amountOfFloors {
				servedFloors = append(servedFloors, floor)
				floor++
			}
		}
		column:= NewColumn(b.columnID,_amountOfFloors,_amountOfElevatorPerColumn,servedFloors,false)
		b.columnsList = append(b.columnsList, *column)
		b.columnID++
	}
}
// func (b *Battery) createColumns(_amountOfColumns int, _amountOfFloors int, _amountOfElevatorPerColumn int) {
// 	var amountOfFloorsPerColumn int = (_amountOfFloors-1)/_amountOfColumns + 1
// 	Floor = 1
// 	for i := 0; i < _amountOfColumns; i++ {
// 		servedFloors := []int{}
// 		for n := 0; n < amountOfFloorsPerColumn; n++ {
// 			if Floor <= _amountOfFloors {
// 				servedFloors = append(servedFloors, Floor)
// 				Floor++
// 			}
// 		}
// 		column := NewColumn(b.columnID, _amountOfFloors, _amountOfElevatorPerColumn, servedFloors, false)
// 		battery.columnsList = append(battery.columnsList, column)
// 		b.columnID++
// 	}
// }

func (b *Battery) createFloorRequestButtons(_amountOfFloors int) {
	buttonFloor := 1
	for i := 0; i < _amountOfFloors; i++ {
		floorRequestButton := NewFloorRequestButton(FloorRequestButtonID, "OFF", buttonFloor, "Up")
		b.floorRequestsButtonsList = append(b.floorRequestsButtonsList, floorRequestButton)
		buttonFloor ++
		FloorRequestButtonID ++
	}
}
func (b *Battery) createBasementFloorRequestButtons(_amountOfBasements int) {
	buttonFloor := -1
	for i := 0; i < _amountOfBasements; i++ {
		floorRequestButton := NewFloorRequestButton(FloorRequestButtonID, "OFF", buttonFloor, "Down")
		b.floorRequestsButtonsList = append(b.floorRequestsButtonsList, floorRequestButton)
		buttonFloor--
		FloorRequestButtonID++
	}
}

func (b *Battery) findBestColumn(_requestedFloor int) *Column {
	for _, column := range b.columnsList {
		if contains(column.servedFloorsList, _requestedFloor) {
			return &column
		}
	}
	return nil
}

//Simulate when a user press a button at the lobby
func (b *Battery) assignElevator(_requestedFloor int, _direction string) (*Column, *Elevator) {
	column := *b.findBestColumn(_requestedFloor)
	elevator := column.findElevator(1,_direction)
	elevator.addNewRequest(1)
	elevator.move()

	elevator.addNewRequest(_requestedFloor)
	elevator.move()
	return &column, elevator
}
