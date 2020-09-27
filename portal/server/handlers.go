package server

import (
	"fmt"
	"net/http"

	"helldb/portal/evaluator"
)

func query(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = fmt.Fprint(w, toJson(newError("method not allowed")))
	}
	if queryBody, ok := r.Form["query"]; ok {
		_, _ = fmt.Fprint(w, toJson(evaluator.Eval(queryBody[0])))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, toJson(newError("query not provided")))
	}
}

func status(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprint(w, "ok")
}

func length(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "%d", evaluator.Store.Len())
}
