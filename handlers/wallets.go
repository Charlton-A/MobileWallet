package handlers

import (
	"log"
	"net/http"

	"github.com/charlton/practs/mwallet/models"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (a *Application)WalletCreate(c *gin.Context){
	ctx := c.Request.Context()

	var w models.Wallet

	err := c.ShouldBindJSON(&w)
	if err != nil {
		log.Printf("error on parsing json %#v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error":"Wrong format or field in data",})
		return
	}
	err = w.Create(ctx, a.DB)
	if err != nil {
		log.Printf("error on wallet create %#v", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}


func (a *Application) WalletUpdate(c *gin.Context){
	ctx := c.Request.Context()

	var w models.Wallet
	var t models.Transaction

	err := c.ShouldBindJSON(&w)
	if err != nil {
		log.Printf("error on parsing json %#v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error":"Wrong format or field in data",})
		return
	}
	if w.Action==4|| w.Action==2 {
		w.Amount=0-w.Amount
	}else if w.Action==1 || w.Action==3{
		w.Amount=w.Amount-0

	}else{
		c.JSON(http.StatusBadRequest, gin.H{"error":"update action not allowed",})
		return

	}
	err = w.UpdateBalance(ctx, a.DB)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code=="23514"{
				c.JSON(http.StatusBadRequest, gin.H{"error":"wallet has no balance",})
				return
			}
        }
		log.Printf("error on storing tranasction %#v", err)
		c.Status(http.StatusInternalServerError)
		return
	}

    t.Amount=w.Amount
	t.Status=1
	t.UserID=w.UserID
	t.WalletID=w.WalletID
	t.Type=w.Action

	err =t.Store(ctx, a.DB ,false)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code=="23503"{
				c.JSON(http.StatusBadRequest, gin.H{"error":"wallet_id or user_id does not exist",})
				return
			}
        }
		log.Printf("error on storing tranasction %#v", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)

}