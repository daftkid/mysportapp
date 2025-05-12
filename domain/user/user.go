package user

import (
	avatar2 "github.com/daftkid/mysportapp/domain/avatar"
	"github.com/daftkid/mysportapp/domain/credentials"
	"github.com/daftkid/mysportapp/domain/target"
)

type User struct {
	id      string
	name    string
	bio     string
	avatar  avatar2.Avatar
	creds   credentials.Credentials
	targets []target.Target
}

type UserRepository interface {
	FindAll() ([]User, error)
}
