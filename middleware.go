package middleware

import (
	"github.com/gin-gonic/gin"
)

func Get() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		WithSportContext(),
	}
}
