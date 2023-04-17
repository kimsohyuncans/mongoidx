package repository

import (
	"context"
	"mongoidx/domains/models"
)

type ConnectHistoryRepository interface {
	AddHistory(ctx context.Context, connectionHistory models.ConnectHistory) error
	ListHistory(ctx context.Context) ([]models.ConnectHistory, error)
}
