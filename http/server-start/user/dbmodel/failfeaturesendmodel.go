package dbmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ FailFeatureSendModel = (*customFailFeatureSendModel)(nil)

type (
	// FailFeatureSendModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFailFeatureSendModel.
	FailFeatureSendModel interface {
		failFeatureSendModel
		withSession(session sqlx.Session) FailFeatureSendModel
	}

	customFailFeatureSendModel struct {
		*defaultFailFeatureSendModel
	}
)

// NewFailFeatureSendModel returns a model for the database table.
func NewFailFeatureSendModel(conn sqlx.SqlConn) FailFeatureSendModel {
	return &customFailFeatureSendModel{
		defaultFailFeatureSendModel: newFailFeatureSendModel(conn),
	}
}

func (m *customFailFeatureSendModel) withSession(session sqlx.Session) FailFeatureSendModel {
	return NewFailFeatureSendModel(sqlx.NewSqlConnFromSession(session))
}
