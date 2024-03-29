package auth

import (
	classEntity "go-management-auth-school/entity/class"
	mappingCourse "go-management-auth-school/entity/mapping_course"
	mapStudent "go-management-auth-school/entity/mapping_student"
	userEntity "go-management-auth-school/entity/user"
	studentEntity "go-management-auth-school/entity/student"
	teacherEntity "go-management-auth-school/entity/teacher"

	"github.com/dgrijalva/jwt-go"
)

type JwtCustomClaimsStudent struct {
	Username      string                        `json:"username"`
	Firstname     string                        `json:"firstname"`
	Lastname      string                        `json:"lastname"`
	Email         string                        `json:"email"`
	Identity      string                        `json:"identity"`
	Type          string                        `json:"type"`
	Phone         string                        `json:"phone"`
	Class         classEntity.Class             `json:"class"`
	MappingCourse []mappingCourse.MappingCourse `json:"mapping_course"`
	jwt.StandardClaims
}

type Auth struct {
	Identity         string `json:"identity"`
	IsActive         bool   `json:"isActive"`
	ExpiredAt        string `json:"expiredAt"`
	RefreshExpiredAt string `json:"refreshExpiredAt"`
	Token            string `json:"token"`
	RefreshToken     string `json:"refreshToken"`
	Type             string `json:"type"`
	SessionToken     string `json:"sessionToken"`
	Status           string `json:"status"`
}

type AuthValidate struct {
	Identity string                    `json:"identity"`
	ExpiredAt int64                   `json:"expiredAt"`
	User     userEntity.User           `json:"user"`
	MpClass  mapStudent.MappingStudent `json:"mp_class"`
	Student studentEntity.Student			`json:"student"`
	Teacher teacherEntity.Teacher			`json:"teacher"`
}
