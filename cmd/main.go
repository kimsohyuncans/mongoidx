package main

import (
	"context"
	"mongoidx/controllers/http"
	"mongoidx/infrastructure"
	"mongoidx/usecases"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:example@mongo:27017"))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic("connect to mongodb timeout after 5 seconds")
	}

	db := client.Database("mongoidx")

	connectionRepo := infrastructure.NewConnectionRepo(db)
	connectHistoryRepo := infrastructure.NewConnectHistory(db)

	usecase := usecases.NewUsecase(connectionRepo, connectHistoryRepo)

	http.NewHttpServer(3000, &usecase).Start()
}
