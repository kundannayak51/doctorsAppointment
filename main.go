package main

import (
	"doctorsAppointment/entity"
	"doctorsAppointment/mode"
	"doctorsAppointment/repository"
	"doctorsAppointment/service"
	"fmt"
)

func main() {

	consolePrint := mode.NewConsolePrint()
	doctorsRepo := repository.NewDoctorsRepo(make(map[int]entity.Doctor), make(map[entity.Specialization][]entity.Doctor))
	patientRepo := repository.NewPatientRepo(make(map[int]entity.Patient))

	doctorService := service.NewDoctorService(*doctorsRepo, consolePrint)
	patientService := service.NewPatientService(*patientRepo, consolePrint)
	bookingService := service.NewBookingService(*doctorsRepo, *patientRepo, consolePrint, make(map[int]entity.Booking), make(map[int][]entity.TimeSlot), make(service.Queue, 0))

	doctor1 := entity.Doctor{
		DoctorId:       1,
		DoctorName:     "Curipus",
		Slots:          make(map[entity.TimeSlot]bool),
		Specialization: entity.Cardiologist,
		Rating:         1,
	}

	doctor2 := entity.Doctor{
		DoctorId:       2,
		DoctorName:     "Dredful",
		Slots:          make(map[entity.TimeSlot]bool),
		Specialization: entity.Dermatologist,
		Rating:         1,
	}

	doctor3 := entity.Doctor{
		DoctorId:       3,
		DoctorName:     "Daring",
		Slots:          make(map[entity.TimeSlot]bool),
		Specialization: entity.Dermatologist,
		Rating:         1,
	}

	patient1 := entity.Patient{
		PatientId:   1,
		PatientName: "A",
	}
	patient2 := entity.Patient{
		PatientId:   2,
		PatientName: "B",
	}
	patient3 := entity.Patient{
		PatientId:   3,
		PatientName: "C",
	}

	err := doctorService.RegisterDoctor(doctor1)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = doctorService.RegisterDoctor(doctor2)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = doctorService.RegisterDoctor(doctor3)
	if err != nil {
		fmt.Println(err.Error())
	}

	//Add invalid avaialability
	err = doctorService.AddAvailability(doctor1.DoctorId, entity.TimeSlot{Start: "9:30", End: "10:30"})
	if err != nil {
		fmt.Println(err.Error())
	}

	//Add valid
	err = doctorService.AddAvailability(doctor1.DoctorId, entity.TimeSlot{Start: "9:30", End: "10:00"})
	if err != nil {
		fmt.Println(err.Error())
	}
	err = doctorService.AddAvailability(doctor1.DoctorId, entity.TimeSlot{Start: "12:30", End: "13:00"})
	if err != nil {
		fmt.Println(err.Error())
	}
	err = doctorService.AddAvailability(doctor1.DoctorId, entity.TimeSlot{Start: "16:00", End: "10:30"})
	if err != nil {
		fmt.Println(err.Error())
	}
	err = doctorService.AddAvailability(doctor2.DoctorId, entity.TimeSlot{Start: "12:30", End: "13:00"})
	if err != nil {
		fmt.Println(err.Error())
	}
	err = doctorService.AddAvailability(doctor2.DoctorId, entity.TimeSlot{Start: "13:07", End: "13:37"})
	if err != nil {
		fmt.Println(err.Error())
	}

	err = doctorService.ShowAvailableSlotsBySpeciality(entity.Cardiologist)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = patientService.RegisterPatient(patient1)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = patientService.RegisterPatient(patient2)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = patientService.RegisterPatient(patient3)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = bookingService.BookAppointment(patient1, doctor1, "12:30")
	if err != nil {
		fmt.Println(err.Error())
	}

	err = doctorService.ShowAvailableSlotsBySpeciality(entity.Cardiologist)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = bookingService.BookAppointment(patient2, doctor1, "12:30")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = bookingService.BookAppointment(patient3, doctor1, "12:30")
	if err != nil {
		fmt.Println(err.Error())
	}

	bookingService.ShowBookedAppointment()

	err = doctorService.ShowAvailableSlotsBySpeciality(entity.Cardiologist)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = bookingService.CancelBooking(1)

	bookingService.ShowBookedAppointment()

	err = bookingService.BookAppointment(patient3, doctor2, "13:07")
	if err != nil {
		fmt.Println(err.Error())
	}

	err = doctorService.ShowAvailableSlotsBySpeciality(entity.Dermatologist)
	if err != nil {
		fmt.Println(err.Error())
	}

}
