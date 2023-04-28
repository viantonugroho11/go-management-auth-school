package major

import (
	"context"
	"net/http"

	majorEntity "go-management-auth-school/entity/major"
	"go-management-auth-school/response"

	"github.com/labstack/echo/v4"
)

type MajorService interface {
	FindAll(ctx context.Context, params *MajorParams) (data []majorEntity.Major, err error)
	SelectAll(ctx context.Context, parameter *MajorParams) (data []majorEntity.Major, err error)
	FindOne(ctx context.Context, params *MajorParams) (data majorEntity.Major, err error)
	Create(ctx context.Context, params *majorEntity.Major) (err error)
}

type majorController struct {
	majorService MajorService
}

func NewMajorController(majorServices MajorService) majorController {
	return majorController{
		majorService: majorServices,
	}
}

func (ctrl majorController) InitializeRoutes(userRouter *echo.Group, adminRouter *echo.Group, staticRouter *echo.Group, authRouter *echo.Group) {
	userRouter.GET("", ctrl.SelectAll())
	userRouter.GET("/:id", ctrl.FindOne())
	userRouter.POST("", ctrl.Create())
}

// get all major
func (ctrl majorController) SelectAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		params := new(MajorParams)
		if err := c.Bind(params); err != nil {
			return err
		}
		data, err := ctrl.majorService.SelectAll(ctx, params)
		if err != nil {
			return response.RespondError(c, 400, err)
		}
		return response.RespondSuccess(c, 200, FromServices(data), nil)

	}
}

// get one major
func (ctrl majorController) FindOne() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		params := new(MajorParams)
		if err := c.Bind(params); err != nil {
			return err
		}
		data, err := ctrl.majorService.FindOne(ctx, params)
		if err != nil {
			return response.RespondError(c, 400, err)
		}
		return response.RespondSuccess(c, 200, FromService(data), nil)

	}
}

// create major
func (ctrl majorController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		params := new(MajorRequest)
		if err := c.Bind(params); err != nil {
			return err
		}

		if err := params.Validate(); err != nil {
			return response.RespondError(c, http.StatusBadRequest, err)
		}

		reqData := params.ToService()

		err := ctrl.majorService.Create(ctx, reqData)
		if err != nil {
			return response.RespondError(c, 400, err)
		}
		return response.RespondSuccess(c, 200, nil, nil)

	}
}
