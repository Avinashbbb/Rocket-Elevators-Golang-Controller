package main
import(
	"math"
)

type Column struct {
	ID                int
	status            string
	amountOfFloors    int
	amountOfElevators int
	elevatorsList     []*Elevator
	callButtonsList   []*CallButton
	servedFloorsList  []int
	elevatorID        int
	isBasement        bool
}

func NewColumn(_id int, _amountOfFloors int, _amountOfElevators int, _servedFloors []int, _isBasement bool) *Column {
	column := Column{
		ID:                _id,
		status:            "_status",
		amountOfFloors:    _amountOfFloors,
		amountOfElevators: _amountOfElevators,
		elevatorsList:     []*Elevator{},
		callButtonsList:   []*CallButton{},
		servedFloorsList:  _servedFloors,
		elevatorID:        1,
		isBasement:        _isBasement,
	}
	column.createElevators(_amountOfFloors, _amountOfElevators)
	column.createCallButtons(_amountOfFloors, _isBasement)
	return &column
}

func (c *Column) createCallButtons(_amountOfFloors int, _isBasement bool) {
	if _isBasement {
		buttonFloor := -1
		for i := 0; i < _amountOfFloors; i++ {
			callButton := NewCallButton(CallButtonID, "Off", buttonFloor, "Up")
			c.callButtonsList = append(c.callButtonsList, callButton)
			buttonFloor--
			//CallButtonID++
		}

	} else {
		buttonFloor := 1
		for i := 0; i < _amountOfFloors; i++ {
			callButton := NewCallButton(CallButtonID, "OFF", buttonFloor, "Down")
			c.callButtonsList = append(c.callButtonsList, callButton)
			buttonFloor++
			//CallButtonID++
		}
	}

}

func (c *Column) createElevators(_amountOfFloors int, _amountOfElevators int) {
	for i := 0; i < _amountOfElevators; i++ {
		elevator := NewElevator(c.elevatorID, "idle", _amountOfFloors, 1)
		c.elevatorsList = append(c.elevatorsList, elevator)
		c.elevatorID++
	}
}

//Simulate when a user press a button on a floor to go back to the first floor
func (c *Column) requestElevator(_requestedFloor int, _direction string) *Elevator {
	elevator := c.findElevator(_requestedFloor, _direction)
	elevator.addNewRequest(_requestedFloor)
	elevator.move()
	elevator.addNewRequest(1)
	elevator.move()
	return elevator
}

func (c *Column) findElevator(requestedFloor int, requestedDirection string) *Elevator {
	var bestElevator *Elevator
	bestScore := 6
	referenceGap := 100000
	//var bestElevatorInformations *BestElevatorInformations
	if requestedFloor == 1 {
		for _, elevator := range c.elevatorsList {
			if elevator.currentFloor == 1 && elevator.status == "stopped" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(1, elevator, bestScore, referenceGap, bestElevator, requestedFloor)
			} else if elevator.currentFloor == 1 && elevator.status == "idle" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(2, elevator, bestScore, referenceGap, bestElevator, requestedFloor)
			} else if 1 > elevator.currentFloor  && elevator.direction == "Up" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(3, elevator, bestScore, referenceGap, bestElevator, requestedFloor)
			} else if  1 < elevator.currentFloor  && elevator.direction == "down" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(3, elevator, bestScore, referenceGap, bestElevator, requestedFloor)
			} else if elevator.status == "idle" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(4, elevator, bestScore, referenceGap, bestElevator, requestedFloor)
			} else {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(5, elevator, bestScore, referenceGap, bestElevator, requestedFloor)
			}
			// bestElevator = bestElevatorInformations.bestElevator
			// bestScore = bestElevatorInformations.bestScore
			// referenceGap = bestElevatorInformations.referenceGap
		}
	} else {
		for _, elevator := range c.elevatorsList {
			if requestedFloor == elevator.currentFloor && elevator.status == "stopped" && requestedDirection == elevator.direction {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(1, elevator, bestScore, referenceGap, bestElevator, requestedFloor)
			} else if requestedFloor > elevator.currentFloor && elevator.direction == "up" && requestedDirection == "up" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(2, elevator, bestScore, referenceGap, bestElevator, requestedFloor)
			} else if requestedFloor < elevator.currentFloor && elevator.direction == "down" && requestedDirection == "down" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(2, elevator, bestScore, referenceGap, bestElevator, requestedFloor)
			} else if elevator.status == "idle" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(4, elevator, bestScore, referenceGap, bestElevator, requestedFloor)
			} else {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(5, elevator, bestScore, referenceGap, bestElevator, requestedFloor)
			}
			// bestElevator = bestElevatorInformations.bestElevator
			// bestScore = bestElevatorInformations.bestScore
			// referenceGap = bestElevatorInformations.referenceGap
		}
	}
	return bestElevator
}
func (c *Column) checkIfElevatorIsBetter(scoreToCheck int, newElevator *Elevator, bestScore int, referenceGap int, bestElevator *Elevator, floor int) (*Elevator,int,int) {
	if scoreToCheck < bestScore {
		bestScore = scoreToCheck
		bestElevator = newElevator
		referenceGap = int(math.Abs(float64(newElevator.currentFloor) - float64(floor)))
	} else if bestScore == scoreToCheck {
		gap := int(math.Abs(float64(newElevator.currentFloor) - float64(floor)))
		if referenceGap > gap {
			bestElevator = newElevator
			referenceGap = gap
		}
	}
	//bestElevatorInformations := NewBestElevatorInformations(bestElevator, bestScore, referenceGap)
	return bestElevator,bestScore,referenceGap//bestElevatorInformations
}