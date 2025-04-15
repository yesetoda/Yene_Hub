package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
	"gorm.io/gorm"
)

type AssistantMessageRepository struct {
	db *gorm.DB
}
func NewAssistantMessageRepository(db *gorm.DB) repository.AssistantMessagesRepository {
	return &AssistantMessageRepository{db: db}
}

func (r *AssistantMessageRepository) Create(assistantMessage *entity.AssistantMessage) error {
	return r.db.Create(assistantMessage).Error
}

func (r *AssistantMessageRepository) GetByID(id uint) (*entity.AssistantMessage, error) {
	var assistantMessage entity.AssistantMessage
	if err := r.db.First(&assistantMessage, id).Error; err != nil {
		return nil, err
	}
	return &assistantMessage, nil
}	

func (r *AssistantMessageRepository) GetByUserID(userID uint) ([]*entity.AssistantMessage, error) {
	var assistantMessages []*entity.AssistantMessage
	if err := r.db.Where("user_id = ?", userID).Find(&assistantMessages).Error; err != nil {
		return nil, err
	}
	return assistantMessages, nil
}

func (r *AssistantMessageRepository) List() ([]*entity.AssistantMessage, error) {
	var assistantMessages []*entity.AssistantMessage
	if err := r.db.Find(&assistantMessages).Error; err != nil {
		return nil, err
	}
	return assistantMessages, nil
}

func (r *AssistantMessageRepository) Update(assistantMessage *entity.AssistantMessage) error {
	return r.db.Save(assistantMessage).Error
}

func (r *AssistantMessageRepository) Delete(id uint) error {
	return r.db.Delete(&entity.AssistantMessage{}, id).Error
}
