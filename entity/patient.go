package entity

type Patient struct {
	PatientId   int                `json:"patient_id"`
	PatientName string             `json:"patient_name"`
	BookedSlots map[int][]TimeSlot `json:"booked_slots"` //doctorId-list of slots
}
