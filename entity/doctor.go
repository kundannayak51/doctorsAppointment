package entity

type Specialization string

const (
	Cardiologist     Specialization = "Cardiologist"
	Dermatologist    Specialization = "Dermatologist"
	Orthopedic       Specialization = "Orthopedic"
	GeneralPhysician Specialization = "GeneralPhysician"
)

type Doctor struct {
	DoctorId       int               `json:"doctor_id"`
	DoctorName     string            `json:"doctor_name"`
	Rating         int               `json:"rating"`
	Slots          map[TimeSlot]bool `json:"slots"`
	Specialization Specialization    `json:"specialization"`
}
