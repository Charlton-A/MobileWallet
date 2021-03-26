package models

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

//User Details
type User struct{
	UserID int  `json:"user_id,omitempty"`
	FirstName string `json:"first_name" binding:"required"`
	LastName string  `json:"last_name" binding:"required"`
	Email string     `json:"email,omitempty" binding:"required"`
	Phone string      `json:"phone,omitempty" binding:"required"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time   `json:"updated_at,omitempty"`
	DeletedAt *time.Time   `json:"deleted_at,omitempty"`
	Wallets []Wallet     `json:"wallets,omitempty"`

}

func(u *User) Create(ctx context.Context, db *sql.DB) (int,error) {
	var id int
	stmt:=`INSERT INTO mobile.users (first_name,last_name,email,phone) VALUES ($1, $2, $3, $4) RETURNING id`
	err := db.QueryRow(stmt,u.FirstName,u.LastName,u.Email,u.Phone).Scan(&id)
	if err != nil {
		return 0,err
	}
	return id, nil
}


func(u *User) Fetch (ctx context.Context, db *sql.DB , userID int) (error) {
	stmt:=`SELECT id, first_name,last_name FROM  mobile.users WHERE id=$1`
	row := db.QueryRowContext(ctx,stmt,userID)
	if row.Err() != nil {
		return row.Err()
	}
	err := row.Scan(&u.UserID, &u.FirstName, &u.LastName)
	if err != nil {
		return  err
	}

	return err
}




