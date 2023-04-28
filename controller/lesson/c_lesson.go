package lesson

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	lessonEntity "go-management-auth-school/entity/lesson"
	"go-management-auth-school/response"
)

type LessonService interface {
	FindAll(ctx context.Context, params *LessonParams) (data []lessonEntity.Lesson, err error)
	SelectAll(ctx context.Context, parameter *LessonParams) (data []lessonEntity.Lesson, err error)
	FindOne(ctx context.Context, params *LessonParams) (data lessonEntity.Lesson, err error)
	Create(ctx context.Context, params *lessonEntity.Lesson) (err error)
}

type lessonController struct {
	lessonService LessonService
}

func NewLessonController(lessonService LessonService) lessonController {
	return lessonController{
		lessonService: lessonService,
	}
}

func (ctrl lessonController) InitializeRoutes(userRouter *echo.Group, adminRouter *echo.Group, staticRouter *echo.Group, authRouter *echo.Group) {

}

// get one lesson
func (ctrl lessonController) FindOne() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		ctx := c.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}

		params := new(LessonParams)
		params.ID = c.Param("id")
		data, err := ctrl.lessonService.FindOne(ctx, params)
		if err != nil {
			return response.RespondError(c, http.StatusBadRequest, err)
		}
		return response.RespondSuccess(c, http.StatusOK, FromService(data), nil)
	}
}

// get all lesson
func (ctrl lessonController) FindAll() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		ctx := c.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}

		params := new(LessonParams)
		params.ID = c.QueryParam("id")

		data, err := ctrl.lessonService.FindAll(ctx, params)
		if err != nil {
			return response.RespondError(c, http.StatusBadRequest, err)
		}
		return response.RespondSuccess(c, http.StatusOK, FromServices(data), nil)
	}
}

// create lesson
func (ctrl lessonController) Create() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		ctx := c.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}

		reqData := new(LessonRequest)
		if err = c.Bind(reqData); err != nil {
			return response.RespondError(c, http.StatusBadRequest, err)
		}

		if err = reqData.Validate(); err != nil {
			return response.RespondError(c, http.StatusBadRequest, err)
		}

		params := reqData.ToService()
		if err = ctrl.lessonService.Create(ctx, params); err != nil {
			return response.RespondError(c, http.StatusInternalServerError, err)
		}
		return response.RespondSuccess(c, http.StatusCreated, nil, nil)
	}
}
