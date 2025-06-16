package usecase

import (
	"context"
	"errors"
	"log"
	"xaia-backend/internal/api/auth/delivery/http/dtos"
	"xaia-backend/internal/api/auth/repository"
	"xaia-backend/internal/util"
)

type authUsecase struct {
	userRepo repository.UserRepo
}

func NewAuthUsecase(repo repository.UserRepo) AuthUsecase {
	return &authUsecase{userRepo: repo}
}

func (a *authUsecase) Login(ctx context.Context, email string, password string) (*dtos.LoginResponse, error) {

	//Retrieve user
	u, err := a.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	//Validate password
	err = util.VerifyPassword(u.Password, password)
	if err != nil {
		return nil, errors.New("invalid password")
	}

	token, err := util.GenerateJwtToken(u)
	if err != nil {
		return nil, err
	}

	userDTO := dtos.UserDTO{
		ID:        u.ID,
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
		Email:     u.Email,
	}

	data := dtos.LoginResponse{
		User:  userDTO,
		Token: token,
	}

	log.Println("Login business login")
	return &data, nil
}
func (a *authUsecase) Register(ctx context.Context, user dtos.RegisterUserPayload) (*dtos.UserDTO, error) {

	hashed, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, errors.New("could not hash password")
	}

	user.Password = string(hashed)

	u, err := a.userRepo.CreateNewUser(ctx, user)
	if err != nil {
		return nil, err
	}
	userDTO := dtos.UserDTO{
		ID:        u.ID,
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
		Email:     u.Email,
	}
	log.Println("Login business login")
	return &userDTO, nil
}
