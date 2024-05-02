package doctor

type MedicalPrescription struct {
	PrescriptionID int    `json:"prescription_id"`
	DoctorID       int    `json:"doctor_id"`
	PatientID      int    `json:"patient_id"`
	Date           string `json:"date"`
	Description    string `json:"description"`
}

type DoctorModel struct {
	DoctorID int    `json:"doctor_id"`
	Crm      string `json:"crm"`
	Username string `json:"username"`
	Password string `json:"password"`
}
