package dto

type UserModel struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

type CredentialsModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthTokenModel struct {
	Token string `json:"token"`
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

func NewCredsModel(username string, password string) (credsModel CredentialsModel) {
	credsModel = CredentialsModel{
		Username: username,
		Password: password,
	}
	return
}

func NewAuthTokenModel(token string) (authTokenModel AuthTokenModel) {
	authTokenModel = AuthTokenModel{
		Token: token,
	}
	return
}
