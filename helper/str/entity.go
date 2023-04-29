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

func IsGender(value int) string {
	return map[int]string{
		1: "Male",
		2: "Female",
	}[value]
}

func IsReligion(value int) string {
	return map[int]string{
		1: "Islam",
		2: "Protestan",
		3: "Katolik",
		4: "Hindu",
		5: "Budha",
		6: "Konghucu",
	}[value]
}

func IsStatus(value int) string {
	return map[int]string{
		0: "Single",
		1: "Married",
	}[value]
}

func IsActiveConvert(value int) string {
	return map[int]string{
		0: "Active",
		1: "Inactive",
	}[value]
}

func IsStatusTeacher(value int) string {
	return map[int]string{
		0: "PNS",
		1: "Non PNS",
	}[value]
}

func IsTypeMapping(value int) string {
	return map[int]string{
		0: "Student",
		1: "Homeroom Teacher",
	}[value]
}

func IsTypeParent(value int) string {
	return map[int]string{
		0: "Parent",
		1: "Guardian",
		2: "Relatives",
		3: "Others",
	}[value]
}

func IsTypeStudent(value int) string {
	return map[int]string{
		0: "Student",
		1: "Alumni",
	}[value]
}

func IsDisability(value int) string {
	return map[int]string{
		0: "No",
		1: "Yes",
	}[value]
}

func IsPermission(value int) string {
	return map[int]string{
		0:"student",
		1:"teacher",
		2:"admin",
		3:"superadmin",
		4:"accounting",
		5:"operator",
	}[value]
}
