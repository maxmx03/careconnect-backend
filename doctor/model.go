package doctor

type DoctorModel struct {
	DoctorID int    `json:"doctor_id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	CRM      string `json:"crm"`
	Username string `json:"username"`
	Password string `json:"password"`
}
