package modal

import "time"

type Content struct {
	ID             int64         `gorm:"column:id;primary_key"`
	Title          string        `gorm:"column:title"`
	Description    string        `gorm:"column:description"`
	Author         string        `gorm:"column:author"` //作者
	VideoUrl       string        `gorm:"column:video_url"`
	CoverUrl       string        `gorm:"column:cover_url"`  //封面url
	Category       string        `gorm:"column:category"`   //內容分類
	Duration       time.Duration `gorm:"column:duration"`   //內容時長
	Resolution     string        `gorm:"column:resolution"` //分辨率 如720p、1080p
	Filesize       int64         `gorm:"column:filesize"`
	Extension      string        `gorm:"column:extension"`  //文件格式 如mp4
	Quality        int           `gorm:"column:created_at"` //視頻質量 1高清 2標清
	ApprovalStatus int           `gorm:"column:created_at"` //審核狀態 1審核中 2審核通過 3審核不通過
	CreatedAt      time.Time     `gorm:"column:created_at"`
	UpdatedAt      time.Time     `gorm:"column:updated_at"`
}

func (a Content) TableName() string {
	return "cms_content.t_content_detail"
}
