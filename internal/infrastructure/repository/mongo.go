package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type RepoMongo struct {
	client *mongo.Client
}

func NewRepoMongo(ctx context.Context, clientOptions *options.ClientOptions) *RepoMongo {
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Panicln(err)
	}
	return &RepoMongo{
		client: client,
	}
}
