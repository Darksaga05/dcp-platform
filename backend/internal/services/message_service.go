package services

import "github.com/dcp-project/backend/internal/models"

func CreateMessage(senderID uint64, receiverID uint64, content string) models.Message {

	message := models.Message{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Content:    content,
		Encrypted:  true,
	}

	return message
}
