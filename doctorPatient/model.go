package doctorpatient

import (
	. "github.com/maxmx03/careconnect-backend/doctor"
	. "github.com/maxmx03/careconnect-backend/patient"
)

type DoctorPatient struct {
	Doctor  DoctorModel
	Patient PatientModel
}
