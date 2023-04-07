package teacher


type TeacherRepo interface {
}


type teacherService struct {
	teacherRepo TeacherRepo
}

func NewTeacherService(repo TeacherRepo) *teacherService {
	return &teacherService{
		teacherRepo: repo,
	}
}