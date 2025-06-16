package repository

import (
	"context"
	"xaia-backend/ent"
	"xaia-backend/internal/api/auth/delivery/http/dtos"
)

type UserRepo interface {
	FindByEmail(ctx context.Context, email string) (*ent.User, error)
	CreateNewUser(cts context.Context, user dtos.RegisterUserPayload) (*ent.User, error)
}
