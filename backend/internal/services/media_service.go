package services

import "github.com/dcp-project/backend/internal/models"

func CreateMedia(messageID uint64, url string, mediaType string, size int64) models.Media {

	media := models.Media{
		MessageID: messageID,
		URL:       url,
		Type:      mediaType,
		Size:      size,
	}

	return media
}
