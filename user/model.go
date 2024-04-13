package user

type UserModel struct {
  ID int
  Username string `json:"username"`
  Email string `json:"email"`
  Password string `json:"password"`
}
