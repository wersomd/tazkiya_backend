package handler

import (
	"backend/internal/service"
	"encoding/json"
	"net/http"
)

type PrayerHandler struct {
	Service *service.PrayerService
}

func (h *PrayerHandler) GetPrayerTimes(w http.ResponseWriter, r *http.Request) {
	timings, err := h.Service.GetTodayPrayerTimes()
	if err != nil {
		http.Error(w, "Failed to get prayer times", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(timings)
}
