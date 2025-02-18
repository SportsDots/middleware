package middleware

import (
	"time"

	"git.sportsdots.ru/sportsdots-backend/go-util/logger.git"
	"git.sportsdots.ru/sportsdots-backend/go-util/sportctx.git"
	"github.com/gin-gonic/gin"
)

// WithSportContext ...
func WithSportContext(opts ...sportctx.Opt) gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		ctxOpts := getOptionsForCtx(ginCtx, opts...)

		newCtx := sportctx.NewContext(
			ctxOpts...,
		)

		ginCtx.Request = ginCtx.Request.WithContext(newCtx)

		ginCtx.Next()
	}
}

// WithSportContextWithTimeout ...
func WithSportContextWithTimeout(timeout time.Duration, opts ...sportctx.Opt) gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		ctxOpts := getOptionsForCtx(ginCtx, opts...)

		newCtx, cancel := sportctx.NewContextWithTimeout(
			timeout,
			ctxOpts...,
		)
		defer cancel()

		ginCtx.Request = ginCtx.Request.WithContext(newCtx)

		ginCtx.Next()
	}
}

func getOptionsForCtx(ginCtx *gin.Context, opts ...sportctx.Opt) []sportctx.Opt {
	request := ginCtx.Request
	prevCtx := request.Context()

	values := make(map[string]string)

	for _, param := range Headers {
		values[param] = request.Header.Get(param)
	}

	ctxOpts := []sportctx.Opt{
		sportctx.WithContext(prevCtx),
		sportctx.WithRequestID(GetXRequestIDHeaderFromRequest(request)),
		sportctx.WithThroughParams(values),
		sportctx.WithLogger(logger.Get().Logger),
	}
	ctxOpts = append(ctxOpts, opts...)

	return ctxOpts
}
