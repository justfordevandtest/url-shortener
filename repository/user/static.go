package user

import (
	"shorturl/entity"
)

type StaticRepo struct {
}

func New() (repo *StaticRepo, err error) {
	return &StaticRepo{}, nil
}

func (repo *StaticRepo) Read(username string) (user *entity.User, err error) {
	user = &entity.User{
		Username: username,
		Password: "a",
	}
	user.HashPassword()

	return user, nil
}
