package main

import (
	"backend/internal/client"
	"backend/internal/handler"
	"backend/internal/service"
	"log"
	"net/http"
)

func main() {
	prayerTimes := &client.Muftyat{
		BaseURL: "http://api.aladhan.com/v1/timingsByCity",
		City:    "Almaty",
		Country: "Kazakhstan",
	}

	svc := &service.PrayerService{Client: prayerTimes}
	h := &handler.PrayerHandler{Service: svc}

	http.HandleFunc("/backend", h.GetPrayerTimes)

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
