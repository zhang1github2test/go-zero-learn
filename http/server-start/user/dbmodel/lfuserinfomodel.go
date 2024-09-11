package dbmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ LfUserInfoModel = (*customLfUserInfoModel)(nil)

type (
	// LfUserInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLfUserInfoModel.
	LfUserInfoModel interface {
		lfUserInfoModel
		withSession(session sqlx.Session) LfUserInfoModel
	}

	customLfUserInfoModel struct {
		*defaultLfUserInfoModel
	}
)

// NewLfUserInfoModel returns a model for the database table.
func NewLfUserInfoModel(conn sqlx.SqlConn) LfUserInfoModel {
	return &customLfUserInfoModel{
		defaultLfUserInfoModel: newLfUserInfoModel(conn),
	}
}

func (m *customLfUserInfoModel) withSession(session sqlx.Session) LfUserInfoModel {
	return NewLfUserInfoModel(sqlx.NewSqlConnFromSession(session))
}
