package mapping_course

import (
	mapCourseEntity "go-management-auth-school/entity/mapping_course"
	helperStr "go-management-auth-school/helper/str"
)

type MappingCourseResponse struct {
	ID        string `json:"id"`
	ClassID   string `json:"class_id"`
	Class     string `json:"class"`
	LessonID  string `json:"lesson_id"`
	Lesson    string `json:"lesson"`
	TeacherID string `json:"teacher_id"`
	Teacher   string `json:"teacher"`
}

func FromServices(res []mapCourseEntity.MappingCourse) (data []MappingCourseResponse) {
	for _, v := range res {
		data = append(data, FromService(v))
	}
	return
}

func FromService(res mapCourseEntity.MappingCourse) (data MappingCourseResponse) {
	data = MappingCourseResponse{
		ID:        res.ID,
		ClassID:   res.ClassID,
		Class:     res.Class.Name,
		LessonID:  res.LessonID,
		Lesson:    res.Lesson.Name,
		TeacherID: res.TeacherID,
		Teacher:   helperStr.GetFullNameTeacher(res.Teacher),
	}
	return
}
