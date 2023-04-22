package service

import (
	"doctorsAppointment/entity"
	"doctorsAppointment/mode"
	"doctorsAppointment/repository"
	"fmt"
)

type PatientService struct {
	PatientRepo repository.PatientRepo
	Print       mode.Print
}

func NewPatientService(patientRepo repository.PatientRepo, print mode.Print) *PatientService {
	return &PatientService{
		PatientRepo: patientRepo,
		Print:       print,
	}
}

func (ps *PatientService) RegisterPatient(patient entity.Patient) error {
	err := ps.PatientRepo.RegisterPatient(patient)
	if err != nil {
		return err
	}
	ps.Print.PrintDataOnConsole(fmt.Sprintf("%s registered successfully", patient.PatientName))
	return nil
}
