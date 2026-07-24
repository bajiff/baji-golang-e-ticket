package domain

import (
	"errors"
	"time"
)

type ConcertEvent struct {
	ID string
	EventName string
	TimeStart time.Time
	MaxCapacity int
}

func NewConcertEvent(concertId string, eventName string, eventTime time.Time, capacity int) (ConcertEvent, error) {

	if capacity < 0 {
		return ConcertEvent{}, errors.New("kapasitas tidak boleh negatif")
	}
	
	if eventTime.Before(time.Now()) {
		return ConcertEvent{}, errors.New("jadwal konser tidak boleh di masa lalu")
	}

	return ConcertEvent{ID: concertId,EventName: eventName,TimeStart: eventTime,MaxCapacity: capacity}, nil
}
