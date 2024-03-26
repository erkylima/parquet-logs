package api

import (
	"fmt"

	"github.com/erkylima/parquet-logs/internal/logs/services"
	"github.com/erkylima/parquet-logs/pkg/database"
	"github.com/erkylima/parquet-logs/pkg/handlers"
)

func (a *Api) InitLogsRouter() {

	repo, err := database.NewParquetLogConnection("duckdb", "data/log.parquet", "data/audit.parquet")
	if err != nil {
		fmt.Println(err)
	}
	logService := services.NewLogService(repo)
	handler := handlers.NewLogsHandler(logService)
	a.Group.POST("/logs", handler.CreateHandler)
	a.Group.GET("/logs", handler.ListHandler)
}
