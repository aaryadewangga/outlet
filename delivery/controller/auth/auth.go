package auth

import (
	"Outlet/delivery/controller/common"
	"Outlet/delivery/middlewares"
	"Outlet/repository/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	repo auth.Auth
}

func New(repo auth.Auth) *AuthController {
	return &AuthController{
		repo: repo,
	}
}

func (ac *AuthController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserLogin := LoginReqFormat{}

		errB := c.Bind(&UserLogin)
		if errB != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There are something problem from input", nil))
		}

		errV := c.Validate(&UserLogin)
		if errV != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		res, err := ac.repo.Login(UserLogin.Email, UserLogin.Password)
		if err != nil {
			var statusCode int
			if err.Error() == "email not found" {
				statusCode = http.StatusUnauthorized
			} else if err.Error() == "invalid password" {
				statusCode = http.StatusUnauthorized
			}
			return c.JSON(statusCode, common.InternalServerError(statusCode, err.Error(), nil))
		}

		token, errT := middlewares.GenerateToken(res)
		if errT != nil {
			return c.JSON(http.StatusNotAcceptable, common.BadRequest(http.StatusNotAcceptable, "error in process token", nil))
		}

		response := UserLoginResponse{}
		response.User_uid = res.User_uid
		response.Name = res.Name
		response.Email = res.Email
		response.Roles = res.Roles
		response.Token = token

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Login Success, Get a new token", response))

	}
}
