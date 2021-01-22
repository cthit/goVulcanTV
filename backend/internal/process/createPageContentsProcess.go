package process

import (
	"github.com/swexbe/govulcantv/internal/db/commands"
	"github.com/swexbe/govulcantv/internal/db/models"
)

func CreatePageContentStruct(youtubeID, description, addedBy string, length uint32) error {
	pageContent := models.PageContent{
		YoutubeID:     youtubeID,
		Enabled:       true,
		Description:   description,
		AddedBy:       addedBy,
		LengthSeconds: length,
	}

	return CreatePageContent(&pageContent)
}

func CreatePageContent(content *models.PageContent) error {
	return commands.InsertPageContent(content)
}