package service

import (
	"context"
	"github.com/Waldemarsch/medods_test/internal/infrastructure"
	"github.com/Waldemarsch/medods_test/internal/models"
	service "github.com/Waldemarsch/medods_test/internal/service/tokenization"
)

type Tokenization interface {
	CreateToken(ctx context.Context, tokens *models.Token) *models.Token
	RefreshToken(ctx context.Context, tokens *models.Token) *models.Token
}

type Service struct {
	Tokenization
}

func NewService(inf *infrastructure.Infrastructure) *Service {
	return &Service{
		Tokenization: service.NewTokenizationService(inf),
	}
}
