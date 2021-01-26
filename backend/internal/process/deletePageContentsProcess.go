package process

import (
	"errors"
	"github.com/google/uuid"
	"github.com/swexbe/govulcantv/internal/db/commands"
	"github.com/swexbe/govulcantv/internal/db/queries"
	"github.com/swexbe/govulcantv/internal/vulcanTvErrors"
	"gorm.io/gorm"
)

func DeletePageContent(pageContentId uuid.UUID) error {
	pageContent, err := queries.GetPageContentById(pageContentId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return vulcanTvErrors.NoSuchPageContent
		}
		return err
	}

	err = commands.DeletePageContent(pageContent)
	return err
}