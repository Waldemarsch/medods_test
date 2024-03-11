package models

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Token struct {
	CreateToken  CreateToken
	RefreshToken RefreshToken
	Access       string
	Refresh      string
	HashRefresh  string
	GUID         string
	UserClaims   UserClaims
}

type CreateToken struct {
	Request struct {
		GUID string `json:"GUID"`
	}
	Response struct {
		Access  string `json:"access"`
		Refresh string `json:"refresh"`
	}
}

type RefreshToken struct {
	Request struct {
		Access  string `json:"access"`
		Refresh string `json:"refresh"`
	}
	Response struct {
		Access  string `json:"access"`
		Refresh string `json:"refresh"`
	}
}

func (t *Token) SetAccess(token interface{}) {
	switch token.(type) {
	case string:
		t.Access = token.(string)
	case []byte:
		t.Access = string(token.([]byte))
	}
	return
}

func (t *Token) SetRefresh(token interface{}) {
	switch token.(type) {
	case string:
		t.Refresh = token.(string)
	case []byte:
		t.Refresh = string(token.([]byte))
	}
	return
}

func (t *Token) CompareHash() error {
	err := bcrypt.CompareHashAndPassword([]byte(t.HashRefresh), []byte(t.Refresh))
	if err != nil {
		return err
	}
	return nil
}

type UserClaims struct {
	jwt.RegisteredClaims
	GUID    string
	Refresh string
}
