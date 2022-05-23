package handler

import (
	"backend/app/common/dto"
	"backend/app/domain/service"
	"bufio"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh/terminal"
)

type IUserHandler interface {
	GetAll() error
	ValidateUser() (dto.UserModel, error)
	Create() error
	Update() error
	Delete() error
}

type UserHandler struct {
	service.IUserService
}

func NewUserHandler(srv service.IUserService) (iUserHandler IUserHandler) {
	iUserHandler = &UserHandler{srv}
	return
}

func (h *UserHandler) GetAll() (err error) {
	userDtos, err := h.IUserService.GetAll()
	if err != nil {
		return
	}
	for _, userDto := range userDtos {
		fmt.Println(userDto.Name)
	}
	return
}

func (h *UserHandler) ValidateUser() (userDto dto.UserModel, err error) {
	fmt.Printf("username: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	username := scanner.Text()
	fmt.Printf("password: ")
	password, err := terminal.ReadPassword(0)
	fmt.Printf("\n")
	if err != nil {
		return
	}

	credsDto := dto.NewCredsModel(username, password)

	userDto, err = h.IUserService.ValidateUser(credsDto)
	if err != nil {
		return
	}
	fmt.Printf("validated user's name: %s\n", credsDto.Username)
	return
}

func (h *UserHandler) Create() (err error) {
	fmt.Printf("username: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	username := scanner.Text()
	fmt.Printf("password: ")
	password, err := terminal.ReadPassword(0)
	fmt.Printf("\n")
	if err != nil {
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword(password, 10)
	stringedHashedPassword := string(hashedPassword)

	userDto := dto.UserModel{
		Name:     username,
		Password: stringedHashedPassword,
	}

	err = h.IUserService.Create(userDto)
	if err != nil {
		fmt.Printf("Error creating user: %s\n", err)
		return
	}
	fmt.Println("new user has been created")
	return
}

func (h *UserHandler) Update() (err error) {
	userDto, err := h.ValidateUser()
	if err != nil {
		panic(err)
	}
	fmt.Printf("change username: u / change password: p; (u/p): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	option := scanner.Text()
	switch option {
	case "u":
		err = h.ChangeUserName(userDto)
	case "p":
		err = h.ChangePassword(userDto)
	default:
		fmt.Println("There is no such option. valid option is (u/p)")
		return
	}
	return
}

func (h *UserHandler) Delete() (err error) {
	userDto, err := h.ValidateUser()
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Are you sure? (y/n): ")
	scanner.Scan()
	switch scanner.Text() {
	case "y":
		err = h.IUserService.Delete(userDto)
		if err != nil {
			return
		}
		fmt.Println("The user has been deleted")
	case "n":
		fmt.Println("To delete user was canceled")
	default:
		fmt.Println("To delete user was canceled")
	}
	return
}

func (h *UserHandler) ChangeUserName(userDto dto.UserModel) (err error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("new username: ")
	scanner.Scan()
	newUsername := scanner.Text()
	userDto.Name = newUsername
	err = h.IUserService.Update(userDto)
	if err != nil {
		return
	}
	fmt.Printf("usesrname has been changed: new username is %s\n", userDto.Name)
	return
}

func (h *UserHandler) ChangePassword(userDto dto.UserModel) (err error) {
	fmt.Printf("new password: ")
	newPassword, err := terminal.ReadPassword(0)
	fmt.Printf("\n")
	if err != nil {
		return
	}

	hashedNewPassword, _ := bcrypt.GenerateFromPassword(newPassword, 10)
	stringedHashedNewPassword := string(hashedNewPassword)

	userDto.Password = stringedHashedNewPassword
	err = h.IUserService.Update(userDto)
	fmt.Printf("%s's password has been changed\n", userDto.Name)
	return
}
