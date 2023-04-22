package service

import (
	"doctorsAppointment/entity"
	"doctorsAppointment/mode"
	"doctorsAppointment/repository"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type Queue []entity.Booking

func (q *Queue) Enqueue(value entity.Booking) {
	*q = append(*q, value)
}

func (q *Queue) Dequeue() entity.Booking {
	if len(*q) == 0 {
		return entity.Booking{}
	}
	value := (*q)[0]
	*q = (*q)[1:]
	return value
}

type BookingService struct {
	DoctorsRepo   repository.DoctorsRepo
	PatientRepo   repository.PatientRepo
	Print         mode.Print
	Bookings      map[int]entity.Booking
	PatientSlots  map[int][]entity.TimeSlot
	WaitListQueue Queue
}

func NewBookingService(doctorsRepo repository.DoctorsRepo, patientRepo repository.PatientRepo, print mode.Print, bookings map[int]entity.Booking, patientSlots map[int][]entity.TimeSlot,
	waitQueue Queue) *BookingService {
	return &BookingService{
		DoctorsRepo:   doctorsRepo,
		PatientRepo:   patientRepo,
		Print:         print,
		Bookings:      bookings,
		PatientSlots:  patientSlots,
		WaitListQueue: waitQueue,
	}
}

func (bs *BookingService) BookAppointment(patient entity.Patient, doctor entity.Doctor, fromSlot string) error {
	if !bs.PatientRepo.IsPatientRegistered(patient.PatientId) {
		return errors.New("Patient not registered")
	}
	if !bs.DoctorsRepo.IsDoctorRegistered(doctor.DoctorId) {
		return errors.New("Doctor not registered")
	}

	//check is patient already booked for that slot
	if _, exists := bs.PatientSlots[patient.PatientId]; exists {
		patientSlots := bs.PatientSlots[patient.PatientId]

		for _, slot := range patientSlots {
			if slot.Start == fromSlot {
				return errors.New("Patient already Occupied")
			}
		}
	} else {
		bs.PatientSlots[patient.PatientId] = make([]entity.TimeSlot, 0)
	}

	//check if doctor available for that slot
	doctorDetails, err := bs.DoctorsRepo.GetDoctorDetails(doctor.DoctorId)
	if err != nil {
		return err
	}
	slots := doctorDetails.Slots

	for key, val := range slots {
		if key.Start == fromSlot && val {
			slots[key] = false
			bs.PatientSlots[patient.PatientId] = append(bs.PatientSlots[patient.PatientId], key)
			bookingId := uuid.New().ClockSequence()
			booking := entity.Booking{
				BookingId: bookingId,
				Doctor:    doctor,
				Patient:   patient,
				Slot:      key,
			}
			bs.Bookings[bookingId] = booking
			bs.Print.PrintDataOnConsole(fmt.Sprintf("Appointment created successfully for ID: %s", bookingId))
			return nil
		}
	}

	bs.Print.PrintDataOnConsole("No Avaialble Slots")
	bookingId := uuid.New().ClockSequence()
	booking := entity.Booking{
		BookingId: bookingId,
		Doctor:    doctor,
		Patient:   patient,
		Slot:      entity.TimeSlot{Start: fromSlot, End: fromSlot},
		WaitList:  true,
	}
	bs.Print.PrintDataOnConsole(fmt.Sprintf("Added to waitList, booking Id: %s", bookingId))
	bs.WaitListQueue.Enqueue(booking)
	return nil

}

func (bs *BookingService) CancelBooking(bookingId int) error {
	if _, exists := bs.Bookings[bookingId]; !exists {
		return errors.New("Booking not Present")
	}

	booking := bs.Bookings[bookingId]
	err := bs.DoctorsRepo.FreeSlot(booking.Doctor.DoctorId, booking.Slot)
	if err != nil {
		return err
	}
	delete(bs.Bookings, bookingId)
	bs.Print.PrintDataOnConsole("Booking Cancelled")
	delete(bs.PatientSlots, booking.Patient.PatientId)
	bs.checkForFreeSlot(booking)
	return nil
}

func (bs *BookingService) ShowBookedAppointment() {
	bookings := bs.Bookings
	for _, booking := range bookings {
		bs.Print.PrintDataOnConsole(fmt.Sprintf("BookingId: %s, Doctor: %s, Patient: %s, SlotStart: %s", booking.BookingId, booking.Doctor.DoctorName, booking.Patient.PatientName, booking.Slot.Start))
	}
}

//After a booking cancelled, confirming a wait list booking
func (bs *BookingService) checkForFreeSlot(booking entity.Booking) {
	waitQueue := bs.WaitListQueue

	for idx, wBooking := range waitQueue {
		if wBooking.Slot.Start == booking.Slot.Start {
			wBooking.Slot.End = booking.Slot.End
			wBooking.WaitList = false
			doctor, err := bs.DoctorsRepo.GetDoctorDetails(booking.Doctor.DoctorId)
			if err != nil {

			}
			doctorSlots := doctor.Slots
			for key, _ := range doctorSlots {
				if key.Start == booking.Slot.Start {
					doctorSlots[key] = false
					break
				}
			}
			bs.Bookings[wBooking.BookingId] = wBooking
			waitQueue = append(waitQueue[:idx], waitQueue[idx+1:]...)
			bs.WaitListQueue = waitQueue
			return
		}
	}
}
