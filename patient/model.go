package patient

type PatientModel struct {
	PatientID   int    `json:"user_id"`
	Cpf         string `json:"name"`
	DateOfBirth string `json:"email"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}
