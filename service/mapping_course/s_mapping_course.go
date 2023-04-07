package mapping_course


type MpCourseRepo interface {
}


type mpCourseService struct {
	mpCourseRepo MpCourseRepo
}

func NewMappingCourseService(repo MpCourseRepo) *mpCourseService {
	return &mpCourseService{
		mpCourseRepo: repo,
	}
}
