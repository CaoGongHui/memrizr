package service

import (
	"context"
	"crypto/rsa"
	"log"

	"github.com/caogonghui/memrizr/account/model"
	"github.com/caogonghui/memrizr/account/model/apperrors"
)

//TokenService 用来注入一个TokenRepository的实现
//TokenRepository是为了用key和密钥来登录JWTS在服务层方法里面
type tokenService struct {
	PrivKey               *rsa.PrivateKey
	PubKey                *rsa.PublicKey
	RefreshSecret         string
	IDExpirationSecs      int64
	RefreshExpirationSecs int64
}

//用来持有repository 最终注入到当前服务层
type TSConfig struct {
	PrivKey               *rsa.PrivateKey
	PubKey                *rsa.PublicKey
	RefreshSecret         string
	IDExpirationSecs      int64
	RefreshExpirationSecs int64
}

//一个工厂方法初始化一个UserService使用它的响应层依赖
func NewTokenService(ts *TSConfig) model.TokenService {
	return &tokenService{
		PrivKey:               ts.PrivKey,
		PubKey:                ts.PubKey,
		RefreshSecret:         ts.RefreshSecret,
		IDExpirationSecs:      ts.IDExpirationSecs,
		RefreshExpirationSecs: ts.RefreshExpirationSecs,
	}
}
func (s *tokenService) NewPairFromUser(ctx context.Context, u *model.User, prevTokenID string) (*model.TokenPair, error) {
	idToken, err := generateIDToken(u, s.PrivKey, s.IDExpirationSecs)

	if err != nil {
		log.Printf("Error generating idToken for uid: %v. Error: %v\n", u.UID, err.Error())
		return nil, apperrors.NewInternal()
	}

	refreshToken, err := generateRefreshToken(u.UID, s.RefreshSecret, s.RefreshExpirationSecs)

	return &model.TokenPair{
		IDToken:      idToken,
		RefreshToken: refreshToken.SS,
	}, nil
}
