package user

import "Outlet/entities"

type User interface {
	Register(u entities.User) (entities.User, error)
}
