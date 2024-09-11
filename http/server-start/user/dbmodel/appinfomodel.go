package dbmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ AppInfoModel = (*customAppInfoModel)(nil)

type (
	// AppInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAppInfoModel.
	AppInfoModel interface {
		appInfoModel
		withSession(session sqlx.Session) AppInfoModel
	}

	customAppInfoModel struct {
		*defaultAppInfoModel
	}
)

// NewAppInfoModel returns a model for the database table.
func NewAppInfoModel(conn sqlx.SqlConn) AppInfoModel {
	return &customAppInfoModel{
		defaultAppInfoModel: newAppInfoModel(conn),
	}
}

func (m *customAppInfoModel) withSession(session sqlx.Session) AppInfoModel {
	return NewAppInfoModel(sqlx.NewSqlConnFromSession(session))
}
