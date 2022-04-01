package user

import (
	"Outlet/entities"
	"Outlet/repository/hash"
	"errors"

	"github.com/lithammer/shortuuid"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Register(u entities.User) (entities.User, error) {

	u.Password, _ = hash.HashPassword(u.Password)
	uid := shortuuid.New()
	u.User_uid = uid
	u.Roles = false

	err := ur.db.Create(&u)
	if err != nil {
		return u, errors.New("email has been created")
	}

	return u, nil
}
