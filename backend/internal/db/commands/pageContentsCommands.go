package commands

import (
	"github.com/google/uuid"
	"github.com/swexbe/govulcantv/internal/db/models"
)

func InsertPageContent(pageContent *models.PageContent) (uuid.UUID, error) {
	db := getDB()
	result := db.Create(pageContent)
	return pageContent.ID, result.Error
}

func DeletePageContent(pageContent *models.PageContent) error {
	db := getDB()
	tx := db.Delete(pageContent)
	return tx.Error
}