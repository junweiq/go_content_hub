package service

import (
	"go_content_hub/internal/util"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type CmsApp struct {
	Db  *sqlx.DB
	Rdb *redis.Client
}

func NewCmsApp() *CmsApp {
	app := &CmsApp{}
	app.Db = util.ConnDB()
	app.Rdb = util.ConnRdb()
	return app
}
