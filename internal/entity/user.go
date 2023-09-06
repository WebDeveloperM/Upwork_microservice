package entity

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"users"`
	Id            int    `bun:"id,pk,autoincrement"`
	Username      string `bun:"username"`
	Password      string `bun:"password"`
	Role          string `bun:"role"`
}
