package quiz

import (
	"time"

	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	StartTime time.Time  `gorm:"not null;" json:"start_time" xml:"start_time" form:"start_time" validate:"required"`
	Questions []Question ``
}

type Question struct {
	gorm.Model
	QuizID uint   `json:"quiz_id" xml:"quiz_id" form:"quiz_id" validate:"required"`
	Quiz   Quiz   `validate:"-"`
	Text   string `gorm:"not null;" json:"text" xml:"text" form:"text" validate:"required"`
}
