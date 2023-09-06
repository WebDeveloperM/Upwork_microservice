package entity

import (
	"github.com/uptrace/bun"
)

type Vacancy struct {
	bun.BaseModel `bun:"vacancy"`
	Id            int     `bun:"id,pk,autoincrement"`
	FullName      string  `bun:"full_name"`
	PhoneNumber   string  `bun:"phone_number"`
	Email         string  `bun:"email"`
	CV            *string `bun:"cv"`
	Rating        int     `bun:"rating"`
	Success       bool    `bun:"success"`
	IsDelete      bool    `bun:"is_delete"`
	//UpdatedId     int       `bun:"updated_id"`
	//UpdatedAt     time.Time `bun:"updated_at"`
	//DeletedId     int       `bun:"deleted_id"`
	//DeletedAt     time.Time `bun:"deleted_at"`
}
