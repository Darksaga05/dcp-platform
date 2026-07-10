package services

import "github.com/dcp-project/backend/internal/models"

func CreateConversation(userIDs []uint64) models.Conversation {

	conversation := models.Conversation{
		UserIDs: userIDs,
	}

	return conversation
}
