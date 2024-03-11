package infrastructure

import (
	"context"
	"github.com/Waldemarsch/medods_test/internal/infrastructure/repository"
	"github.com/Waldemarsch/medods_test/internal/models"
)

type Repository interface {
	CreateToken(context.Context, *models.Token)
	GetToken(context.Context, *models.Token) *models.Token
	RefreshToken(context.Context, *models.Token) *models.Token
	CloseDB(ctx context.Context) error
}

type Infrastructure struct {
	Repository
}

func NewInfrastructure(uri string) *Infrastructure {
	return &Infrastructure{
		Repository: repository.NewRepoMongo(uri),
	}
}
