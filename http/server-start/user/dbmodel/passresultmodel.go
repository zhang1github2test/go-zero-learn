package dbmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PassResultModel = (*customPassResultModel)(nil)

type (
	// PassResultModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPassResultModel.
	PassResultModel interface {
		passResultModel
		withSession(session sqlx.Session) PassResultModel
	}

	customPassResultModel struct {
		*defaultPassResultModel
	}
)

// NewPassResultModel returns a model for the database table.
func NewPassResultModel(conn sqlx.SqlConn) PassResultModel {
	return &customPassResultModel{
		defaultPassResultModel: newPassResultModel(conn),
	}
}

func (m *customPassResultModel) withSession(session sqlx.Session) PassResultModel {
	return NewPassResultModel(sqlx.NewSqlConnFromSession(session))
}
