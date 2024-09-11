package dbmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ FailPassResultModel = (*customFailPassResultModel)(nil)

type (
	// FailPassResultModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFailPassResultModel.
	FailPassResultModel interface {
		failPassResultModel
		withSession(session sqlx.Session) FailPassResultModel
	}

	customFailPassResultModel struct {
		*defaultFailPassResultModel
	}
)

// NewFailPassResultModel returns a model for the database table.
func NewFailPassResultModel(conn sqlx.SqlConn) FailPassResultModel {
	return &customFailPassResultModel{
		defaultFailPassResultModel: newFailPassResultModel(conn),
	}
}

func (m *customFailPassResultModel) withSession(session sqlx.Session) FailPassResultModel {
	return NewFailPassResultModel(sqlx.NewSqlConnFromSession(session))
}
