package domain

import (
	"errors"
	"time"
)

type TicketPrice struct {
	value int
}

type TicketType struct {
	ID string
	EventID string
	Name string
	Price TicketPrice
	Quota int
	FlashSaleStart time.Time
	FlashSaleEnd time.Time
}

// Value Object
type OrderItem struct {
	ID string
	TicketTypeID string
	Quantity int
	Subtotal int
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
		ID: "",TicketTypeID: ticketTypeID, Quantity: quantity, Subtotal: subTotal,
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

func NewTypeTicket(id string, eventId string, name string, price TicketPrice, quota int, flashSaleStart time.Time, flashSaleEnd time.Time) (TicketType, error) {

	if quota < 0 {
		return TicketType{}, errors.New("kuota tiket tidak boleh negatif")
	}

	return TicketType{ID: id, EventID: eventId, Name: name, Price: price, Quota: quota, FlashSaleStart: flashSaleStart, FlashSaleEnd: flashSaleEnd}, nil
}

func (t *TicketType) IsFlashSaleActive() error {
	rightNow := time.Now()

	if rightNow.Before(t.FlashSaleStart) {
		return errors.New("flash sale belum dimulai")
	}
	
	if rightNow.After(t.FlashSaleEnd) {
		return errors.New("flash sale sudah berakhir")
	}
	
	return  nil
}
