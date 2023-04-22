package repository

import (
	"doctorsAppointment/entity"
	"errors"
)

type DoctorsRepo struct {
	Doctors                 map[int]entity.Doctor
	DoctorForSpecialization map[entity.Specialization][]entity.Doctor
}

func NewDoctorsRepo(doctors map[int]entity.Doctor, doctorForSpecialization map[entity.Specialization][]entity.Doctor) *DoctorsRepo {
	return &DoctorsRepo{
		Doctors:                 doctors,
		DoctorForSpecialization: doctorForSpecialization,
	}
}

func (dr *DoctorsRepo) RegisterDoctor(doctor entity.Doctor) error {
	if _, exist := dr.Doctors[doctor.DoctorId]; exist {
		return errors.New("Doctor Already Registered!")
	}
	dr.Doctors[doctor.DoctorId] = doctor

	if _, exist := dr.DoctorForSpecialization[doctor.Specialization]; !exist {
		dr.DoctorForSpecialization[doctor.Specialization] = make([]entity.Doctor, 0)
	}
	dr.DoctorForSpecialization[doctor.Specialization] = append(dr.DoctorForSpecialization[doctor.Specialization], doctor)
	return nil
}

func (dr *DoctorsRepo) AddAvailability(doctorId int, timeSlot entity.TimeSlot) error {
	if _, exist := dr.Doctors[doctorId]; !exist {
		return errors.New("Doctor Doesn't exit!")
	}

	doctor := dr.Doctors[doctorId]
	slots := doctor.Slots
	slots[timeSlot] = true
	doctor.Slots = slots
	dr.Doctors[doctorId] = doctor
	return nil

}

func (dr *DoctorsRepo) GetDoctorsBySpeciality(specialization entity.Specialization) ([]entity.Doctor, error) {
	if _, exists := dr.DoctorForSpecialization[specialization]; !exists {
		return nil, errors.New("No Doctors for this specialization exists")
	}

	return dr.DoctorForSpecialization[specialization], nil
}

func (dr *DoctorsRepo) GetAvailableTimeSlotsForAllDoctorsForSpecialization(specializedDoctors []entity.Doctor) ([]entity.AvailableDoctor, error) {
	doctorsWithAvaialableSlots := make([]entity.AvailableDoctor, 0)

	for _, doctor := range specializedDoctors {
		var availableDoctor entity.AvailableDoctor
		aSlots := make([]entity.TimeSlot, 0)

		availableDoctor.Doctor = doctor
		slots := doctor.Slots

		for key, val := range slots {
			if val {
				aSlots = append(aSlots, key)
			}
		}
		availableDoctor.SlotList = aSlots
		doctorsWithAvaialableSlots = append(doctorsWithAvaialableSlots, availableDoctor)
	}
	return doctorsWithAvaialableSlots, nil
}

func (dr *DoctorsRepo) IsDoctorRegistered(doctorId int) bool {
	if _, exits := dr.Doctors[doctorId]; exits {
		return true
	}
	return false
}

func (dr *DoctorsRepo) GetDoctorDetails(doctorId int) (*entity.Doctor, error) {
	if _, exits := dr.Doctors[doctorId]; !exits {
		return nil, errors.New("Doctor not Registered")
	}
	doctor := dr.Doctors[doctorId]
	return &doctor, nil
}

func (dr *DoctorsRepo) FreeSlot(doctorId int, slot entity.TimeSlot) error {
	timeSlots := dr.Doctors[doctorId].Slots
	found := false
	for key, _ := range timeSlots {
		if key.Start == slot.Start && key.End == slot.End {
			timeSlots[key] = true
			found = true
			break
		}
	}

	if !found {
		return errors.New("Slot Not Found")
	}
	return nil
}
