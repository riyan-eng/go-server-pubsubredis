package datastruct

import "time"

type Example struct {
	ID        int       `db:"id" json:"-"`
	UUID      string    `db:"uuid" json:"id"`
	Nama      any       `db:"nama" json:"nama"`
	Detail    any       `db:"detail" json:"detail"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	TotalRows int       `db:"total_rows" json:"-"`
}

type DetailExample struct {
	UUID   string `db:"uuid" json:"id"`
	Name   string `db:"nama" json:"name"`
	Detail any    `db:"detail" json:"detail"`
}

type ListExample struct {
	ID        int       `db:"id" json:"-"`
	UUID      string    `db:"uuid" json:"id"`
	Name      any       `db:"nama" json:"name"`
	Detail    any       `db:"detail" json:"detail"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	TotalRows int       `db:"total_rows" json:"-"`
}
