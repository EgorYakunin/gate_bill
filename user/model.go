package user

import (
	"time"
)

type User struct {
    Id          int       `json:"id"`
    Email       string    `json:"email"`
    Username    string    `json:"username"`
    CreatedAt   time.Time `json:"created_at"`
    TelegramId  string    `json:"telegram_id"`
    AvatarPath  string    `json:"avatar_path"`
    FamilyId    int       `json:"family_id"`
}
