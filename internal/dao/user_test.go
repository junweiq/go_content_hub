package dao

import (
	"go_content_hub/internal/modal"
	"go_content_hub/internal/util"
	"testing"

	"github.com/jmoiron/sqlx"
)

func TestUserDao_Create(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		user *modal.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "用戶創建",
			fields: fields{
				db: util.ConnDB(),
			},
			args: args{
				user: &modal.User{
					Username: "test",
					Password: "test",
					Nickname: "test",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserDao{
				db: tt.fields.db,
			}
			if err := u.Create(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
