package dbmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ FailUserRealNameModel = (*customFailUserRealNameModel)(nil)

type (
	// FailUserRealNameModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFailUserRealNameModel.
	FailUserRealNameModel interface {
		failUserRealNameModel
		withSession(session sqlx.Session) FailUserRealNameModel
	}

	customFailUserRealNameModel struct {
		*defaultFailUserRealNameModel
	}
)

// NewFailUserRealNameModel returns a model for the database table.
func NewFailUserRealNameModel(conn sqlx.SqlConn) FailUserRealNameModel {
	return &customFailUserRealNameModel{
		defaultFailUserRealNameModel: newFailUserRealNameModel(conn),
	}
}

func (m *customFailUserRealNameModel) withSession(session sqlx.Session) FailUserRealNameModel {
	return NewFailUserRealNameModel(sqlx.NewSqlConnFromSession(session))
}
