package userRepo

import (
	"connup/entity/user"
	"github.com/jackc/pgx/v5"
)

type psqlRepo struct {
	client *pgx.Conn
}

func (p *psqlRepo) Save(user *user.User) (*user.AccessUser, error) {
	//TODO implement me
	panic("implement me")
}

func (p *psqlRepo) FindById(id int) (*user.AccessUser, error) {
	//TODO implement me
	panic("implement me")
}

func (p *psqlRepo) FindByEmail(email string) (*user.AccessUser, error) {
	//TODO implement me
	panic("implement me")
}

func (p *psqlRepo) DeactivateAccount(id int) error {
	//TODO implement me
	panic("implement me")
}

func (p *psqlRepo) VerifyAccount(id int) error {
	//TODO implement me
	panic("implement me")
}

func NewPsqlRepo(client *pgx.Conn) UserRepo {
	return &psqlRepo{
		client: client,
	}
}
