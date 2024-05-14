package message

type MessageModel struct { 
	MessageID      int    `json:"message_id"`
	DateTime       string `json:"date"`
	Content        string `json:"content"`
	DoctorID       int    `json:"doctor_id"`
	DoctorName     string `json:"doctor_name"`
	DoctorSurname  string `json:"doctor_surname"`
	PatientID      int    `json:"patient_id"`
	PatientName    string `json:"patient_name"`
	PatientSurname string `json:"patient_surname"`
}
