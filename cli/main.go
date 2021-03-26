package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/charlton/practs/mwallet/handlers"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
)

func SetUp() (*gin.Engine, *handlers.Application) {
	router := gin.Default()

	appCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	db, err  :=handlers.NewDB(appCtx)
	if err!=nil{
		log.Panicf("unable to initialize db due to error %s", err)

	}

	app:=handlers.Application{AppCtx: appCtx ,DB: db}

	if err!=nil{
		log.Panicf("unable to initialize application due to error %s", err)

	}

	apiKey:=os.Getenv("APP_KEY")
	apiPass:=os.Getenv("APP_PASS")

	v1 := router.Group("api/v1/users/" ,gin.BasicAuth(gin.Accounts{
		apiKey: apiPass,}) )
	{
		v1.GET("/ping",app.PingHandler)
		v1.POST("/create", app.UserCreate)
		v1.GET("/balance/:user_id", app.UserBalance)
		v1.POST("wallet/create", app.WalletCreate)
		v1.POST("/wallet/update", app.WalletUpdate)
		v1.POST("/wallet/transfer", app.Transfer)
		v1.GET("/transactions/:user_id", app.UserTransaction)
	}
	return router,&app
}

func init() {
	l := &lumberjack.Logger{
		Compress:   false,
		Filename:   "api-logs.log",
		MaxSize:    35,
		MaxBackups: 10,
	}
	log.SetOutput(io.MultiWriter(gin.DefaultWriter, l))
	log.SetPrefix("[WALLET-API] ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)

}

func main(){
	r, app := SetUp()
	log.Println("Apllication starting")
	defer app.DB.Close()
	r.Run()

}