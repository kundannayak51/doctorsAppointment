package entity

type AvailableDoctor struct {
	Doctor   Doctor     `json:"doctor"`
	SlotList []TimeSlot `json:"slotList"`
}
