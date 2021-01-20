package models

type PageContent struct {
	ID uint64 `gorm:"autoIncrement"`
	YoutubeID string `gorm:"unique;not null" json:"youtubeID"`
	Enabled bool `gorm:"not null" json:"enabled"`
	Description string `gorm:"not null" json:"description"`
	AddedBy string `gorm:"not null" json:"addedBy"`
	LengthSeconds uint32 `gorm:"not null" json:"length"`
}

func (_ PageContent) StructName() string {
	return "Page Content"
}