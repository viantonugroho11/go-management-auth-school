package student

import (
	"context"
	"log"
	"net/http"

	// "go-management-auth-school/response"

	studentEntity "go-management-auth-school/entity/student"
	"go-management-auth-school/response"

	"github.com/labstack/echo/v4"
)


type StudentService interface {

	SelectAll(ctx context.Context, parameter *StudentParams) (data []studentEntity.Student, err error)
	FindOne(ctx context.Context, parameter *StudentParams) (data studentEntity.Student, err error)
	
	
}

type studentController struct {
	studentServices StudentService
}

func NewStudentController(studentServices StudentService) studentController {
	return studentController{
		studentServices: studentServices,
	}
}

func (ctrl studentController) InitializeRoutes(userRouter *echo.Group, adminRouter *echo.Group, staticRouter *echo.Group, authRouter *echo.Group) {
	userRouter.GET("/student", ctrl.GetStudent())
}

func (ctrl studentController) GetStudent() echo.HandlerFunc{
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}
		
		params := new(StudentParams)
		params.ID = c.QueryParam("id")
		params.Nik = c.QueryParam("nik")
		params.Nisn = c.QueryParam("nisn")
		params.Nis = c.QueryParam("nis")
		params.FirstName = c.QueryParam("first_name")
		params.LastName = c.QueryParam("last_name")
		log.Println(params)
		data, err := ctrl.studentServices.SelectAll(ctx, params)
		if err != nil {
			return err
		}
		return response.RespondSuccess(c,http.StatusAlreadyReported, FromServices(data),nil)
	}
		
} 