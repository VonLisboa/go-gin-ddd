package http

import (
	"go-gin-ddd/src/domain/agendas/model"
	service "go-gin-ddd/src/domain/agendas/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

type AgendaServer interface {
	Get(ctx *gin.Context)
	GetBy(ctx *gin.Context)
	Create(ctx *gin.Context)
}

type AgendaHandler struct {
	service service.AgendaService
}

func NewHandler(agenda service.AgendaService) AgendaServer {
	return &AgendaHandler{
		service: agenda,
	}
}

// Create implements AgendaServer
func (handler *AgendaHandler) Create(ctx *gin.Context) {
	var data model.Agenda
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	if _, err := handler.service.Create(data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, data)
}

// Get implements AgendaServer
func (handler *AgendaHandler) Get(ctx *gin.Context) {
	data, err := handler.service.Get()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// GetBy implements AgendaServer
func (handler *AgendaHandler) GetBy(ctx *gin.Context) {
	data, err := handler.service.GetBy(ctx.Param("disponibilidade"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}
