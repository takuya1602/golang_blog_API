package dto

type UserModel struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

type CredentialsModel struct {
	Username string
	Password []byte
}

func NewUserModel(id int, name string, password string) (userModel UserModel) {
	userModel = UserModel{
		Id:       id,
		Name:     name,
		Password: password,
		IsAdmin:  true,
	}
	return
}

func NewCredsModel(username string, password []byte) (credsModel CredentialsModel) {
	credsModel = CredentialsModel{
		Username: username,
		Password: password,
	}
	return
}
