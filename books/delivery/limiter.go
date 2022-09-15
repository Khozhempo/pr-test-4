package delivery

import (
	"github.com/gin-gonic/gin"
)

var limiter = make(chan struct{}, 10)

func RpsLimiter(c *gin.Context) {
	limiterGetQuota()
	c.Next()
	limiterFreeQuota()
}

func limiterGetQuota() {
	var s struct{}
	limiter <- s
}

func limiterFreeQuota() {
	<-limiter
}
