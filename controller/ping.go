package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var html = `
<html>
<body>
  Pong! From Gitran API v1.<hr>
  <a href="/api/v1/auth/github"><Button>/api/v1/auth/github</Button></a>
</body>
</html>
`

func PingV1(ctx *gin.Context) {
	// ctx.String(http.StatusOK, "Pong! from Gitran API v1.")
	ctx.Header("Content-Type", "text/html")
	ctx.String(http.StatusOK, html)
}
