package auth

import "Outlet/entities"

type Auth interface {
	Login(email, password string) (entities.User, error)
}
