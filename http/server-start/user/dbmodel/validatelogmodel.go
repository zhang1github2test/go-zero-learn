package dbmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ValidateLogModel = (*customValidateLogModel)(nil)

type (
	// ValidateLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customValidateLogModel.
	ValidateLogModel interface {
		validateLogModel
		withSession(session sqlx.Session) ValidateLogModel
	}

	customValidateLogModel struct {
		*defaultValidateLogModel
	}
)

// NewValidateLogModel returns a model for the database table.
func NewValidateLogModel(conn sqlx.SqlConn) ValidateLogModel {
	return &customValidateLogModel{
		defaultValidateLogModel: newValidateLogModel(conn),
	}
}

func (m *customValidateLogModel) withSession(session sqlx.Session) ValidateLogModel {
	return NewValidateLogModel(sqlx.NewSqlConnFromSession(session))
}
