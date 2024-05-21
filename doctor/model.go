package doctor

type DoctorModel struct {
	DoctorID int    `json:"doctor_id"`
	UserID   int    `json:"user_id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	CRM      string `json:"crm"`
}
