package utils

import (
	"context"

	"gorm.io/gorm"
)

func WithDbContext(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, "db", db)
}
