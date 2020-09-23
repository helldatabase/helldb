package evaluator

import "helldb/engine/types"

type Response struct {
	Errors  []string         `json:"errors"`
	Results []types.BaseType `json:"results"`
}
