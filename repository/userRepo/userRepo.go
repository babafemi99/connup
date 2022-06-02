package userRepo

import "connup/entity/user"

type UserRepo interface {
	Save(user *user.User) (*user.AccessUser, error)
	FindById(id int) (*user.AccessUser, error)
	FindByEmail(email string) (*user.AccessUser, error)
	DeactivateAccount(id int) error
	VerifyAccount(id int) error
}
