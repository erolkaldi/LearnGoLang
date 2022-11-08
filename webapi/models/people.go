package models

import "time"

type person struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Birthdate time.Time `json:"birthdate"`
	Weight    int       `json:"weight"`
	Height    int       `json:"height"`
}
