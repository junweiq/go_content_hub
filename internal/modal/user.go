package modal

import "time"

type User struct {
	ID        int64
	Username  string
	Password  string
	Nickname  string
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
