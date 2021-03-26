module github.com/charlton/practs/mwallet/cli

go 1.16

replace github.com/charlton/practs/mwallet/handlers => ../handlers

replace github.com/charlton/practs/mwallet/models => ../models

replace github.com/charlton/practs/mwallet/utils => ../utils

replace github.com/charlton/practs/mwallet/forms => ../forms

require (
	github.com/charlton/practs/mwallet/handlers v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.6.3
	github.com/natefinch/lumberjack v2.0.0+incompatible
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect

)
