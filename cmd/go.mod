module github.com/charlton/practs/mwallet/cmd

go 1.16

replace github.com/charlton/practs/mwallet/handlers => ../handlers

replace github.com/charlton/practs/mwallet/models => ../models

replace github.com/charlton/practs/mwallet/utils => ../utils

replace github.com/charlton/practs/mwallet/forms => ../forms

require (
	github.com/charlton/practs/mwallet/forms v0.0.0-00010101000000-000000000000 // indirect
	github.com/charlton/practs/mwallet/handlers v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.6.3

)
