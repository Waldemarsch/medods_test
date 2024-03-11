package repository

import (
	"context"
	"fmt"
	"github.com/Waldemarsch/medods_test/internal/entities"
	"github.com/Waldemarsch/medods_test/internal/models"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

type RepoMongo struct {
	client *mongo.Client
}

func NewRepoMongo(uri string) *RepoMongo {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Panicln("Repository: ", err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Panicln("Repository: ", err)
	}
	log.Println("Connected to the Mongo DB!")
	if _, err := client.Database("test").Collection("tokens").DeleteMany(context.Background(), bson.M{}); err != nil {
		panic(err)
	}
	return &RepoMongo{
		client: client,
	}
}

func (r *RepoMongo) CloseDB(ctx context.Context) error {
	return r.client.Disconnect(ctx)
}

func (r *RepoMongo) CreateToken(ctx context.Context, tokens *models.Token) {
	collection := r.client.Database("test").Collection("tokens")
	tokenEntity := new(entities.Token)
	tokenEntity.GUID = tokens.GUID
	tokenEntity.SetToken(tokens)
	insResult, err := collection.InsertOne(ctx, tokenEntity)
	if err != nil {
		log.Fatal("Repository: ", err)
	}

	fmt.Println("Inserted a document: ", insResult.InsertedID)

	return
}

func (r *RepoMongo) GetToken(ctx context.Context, tokens *models.Token) *models.Token {
	collection := r.client.Database("test").Collection("tokens")
	tokenEntity := new(entities.Token)
	tokenEntity.GUID = tokens.GUID
	filter := bson.D{{"GUID", tokenEntity.GUID}}
	err := collection.FindOne(ctx, filter).Decode(&tokenEntity)
	if err != nil {
		logrus.Infoln("Repository: for GUID:", tokenEntity.GUID, err)
		return nil
	}
	tokens.HashRefresh = tokenEntity.Refresh
	return tokens
}

func (r *RepoMongo) RefreshToken(ctx context.Context, tokens *models.Token) *models.Token {
	collection := r.client.Database("test").Collection("tokens")
	tokenEntity := new(entities.Token)
	tokenEntity.GUID = tokens.GUID
	tokenEntity.SetToken(tokens)
	filter := bson.D{{"GUID", tokenEntity.GUID}}
	update := bson.D{{"$set", bson.D{{"refresh", tokenEntity.Refresh}}}}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal("Repository: ", err)
	}
	fmt.Printf("Successfully updated refresh token!\n")
	return tokens
}
