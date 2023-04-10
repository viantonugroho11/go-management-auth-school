package auth

import (
	"context"
	"errors"

	authEntity "go-management-auth-school/entity/auth"
	userEntity "go-management-auth-school/entity/user"
	"go-management-auth-school/response"

	"github.com/labstack/echo/v4"
)



type authService interface {
	Login(ctx context.Context, input *LoginRequest) (data authEntity.Auth, err error)
	RegisterStudent(ctx context.Context, input *userEntity.User) (data authEntity.Auth, err error)
	ValidateToken(ctx context.Context, token string) (data authEntity.AuthValidate, err error)
}

type authController struct {
	service authService
}

func NewAuthController(service authService) *authController {
	return &authController{service: service}
}

func (ctrl authController) InitializeRoutes(userRouter *echo.Group, adminRouter *echo.Group, staticRouter *echo.Group, authRouter *echo.Group) {
	authRouter.POST("/login", ctrl.Login())
	authRouter.POST("/register", ctrl.RegisterStudent())
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

		dataService := reqRegister.ToService()
		data, err := ctrl.service.RegisterStudent(c.Request().Context(), dataService)
		if err != nil {
			return err
		}
		return response.RespondSuccess(c, 200, FromServiceLogin(data), nil)
	}
}

// validate JWT
func (ctrl authController) ValidateJWT() echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error

		ctx := c.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}

		// get token from header
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			err = errors.New("Token is required")
			return response.RespondError(c, 401, err)
		}

		// validate token
		_, err = ctrl.service.ValidateToken(ctx, token)
		if err != nil {
			return response.RespondError(c, 401, err)
		}

		return nil
	}
}

