package doctorconsultation

type MedicalConsultationModel struct {
	ConsultationID int    `json:"consultation_id"`
	DoctorID       string `json:"doctor_crm"`
	PatientID      int    `json:"patient_id"`
	DateTime       string `json:"date_time"`
	Description    string `json:"description"`
}
