package class

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	classEntity "go-management-auth-school/entity/class"
	"go-management-auth-school/response"
)

type ClassService interface {
	FindAll(ctx context.Context, params *ClassParams) (data []classEntity.Class, err error)
	SelectAll(ctx context.Context, parameter *ClassParams) (data []classEntity.Class, err error)
	FindOne(ctx context.Context, params *ClassParams) (data classEntity.Class, err error)
	Create(ctx context.Context, params *classEntity.Class) (err error)
}

type classController struct {
	classService ClassService
}

func NewClassController(classService ClassService) classController {
	return classController{
		classService: classService,
	}
}

func (ctrl classController) InitializeRoutes(userRouter *echo.Group, adminRouter *echo.Group, staticRouter *echo.Group, authRouter *echo.Group) {
	userRouter.GET("/:id", ctrl.FindOne())
	userRouter.GET("/all", ctrl.SelectAll())
	userRouter.POST("", ctrl.Create())
}

// get one
func (ctrl classController) FindOne() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		ctx := c.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}

		params := new(ClassParams)
		params.ID = c.Param("id")
		data, err := ctrl.classService.FindOne(ctx, params)
		if err != nil {
			return response.RespondError(c, http.StatusBadRequest, err)
		}
		return response.RespondSuccess(c, http.StatusOK, FromService(data), nil)
	}
}

// get all
func (ctrl classController) SelectAll() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		ctx := c.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}

		params := new(ClassParams)
		params.ID = c.Param("id")
		data, err := ctrl.classService.SelectAll(ctx, params)
		if err != nil {
			return response.RespondError(c, http.StatusBadRequest, err)
		}
		return response.RespondSuccess(c, http.StatusOK, FromServices(data), nil)
	}
}

// create
func (ctrl classController) Create() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		ctx := c.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}
		reqData := new(ClassRequest)
		if err = c.Bind(reqData); err != nil {
			return response.RespondError(c, http.StatusBadRequest, err)
		}

		if err = reqData.Validate(); err != nil {
			return response.RespondError(c, http.StatusBadRequest, err)
		}

		params := reqData.ToService()
		err = ctrl.classService.Create(ctx, params)
		if err != nil {
			return response.RespondError(c, http.StatusInternalServerError, err)
		}
		return response.RespondSuccess(c, http.StatusCreated, nil, nil)
	}
}
