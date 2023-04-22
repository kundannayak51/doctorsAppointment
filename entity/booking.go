package entity

type Booking struct {
	BookingId int      `json:"booking_id"`
	Doctor    Doctor   `json:"doctor"`
	Patient   Patient  `json:"patient"`
	Slot      TimeSlot `json:"slot"`
	WaitList  bool     `json:"waitList"`
}
