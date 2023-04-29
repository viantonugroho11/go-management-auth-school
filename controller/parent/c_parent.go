package parent

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	parentEntity "go-management-auth-school/entity/parent"
	"go-management-auth-school/response"
)

type ParentService interface {
	FindAll(ctx context.Context, params *ParentParams) (data []parentEntity.Parent, err error)
	SelectAll(ctx context.Context, parameter *ParentParams) (data []parentEntity.Parent, err error)
	FindOne(ctx context.Context, params *ParentParams) (data parentEntity.Parent, err error)
	Create(ctx context.Context, params *parentEntity.Parent) (err error)
}

type parentController struct {
	parentService ParentService
}

func NewParentController(parentService ParentService) parentController {
	return parentController{
		parentService: parentService,
	}
}

func (ctrl parentController) InitializeRoutes(userRouter *echo.Group, adminRouter *echo.Group, staticRouter *echo.Group, authRouter *echo.Group) {
	userRouter.GET("/all", ctrl.SelectAllParent())
	userRouter.POST("", ctrl.CreateParent())
	userRouter.GET("/:id", ctrl.FindOneParent())
}

// get all parent
func (ctrl parentController) SelectAllParent() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}

		params := new(ParentParams)
		params.ID = c.QueryParam("id")
		params.NIK = c.QueryParam("nik")
		params.FirstName = c.QueryParam("first_name")
		params.LastName = c.QueryParam("last_name")
		params.StudentID = c.QueryParam("student_id")
		data, err := ctrl.parentService.SelectAll(ctx, params)
		if err != nil {
			return response.RespondError(c, http.StatusBadRequest, err)
		}
		return response.RespondSuccess(c, http.StatusOK, FromServices(data), nil)
	}
}

// get one parent
func (ctrl parentController) FindOneParent() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}

		params := new(ParentParams)
		params.ID = c.Param("id")
		data, err := ctrl.parentService.FindOne(ctx, params)
		if err != nil {
			return response.RespondError(c, http.StatusBadRequest, err)
		}
		return response.RespondSuccess(c, http.StatusOK, FromService(data), nil)
	}
}

// create parent
func (ctrl parentController) CreateParent() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}

		reqParent := new(ParentRequest)
		if err := c.Bind(reqParent); err != nil {
			return err
		}

		if err := reqParent.Validate(); err != nil {
			return response.RespondError(c, http.StatusBadRequest, err)
		}

		reqData := reqParent.ToService()

		if err := ctrl.parentService.Create(ctx, reqData); err != nil {
			return response.RespondError(c, http.StatusBadRequest, err)
		}
		return response.RespondSuccess(c, http.StatusCreated, nil, nil)
	}
}
