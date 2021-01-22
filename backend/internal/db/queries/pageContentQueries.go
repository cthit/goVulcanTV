package queries

import "github.com/swexbe/govulcantv/internal/db/models"

func GetPageContents() ([]*models.PageContent, error) {
	db := getDB()
	var pageContents []*models.PageContent
	tx := db.Find(&pageContents)
	return pageContents, tx.Error
}

func GetPageContentByIdLength(youtubeId string, length uint32) (*models.PageContent, error) {
	db := getDB()
	query := models.PageContent{
		YoutubeID: youtubeId,
		LengthSeconds: length,
	}
	var pageContent models.PageContent
	tx := db.Where(&query, "YoutubeID", "LengthSeconds").First(&pageContent)
	return &pageContent, tx.Error
}