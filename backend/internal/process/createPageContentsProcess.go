package process

import (
	"errors"
	"github.com/google/uuid"
	"github.com/swexbe/govulcantv/internal/db/commands"
	"github.com/swexbe/govulcantv/internal/db/models"
	"github.com/swexbe/govulcantv/internal/db/queries"
	"github.com/swexbe/govulcantv/internal/vulcanTvErrors"
	"gorm.io/gorm"
)

func CreatePageContent(content *models.PageContent) (uuid.UUID, error) {
	_, err := queries.GetPageContentByIdLength(content.YoutubeID, content.LengthSeconds)
	var id uuid.UUID
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			id, err = commands.InsertPageContent(content)
		}

		return id, err
	}

	return id, vulcanTvErrors.ErrAlreadyExists
}