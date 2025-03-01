// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package repository

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        int32     `json:"id"`
	Pid       uuid.UUID `json:"pid"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}
