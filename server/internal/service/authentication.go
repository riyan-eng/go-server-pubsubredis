package service

import (
	"database/sql"
	"strconv"

	"server/config"
	"server/env"
	"server/internal/datastruct"
	dtorepository "server/internal/dto_repository"
	"server/internal/entity"
	"server/internal/model"
	"server/internal/repository"
	"server/pkg/util"

	"github.com/blockloop/scan/v2"
)

type AuthenticationService interface {
	RefreshToken(req entity.AuthenticationRefreshTokenReq) (res entity.AuthenticationRefreshTokenRes)
	ResetPassword(req entity.AuthenticationResetPasswordReq)
	Login(req entity.AuthenticationLoginReq) (res entity.AuthenticationLoginRes)
	Register(req entity.AuthenticationRegisterReq)
	Logout(req entity.AuthenticationLogoutReq)
	Me(req entity.AuthenticationMeReq) (res entity.AuthenticationMeRes)
	RequestResetToken(req entity.AuthenticationRequestResetToken)
	ValidateResetToken(req entity.AuthenticationValidateResetTokenReq)
}

type authenticationService struct {
	dao repository.DAO
}

func NewAuthenticationService(dao repository.DAO) AuthenticationService {
	return &authenticationService{
		dao: dao,
	}
}

func (a *authenticationService) RefreshToken(req entity.AuthenticationRefreshTokenReq) (res entity.AuthenticationRefreshTokenRes) {
	claim, err := util.ParseToken(req.RefreshToken, env.NewEnvironment().JWT_SECRET_REFRESH)
	if err != nil {
		util.PanicIfNeeded(util.CustomBadRequest{
			Messages:    "Invalid refresh token.",
			StatusCodes: 401,
		})

	}
	if err := util.ValidateToken(claim, "refresh"); err != nil {
		util.PanicIfNeeded(util.CustomBadRequest{
			Messages:    "Invalid refresh token.",
			StatusCodes: 401,
		})
	}
	genJwt := util.GenerateJwt(claim.UserUUID, claim.RoleCode, req.Issuer)
	res.AccessToken = genJwt.AccessToken
	res.RefreshToken = genJwt.RefreshToken
	res.ExpiredAt = genJwt.ExpiredAt
	return
}

func (a *authenticationService) ResetPassword(req entity.AuthenticationResetPasswordReq) {
	claim, err := util.ParseToken(req.ResetToken, env.NewEnvironment().JWT_SECRET_RESET)
	if err != nil {
		util.PanicIfNeeded(util.CustomBadRequest{
			Messages:    "Invalid token.",
			StatusCodes: 401,
		})

	}
	if err := util.ValidateToken(claim, "resetPwd"); err != nil {
		util.PanicIfNeeded(util.CustomBadRequest{
			Messages:    "Invalid token.",
			StatusCodes: 401,
		})
	}
	// change password
	password := util.GenerateHash(req.Password)
	tabelUser := model.User{
		UUID:     claim.UserUUID,
		Password: sql.NullString{String: password, Valid: util.IsValid(password)},
	}
	a.dao.NewAuthenticationQuery().ResetPassword(dtorepository.AuthenticationResetPasswordReq{
		TabelUser: tabelUser,
	})

	// delete token
	a.dao.NewAuthenticationQuery().Logout(dtorepository.AuthenticationLogoutReq{
		UserUUID: claim.UserUUID,
	})
}

func (a *authenticationService) Login(req entity.AuthenticationLoginReq) (res entity.AuthenticationLoginRes) {
	sqlRows := a.dao.NewAuthenticationQuery().Login(dtorepository.AuthenticationLoginReq{
		Email: req.Email,
	})
	var user datastruct.AuthenticationLogin
	err := scan.Row(&user, sqlRows)
	if err == sql.ErrNoRows {
		return
	} else {
		util.PanicIfNeeded(err)
	}

	enforce := config.NewEnforcer()
	enforce.AddRoleForUser(strconv.Itoa(user.ID), user.KodeRole)

	ok := util.VerifyHash(user.Password, req.Password)
	if ok {
		if !user.IsAktif {
			util.PanicIfNeeded(util.CustomBadRequest{
				Messages: "User is not active.",
			})
		}
		res.Match = true
		genJwt := util.GenerateJwt(user.UUID, user.KodeRole, req.Issuer)
		res.AccessToken = genJwt.AccessToken
		res.RefreshToken = genJwt.RefreshToken
		res.ExpiredAt = genJwt.ExpiredAt
		return
	}
	return
}

func (a *authenticationService) Register(req entity.AuthenticationRegisterReq) {
	password := util.GenerateHash(req.Password)
	tabelUser := model.User{
		UUID:     req.UUIDUser,
		Email:    sql.NullString{String: req.Email, Valid: util.IsValid(req.Email)},
		Password: sql.NullString{String: password, Valid: util.IsValid(password)},
		Role:     sql.NullString{String: req.KodeRole, Valid: util.IsValid(req.KodeRole)},
		UserData: sql.NullString{String: req.UUIDUserData, Valid: util.IsValid(req.UUIDUserData)},
		IsAktif:  sql.NullBool{Bool: true, Valid: true},
	}
	tableUserData := model.UserData{
		UUID:         req.UUIDUserData,
		Nama:         sql.NullString{String: req.Nama, Valid: util.IsValid(req.Nama)},
		NIK:          sql.NullString{String: req.NIK, Valid: util.IsValid(req.NIK)},
		NomorTelepon: sql.NullString{String: req.NomorTelepon, Valid: util.IsValid(req.NomorTelepon)},
	}
	a.dao.NewAuthenticationQuery().Register(dtorepository.AuthenticationRegisterReq{
		TabelUser:     tabelUser,
		TabelUserData: tableUserData,
	})
}

func (a *authenticationService) Logout(req entity.AuthenticationLogoutReq) {
	a.dao.NewAuthenticationQuery().Logout(dtorepository.AuthenticationLogoutReq{
		UserUUID: req.UserUUID,
	})
}

func (a *authenticationService) Me(req entity.AuthenticationMeReq) (res entity.AuthenticationMeRes) {
	sqlRows := a.dao.NewAuthenticationQuery().Me(dtorepository.AuthenticationMeReq{
		UserUUID: req.UserUUID,
	})

	err := scan.Row(&res.Data, sqlRows)
	util.PanicIfNeeded(err)
	return
}

func (a *authenticationService) RequestResetToken(req entity.AuthenticationRequestResetToken) {
	var user datastruct.AuthenticationRequestResetToken
	sqlrows := a.dao.NewAuthenticationQuery().RequestResetToken(dtorepository.AuthenticationRequestResetTokenReq{
		Email: req.Email,
	})
	err := scan.Row(&user, sqlrows)
	if err == sql.ErrNoRows {
		util.PanicIfNeeded(util.CustomBadRequest{
			Messages: "Email tidak terdaftar.",
		})
	} else {
		util.PanicIfNeeded(err)
	}

	genToken := util.GenerateTokenResetPassword(strconv.Itoa(user.ID), user.KodeRole, req.Issuer)
	go func() {
		sender := util.NewGmailSender("SIPENTA", env.NewEnvironment().SMTP_EMAIL, env.NewEnvironment().SMTP_PASSWORD)
		subject := "Reset Password Verification"
		content := util.NewTemplate().EmailResetPassword(genToken.ResetPwdToken, genToken.ExpiredAt)
		to := []string{user.Email}
		err = sender.SendEmail(subject, content, to, nil, nil, nil)

	}()

}

func (a *authenticationService) ValidateResetToken(req entity.AuthenticationValidateResetTokenReq) {
	claim, err := util.ParseToken(req.ResetToken, env.NewEnvironment().JWT_SECRET_RESET)
	if err != nil {
		util.PanicIfNeeded(util.CustomBadRequest{
			Messages:    "Invalid token.",
			StatusCodes: 401,
		})

	}
	if err := util.ValidateToken(claim, "resetPwd"); err != nil {
		util.PanicIfNeeded(util.CustomBadRequest{
			Messages:    "Invalid token.",
			StatusCodes: 401,
		})
	}
}
