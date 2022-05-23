package entity

type User struct {
	Id       int
	Name     string
	Password string
	IsAdmin  bool
}

type Credentials struct {
	Username string
	Password []byte
}

func NewUser(id int, name string, password string) (user User) {
	user = User{
		Id:       id,
		Name:     name,
		Password: password,
		IsAdmin:  true,
	}
	return
}

func NewCreds(username string, password []byte) (creds Credentials) {
	creds = Credentials{
		Username: username,
		Password: password,
	}
	return
}
