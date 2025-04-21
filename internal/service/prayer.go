package service

import "backend/internal/client"

type PrayerService struct {
	Client *client.Muftyat
}

func (s *PrayerService) GetTodayPrayerTimes() (map[string]string, error) {
	return s.Client.GetPrayerTimes()
}
