package models

import (
	"encoding/json"
)

type ListOfHolders struct {
	Holders []Holder `json:"holders"`
}

type Holder struct {
	Address string `json:"address"`
	Balance json.Number `json:"balance"`
	Share float32 `json:"share"`
}