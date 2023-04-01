package auth

import "github.com/labstack/echo/v4"



type authService interface {
	Login(request LoginRequest) (string, error)
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
		request := new(LoginRequest)
		if err := c.Bind(request); err != nil {
			return err
		}
		token, err := ctrl.service.Login(*request)
		if err != nil {
			return err
		}
		return c.JSON(200, map[string]string{
			"token": token,
		})
	}
}

