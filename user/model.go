package user

type UserModel struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
}
