package repository

import (
	"context"
	"log"
	"xaia-backend/ent"
	"xaia-backend/ent/user"
	"xaia-backend/internal/api/auth/delivery/http/dtos"
)

type entUserRepo struct {
	client *ent.Client
}

func NewEntUserRepo(client *ent.Client) UserRepo {
	return &entUserRepo{client: client}
}

func (r *entUserRepo) FindByEmail(ctx context.Context, email string) (*ent.User, error) {

	u, err := r.client.User.Query().Where(user.Email(email)).Only(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return u, nil
}

func (r *entUserRepo) CreateNewUser(ctx context.Context, user dtos.RegisterUserPayload) (*ent.User, error) {

	u, err := r.client.User.
		Create().
		SetFirstname(user.Firstname).
		SetLastname(user.Lastname).
		SetPassword(user.Password).
		SetEmail(user.Email).
		Save(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return u, nil
}
