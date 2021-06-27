package app

import (
	"github.com/osdifera/report_research/controllers/ethplorer"
	"github.com/osdifera/report_research/controllers/ping"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/holders/:address", ethplorer.GetTopHolders)
}
