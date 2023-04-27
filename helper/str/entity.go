package str

import (
	teacherEntity "go-management-auth-school/entity/teacher"
	studentEntity "go-management-auth-school/entity/student"
	parentEntity "go-management-auth-school/entity/parent"
)

func GetFullNameTeacher(teacher teacherEntity.Teacher) string {
	if teacher.LastName == "" {
		return teacher.FirstName
	}
	return teacher.FirstName + " " + teacher.LastName
}

func GetFullNameStudent(student studentEntity.Student) string {
	if student.LastName == "" {
		return student.FirstName
	}
	return student.FirstName + " " + student.LastName
}

func GetFullNameParent(parent parentEntity.Parent) string {
	if parent.LastName == "" {
		return parent.FirstName
	}
	return parent.FirstName + " " + parent.LastName
}
