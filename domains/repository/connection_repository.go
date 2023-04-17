package repository

import (
	"context"
	"mongoidx/domains/models"
)

type ConnectionRepository interface {
	AddConnection(ctx context.Context, conn models.Connection) (connectionID string, err error)
	ListConnection(ctx context.Context) ([]models.Connection, error)
	GetConnection(ctx context.Context, ID string) (*models.Connection, error)
	DeleteConnection(ctx context.Context, ID string) error
}
