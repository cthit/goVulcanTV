package process

import (
	"errors"
	"github.com/swexbe/govulcantv/internal/common"
	"github.com/swexbe/govulcantv/internal/db/queries"
	"github.com/swexbe/govulcantv/internal/player"
	"github.com/swexbe/govulcantv/internal/vulcanTvErrors"
	"gorm.io/gorm"
)

func OverrideVideo(pageContentId uint64) error {
	pageContent, err := queries.GetPageContentById(pageContentId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return vulcanTvErrors.NoSuchPageContent
		}
		return err
	}

	player.ForcePlayVideo(&common.Video{
		Id: pageContent.YoutubeID,
		LengthSeconds: pageContent.LengthSeconds,
	})

	return nil
}
