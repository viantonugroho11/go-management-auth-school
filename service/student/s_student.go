package student

import (
	"context"
	"errors"

	// "log"

	studentRequset "go-management-auth-school/controller/student"
	studentEntity "go-management-auth-school/entity/student"

	parentController "go-management-auth-school/controller/parent"
	parentRepository "go-management-auth-school/service/parent"

	"github.com/jmoiron/sqlx"
)

type StudentRepo interface {
	SelectAll(ctx context.Context, parameter *studentRequset.StudentParams) (data []studentEntity.Student, err error)
	FindOne(ctx context.Context, parameter *studentRequset.StudentParams) (data studentEntity.Student, err error)
	Create(ctx context.Context, tx *sqlx.Tx, input *studentEntity.Student) (res string, err error)
	CreateTx(ctx context.Context) (tx *sqlx.Tx, err error)
}

type studentService struct {
	studentRepo StudentRepo
	parentRepo parentRepository.ParentRepo
}

func NewStudentService(repo StudentRepo, parentRepo parentRepository.ParentRepo) *studentService {
	return &studentService{
		studentRepo: repo,
		parentRepo: parentRepo,
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

	// add parent
	for i, v := range data {
		data[i].Parent, err = service.parentRepo.SelectAll(ctx, &parentController.ParentParams{
			StudentID: v.Nis,
		})
	}
	
	return
}

func (service studentService) FindOne(ctx context.Context, parameter *studentRequset.StudentParams) (data studentEntity.Student, err error) {
	data, err = service.studentRepo.FindOne(ctx, parameter)
	if err != nil {
		// logger.ErrorWithStack(ctx, err, "select all user query")
		return
	}
	data.Parent, err = service.parentRepo.SelectAll(ctx, &parentController.ParentParams{
		StudentID: data.Nis,
	})
	return
}

func (service studentService) Create(ctx context.Context, input *studentEntity.Student) (err error) {
	checkStudent, err := service.studentRepo.FindOne(ctx, &studentRequset.StudentParams{
		Nis:  input.Nis,
		Nisn: input.Nisn,
	})
	if err != nil {
		return
	}

	if checkStudent.Nis != "" || checkStudent.Nisn != "" {
		err = errors.New("NIS already exist")
		return
	}

	tx, err := service.studentRepo.CreateTx(ctx)
	if err != nil {
		return
	}
	_, err = service.studentRepo.Create(ctx, tx, input)
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}
