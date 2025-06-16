package usecase

import (
	"context"
	"xaia-backend/internal/api/auth/delivery/http/dtos"
)

/*
AuthUsecase is an interface that defines authentication-related use cases.
Any struct that implements this interface should provide logic for user login and registration.

Methods:
- Login: Authenticates a user using email and password.
- Register: Registers a new user with the provided user information.
*/
type AuthUsecase interface {
	Login(ctx context.Context, email string, password string) (*dtos.LoginResponse, error)
	Register(ctx context.Context, user dtos.RegisterUserPayload) (*dtos.UserDTO, error)
}
