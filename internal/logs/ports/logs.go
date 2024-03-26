package ports

import (
	"github.com/erkylima/parquet-logs/internal/logs/domains"
	"github.com/erkylima/parquet-logs/internal/shared/models"
)

type LogService interface {
	CreateLog(Log *domains.Log) (string, error)
	ListLogs(filter models.Filter) ([]domains.Log, error)
}
