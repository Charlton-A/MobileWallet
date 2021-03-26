package handlers

import "github.com/gin-gonic/gin"


func (a *Application)PingHandler(c *gin.Context){
	c.JSON(200, gin.H{"message": "pong"})
}

