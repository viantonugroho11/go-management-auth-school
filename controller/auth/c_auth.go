package auth

import (
	"context"

	userEntity "go-management-auth-school/entity/user"
	authEntity "go-management-auth-school/entity/auth"
	"go-management-auth-school/response"

	"github.com/labstack/echo/v4"
)



type authService interface {
	Login(ctx context.Context, input *LoginRequest) (data authEntity.Auth, err error)
	RegisterStudent(ctx context.Context, input *userEntity.User) (data string, err error)
}

type authController struct {
	service authService
}

func NewAuthController(service authService) *authController {
	return &authController{service: service}
}

func (ctrl authController) InitializeRoutes(userRouter *echo.Group, adminRouter *echo.Group, staticRouter *echo.Group, authRouter *echo.Group) {
	userRouter.POST("/login", ctrl.Login())
}

func (ctrl authController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {

		var err error
		reqLogin := new(LoginRequest)
		if err := c.Bind(reqLogin); err != nil {
			return err
		}

		if err = reqLogin.Validate(); err != nil {
			return err
		}

		data, err := ctrl.service.Login(c.Request().Context(), reqLogin)
		if err != nil {
			return err
		}
		return response.RespondSuccess(c, 200,FromServiceLogin(data),nil)
	}
}

func (ctrl authController) RegisterStudent() echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error
		reqRegister := new(RegisterRequest)
		if err := c.Bind(reqRegister); err != nil {
			return err
		}
		if err = reqRegister.Validate(); err != nil {
			return err
		}

		data := reqRegister.ToService()
		// token, err := ctrl.service.RegisterStudent(c.Request().Context(), reqRegister)
		// if err != nil {
		// 	return err
		// }
		return c.JSON(200, map[string]string{
			"token": token,
		})
	}
}

