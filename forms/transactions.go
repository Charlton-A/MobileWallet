package forms

import (
	"time"

	"github.com/charlton/practs/mwallet/models"
)



type UserTranasctionResp struct{
	UserID int  `json:"user_id,omitempty"`
	FirstName string `json:"first_name"`
	LastName string  `json:"last_name"`
	Email string     `json:"email,omitempty"`
	Phone string      `json:"phone,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time   `json:"updated_at,omitempty"`
	DeletedAt *time.Time   `json:"deleted_at,omitempty"`
	Transactions  []  models.Transaction     `json:"transactions,omitempty"`

}