package mapping_student


type MpStudentRepo interface {
}


type mpStudentService struct {
	mpStudentRepo MpStudentRepo
}

func NewMappingStudentService(repo MpStudentRepo) *mpStudentService {
	return &mpStudentService{
		mpStudentRepo: repo,
	}
}