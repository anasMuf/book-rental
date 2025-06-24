package service

import (
	"book-rental-app/model"
	"book-rental-app/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetAllUser() ([]model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	CreateUser(req model.User) (*model.User, error)
	Login(email, password string) (*model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo}
}

func (s *userService) GetAllUser() ([]model.User, error) {
	return s.userRepo.GetAll()
}

func (s *userService) GetUserByEmail(email string) (*model.User, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("user tidak ditemukan")
	}
	return user, nil
}

func (s *userService) CreateUser(req model.User) (*model.User, error) {
	_, err := s.userRepo.FindByEmail(req.Email)
	if err == nil {
		return nil, errors.New("email sudah terdaftar")
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	req.Password = string(hash)

	err = s.userRepo.Create(&req)
	return &req, err
}

func (s *userService) Login(email, password string) (*model.User, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("user tidak ditemukan")
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, errors.New("password salah")
	}
	return user, nil
}
