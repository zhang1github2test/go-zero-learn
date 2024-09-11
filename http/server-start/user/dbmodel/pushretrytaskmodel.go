package dbmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PushRetryTaskModel = (*customPushRetryTaskModel)(nil)

type (
	// PushRetryTaskModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPushRetryTaskModel.
	PushRetryTaskModel interface {
		pushRetryTaskModel
		withSession(session sqlx.Session) PushRetryTaskModel
	}

	customPushRetryTaskModel struct {
		*defaultPushRetryTaskModel
	}
)

// NewPushRetryTaskModel returns a model for the database table.
func NewPushRetryTaskModel(conn sqlx.SqlConn) PushRetryTaskModel {
	return &customPushRetryTaskModel{
		defaultPushRetryTaskModel: newPushRetryTaskModel(conn),
	}
}

func (m *customPushRetryTaskModel) withSession(session sqlx.Session) PushRetryTaskModel {
	return NewPushRetryTaskModel(sqlx.NewSqlConnFromSession(session))
}
