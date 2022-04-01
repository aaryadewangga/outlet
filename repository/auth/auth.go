package auth

import (
	"Outlet/entities"
	"Outlet/repository/hash"
	"errors"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (ar *AuthRepository) Login(email, password string) (entities.User, error) {

	user := entities.User{}

	err := ar.db.Model(&user).Where("email = ?", email).First(&user)
	if err != nil {
		return user, errors.New("email nof found")
	}

	match := hash.CheckPasswordHash(password, user.Password)
	if !match {
		return user, errors.New("invalid password")
	}

	return user, nil
}
