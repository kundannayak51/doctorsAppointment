package service

import (
	"doctorsAppointment/entity"
	"doctorsAppointment/mode"
	"doctorsAppointment/repository"
	"doctorsAppointment/utils"
	"errors"
	"fmt"
)

type DoctorService struct {
	DoctorsRepo repository.DoctorsRepo
	Print       mode.Print
}

func NewDoctorService(doctorsRepo repository.DoctorsRepo, print mode.Print) *DoctorService {
	return &DoctorService{
		DoctorsRepo: doctorsRepo,
		Print:       print,
	}
}

func (ds *DoctorService) RegisterDoctor(doctor entity.Doctor) error {
	err := ds.DoctorsRepo.RegisterDoctor(doctor)
	if err != nil {
		return err
	}
	ds.Print.PrintDataOnConsole(fmt.Sprintf("Welcome %s", doctor.DoctorName))
	return nil
}

func (ds *DoctorService) AddAvailability(doctorId int, timeSlot entity.TimeSlot) error {
	startTime, err := utils.ConvertStringToTime(timeSlot.Start)
	if err != nil {

	}
	endTime, err := utils.ConvertStringToTime(timeSlot.End)
	if err != nil {
		return err
	}

	duration := endTime.Sub(startTime)
	minutes := int(duration.Minutes())

	if minutes == 30 {
		err = ds.DoctorsRepo.AddAvailability(doctorId, timeSlot)
		if err != nil {
			return err
		}
		ds.Print.PrintDataOnConsole("Slot added successfully")
	} else {
		ds.Print.PrintDataOnConsole("Slots are not of 30 minutes")
	}
	return nil

}

func (ds *DoctorService) ShowAvailableSlotsBySpeciality(specialization entity.Specialization) error {
	specializedDoctors, err := ds.DoctorsRepo.GetDoctorsBySpeciality(specialization)
	if err != nil {
		return err
	}
	availableTimeSlots, err := ds.DoctorsRepo.GetAvailableTimeSlotsForAllDoctorsForSpecialization(specializedDoctors)
	if err != nil {
		return err
	}

	if len(availableTimeSlots) == 0 {
		ds.Print.PrintDataOnConsole("No Slots available")
		return errors.New("No Slots Avaialble")
	}

	for _, avaialbleDoctors := range availableTimeSlots {
		for _, slot := range avaialbleDoctors.SlotList {
			ds.Print.PrintDataOnConsole(fmt.Sprintf("Doctor: %s, startTime: %s, endTime: %s", avaialbleDoctors.Doctor.DoctorName, slot.Start, slot.End))
		}
	}
	return nil
}
