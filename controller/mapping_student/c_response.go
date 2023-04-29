package mapping_student

import (
	mapCourseEntity "go-management-auth-school/entity/mapping_course"
	mapStudentEntity "go-management-auth-school/entity/mapping_student"
	helperStr "go-management-auth-school/helper/str"
)

type MappingStudentResponse struct {
	ID        string              `json:"id"`
	ClassID   string           `json:"class_id"`
	Class     string           `json:"class"`
	StudentID string           `json:"student_id"`
	Student   string           `json:"student"`
	TeacherID string           `json:"teacher_id"`
	Teacher   string           `json:"teacher"`
	Type      string           `json:"type"`
	Course    []CourseResponse `json:"course"`
}

type CourseResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func FromServices(res []mapStudentEntity.MappingStudent) (data []MappingStudentResponse) {
	for _, v := range res {
		data = append(data, FromService(v))
	}
	return
}

func FromService(res mapStudentEntity.MappingStudent) (data MappingStudentResponse) {
	switch res.Type {
	case "0":
		data = MappingStudentResponse{
		ID:        res.ID,
		ClassID:   res.ClassID,
		Class:     res.Class.Name,
		StudentID: res.Identity,
		Student:   helperStr.GetFullNameStudent(res.Student),
		Type:      res.Type,
		Course:    FromServicesCourse(res.MpCourse),
	}
	case "1":
		data = MappingStudentResponse{
		ID:        res.ID,
		ClassID:   res.ClassID,
		Class:     res.Class.Name,
		TeacherID: res.Identity,
		Teacher:   helperStr.GetFullNameTeacher(res.Teacher),
		Type:      res.Type,
		Course:    FromServicesCourse(res.MpCourse),
		}
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
		ID:   res.ID,
		Name: res.Lesson.Name,
	}
	return
}
