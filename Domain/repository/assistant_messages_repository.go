package repository

import "a2sv.org/hub/Domain/entity"

// AssistantMessagesRepository defines methods for assistant messages database operations
type AssistantMessagesRepository interface {
	Create(message *entity.AssistantMessage) error
	GetByID(id uint) (*entity.AssistantMessage, error)
	GetByUserID(userID uint) ([]*entity.AssistantMessage, error)
	Update(message *entity.AssistantMessage) error
	Delete(id uint) error
	List() ([]*entity.AssistantMessage, error)
}
