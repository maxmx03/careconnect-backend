package patient

type PatientModel struct {
	PatientID   int    `json:"patient_id"`
	Name        string `json:"name"`
	CPF         string `json:"cpf"`
	DateOfBirth string `json:"date_of_birth"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}
