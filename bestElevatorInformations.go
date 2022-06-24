package main

type BestElevatorInformations struct {
	bestElevator *Elevator 
	bestScore    int
	referenceGap int
}

func NewBestElevatorInformations(_bestElevator *Elevator, _bestScore int,_referenceGap int)* BestElevatorInformations{
	bestElevatorInformations := BestElevatorInformations{
		bestElevator: _bestElevator,
		bestScore: _bestScore,
		referenceGap: _referenceGap,
	}
	return &bestElevatorInformations
}