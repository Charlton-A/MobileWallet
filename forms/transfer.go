package forms


type Transfer struct{
	SourceUserID int `json:"source_user_id" binding:"required"`
	SourceWalletID int `json:"source_wallet_id" binding:"required"`
	DestUserID int     `json:"dest_user_id" binding:"required"`
	DestWalletID int    `json:"dest_wallet_id" binding:"required"`
	Amount float64        `json:"amount" binding:"required"`
}


