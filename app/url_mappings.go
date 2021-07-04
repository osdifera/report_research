package app

import (
	"github.com/osdifera/report_research/controllers/ethplorer"
	"github.com/osdifera/report_research/controllers/ping"
	"github.com/osdifera/report_research/controllers/thegraph"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/holders/:address", ethplorer.GetTopHolders)
	router.POST("/liquidity/:address", thegraph.GetLiquidity)
}
