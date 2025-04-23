package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Muftyat struct {
	BaseURL string
	City    string
}

type PrayerEntry struct {
	Fajr    string `json:"fajr"`
	Sunrise string `json:"sunrise"`
	Dhuhr   string `json:"dhuhr"`
	Asr     string `json:"asr"`
	Maghrib string `json:"maghrib"`
	Isha    string `json:"isha"`
	Date    string `json:"Date"`
}

type PrayerResponse struct {
	Result []PrayerEntry `json:"result"`
}

func (c *Muftyat) GetPrayerTimes() ([]PrayerEntry, error) {
	resp, err := http.Get(c.BaseURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res PrayerResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	if len(res.Result) == 0 {
		return nil, fmt.Errorf("no prayer times found")
	}

	return res.Result, nil
}
