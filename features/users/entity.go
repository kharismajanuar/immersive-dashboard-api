package users

import "github.com/labstack/echo/v4"

type Core struct {
	Id       uint
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	Role     string `validate:"required"`
	Team     string `validate:"required"`
	Status   string `validate:"required"`
}

type UserDelivery interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
}

type UserService interface {
	RegisterSrv(newUser Core) error
	LoginSrv(email, password string) (string, Core, error)
}

type UserData interface {
	RegisterData(newUser Core) error
	LoginData(email string) (Core, error)
}