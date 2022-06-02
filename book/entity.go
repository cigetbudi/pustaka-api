package book

import "time"

type Book struct {
	ID          int
	Title       string
	Decsription string
	Price       int
	Rating      int
	Discount    int
	CreatedAt   time.Time
	Updateat    time.Time
}
