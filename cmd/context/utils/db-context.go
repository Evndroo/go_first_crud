package utils

import (
	"context"
	"log"

	"gorm.io/gorm"
)

func WithDbContext(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, "db", db)
}

func GetDbFromContext(ctx context.Context) (*gorm.DB, bool) {
	db, success := ctx.Value("db").(*gorm.DB)

	if !success {
		log.Fatalln("Error getting db from context")
		return nil, false
	}

	return db, true

}
