package main

type Door struct {
	ID     int
	status string
}

func NewDoor(_id int, _status string) *Door {
	door := Door{
		ID:     _id,
		status: _status,
	}
	return &door
}
