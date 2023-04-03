package auth

import (
	classEntity "go-management-auth-school/entity/class"
	mappingCourse "go-management-auth-school/entity/mapping_course"

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
		MappingCourse []mappingCourse.MappingCourse `json:"mapping_course"`
		jwt.StandardClaims
}

type Auth struct {
	Indentity string `json:"indentity"`
	IsActive	bool   `json:"isActive"`
	ExpiredAt string `json:"expiredAt"`
	RefreshExpiredAt string `json:"refreshExpiredAt"`
	Token string `json:"token"`
	RefreshToken string `json:"refreshToken"`
	Type string `json:"type"`
	SessionToken string `json:"sessionToken"`
	Status string `json:"status"`
}
