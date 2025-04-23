package service

import "backend/internal/client"

type PrayerService struct {
	Client *client.Muftyat
}

func (s *PrayerService) GetTodayPrayerTimes() ([]client.PrayerEntry, error) {
	return s.Client.GetPrayerTimes()
}
