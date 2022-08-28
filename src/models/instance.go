package models

import (
	"gorm.io/gorm"
	"time"
)

type Instance struct {
	gorm.Model
	ID             string `json:"id" gorm:"primary_key"`
	Users          uint   `json:"users"`
	PublicQuizzes  uint   `json:"public_quizzes"`
	PrivateQuizzes uint   `json:"private_quizzes"`
	IP             string `json:"ip"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
