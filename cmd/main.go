package main

import (
	"backend/internal/client"
	"backend/internal/handler"
	"backend/internal/service"
	"fmt"
	"log"
	"net/http"
)

func main() {
	prayerTimes := &client.Muftyat{
		BaseURL: "https://api.muftyat.kz/prayer-times/2025/42.368009/69.612769",
		City:    "Shymkent",
	}

	prayers, err := prayerTimes.GetPrayerTimes()
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range prayers {
		fmt.Printf("Дата: %s | Фаджр: %s | Зухр: %s | Аср: %s | Магриб: %s | Иша: %s\n",
			p.Date, p.Fajr, p.Dhuhr, p.Asr, p.Maghrib, p.Isha)
	}
	svc := &service.PrayerService{
		Client: prayerTimes,
	}

	h := &handler.PrayerHandler{Service: svc}

	http.HandleFunc("/backend", h.GetPrayerTimes)

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
