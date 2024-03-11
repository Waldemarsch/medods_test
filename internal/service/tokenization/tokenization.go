package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"github.com/Waldemarsch/medods_test/internal/infrastructure"
	"github.com/Waldemarsch/medods_test/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

type Tokenization struct {
	inf          *infrastructure.Infrastructure
	secretString string
}

func NewTokenizationService(inf *infrastructure.Infrastructure) *Tokenization {
	return &Tokenization{
		inf:          inf,
		secretString: "test_task",
	}
}

func (s *Tokenization) CreateToken(ctx context.Context, tokens *models.Token) *models.Token {

	tokens.GUID = tokens.CreateToken.Request.GUID

	if s.inf.Repository.GetToken(ctx, tokens) != nil {
		logrus.Errorf("Tokens for user %s already exist!", tokens.GUID)
		return nil
	}

	tokenRefresh := make([]byte, 16)
	rand.Read(tokenRefresh)
	tokens.SetRefresh(tokenRefresh)
	tokens.CreateToken.Response.Refresh = base64.URLEncoding.EncodeToString([]byte(tokens.Refresh))

	payload := jwt.MapClaims{
		"guid": tokens.CreateToken.Request.GUID,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	}
	tokenAccess := jwt.NewWithClaims(jwt.SigningMethodHS512, payload)
	tokenAccessString, err := tokenAccess.SignedString([]byte(s.secretString))
	if err != nil {
		log.Fatal("Service: ", err)
	}
	tokens.SetAccess([]byte(tokenAccessString))

	s.inf.Repository.CreateToken(ctx, tokens)

	tokens.CreateToken.Response.Access = tokens.Access

	return tokens
}

func (s *Tokenization) RefreshToken(ctx context.Context, tokens *models.Token) *models.Token {

	tokens.SetAccess(tokens.RefreshToken.Request.Access)
	acToken, err := base64.URLEncoding.DecodeString(tokens.RefreshToken.Request.Refresh)
	if err != nil {
		log.Fatal(err)
	}
	tokens.SetRefresh(acToken)

	_, err = jwt.ParseWithClaims(tokens.Access, &tokens.UserClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretString), nil
	})

	tokens.GUID = tokens.UserClaims.GUID

	if s.inf.Repository.GetToken(ctx, tokens) == nil {
		return nil
	}

	if err := tokens.CompareHash(); err != nil {
		logrus.Errorln(err)
		return nil
	}

	tokenRefresh := make([]byte, 16)
	rand.Read(tokenRefresh)
	tokens.SetRefresh(tokenRefresh)
	tokens.RefreshToken.Response.Refresh = base64.URLEncoding.EncodeToString([]byte(tokens.Refresh))

	payload := jwt.MapClaims{
		"guid": tokens.GUID,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	}
	tokenAccess := jwt.NewWithClaims(jwt.SigningMethodHS512, payload)
	tokenAccessString, err := tokenAccess.SignedString([]byte(s.secretString))
	if err != nil {
		log.Fatal("Service: ", err)
	}
	tokens.SetAccess([]byte(tokenAccessString))

	tokens = s.inf.Repository.RefreshToken(ctx, tokens)

	tokens.RefreshToken.Response.Access = tokens.Access

	return tokens
}
