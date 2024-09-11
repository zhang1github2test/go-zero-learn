package dbmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PlatformInfoModel = (*customPlatformInfoModel)(nil)

type (
	// PlatformInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPlatformInfoModel.
	PlatformInfoModel interface {
		platformInfoModel
		withSession(session sqlx.Session) PlatformInfoModel
	}

	customPlatformInfoModel struct {
		*defaultPlatformInfoModel
	}
)

// NewPlatformInfoModel returns a model for the database table.
func NewPlatformInfoModel(conn sqlx.SqlConn) PlatformInfoModel {
	return &customPlatformInfoModel{
		defaultPlatformInfoModel: newPlatformInfoModel(conn),
	}
}

func (m *customPlatformInfoModel) withSession(session sqlx.Session) PlatformInfoModel {
	return NewPlatformInfoModel(sqlx.NewSqlConnFromSession(session))
}
