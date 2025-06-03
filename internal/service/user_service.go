package service

import (
	"errors"

	"github.com/guide-backend/internal/model"
	"github.com/guide-backend/internal/repository"
	"github.com/guide-backend/pkg/helpers"
	"github.com/guide-backend/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo repository.UserRepo
	jwt      jwt.JWTService
}

func NewUserService(userRepo repository.UserRepo, jwt jwt.JWTService) UserService {
	return UserService{
		userRepo: userRepo,
		jwt:      jwt,
	}
}

func (s UserService) CreateUser(req CreateUserRequest) (CreateUserResponse, error) {
	hashedPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		return CreateUserResponse{}, err
	}
	user := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Balance:  100,
	}

	createdUser, err := s.userRepo.CreateUser(user)
	if err != nil {
		return CreateUserResponse{}, err
	}

	resp := CreateUserResponse{
		Name:      createdUser.Name,
		Email:     createdUser.Email,
		Balance:   createdUser.Balance,
		CreatedAt: createdUser.CreatedAt,
	}

	return resp, nil
}

func (s UserService) Login(req LoginRequest) (LoginResponse, error) {
	user, err := s.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		return LoginResponse{}, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return LoginResponse{}, errors.New("invalid email or password")
	}

	token, err := s.jwt.GenerateToken(int(user.ID))
	if err != nil {
		return LoginResponse{}, err
	}

	return LoginResponse{Token: token}, nil
}

func (s UserService) ListAllUsers() ([]ListUserResponse, error) {
	users, err := s.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var responses []ListUserResponse
	for _, user := range users {
		responses = append(responses, ListUserResponse{
			ID:      user.ID,
			Name:    user.Name,
			Email:   user.Email,
			Balance: user.Balance,
		})
	}
	return responses, nil
}

func (s UserService) GetUserByID(id uint) (ListUserResponse, error) {
	user, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return ListUserResponse{}, err
	}

	listUserResponse := ListUserResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Balance: user.Balance,
	}
	return listUserResponse, nil
}
