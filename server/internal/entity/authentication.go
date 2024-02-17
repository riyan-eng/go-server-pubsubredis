package entity

import (
	"server/internal/datastruct"

	"github.com/golang-jwt/jwt/v5"
)

type AuthenticationLoginReq struct {
	Email    string
	Password string
	Issuer   string
}

type AuthenticationLoginRes struct {
	AccessToken  string
	RefreshToken string
	ExpiredAt    *jwt.NumericDate
	Match        bool
}

type AuthenticationRegisterReq struct {
	UUIDUser     string
	UUIDUserData string
	Email        string
	Password     string
	Nama         string
	NIK          string
	KodeRole     string
	NomorTelepon string
}

type AuthenticationRefreshTokenReq struct {
	RefreshToken string
	Issuer       string
}

type AuthenticationRefreshTokenRes struct {
	AccessToken  string
	RefreshToken string
	ExpiredAt    *jwt.NumericDate
}

type AuthenticationValidateResetTokenReq struct {
	ResetToken string
}

type AuthenticationRequestResetToken struct {
	Email  string
	Issuer string
}

type AuthenticationResetPasswordReq struct {
	ResetToken string
	Password   string
}

type AuthenticationLogoutReq struct {
	UserUUID string
}

type AuthenticationMeReq struct {
	UserUUID string
}

type AuthenticationMeRes struct {
	Data datastruct.AuthenticationMe
}
