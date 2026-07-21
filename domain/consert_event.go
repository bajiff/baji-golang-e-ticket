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

/* 
Buat Fungsi Factory NewEventKonser:

Fungsi ini harus menerima 4 parameter input: id, nama, waktu, dan kapasitas.

Validasi 1 (Kapasitas): Cek apakah kapasitas < 0. Jika ya, kembalikan EventKonser{} kosong dan error "kapasitas tidak boleh negatif".

Validasi 2 (Waktu): Cek apakah input waktu terjadi sebelum saat ini menggunakan waktu.Before(time.Now()). Jika ya, kembalikan error "jadwal konser tidak boleh di masa lalu".

Jika semua validasi berhasil dilewati, kembalikan struct EventKonser yang sudah diisi penuh beserta nil untuk error.

*/

func NewConcertEvent(concertId string, eventName string, eventTime time.Time, capacity int) (ConcertEvent, error) {

	if capacity < 0 {
		return ConcertEvent{}, errors.New("kapasitas tidak boleh negatif")
	}
	
	if eventTime.Before(time.Now()) {
		return ConcertEvent{}, errors.New("jadwal konser tidak boleh di masa lalu")
	}

	return ConcertEvent{ID: concertId,EventName: eventName,TimeStart: eventTime,MaxCapacity: capacity}, nil
}
