package usecases

import (
	"context"
	"fmt"
	"mongoidx/domains/models"
	"mongoidx/domains/repository"
	"time"
)

type Usecase struct {
	connectionRepository     repository.ConnectionRepository
	connectHistoryRepository repository.ConnectHistoryRepository
}

func NewUsecase(connectionRepository repository.ConnectionRepository, connectHistoryRepository repository.ConnectHistoryRepository) Usecase {
	return Usecase{
		connectionRepository:     connectionRepository,
		connectHistoryRepository: connectHistoryRepository,
	}
}

func (u *Usecase) CreateConnection(uri string) (connection models.Connection, err error) {

	ctx := context.TODO()

	connection = models.Connection{
		URI: uri,
	}

	conCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = connection.Connect(conCtx)
	if err != nil {
		return connection, fmt.Errorf("failed to connect to mongodb: %w", err)
	}

	defer connection.Disconnect(ctx)

	connectionID, err := u.connectionRepository.AddConnection(ctx, connection)
	if err != nil {
		return connection, err
	}

	connection.ID = connectionID

	return connection, nil
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

func (u *Usecase) ListConnectHistory() ([]models.ConnectHistory, error) {
	return u.connectHistoryRepository.ListHistory(context.TODO())
}
