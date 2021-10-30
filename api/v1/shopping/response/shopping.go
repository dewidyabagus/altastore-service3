package response

import "time"

type ShoppData struct {
	ID         string    `json:"id"`
	IsCheckOut bool      `json:"ischeckout"`
	UpdatedAt  time.Time `json:"updated_at"`
}
