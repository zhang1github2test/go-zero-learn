package dbmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ FailValidateLogModel = (*customFailValidateLogModel)(nil)

type (
	// FailValidateLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFailValidateLogModel.
	FailValidateLogModel interface {
		failValidateLogModel
		withSession(session sqlx.Session) FailValidateLogModel
	}

	customFailValidateLogModel struct {
		*defaultFailValidateLogModel
	}
)

// NewFailValidateLogModel returns a model for the database table.
func NewFailValidateLogModel(conn sqlx.SqlConn) FailValidateLogModel {
	return &customFailValidateLogModel{
		defaultFailValidateLogModel: newFailValidateLogModel(conn),
	}
}

func (m *customFailValidateLogModel) withSession(session sqlx.Session) FailValidateLogModel {
	return NewFailValidateLogModel(sqlx.NewSqlConnFromSession(session))
}
