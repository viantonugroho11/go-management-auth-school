package mapping_course

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	mapCourseEntity "go-management-auth-school/entity/mapping_course"
	"go-management-auth-school/response"
)

type MappingCourseService interface {
	FindAll(ctx context.Context, params *MappingCourseParams) (data []mapCourseEntity.MappingCourse, err error)
	SelectAll(ctx context.Context, parameter *MappingCourseParams) (data []mapCourseEntity.MappingCourse, err error)
	FindOne(ctx context.Context, params *MappingCourseParams) (data mapCourseEntity.MappingCourse, err error)
	Create(ctx context.Context, params *mapCourseEntity.MappingCourseReq) (err error)
}

type mappingCourseController struct {
	mappingCourseService MappingCourseService
}

func NewMappingCourseController(mappingCourseService MappingCourseService) mappingCourseController {
	return mappingCourseController{
		mappingCourseService: mappingCourseService,
	}
}

func (ctrl mappingCourseController) InitializeRoutes(userRouter *echo.Group, adminRouter *echo.Group, staticRouter *echo.Group, authRouter *echo.Group) {
	userRouter.POST("", ctrl.Create())
	userRouter.GET("/all", ctrl.SelectAll())
	userRouter.GET("/:id", ctrl.FindOne())
}

// func (ctrl mappingCourseController) FindAll(c echo.Context) echo.HandlerFunc {
// 	return
// }

func (ctrl mappingCourseController) SelectAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}

		params := new(MappingCourseParams)
		params.ID = c.QueryParam("id")
		// params. = c.QueryParam("student_id")
		params.ClassID = c.QueryParam("class_id")
		data, err := ctrl.mappingCourseService.SelectAll(ctx, params)
		if err != nil {
			return response.RespondError(c, http.StatusBadRequest, err)
		}
		return response.RespondSuccess(c, http.StatusOK, FromServices(data), nil)
	}
}

func (ctrl mappingCourseController) FindOne() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}

		params := new(MappingCourseParams)
		params.ID = c.QueryParam("id")
		// params. = c.QueryParam("student_id")
		params.ClassID = c.QueryParam("class_id")
		data, err := ctrl.mappingCourseService.FindOne(ctx, params)
		if err != nil {
			return response.RespondError(c, http.StatusBadRequest, err)
		}
		return response.RespondSuccess(c, http.StatusOK, FromService(data), nil)
	}
}

func (ctrl mappingCourseController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}

		reqData := new(MappingCourseRequest)
		if err := c.Bind(reqData); err != nil {
			return err
		}

		if err := reqData.Validate(); err != nil {
			return err
		}

		params := reqData.ToService()

		err := ctrl.mappingCourseService.Create(ctx, params)
		if err != nil {
			return response.RespondError(c, http.StatusBadRequest, err)
		}
		return response.RespondSuccess(c, http.StatusCreated, nil, nil)
	}
}
