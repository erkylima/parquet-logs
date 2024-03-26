package ports

import (
	"github.com/erkylima/parquet-logs/internal/logs/domains"
	"github.com/erkylima/parquet-logs/internal/shared/models"
)

type Repository interface {
	Create(entity *domains.Log) (string, error)
	List(filter models.Filter) ([]domains.Log, error)
}
