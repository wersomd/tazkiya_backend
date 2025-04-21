package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Muftyat struct {
	BaseURL string
	City    string
	Country string
}

type PrayerResponse struct {
	Data struct {
		Timings map[string]string `json:"timings"`
	} `json:"data"`
}

func (c *Muftyat) GetPrayerTimes() (map[string]string, error) {
	url := fmt.Sprintf("%s?city=%s&country=%s&method=2", c.BaseURL, c.City, c.Country)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res PrayerResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	return res.Data.Timings, nil
}
