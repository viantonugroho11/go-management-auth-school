package lesson

type LessonRepo interface {
}


type lessonService struct {
	lessonRepo LessonRepo
}

func NewLessonService(repo LessonRepo) *lessonService {
	return &lessonService{
		lessonRepo: repo,
	}
}
