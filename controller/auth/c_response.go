package auth

import (
	authEntity "go-management-auth-school/entity/auth"
)

type LoginResponse struct {
	ExpiredDate string `json:"expiredDate"`
	Token       string `json:"token"`
	Status 		bool `json:"status"`
	ExpiredRefreshToken string `json:"expiredRefreshToken"`
	ResfreshToken string `json:"resfreshToken"`
	SessionToken string `json:"sessionToken"`
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
