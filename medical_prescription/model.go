package medicalprescription

type MedicalPrescriptionModel struct {
	PrescriptionID int    `json:"prescription_id"`
	Date           string `json:"date"`
	Description    string `json:"description"`
	DoctorID       int    `json:"doctor_id"`
	DoctorName     string `json:"doctor_name"`
	DoctorSurname  string `json:"doctor_surname"`
	PatientID      int    `json:"patient_id"`
	PatientName    string `json:"patient_name"`
	PatientSurname string `json:"patient_surname"`
}
