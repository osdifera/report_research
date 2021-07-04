package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/osdifera/report_research/models"
)

func GetTopHolders(address string) []models.Holder {

	var(
		listOfholders models.ListOfHolders
	)
	
	endpointURL := "https://api.ethplorer.io/getTopTokenHolders/"+address+"?apiKey=EK-xy1hm-wXaD35q-oNbjf&limit=15"
	resp, err := http.Get(endpointURL)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	
	if err := json.Unmarshal([]byte(bodyBytes), &listOfholders); err != nil {
  		fmt.Println("failed to unmarshal:", err)
	}

	return listOfholders.Holders
}