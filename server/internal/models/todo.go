package models

import "time"

type Todo struct {
    ID          string    `gorm:"primaryKey" json:"id"`
    Title       string    `gorm:"not null" json:"title"`
    Description string    `json:"description"`
    Completed   bool      `gorm:"default:false" json:"completed"`
    CreatedAt   time.Time `gorm:"autoCreateTime" json:"createdAt"`
    UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}