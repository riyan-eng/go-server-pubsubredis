package util

import (
	"errors"
	"fmt"
	"server/env"
	"server/infrastructure"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	json "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
	"golang.org/x/sync/errgroup"
)

var mySigningKey = []byte("AllYourBase")

type AccessTokenClaims struct {
	UserUUID string `json:"user_id"`
	RoleCode string `json:"role_code"`
	UUID     string `json:"id"`
	jwt.RegisteredClaims
}

type AccessTokenCached struct {
	AccessUID string `json:"access"`
}

type RefreshTokenClaims struct {
	UserUUID string `json:"user_id"`
	RoleCode string `json:"role_code"`
	UUID     string `json:"id"`
	jwt.RegisteredClaims
}

type RefreshTokenCached struct {
	RefreshUID string `json:"refresh"`
}

type tokenStruct struct{}

func NewToken() *tokenStruct {
	return &tokenStruct{}
}

func (t *tokenStruct) CreateAccess(ctx *fasthttp.RequestCtx, userUUID string) (string, *jwt.NumericDate, error) {
	expiredTime := time.Minute * env.NewEnvironment().JWT_EXPIRED_ACCESS
	tokenUUID := uuid.NewString()
	claims := AccessTokenClaims{
		userUUID,
		"ADMIN",
		tokenUUID,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiredTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", nil, fmt.Errorf("can't signed the token")
	}

	cachedJson, err := json.Marshal(AccessTokenCached{
		AccessUID: claims.UUID,
	})
	if err != nil {
		return "", nil, fmt.Errorf("can't marshal access token")
	}

	if err := infrastructure.Redis.Set(ctx, fmt.Sprintf("access-token-%s", userUUID), string(cachedJson), expiredTime).Err(); err != nil {
		return "", nil, fmt.Errorf("can't cached access token")
	}

	return ss, claims.ExpiresAt, nil
}

func (t *tokenStruct) ParseAccess(tokenString string) (*AccessTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*AccessTokenClaims); ok {
		return claims, nil
	} else {
		return nil, fmt.Errorf("unknown claims type, cannot proceed")
	}
}

func (t *tokenStruct) ValidateAccess(ctx *fasthttp.RequestCtx, claims *AccessTokenClaims) error {
	g := new(errgroup.Group)
	g.Go(func() error {
		cacheJSON, err := infrastructure.Redis.Get(ctx, fmt.Sprintf("access-token-%s", claims.UserUUID)).Result()
		if err != nil {
			return fmt.Errorf("token not found")
		}
		cachedTokens := new(AccessTokenCached)
		err = json.Unmarshal([]byte(cacheJSON), cachedTokens)
		var tokenUID string = cachedTokens.AccessUID
		if err != nil || tokenUID != claims.UUID {
			return errors.New("token not found")
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}

func (t *tokenStruct) CreateRefresh(ctx *fasthttp.RequestCtx, userUUID string) (string, *jwt.NumericDate, error) {
	expiredTime := time.Minute * env.NewEnvironment().JWT_EXPIRED_REFRESH
	tokenUUID := uuid.NewString()
	claims := RefreshTokenClaims{
		userUUID,
		"ADMIN",
		tokenUUID,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiredTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", nil, fmt.Errorf("can't signed the token")
	}

	cachedJson, err := json.Marshal(RefreshTokenCached{
		RefreshUID: claims.UUID,
	})
	if err != nil {
		return "", nil, fmt.Errorf("can't marshal refresh token")
	}

	if err := infrastructure.Redis.Set(ctx, fmt.Sprintf("refresh-token-%s", userUUID), string(cachedJson), expiredTime).Err(); err != nil {
		return "", nil, fmt.Errorf("can't cached refresh token")
	}

	return ss, claims.ExpiresAt, nil
}

func (t *tokenStruct) ParseRefresh(tokenString string) (*RefreshTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &RefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*RefreshTokenClaims); ok {
		return claims, nil
	} else {
		return nil, fmt.Errorf("unknown claims type, cannot proceed")
	}
}

func (t *tokenStruct) ValidateRefresh(ctx *fasthttp.RequestCtx, claims *RefreshTokenClaims) error {
	g := new(errgroup.Group)
	g.Go(func() error {
		cacheJSON, err := infrastructure.Redis.Get(ctx, fmt.Sprintf("refresh-token-%s", claims.UserUUID)).Result()
		if err != nil {
			return fmt.Errorf("token not found")
		}
		cachedTokens := new(RefreshTokenCached)
		err = json.Unmarshal([]byte(cacheJSON), cachedTokens)
		var tokenUID string = cachedTokens.RefreshUID
		if err != nil || tokenUID != claims.UUID {
			return errors.New("token not found")
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}
