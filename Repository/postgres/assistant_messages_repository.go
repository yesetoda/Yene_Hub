package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
	"gorm.io/gorm"
)

type assistantMessagesRepository struct {
	BaseRepository
}

func NewAssistantMessagesRepository(db *gorm.DB) repository.AssistantMessagesRepository {
	return &assistantMessagesRepository{
		BaseRepository: NewBaseRepository(db, "assistant_message"),
	}
}

func (r *assistantMessagesRepository) Create(message *entity.AssistantMessage) error {
	err := r.db.Create(message).Error
	if err != nil {
		return err
	}

	// Cache the newly created message
	_ = r.cacheDetail("byid", message, message.ID)
	_ = r.cacheDetail("byuser", message, message.UserID)

	// Invalidate list caches
	r.invalidateCache("list")
	
	return nil
}

func (r *assistantMessagesRepository) GetByID(id uint) (*entity.AssistantMessage, error) {
	var message entity.AssistantMessage
	err := r.getCachedDetail("byid", &message, func() error {
		return r.db.First(&message, id).Error
	}, id)
	
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func (r *assistantMessagesRepository) GetByUserID(userID uint) ([]*entity.AssistantMessage, error) {
	var messages []*entity.AssistantMessage
	err := r.getCachedList("byuser", &messages, func() error {
		return r.db.Where("user_id = ?", userID).Find(&messages).Error
	}, userID)
	
	return messages, err
}

func (r *assistantMessagesRepository) List() ([]*entity.AssistantMessage, error) {
	var messages []*entity.AssistantMessage
	err := r.getCachedList("list", &messages, func() error {
		return r.db.Find(&messages).Error
	})
	
	return messages, err
}

func (r *assistantMessagesRepository) Update(message *entity.AssistantMessage) error {
	err := r.db.Save(message).Error
	if err != nil {
		return err
	}

	// Update detail caches
	_ = r.cacheDetail("byid", message, message.ID)
	_ = r.cacheDetail("byuser", message, message.UserID)

	// Invalidate list caches
	r.invalidateCache("list")
	
	return nil
}

func (r *assistantMessagesRepository) Delete(id uint) error {
	var message entity.AssistantMessage
	if err := r.db.First(&message, id).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&message).Error; err != nil {
		return err
	}

	// Invalidate all caches for this message
	r.invalidateAllCache()
	
	return nil
}
