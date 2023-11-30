package dtorepository

import "server/internal/model"

type AuthenticationRegisterReq struct {
	TabelUser     model.User
	TabelUserData model.UserData
}

type AuthenticationLoginReq struct {
	Email string
}

type AuthenticationLogoutReq struct {
	UserUUID string
}

type AuthenticationRequestResetTokenReq struct {
	Email string
}

type AuthenticationResetPasswordReq struct {
	TabelUser model.User
}

type AuthenticationMeReq struct {
	UserUUID string
}
