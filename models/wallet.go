package models

import (
	"context"
	"database/sql"
	"time"
)


type Wallet struct{
	WalletID int `json:"wallet_id,omitempty"`
	UserID int  `json:"user_id,omitempty" binding:"required"`
	Balance float64  `json:"balance"`
	Amount float64   `json:"amount,omitempty"`
	Action int `json:"action,omitempty"`
	Currency string    `json:"currency"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time   `json:"updated_at,omitempty"`
	DeletedAt *time.Time   `json:"deleted_at,omitempty"`
}



func(w *Wallet) Create (ctx context.Context, db *sql.DB) error {
	stmt:=`INSERT INTO mobile.wallets (user_id ,balance) VALUES ($1, $2)`
	_, err := db.Exec(stmt,w.UserID ,w.Balance)
	if err != nil {
		return err
	}

	return nil
}

func FetchBalances  (ctx context.Context, db *sql.DB , userID int) ([]Wallet,error) {
	var Balances []Wallet
	stmt:=`SELECT id,balance,currency,created_at FROM  mobile.wallets WHERE user_id=$1`
	rows, err := db.QueryContext(ctx,  stmt,userID)
	if err != nil {
		return Balances ,err
	}
	defer rows.Close()
	for rows.Next() {
		var w Wallet

		err = rows.Scan(&w.WalletID, &w.Balance, &w.Currency ,&w.CreatedAt)
		if err != nil {
			return Balances , err
	    }
		Balances =  append(Balances, w)
	}
	return Balances, err
}

func FetchBalance  (ctx context.Context, db *sql.DB ,walletID int) (float32 ,error) {
	var balance float32
	stmt:=`SELECT balance  FROM  mobile.wallets WHERE id= $1`
	err := db.QueryRowContext(ctx, stmt, walletID).Scan(&balance)
	if err != nil {
		return 0 ,err
	}
	return 0, err

}

func (w *Wallet)UpdateBalance (ctx context.Context, db *sql.DB) error {

	stmt:=`UPDATE mobile.wallets SET balance=balance+$1 ,updated_at=now() WHERE id= $2 AND user_id=$3`

	_, err := db.ExecContext(ctx, stmt, w.Amount, w.WalletID ,w.UserID)
	if err != nil {
		return err
	}
	return  err

}
