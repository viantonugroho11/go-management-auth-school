package auth

import (
	authEntity "go-management-auth-school/entity/auth"
	helperStr "go-management-auth-school/helper/str"
	"time"
)

type LoginResponse struct {
	ExpiredDate string `json:"expired_date"`
	Token       string `json:"token"`
	Status 		bool `json:"status"`
	ExpiredRefreshToken string `json:"expired_refresh_token"`
	ResfreshToken string `json:"resfresh_token"`
	SessionToken string `json:"session_token"`
}

type ValidateResponse struct {
	IdentityID string `json:"identityId"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Permission string `json:"permission"`
	IsActive bool `json:"is_active"`
	ExpiredAt time.Time `json:"expired_at"`
	RefreshExpiredAt string `json:"refresh_expired_at"`
	Class string `json:"class"`
	StudentID string `json:"student_id"`
	TeacherID string `json:"teacher_id"`
}

func FromServiceLogin(res authEntity.Auth) *LoginResponse {
	return &LoginResponse{
		ExpiredDate: res.ExpiredAt,
		Token:       res.Token,
		Status:      res.IsActive,
		ExpiredRefreshToken: res.RefreshExpiredAt,
		ResfreshToken: res.RefreshToken,
		SessionToken: res.SessionToken,
	}
}

func FromServiceValidate(res authEntity.AuthValidate) *ValidateResponse {
	return &ValidateResponse{
		IdentityID: res.Identity,
		Username: res.User.Username,
		Fullname: helperStr.GetFullNameStudent(res.Student),
		Email: res.Student.Email,
		Phone: res.Student.Phone,
		Permission: helperStr.IsPermission(res.User.Permission),
		IsActive: res.User.Status,
		ExpiredAt: helperStr.ConvertUnixToTime(res.ExpiredAt),
	}
}
