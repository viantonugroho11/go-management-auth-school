package lesson

import (
	"context"

	"github.com/labstack/echo/v4"

	lessonEntity "go-management-auth-school/entity/lesson"
)

type LessonService interface {
	FindAll(ctx context.Context, params *LessonParams) (data []lessonEntity.Lesson, err error)
	SelectAll(ctx context.Context, parameter *LessonParams) (data []lessonEntity.Lesson, err error)
	FindOne(ctx context.Context, params *LessonParams) (data lessonEntity.Lesson, err error)
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