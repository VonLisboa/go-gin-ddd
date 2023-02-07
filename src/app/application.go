package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	gin.SetMode("DEBUG")
	router := gin.Default()

	_ = router.Run(":3000")
}
