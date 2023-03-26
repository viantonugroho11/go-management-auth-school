package student

import (
	"context"

	studentRequset "go-management-auth-school/controller/student"
	studentEntity "go-management-auth-school/entity/student"
)


type StudentRepo interface {
	SelectAll(ctx context.Context, parameter *studentRequset.StudentParams) (data []studentEntity.Student, err error)
}

type studentService struct {
	studentRepo StudentRepo
}

func NewStudentService(repo StudentRepo) *studentService {
	return &studentService{
		studentRepo: repo,
	}
}

func (service studentService)SelectAll(ctx context.Context, parameter *studentRequset.StudentParams) (data []studentEntity.Student, err error) {
	// parameter.Offset, parameter.Limit, parameter.Page, parameter.OrderBy, parameter.Sort =
	// 	service.SetPaginationParameter(parameter.Page, parameter.Limit, studentEntity.MapOrderBy[parameter.OrderBy], parameter.Sort, studentEntity.OrderBy, studentEntity.OrderByString)

	data, err = service.studentRepo.SelectAll(ctx, parameter)
	if err != nil {
		// logger.ErrorWithStack(ctx, err, "select all user query")
		return
	}
	return
}