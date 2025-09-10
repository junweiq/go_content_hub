package modal

import "time"

type Content struct {
	ID             int64
	Title          string
	Description    string
	Author         string        //作者
	VideoUrl       string        `db:"video_url"`
	CoverUrl       string        `db:"cover_url"` //封面url
	Category       string        //內容分類
	Duration       time.Duration //內容時長
	Resolution     string        //分辨率 如720p、1080p
	Filesize       int64
	Extension      string    //文件格式 如mp4
	Quality        int       //視頻質量 1高清 2標清
	ApprovalStatus int       `db:"approval_status"` //審核狀態 1審核中 2審核通過 3審核不通過
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

func (a Content) TableName() string {
	return "cms_content.t_content_detail"
}
