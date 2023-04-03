package http

import (
	"encoding/json"
	"net/http"
	"timestamps-assignment/api/app"
)

type TimestampDataHandler struct {
	timestampDataService app.TimestampDataService
}

func NewTimestampDataHandler(t app.TimestampDataService) *TimestampDataHandler {
	return &TimestampDataHandler{
		timestampDataService: t,
	}
}

func (h *TimestampDataHandler) TimestampsMatching(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	period := r.URL.Query().Get("period")
	timezone := r.URL.Query().Get("tz")
	start := r.URL.Query().Get("t1")
	end := r.URL.Query().Get("t2")

	if period == "" || timezone == "" || start == "" || end == "" {
		http.Error(w, "Invalid/Missing parameters", http.StatusBadRequest)
	}

	result, err := h.timestampDataService.TimestampsCalculation(period, timezone, start, end)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
