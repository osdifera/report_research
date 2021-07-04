package services

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/osdifera/report_research/models"
)

func GetLiquidity(address string) models.LiquidityData{

	//TODO: write function to replace address with url param
	values := map[string]string{"query": "{pair(id: \"0xa478c2975ab1ea89e8196811f51a7b7ade33eb11\"){token0{id symbol} token1 {id symbol} reserve0 reserve1 reserveUSD token0Price token1Price }}"}
	json_data, err := json.Marshal(values)
	
	 if err != nil {
        log.Fatal(err)
    }

	endpoint := "https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2"
	resp, err := http.Post(endpoint,"application/json", bytes.NewBuffer(json_data))
	
	if err != nil {
		log.Fatalln(err)
	}

	var response models.LiquidityWrapper
    json.NewDecoder(resp.Body).Decode(&response)

	return response.LiquidityData
}