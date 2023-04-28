package mapping_course

import (
	"context"
	"errors"

	"go-management-auth-school/controller"
	mapCourseController "go-management-auth-school/controller/mapping_course"
	mapCourseEntity "go-management-auth-school/entity/mapping_course"

	classServices "go-management-auth-school/controller/class"
	lessonServices "go-management-auth-school/controller/lesson"
	teacherServices "go-management-auth-school/controller/teacher"

	"github.com/jmoiron/sqlx"
)

type MpCourseRepo interface {
	FindAll(ctx context.Context, params *mapCourseController.MappingCourseParams) (data []mapCourseEntity.MappingCourse, err error)
	SelectAll(ctx context.Context, parameter *mapCourseController.MappingCourseParams) (data []mapCourseEntity.MappingCourse, err error)
	FindOne(ctx context.Context, params *mapCourseController.MappingCourseParams) (data mapCourseEntity.MappingCourse, err error)
	Create(ctx context.Context,tx *sqlx.Tx, params *mapCourseEntity.MappingCourseReq) (err error)
	CreateTx(ctx context.Context) (tx *sqlx.Tx, err error)
}

type mpCourseService struct {
	mpCourseRepo MpCourseRepo
	lessonServices lessonServices.LessonService
	classServices classServices.ClassService
	teacherServices teacherServices.TeacherService
}

func NewMappingCourseService(repo MpCourseRepo, lessonService lessonServices.LessonService,
	classService classServices.ClassService,
	teacherService teacherServices.TeacherService) *mpCourseService {
	return &mpCourseService{
		mpCourseRepo: repo,
		lessonServices: lessonService,
		classServices: classService,
		teacherServices: teacherService,
	}
}

func (service mpCourseService) FindAll(ctx context.Context, params *mapCourseController.MappingCourseParams) (data []mapCourseEntity.MappingCourse, err error) {
	//logic here
	return
}

func (service mpCourseService) SelectAll(ctx context.Context, parameter *mapCourseController.MappingCourseParams) (data []mapCourseEntity.MappingCourse, err error) {
	data , err = service.mpCourseRepo.SelectAll(ctx, parameter)
	if err != nil {
		return nil,err
	}
	return
}

func (service mpCourseService) FindOne(ctx context.Context, params *mapCourseController.MappingCourseParams) (data mapCourseEntity.MappingCourse, err error) {
	//logic here
	data , err = service.mpCourseRepo.FindOne(ctx, params)
	if err != nil {
		return data,err
	}
	return
}

func (service mpCourseService) Create(ctx context.Context, params *mapCourseEntity.MappingCourseReq) (err error) {
	//check lesson
	check, err := service.validateData(ctx, params)
	if check {
		return err
	}

	// check duplicate
	duplicate, err := service.mpCourseRepo.FindOne(ctx, &mapCourseController.MappingCourseParams{
		ClassID:	 params.ClassID,
		LessonID:	 params.LessonID,
		TeacherID:	 params.TeacherID,
	})
	if err != nil {
		return err
	}
	if duplicate.ID != ""{
		return errors.New("duplicate mapping course")
	}

	//create
	tx, err := service.mpCourseRepo.CreateTx(ctx)
	if err != nil {
		return err
	}

	err = service.mpCourseRepo.Create(ctx, tx, params)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	


	return
}

func (service mpCourseService) validateData(ctx context.Context, params *mapCourseEntity.MappingCourseReq) (bool, error) {
	lesson, err := service.lessonServices.FindOne(ctx, &lessonServices.LessonParams{
		DefaultParameter: controller.DefaultParameter{
			ID: params.LessonID,
		},
	})
	if err != nil || lesson.ID == 0 {
		return true, err
	}

	class, err := service.classServices.FindOne(ctx, &classServices.ClassParams{
		DefaultParameter: controller.DefaultParameter{
			ID: params.ClassID,
		},
	})
	if err != nil || class.ID == 0 {
		return true, err
	}

	teacher, err := service.teacherServices.FindOne(ctx, &teacherServices.TeacherParams{
		DefaultParameter: controller.DefaultParameter{
			ID: params.TeacherID,
		},
	})
	if err != nil || teacher.ID == "" {
		return true, err
	}
	return false, nil
}
