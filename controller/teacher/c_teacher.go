package teacher

import (
	"context"

	"github.com/labstack/echo/v4"

	teacherEntity "go-management-auth-school/entity/teacher"
)



type TeacherService interface {
	FindAll(ctx context.Context, params *TeacherParams) (data []teacherEntity.Teacher, err error)
	SelectAll(ctx context.Context, parameter *TeacherParams) (data []teacherEntity.Teacher, err error)
	FindOne(ctx context.Context, params *TeacherParams) (data teacherEntity.Teacher, err error)
}

type teacherController struct {
	teacherService TeacherService
}

func NewTeacherController(service TeacherService) *teacherController {
	return &teacherController{
		teacherService: service,
	}
}

func (ctrl teacherController) InitializeRoutes(userRouter *echo.Group, adminRouter *echo.Group, staticRouter *echo.Group, authRouter *echo.Group) {

}


