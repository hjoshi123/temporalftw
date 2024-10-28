package server

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/hjoshi123/temporal-loan-app/internal/logging"
	"net/http"
)

type Input struct {
	ID        string
	W         http.ResponseWriter
	R         *http.Request
	GetParams map[string][]string
}

type Output struct {
	Output interface{}
}

type CustomHandler func(ctx context.Context, input Input) (Output, error)

func (c CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	input := Input{}
	input.W = w
	input.R = r
	input.ID = mux.Vars(r)["id"]
	input.GetParams = make(map[string][]string)

	for k, v := range r.URL.Query() {
		input.GetParams[k] = v
	}

	output, err := c(ctx, input)
	if err != nil {
		logging.FromContext(ctx).Errorw("failed to execute handler", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if output.Output != nil {
		b, err := json.Marshal(output.Output)
		if err != nil {
			logging.FromContext(ctx).Errorw("failed to marshal response", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}

	w.WriteHeader(http.StatusOK)
}
