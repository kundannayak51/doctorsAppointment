package repository

import (
	"doctorsAppointment/entity"
	"errors"
)

type PatientRepo struct {
	Patients map[int]entity.Patient
}

func NewPatientRepo(patients map[int]entity.Patient) *PatientRepo {
	return &PatientRepo{
		Patients: patients,
	}
}

func (pr *PatientRepo) RegisterPatient(patient entity.Patient) error {
	if _, exits := pr.Patients[patient.PatientId]; exits {
		return errors.New("Patient already registered")
	}
	pr.Patients[patient.PatientId] = patient
	return nil
}

func (pr *PatientRepo) IsPatientRegistered(patientId int) bool {
	_, exists := pr.Patients[patientId]
	return exists
}
