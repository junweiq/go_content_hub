package modal

import "time"

type Content struct {
	ID             int64
	Title          string
	Description    string
	Author         string
	VideoUrl       string `db:"video_url"`
	CoverUrl       string `db:"cover_url"`
	Category       string
	Duration       time.Duration
	Resolution     string
	Filesize       int64
	Extension      string
	Quality        int
	ApprovalStatus int       `db:"approval_status"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}
