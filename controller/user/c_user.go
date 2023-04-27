package user

import (
	"context"
	"log"
	"net/http"

	userEntity "go-management-auth-school/entity/user"
	"go-management-auth-school/response"

	"github.com/labstack/echo/v4"
)

type UserService interface {
	SelectAll(ctx context.Context, parameter *UserParams) (data []userEntity.User, err error)
	FindOne(ctx context.Context, parameter *UserParams) (data userEntity.User, err error)
	Create(ctx context.Context, parameter *userEntity.User) (data userEntity.User, err error)
	UpdateUsername(ctx context.Context, parameter *userEntity.User) (err error)
}

type userController struct {
	userServices UserService
}

func NewUserController(service UserService) *userController {
	return &userController{
		userServices: service,
	}
}

func (ctrl userController) InitializeRoutes(userRouter *echo.Group, adminRouter *echo.Group, staticRouter *echo.Group, authRouter *echo.Group) {
	userRouter.GET("", ctrl.GetUser())
}

func (ctrl userController) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}

		params := new(UserParams)
		log.Println(params)
		data, err := ctrl.userServices.SelectAll(ctx, params)
		if err != nil {
			return err
		}
		return response.RespondSuccess(c, http.StatusAlreadyReported, data, nil)
	}
}
