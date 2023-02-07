package app

import (
	"go-gin-ddd/src/database"
	agendaService "go-gin-ddd/src/domain/agendas/services"

	"github.com/gin-gonic/gin"

	"go-gin-ddd/src/api/http"
)

var (
	router = gin.Default()
)

func StartApp() {
	gin.SetMode("debug")

	database.Init()
	agendaHandler := http.NewHandler(agendaService.NewService(database.NewAgendaDB()))

	router.POST("/agendas", agendaHandler.Create)

	_ = router.Run(":3000")
}
