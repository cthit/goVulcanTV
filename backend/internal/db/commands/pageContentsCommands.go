package commands

import (
	"github.com/swexbe/govulcantv/internal/db/models"
)

func InsertPageContent(pageContent *models.PageContent) error {
	db := getDB()
	result := db.Create(pageContent)
	return result.Error
}