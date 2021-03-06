package service

import (
	"backend/app/common/dto"
	"backend/app/domain/entity"
	"backend/app/domain/repository"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type IUserService interface {
	GetAll() ([]dto.UserModel, error)
	ValidateUser(dto.CredentialsModel) (dto.UserModel, error)
	Create(dto.UserModel) error
	Update(dto.UserModel) error
	Delete(dto.UserModel) error
	IssueToken(int) (dto.AuthTokenModel, error)
	ValidateToken(dto.AuthTokenModel) (bool, error)
}

type UserService struct {
	repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) (userService IUserService) {
	userService = &UserService{repo}
	return
}

func (s *UserService) convertToDtosFromEntities(users []entity.User) (userDtos []dto.UserModel) {
	for _, user := range users {
		userDto := dto.NewUserModel(user.Id, user.Name, user.Password)
		userDtos = append(userDtos, userDto)
	}
	return
}

func (s *UserService) convertToDtoFromEntity(user entity.User) (userDto dto.UserModel) {
	userDto = dto.NewUserModel(user.Id, user.Name, user.Password)
	return
}

func (s *UserService) convertToEntityFromDto(userDto dto.UserModel) (user entity.User) {
	user = entity.NewUser(userDto.Id, userDto.Name, userDto.Password)
	return
}

func (s *UserService) convertToEntityFromDtoCreds(credsDto dto.CredentialsModel) (creds entity.Credentials) {
	creds = entity.NewCreds(credsDto.Username, credsDto.Password)
	return
}

func (s *UserService) GetAll() (userDtos []dto.UserModel, err error) {
	users, err := s.IUserRepository.GetAll()
	if err != nil {
		return
	}
	userDtos = s.convertToDtosFromEntities(users)
	return
}

func (s *UserService) ValidateUser(credsDto dto.CredentialsModel) (userDto dto.UserModel, err error) {
	creds := s.convertToEntityFromDtoCreds(credsDto)
	user, err := s.IUserRepository.ValidateUser(creds)
	userDto = s.convertToDtoFromEntity(user)
	return
}

func (s *UserService) Create(userDto dto.UserModel) (err error) {
	user := s.convertToEntityFromDto(userDto)
	err = s.IUserRepository.Create(user)
	return
}

func (s *UserService) Update(userDto dto.UserModel) (err error) {
	user := s.convertToEntityFromDto(userDto)
	err = s.IUserRepository.Update(user)
	return
}

func (s *UserService) Delete(userDto dto.UserModel) (err error) {
	user := s.convertToEntityFromDto(userDto)
	err = s.IUserRepository.Delete(user)
	return
}

func (s *UserService) IssueToken(userId int) (authTokenModel dto.AuthTokenModel, err error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	err = godotenv.Load(fmt.Sprint(".env", os.Getenv("GO_ENV")))
	if err != nil {
		panic(err)
	}

	secret := os.Getenv("SECRET_KEY")

	tokenString, _ := token.SignedString([]byte(secret))

	authTokenModel = dto.AuthTokenModel{
		Token: tokenString,
	}
	return
}

func (s *UserService) ValidateToken(authTokenDto dto.AuthTokenModel) (isAdmin bool, err error) {
	err = godotenv.Load(fmt.Sprint(".env", os.Getenv("GO_ENV")))
	if err != nil {
		panic(err)
	}

	secret := os.Getenv("SECRET_KEY")

	authToken, err := jwt.Parse(authTokenDto.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if claims, ok := authToken.Claims.(jwt.MapClaims); ok && authToken.Valid {
		userId := int(claims["user_id"].(float64))
		isAdmin, err = s.IUserRepository.IsAdmin(userId)
		if err != nil {
			return
		}
	} else {
		fmt.Println(err)
	}
	return

}
