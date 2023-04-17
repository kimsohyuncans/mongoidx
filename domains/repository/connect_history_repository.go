package repository

import (
	"context"
	"mongoidx/domains/models"
)

type ConnectHistoryRepository interface {
	AddHistory(ctx context.Context, connectionHistory models.ConnectHistory) error
	GetHistoryByID(ctx context.Context, ID string) (*models.ConnectHistory, error)
}
