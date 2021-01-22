package commands

import (
	"github.com/swexbe/govulcantv/internal/db/models"
)

func InsertPageContent(pageContent *models.PageContent) (uint64, error) {
	db := getDB()
	result := db.Create(pageContent)
	return pageContent.ID, result.Error
}