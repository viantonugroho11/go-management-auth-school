package mapping_course


type MappingCourse struct {
	ID int `json:"id"`
	CourseID int `json:"course_id"`
	TeacherID int `json:"teacher_id"`
	LessonID int `json:"lesson_id"`
}