package handlers

import (
	"context"
	"database/sql"
	"os"
	"time"

	_ "github.com/lib/pq"
)


type Application struct{
	AppCtx context.Context
	DB  *sql.DB
}

func NewDB(ctx context.Context) (*sql.DB,error ){
	db, err := sql.Open("postgres", os.Getenv("DB_DSN"))
	if err != nil {
		return nil,err
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil ,err
	}
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(1*time.Hour)
	return db,nil


}