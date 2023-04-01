package lesson

type LessonService interface {
}

type lessonController struct {
	lessonService LessonService
}

func NewLessonController(lessonService LessonService) lessonController {
	return lessonController{
		lessonService: lessonService,
	}
}