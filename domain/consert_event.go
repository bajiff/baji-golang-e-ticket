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

type TicketPrice struct {
	value int
}

type TicketType struct {
	ID string
	EventID string
	Name string
	Price TicketPrice
	Quota int
}

// Value Object
type OrderItem struct {
	id string
	ticketTypeID string
	quantity int
	subtotal int
}

// Aggregrat Root
type Order struct {
	ID string
	UserID string
	Items []OrderItem
	TotalAmount int
}

func (o *Order) AddItem(ticketTypeID string, quantity int, pricePerItem int) {
	subTotal := quantity * pricePerItem

	newItem := OrderItem{
		id: "",ticketTypeID: ticketTypeID,quantity: quantity,subtotal: subTotal,
	}
	
	o.Items = append(o.Items, newItem)
	o.TotalAmount += subTotal
}



func NewTicketPrice(amount int) (TicketPrice, error) {

	if amount < 0 {
		return TicketPrice{}, errors.New("harga tiket tidak boleh negatif")
	}

	return TicketPrice{value: amount}, nil
}

func NewTypeTicket(id string, eventId string, name string, price TicketPrice, quota int) (TicketType, error) {

	if quota < 0 {
		return TicketType{}, errors.New("kuota tiket tidak boleh negatif")
	}

	return TicketType{ID: id, EventID: eventId, Name: name, Price: price, Quota: quota}, nil
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
