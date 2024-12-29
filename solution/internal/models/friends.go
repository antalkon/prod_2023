package models

import "time"

type Friends struct {
	ID       int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID   int       `gorm:"not null;index" json:"userId"`   // Индексирование и запрет на NULL
	FriendID int       `gorm:"not null;index" json:"friendId"` // Индексирование и запрет на NULL
	AddedAt  time.Time `gorm:"autoCreateTime" json:"addedAt"`  // Автоматическое добавление времени
}
type AddFriend struct {
	Login string `json:"login"`
}

type FriendResponse struct {
	FriendID int       `json:"-"`       // ID друга для внутреннего использования
	AddedAt  time.Time `json:"addedAt"` // Время добавления друга
}
