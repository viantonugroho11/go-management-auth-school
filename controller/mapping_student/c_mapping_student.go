package mapping_student

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"context"

	mapStudentEntity "go-management-auth-school/entity/mapping_student"
	"go-management-auth-school/response"
)

type MappingStudentService interface {
	FindAll(ctx context.Context, params *MappingStudentParams) (data []mapStudentEntity.MappingStudent, err error)
	SelectAll(ctx context.Context, parameter *MappingStudentParams) (data []mapStudentEntity.MappingStudent, err error)
	FindOne(ctx context.Context, params *MappingStudentParams) (data mapStudentEntity.MappingStudent, err error)
	Create(ctx context.Context, params *mapStudentEntity.MappingStudentReq) (err error)
}

type mappingStudentController struct {
	mappingStudentService MappingStudentService
}

func NewMappingStudentController(mappingStudentService MappingStudentService) mappingStudentController {
	return mappingStudentController{
		mappingStudentService: mappingStudentService,
	}
}

func (ctrl mappingStudentController) InitializeRoutes(userRouter *echo.Group, adminRouter *echo.Group, staticRouter *echo.Group, authRouter *echo.Group) {
	userRouter.GET("/all", ctrl.SelectAllMappingStudent())
	userRouter.POST("", ctrl.CreateMappingStudent())
	userRouter.GET("/one", ctrl.FindOneMappingStudent())
}

// get all mapping student
func (ctrl mappingStudentController) SelectAllMappingStudent() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}

		params := new(MappingStudentParams)
		params.ID = c.QueryParam("id")
		// params. = c.QueryParam("student_id")
		params.ClassID = c.QueryParam("class_id")
		data, err := ctrl.mappingStudentService.SelectAll(ctx, params)
		if err != nil {
			return response.RespondError(c, http.StatusBadRequest, err)
		}
		return response.RespondSuccess(c, http.StatusOK, FromServices(data), nil)
	}
}

// get one mapping student
func (ctrl mappingStudentController) FindOneMappingStudent() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}

		params := new(MappingStudentParams)
		params.ID = c.QueryParam("id")
		data, err := ctrl.mappingStudentService.FindOne(ctx, params)
		if err != nil {
			return response.RespondError(c, http.StatusBadRequest, err)
		}
		return response.RespondSuccess(c, http.StatusOK, FromService(data), nil)
	}
}

// create mapping student
func (ctrl mappingStudentController) CreateMappingStudent() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}
		reqMapStudent := new(MappingStudentRequest)
		if err := c.Bind(reqMapStudent); err != nil {
			return err
		}

		if err := reqMapStudent.Validate(); err != nil {
			return response.RespondError(c, http.StatusBadRequest, err)
		}

		reqData := reqMapStudent.ToService()
		if err := ctrl.mappingStudentService.Create(ctx, reqData); err != nil {
			return response.RespondError(c, http.StatusBadRequest, err)
		}
		return response.RespondSuccess(c, http.StatusCreated, nil, nil)
	}
}
