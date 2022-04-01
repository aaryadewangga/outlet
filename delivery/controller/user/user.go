package user

import (
	"Outlet/delivery/controller/common"
	"Outlet/entities"
	"Outlet/repository/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	repo user.User
}

func New(repo user.User) *UserController {
	return &UserController{
		repo: repo,
	}
}

func (uc *UserController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := CreateUserRequestFormat{}

		errB := c.Bind(user)
		if errB != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There are something problem from input", nil))
		}

		errV := c.Validate(&user)
		if errV != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There are something problem from input", nil))
		}

		res, err := uc.repo.Register(entities.User{
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		})

		if err != nil {
			return c.JSON(http.StatusConflict, common.InternalServerError(http.StatusConflict, err.Error(), nil))
		}

		response := UserCreateResponse{}
		response.User_uid = res.User_uid
		response.Name = res.Name
		response.Email = res.Email
		response.Roles = res.Roles

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success Create User", response))
	}
}
