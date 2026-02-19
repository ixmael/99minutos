package domain

import "time"

type EventRequest struct {
	TrackingNumber string    `json:"tracking_number"`
	Status         string    `json:"status"`
	Timestamp      time.Time `json:"timestamp"`
	Source         string    `json:"source"`
	Location       struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"location"`
}
