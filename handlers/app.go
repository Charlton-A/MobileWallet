package handlers

import (
	"context"
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)


type Application struct{
	AppCtx context.Context
	DB  *sql.DB
}

func NewDB(ctx context.Context) (error, *sql.DB ){
	db, err := sql.Open("postgres", os.Getenv("DB_DSN"))
	if err != nil {
		return err,nil
	}

	err = db.PingContext(ctx)
	if err != nil {
		return err,nil
	}
	db.SetMaxOpenConns(5)
	return nil, db


}