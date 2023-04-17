package usecases

import (
	"context"
	"fmt"
	"mongoidx/domains/models"
	"mongoidx/domains/repository"
)

type Usecase struct {
	connectionRepository     repository.ConnectionRepository
	connectHistoryRepository repository.ConnectHistoryRepository
}

func NewUsecase(connectionRepository repository.ConnectionRepository, connectHistoryRepository repository.ConnectHistoryRepository) *Usecase {
	return &Usecase{
		connectionRepository:     connectionRepository,
		connectHistoryRepository: connectHistoryRepository,
	}
}

func (u *Usecase) CreateConnection(uri string) (connectionID string, err error) {
	ctx := context.TODO()

	connection := models.Connection{
		URI: uri,
	}

	err = connection.Connect(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to connect to mongodb: %w", err)
	}

	defer connection.Disconnect(ctx)

	connectionID, err = u.connectionRepository.AddConnection(ctx, connection)
	if err != nil {
		return "", err
	}

	return connectionID, nil
}

func (u *Usecase) ListConnection() ([]models.Connection, error) {
	return u.connectionRepository.ListConnection(context.TODO())
}

func (u *Usecase) GetConnection(connectionID string) (*models.Connection, error) {
	return u.connectionRepository.GetConnection(context.TODO(), connectionID)
}

func (u *Usecase) DeleteConnection(connectionID string) error {
	return u.connectionRepository.DeleteConnection(context.TODO(), connectionID)
}
