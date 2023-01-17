package middleware

import (
	"context"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DbKeyType string

var DbKey DbKeyType = "DB"

var CreateDbContext = func(database *gorm.DB, h *handler.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), DbKey, database)
		c.Request = c.Request.WithContext(ctx)
		h.ServeHTTP(c.Writer, c.Request)
	}
}

var GetDbFromContext = func(ctx context.Context) *gorm.DB {
	return ctx.Value(DbKey).(*gorm.DB)
}
