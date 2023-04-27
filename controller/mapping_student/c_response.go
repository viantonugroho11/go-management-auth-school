package mapping_student

import (
	mapStudentEntity "go-management-auth-school/entity/mapping_student"
	mapCourseEntity "go-management-auth-school/entity/mapping_course"
	helperStr "go-management-auth-school/helper/str"
)

type MappingStudentResponse struct {
	ID        int    `json:"id"`
	ClassID   string `json:"class_id"`
	Class     string `json:"class"`
	StudentID string `json:"student_id"`
	Student   string `json:"student"`
	TeacherID string `json:"teacher_id"`
	Teacher   string `json:"teacher"`
	Type 		  string `json:"type"`
	Course 		[]CourseResponse `json:"course"`
}

type CourseResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
}

func FromServices(res []mapStudentEntity.MappingStudent) (data []MappingStudentResponse) {
	for _, v := range res {
		data = append(data, FromService(v))
	}
	return
}

func FromService(res mapStudentEntity.MappingStudent) (data MappingStudentResponse) {
	data = MappingStudentResponse{
		ID:        res.ID,
		ClassID:   res.ClassID,
		Class:     res.Class.Name,
		StudentID: res.Student.ID,
		Student:   helperStr.GetFullNameStudent(res.Student),
		TeacherID: res.Teacher.ID,
		Teacher:   helperStr.GetFullNameTeacher(res.Teacher),
		Type: 		 res.Type,
		Course: 	 FromServicesCourse(res.MpCourse),
	}
	return
}

func FromServicesCourse(res []mapCourseEntity.MappingCourse) (data []CourseResponse) {
	for _, v := range res {
		data = append(data, FromServiceCourse(v))
	}
	return
}

func FromServiceCourse(res mapCourseEntity.MappingCourse) (data CourseResponse) {
	data = CourseResponse{
		ID:        res.ID,
		Name:      res.Lesson.Name,
	}
	return
}
