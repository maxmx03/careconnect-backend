package medicalconsultation

type MedicalConsultationModel struct {
	ConsultationID int    `json:"consultation_id"`
	DoctorID       int    `json:"doctor_id"`
	PatientID      int    `json:"patient_id"`
	DateTime       string `json:"date_time"`
	Description    string `json:"description"`
	DoctorName     string `json:"doctor_name"`
	DoctorSurname  string `json:"doctor_surname"`
	PatientName    string `json:"patient_name"`
	PatientSurname string `json:"patient_surname"`
}
