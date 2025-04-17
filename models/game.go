package models

type Game struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	UserID      uint    `json:"user_id"`                    // Внешний ключ
	User        User    `json:"-" gorm:"foreignKey:UserID"` // Связь с пользователем
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
