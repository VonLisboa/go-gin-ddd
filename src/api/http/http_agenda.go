package http

import (
	service "go-gin-ddd/src/domain/agendas/services"

	"github.com/gin-gonic/gin"
)

type AgendaInterface interface {
	Get(ctx *gin.Context)
	GetBy(ctx *gin.Context)
	Create(ctx *gin.Context)
}

type AgendaHandler struct {
	service service.AgendaService
}

func NewHandler(agenda service.AgendaService) AgendaInterface {
	return &AgendaHandler{
		service: agenda,
	}
}

// Create implements AgendaInterface
func (handler *AgendaHandler) Create(ctx *gin.Context) {
	panic("unimplemented")
}

// Get implements AgendaInterface
func (handler *AgendaHandler) Get(ctx *gin.Context) {
	panic("unimplemented")
}

// GetBy implements AgendaInterface
func (handler *AgendaHandler) GetBy(ctx *gin.Context) {
	panic("unimplemented")
}
