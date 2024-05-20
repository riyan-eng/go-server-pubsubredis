package model

import (
	"database/sql"
	"time"
)

type Example struct {
	UUID      string         `gorm:"column:uuid"`
	Nama      sql.NullString `gorm:"column:nama"`
	Detail    sql.NullString `gorm:"column:detail"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
}

func (Example) TableName() string {
	return "example"
}

type ModelExample struct {
	UUID   string         `db:"uuid"`
	Name   sql.NullString `db:"nama"`
	Detail sql.NullString `db:"detail"`
}
