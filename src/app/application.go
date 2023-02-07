package app

import (
	"go-gin-ddd/src/database"
	service "go-gin-ddd/src/domain/agendas/services"

	"github.com/gin-gonic/gin"

	"go-gin-ddd/src/api/http"
)

var (
	router = gin.Default()
)

func StartApp() {
	gin.SetMode("DEBUG")

	database.Init()
	agendaHandler := http.NewHandler(service.NewService(database.NewAgendaDB()))

	router.POST("/agendas", agendaHandler.Create)

	_ = router.Run(":3000")
}
