package thegraph

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/osdifera/report_research/services"
)

func GetLiquidity(c *gin.Context) {
	address := c.Param("address")
	liquidity := services.GetLiquidity(address)
	c.JSON(http.StatusOK, liquidity)
}