package app

import (
	"context"
	"github.com/go-logr/logr"
	"github.com/lits01/xiaozhan/pkg/configs"

	"github.com/lits01/xiaozhan/domain/user"
)

type User struct {
	config configs.Configuration
	log    *logr.Logger
	user   *user.User
}

func NewUser(user *user.User, config configs.Configuration, log *logr.Logger) *User {
	return &User{
		config: config,
		log:    log,
		user:   user,
	}
}

func (m *User) Create(ctx context.Context, req *UserCreateRequest) (*UserCreateResponse, error) {
	resp := &UserCreateResponse{}

	return resp, nil
}
