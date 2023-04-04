package demo_user

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
)

var _ DemoUserModel = (*customDemoUserModel)(nil)

type (
	// DemoUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDemoUserModel.
	DemoUserModel interface {
		demoUserModel
		GetList(ctx context.Context) ([]*DemoUser, error)
		Create(ctx context.Context, orm *gorm.DB, user *DemoUser) error
		GetInfoByUsername(ctx context.Context, orm *gorm.DB, username string) (*DemoUser, error)
		GetInfoById(ctx context.Context, orm *gorm.DB, uid int64) (*DemoUser, error)
	}

	customDemoUserModel struct {
		*defaultDemoUserModel
	}
)

// NewDemoUserModel returns a model for the database table.
func NewDemoUserModel(conn sqlx.SqlConn) DemoUserModel {
	return &customDemoUserModel{
		defaultDemoUserModel: newDemoUserModel(conn),
	}
}

func (m *customDemoUserModel) GetList(ctx context.Context) ([]*DemoUser, error) {
	var resp []*DemoUser
	query := fmt.Sprintf("select %s from %s", demoUserRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customDemoUserModel) Create(ctx context.Context, orm *gorm.DB, user *DemoUser) error {
	if err := orm.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (m *customDemoUserModel) GetInfoByUsername(ctx context.Context, orm *gorm.DB, username string) (*DemoUser, error) {
	var resp *DemoUser
	db := orm.Model(&DemoUser{}).Where("`user_name` = ?", username)
	err := db.First(&resp).Error
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (m *customDemoUserModel) GetInfoById(ctx context.Context, orm *gorm.DB, uid int64) (*DemoUser, error) {
	var resp *DemoUser
	db := orm.Model(&DemoUser{}).Where("`id` = ?", uid)
	err := db.First(&resp).Error
	if err != nil {
		return resp, err
	}
	return resp, nil
}
