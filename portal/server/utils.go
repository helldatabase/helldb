package server

import "encoding/json"

type errors struct {
	Errors []string `json:"errors"`
}

func newError(err string) errors {
	return errors{
		[]string{err},
	}
}

func toJson(resp interface{}) string {
	b, err := json.Marshal(resp)
	if err != nil {
		return ""
	}
	return string(b)
}
