package schemas

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Lead struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;" json:"id"`
	FirstName string     `gorm:"size:100;not null" json:"first_name"`
	LastName  string     `gorm:"size:100;not null" json:"last_name"`
	Email     string     `gorm:"size:100;uniqueIndex;not null" json:"email"`
	Phone     string     `gorm:"size:20;notnull" json:"phone"`
	Company   string     `gorm:"size:100;not null" json:"company"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (s *Lead) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New()
	s.CreatedAt = time.Now()
	return
}

func (s *Lead) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	s.UpdatedAt = &now
	return
}

type UpdateLeadRequest struct {
	ID   uuid.UUID              `json:"id"`
	Data map[string]interface{} `json:"data"`
}
