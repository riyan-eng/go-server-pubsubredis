package util

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"server/env"
	"server/infrastructure"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

type CachedToken struct {
	AccessUID   string `json:"access"`
	RefreshUID  string `json:"refresh"`
	ResetPwdUID string `json:"reset_pwd"`
}

type CustomClaim struct {
	UserUUID string `json:"user_id"`
	RoleCode string `json:"role_code"`
	UID      string `json:"uid"`
	jwt.RegisteredClaims
}

type tokenResult struct {
	token string
	expat *jwt.NumericDate
	uid   string
}

type JwtResult struct {
	AccessToken   string
	RefreshToken  string
	ResetPwdToken string
	ExpiredAt     *jwt.NumericDate
}

func GenerateJwt(userUUID, roleCode, issuer string) JwtResult {
	access := createToken(userUUID, roleCode, issuer, env.NewEnvironment().JWT_SECRET_ACCESS, env.NewEnvironment().JWT_EXPIRED_ACCESS)
	refresh := createToken(userUUID, roleCode, issuer, env.NewEnvironment().JWT_SECRET_REFRESH, env.NewEnvironment().JWT_EXPIRED_REFRESH)
	cachedJson, err := json.Marshal(CachedToken{
		AccessUID:  access.uid,
		RefreshUID: refresh.uid,
	})
	PanicIfNeeded(err)
	ctx := context.Background()
	infrastructure.Redis.Set(ctx, fmt.Sprintf("token-%s", userUUID), string(cachedJson), time.Minute*env.NewEnvironment().JWT_EXPIRED_LOGOFF)
	return JwtResult{
		AccessToken:  access.token,
		RefreshToken: refresh.token,
		ExpiredAt:    access.expat,
	}
}

func GenerateTokenResetPassword(userUUID, roleCode, issuer string) JwtResult {
	tokenResetPwd := createToken(userUUID, roleCode, issuer, env.NewEnvironment().JWT_SECRET_RESET, env.NewEnvironment().JWT_EXPIRED_RESET)
	cachedJson, err := json.Marshal(CachedToken{
		ResetPwdUID: tokenResetPwd.uid,
	})
	PanicIfNeeded(err)
	ctx := context.Background()
	infrastructure.Redis.Set(ctx, fmt.Sprintf("token-%s", userUUID), string(cachedJson), time.Minute*env.NewEnvironment().JWT_EXPIRED_LOGOFF)
	return JwtResult{
		ResetPwdToken: tokenResetPwd.token,
		ExpiredAt:     tokenResetPwd.expat,
	}
}

func createToken(userUUID, roleCode, issuer, secret string, expMinute time.Duration) tokenResult {
	uid := uuid.NewString()
	expat := jwt.NewNumericDate(time.Now().Add(expMinute * time.Minute))
	claims := CustomClaim{
		userUUID,
		roleCode,
		uid,
		jwt.RegisteredClaims{
			ExpiresAt: expat,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    issuer,
		},
	}
	mySigningKey := []byte(secret)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	PanicIfNeeded(err)
	return tokenResult{
		token: ss,
		expat: expat,
		uid:   uid,
	}
}

func ParseToken(tokenString string, secret string) (claims *CustomClaim, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return
	}
	claims = token.Claims.(*CustomClaim)
	return
}

func ValidateToken(claims *CustomClaim, tokenType string) (err error) {
	ctx := context.Background()
	g := new(errgroup.Group)
	g.Go(func() error {
		cacheJSON, _ := infrastructure.Redis.Get(ctx, fmt.Sprintf("token-%s", claims.UserUUID)).Result()
		cachedTokens := new(CachedToken)
		err := json.Unmarshal([]byte(cacheJSON), cachedTokens)
		var tokenUID string
		if tokenType == "access" {
			tokenUID = cachedTokens.AccessUID
		} else if tokenType == "refresh" {
			tokenUID = cachedTokens.RefreshUID
		} else {
			tokenUID = cachedTokens.ResetPwdUID
		}
		if err != nil || tokenUID != claims.UID {
			return errors.New("token not found")
		}
		return nil
	})

	err = g.Wait()
	return
}
