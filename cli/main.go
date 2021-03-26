package main

import (
	"context"
	"log"
	"os"

	"github.com/charlton/practs/mwallet/handlers"
	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine{
	router := gin.Default()

	appCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err, db  :=handlers.NewDB(appCtx)
	if err!=nil{
		log.Panicf("unable to initialize db due to error %s", err)

	}
	defer db.Close()
	app:=handlers.Application{AppCtx: appCtx ,DB: db}

	if err!=nil{
		log.Panicf("unable to initialize application due to error %s", err)

	}

	apiKey:=os.Getenv("APP_KEY")
	apiPass:=os.Getenv("APP_PASS")
	log.Printf("api pass %#v" , apiPass)

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
	return router
}

func main(){
	r := SetUp()
	r.Run()

}