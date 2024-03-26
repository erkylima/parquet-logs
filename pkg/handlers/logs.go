package handlers

import (
	"net/http"

	"github.com/erkylima/parquet-logs/internal/logs/domains"
	"github.com/erkylima/parquet-logs/internal/logs/ports"
	"github.com/erkylima/parquet-logs/internal/shared/models"
	"github.com/gin-gonic/gin"
)

type logsHandler struct {
	svc ports.LogService
}

func NewLogsHandler(svc ports.LogService) *logsHandler {
	return &logsHandler{svc: svc}
}

func (h *logsHandler) CreateHandler(ctx *gin.Context) {

	var log domains.Log
	if err := ctx.ShouldBindJSON(&log); err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}
	_, err := h.svc.CreateLog(&log)
	if err != nil {
		HandleError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *logsHandler) ListHandler(ctx *gin.Context) {
	filters := models.Filter{}
	if err := ctx.ShouldBindJSON(&filters); err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}
	logs, err := h.svc.ListLogs(filters)
	if err != nil {
		HandleError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, logs)
}
