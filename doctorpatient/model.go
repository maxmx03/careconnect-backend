package doctorpatient

type DoctorPatientModel struct {
	DoctorID           int    `json:"doctor_id"`
	DoctorCrm          string `json:"doctor_crm"`
	DoctorUsername     string `json:"doctor_username"`
	PatientID          int    `json:"patient_id"`
	PatientDateOfBirth string `json:"patient_date_of_birth"`
	PatientUsername    string `json:"patient_username"`
}
