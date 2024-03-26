package services

import (
	"github.com/erkylima/parquet-logs/internal/logs/domains"
	"github.com/erkylima/parquet-logs/internal/shared/models"
	"github.com/erkylima/parquet-logs/internal/shared/ports"
)

type LogService struct {
	repo ports.Repository
}

func NewLogService(repo ports.Repository) *LogService {
	return &LogService{repo: repo}
}

func (ps *LogService) CreateLog(Log *domains.Log) (string, error) {
	result, err := ps.repo.Create(Log)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (ps *LogService) ListLogs(filter models.Filter) ([]domains.Log, error) {

	Logs, err := ps.repo.List(filter)
	if err != nil {
		return nil, err
	}

	return Logs, nil
}
