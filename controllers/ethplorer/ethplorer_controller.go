package ethplorer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/osdifera/report_research/services"
)

func GetTopHolders(c *gin.Context){
	address := c.Param("address")
	topHolders := services.GetTopHolders(address)
	c.JSON(http.StatusOK, topHolders)
}