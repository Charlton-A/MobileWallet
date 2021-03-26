package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/charlton/practs/mwallet/utils"
)


type Transaction struct{
	ID  int    `json:"id,omitempty"`
    UserID int  `json:"-"`
    WalletID int `json:"wallet_id,omitempty"`
	Amount float64 `json:"amount,omitempty"`
	TransKey string `json:"trans_key,omitempty"`
    Status int  `json:"-"`
	Type int  `json:"-"`
	StatusDesc string `json:"status_desc,omitempty"`
	TypeDesc string  `json:"type_desc,omitempty"`
    Currency string  `json:"currency,omitempty"`
    CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}


func(t *Transaction) Store (ctx context.Context, db *sql.DB , shareKey bool) error {
	var transKey string
	var err error
	if !shareKey{
		transKey, err= utils.KeyGen(int(t.WalletID))
		if err != nil {
		return err
	}
	}else{
		transKey=t.TransKey
	}

	stmt:=`INSERT INTO mobile.transactions (trans_key, user_id, wallet_id, t_status, t_type, amount) VALUES ($1,$2,$3,$4,$5,$6)`
	_, err = db.Exec(stmt,transKey, t.UserID ,t.WalletID ,t.Status ,t.Type, t.Amount)
	if err != nil {
		return err
	}

	return nil
}

func FetchTransactions(ctx context.Context, db *sql.DB , userID int) ([]Transaction, error) {
	var list []Transaction

	stmt:=`SELECT w.id, w.trans_key, t.name, s.name, w.user_id ,w.wallet_id, w.amount ,w.created_at FROM  mobile.transactions AS w LEFT JOIN mobile.transaction_type AS t ON w.t_type=t.id  LEFT JOIN mobile.transaction_status AS s ON w.t_status=s.id WHERE  w.user_id=$1  ORDER BY w.id DESC`
	rows, err := db.QueryContext(ctx,  stmt, userID)
	if err != nil {
		return list ,err
	}
	defer rows.Close()
	for rows.Next() {
		var w Transaction
		err = rows.Scan(&w.ID, &w.TransKey, &w.TypeDesc ,&w.StatusDesc ,&w.UserID ,&w.WalletID ,&w.Amount ,&w.CreatedAt)
		if err != nil {
			return list ,err
	    }
		list =  append(list,w)
	}
	return list, nil


}






