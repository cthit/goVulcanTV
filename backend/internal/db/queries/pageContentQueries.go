package queries

import "github.com/swexbe/govulcantv/internal/db/models"

func GetPageContents() ([]*models.PageContent, error) {
	db := getDB()
	var pageContents []*models.PageContent
	tx := db.Find(&pageContents)
	return pageContents, tx.Error
}