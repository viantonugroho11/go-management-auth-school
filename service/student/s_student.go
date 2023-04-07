package student

import (
	"context"

	studentRequset "go-management-auth-school/controller/student"
	studentEntity "go-management-auth-school/entity/student"

	"github.com/jmoiron/sqlx"
)


type StudentRepo interface {
	SelectAll(ctx context.Context, parameter *studentRequset.StudentParams) (data []studentEntity.Student, err error)
	FindOne(ctx context.Context, parameter *studentRequset.StudentParams) (data studentEntity.Student, err error)
	Create(ctx context.Context,tx *sqlx.Tx, input *studentEntity.Student) (res string, err error)
	CreateTx(ctx context.Context) (tx *sqlx.Tx, err error)
}

type studentService struct {
	studentRepo StudentRepo
}

func NewStudentService(repo StudentRepo) *studentService {
	return &studentService{
		studentRepo: repo,
	}
}

func (service studentService) SelectAll(ctx context.Context, parameter *studentRequset.StudentParams) (data []studentEntity.Student, err error) {
	// parameter.Offset, parameter.Limit, parameter.Page, parameter.OrderBy, parameter.Sort =
	// 	service.SetPaginationParameter(parameter.Page, parameter.Limit, studentEntity.MapOrderBy[parameter.OrderBy], parameter.Sort, studentEntity.OrderBy, studentEntity.OrderByString)

	data, err = service.studentRepo.SelectAll(ctx, parameter)
	if err != nil {
		// logger.ErrorWithStack(ctx, err, "select all user query")
		return
	}
	return
}


func (service studentService) FindOne(ctx context.Context, parameter *studentRequset.StudentParams) (data studentEntity.Student, err error) {
	data, err = service.studentRepo.FindOne(ctx, parameter)
	if err != nil {
		// logger.ErrorWithStack(ctx, err, "select all user query")
		return
	}
	return
}