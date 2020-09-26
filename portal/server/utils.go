package server

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

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

func hash(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	res := fmt.Sprint(h.Sum(nil))
	return res
}
