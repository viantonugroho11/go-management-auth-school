package auth

import (
	classEntity "go-management-auth-school/entity/class"
	mappingCourse "go-management-auth-school/entity/mapping_course_teacher"

	"github.com/dgrijalva/jwt-go"
)


type JwtCustomClaimsStudent struct {
    Username  string `json:"username"`
    Firstname string `json:"firstname"`
    Lastname  string `json:"lastname"`
    Email    string `json:"email"`
    Indentity string `json:"indentity"`
    Type     string `json:"type"`
    Phone    string `json:"phone"`
    Class    classEntity.Class `json:"class"`
		MappingCourse []mappingCourse.MappingCourseTeacher `json:"mapping_course"`
		jwt.StandardClaims
}