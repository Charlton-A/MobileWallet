package handlers

import (
	"log"
	"net/http"

	"github.com/charlton/practs/mwallet/forms"
	"github.com/charlton/practs/mwallet/models"
	"github.com/charlton/practs/mwallet/utils"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)




func (a *Application) Transfer(c *gin.Context){

    ctx:= c.Request.Context()
	var t models.Transaction
	var f forms.Transfer


	err := c.ShouldBindJSON(&f)
	if err != nil {
		log.Printf("error on parsing json %#v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error":"Wrong format or field in data",})
		return
	}

 	tX, err:=  a.DB.BeginTx(ctx,nil)
	if err != nil {
		log.Printf("error on transaction create %#v", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	source_debit := `UPDATE mobile.wallets SET balance=balance-$1 ,updated_at=now() WHERE id= $2 AND user_id=$3`
	res , err := tX.ExecContext(ctx, source_debit,f.Amount, f.SourceWalletID, f.SourceUserID)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code=="23514"{
				c.JSON(http.StatusBadRequest, gin.H{"error":"source wallet does not have enough balance to perform transfer",})
				return
			}
        }
		log.Printf("error on storing tranasction %#v", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	rows ,_:=res.RowsAffected()
	if rows<=0{
		c.JSON(http.StatusBadRequest, gin.H{"error":"source_user_id or source_wallte_id does not exists",})
		tX.Rollback()
		return

	}

    transKey, _ :=utils.KeyGen(f.DestUserID)

	dest_credit := `UPDATE mobile.wallets SET balance=balance+$1 , updated_at=now() WHERE id= $2 AND user_id=$3`
	res  , err = tX.ExecContext(ctx, dest_credit,f.Amount, f.DestWalletID, f.DestUserID)
	if err != nil {
		log.Printf("error on transaction create %#v", err)
		tX.Rollback()
		c.Status(http.StatusInternalServerError)
		return
	}
	rows ,_=res.RowsAffected()
	if rows<=0{
		c.JSON(http.StatusBadRequest, gin.H{"error":"dest_user_id or dest_wallte_id does not exists",})
		tX.Rollback()
		return

	}

	err = tX.Commit()
	if err != nil {
		log.Printf("error on transfer transaction commit %#v", err)
		tX.Rollback()
		c.Status(http.StatusInternalServerError)
		return
	}

	t.Amount=-f.Amount
	t.Status=1
	t.Type=2
	t.UserID=f.SourceUserID
	t.WalletID=f.SourceWalletID
	t.TransKey=transKey
	err = t.Store(ctx, a.DB ,true)
	if err != nil {
		log.Printf("error on transaction store for debit %#v", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	t.Amount=f.Amount
	t.Status=1
	t.Type=3
	t.UserID=f.DestUserID
	t.WalletID=f.DestWalletID
	t.TransKey=transKey
	err = t.Store(ctx, a.DB ,true)
	if err != nil {
		log.Printf("error on transaction store for credit %#v", err)
		c.Status(http.StatusInternalServerError)
		return
	}








}
