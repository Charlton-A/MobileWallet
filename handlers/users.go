package handlers

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/charlton/practs/mwallet/forms"
	"github.com/charlton/practs/mwallet/models"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)


func (a *Application)UserCreate(c *gin.Context){
	var u models.User
	var w models.Wallet
	err := c.ShouldBindJSON(&u)
	if err != nil {
		log.Printf("error on parsing json %#v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error":"Wrong format or field in data",})
		return
	}

	ctx := c.Request.Context()
	userID,err :=u.Create(ctx , a.DB)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code=="23505"{
				c.JSON(http.StatusBadRequest, gin.H{"error":"Email or phone already registred",})
				return
			}
        }
		log.Printf("error on user create %#v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	w.Balance=0.0
	w.UserID=userID
	err = w.Create(ctx, a.DB)
	if err != nil {
		log.Printf("error on wallet create %#v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusCreated)

}

func (a *Application)UserBalance(c *gin.Context){
	ctx := c.Request.Context()
	var u models.User

	user := c.Param("user_id")
	userID, err := strconv.Atoi(user)
	if err != nil {
		log.Printf("error on user it conversion to int %#v", err)
		c.Status(http.StatusNotFound)
		return
	}
	err =u.Fetch(c , a.DB ,userID)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows){
			c.JSON(http.StatusBadRequest, gin.H{"error":"user does not exist",})
			return
		}
		log.Printf("error on db user fetch %#v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	wallets ,err := models.FetchBalances(ctx, a.DB, u.UserID)
	if err != nil {
		log.Printf("error on db wallets fetch %#v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	u.Wallets=wallets
	c.JSON(http.StatusOK, u)

}

func (a *Application)UserTransaction(c *gin.Context){
	ctx := c.Request.Context()
	var uT forms.UserTranasctionResp
	var u models.User

	user := c.Param("user_id")
	userID, err := strconv.Atoi(user)
	if err != nil {
		log.Printf("error on user it conversion to int %#v", err)
		c.Status(http.StatusNotFound)
		return
	}
	err =u.Fetch(c, a.DB ,userID)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows){
			c.JSON(http.StatusBadRequest, gin.H{"error":"user does not exist",})
			return
		}
		log.Printf("error on db user fetch %#v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	transactions ,err := models.FetchTransactions(ctx, a.DB, u.UserID)
	if err != nil {
		log.Printf("error on db transactions fetch %#v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	uT.FirstName=u.FirstName
	uT.LastName=u.LastName
	uT.UserID =u.UserID
	
	uT.Transactions =transactions
	c.JSON(http.StatusOK, uT)

}

