package entities

import (
	"github.com/Waldemarsch/medods_test/internal/models"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type Token struct {
	GUID    string `bson:"GUID"`
	Refresh string `bson:"refresh"`
}

func (t *Token) SetToken(token *models.Token) {
	bcryptToken, err := bcrypt.GenerateFromPassword([]byte(token.Refresh), 15)
	if err != nil {
		log.Fatal(err)
	}
	t.Refresh = string(bcryptToken)
}
