package process

import (
	"errors"
	"github.com/swexbe/govulcantv/internal/db/commands"
	"github.com/swexbe/govulcantv/internal/db/models"
	"github.com/swexbe/govulcantv/internal/db/queries"
	"github.com/swexbe/govulcantv/internal/vulcanTvErrors"
	"gorm.io/gorm"
)

func CreatePageContent(content *models.PageContent) (uint64, error) {
	_, err := queries.GetPageContentByIdLength(content.YoutubeID, content.LengthSeconds)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			content.ID = 0 // Make sure it's auto-incremented
			return commands.InsertPageContent(content)
		}

		return 0, err
	}

	return 0, vulcanTvErrors.ErrAlreadyExists
}