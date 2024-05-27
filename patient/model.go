package patient

type PatientModel struct {
	PatientID   int    `json:"patient_id"`
	UserID      int    `json:"user_id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	CPF         string `json:"cpf"`
	DateOfBirth string `json:"date_of_birth"`
  Description string `json:"description"`
}
