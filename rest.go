package middleware

import (
	"net/http"
	"time"

	"git.sportsdots.ru/sportsdots-backend/go-util/logger.git"
	"git.sportsdots.ru/sportsdots-backend/go-util/sportctx.git"
	"github.com/gin-gonic/gin"
)

func NewRestBuilder() *RestBuilder {
	return &RestBuilder{}
}

type RestBuilder struct {
	logger  gin.HandlerFunc
	timeout time.Duration
	options []sportctx.Opt
	rfunc   gin.RecoveryFunc
	mws     []gin.HandlerFunc
}

func (b *RestBuilder) WithLogger(logger gin.HandlerFunc) *RestBuilder {
	b.logger = logger

	return b
}

func (b *RestBuilder) WithContextTimeout(timeout time.Duration) *RestBuilder {
	b.timeout = timeout

	return b
}

func (b *RestBuilder) WithContextOptions(options ...sportctx.Opt) *RestBuilder {
	b.options = options

	return b
}

func (b *RestBuilder) WithRecoveryFunc(rfunc gin.RecoveryFunc) *RestBuilder {
	b.rfunc = rfunc

	return b
}

func (b *RestBuilder) WithAdditionalMiddlewares(mws ...gin.HandlerFunc) *RestBuilder {
	b.mws = mws

	return b
}

func (b *RestBuilder) Build(router *gin.Engine) []gin.HandlerFunc {
	res := make([]gin.HandlerFunc, 0, 6+len(b.mws))

	res = append(res, gin.Recovery())

	if b.logger != nil {
		res = append(res, b.logger)
	}

	if b.timeout > 0 {
		res = append(res, WithSportContextWithTimeout(b.timeout, b.options...))
	} else {
		res = append(res, WithSportContext(b.options...))
	}

	if b.rfunc == nil {
		b.rfunc = defaultRecoveryFunc
	}

	res = append(res, gin.CustomRecovery(b.rfunc))

	if len(b.mws) > 0 {
		res = append(res, b.mws...)
	}

	return res
}

func defaultRecoveryFunc(c *gin.Context, p any) {
	logger.Error(c.Request.Context()).Msgf("panic: %v", p)

	c.AbortWithStatus(http.StatusInternalServerError)
}
