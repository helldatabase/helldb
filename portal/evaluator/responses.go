package evaluator

import (
	"encoding/json"

	"helldb/engine/types"
)

func toJson(data interface{}) string {
	b, _ := json.Marshal(data)
	return string(b)
}

type Response struct {
	Errors  []string           `json:"errors"`
	Results [][]types.BaseType `json:"results"`
}
