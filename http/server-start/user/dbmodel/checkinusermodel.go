package dbmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CheckInUserModel = (*customCheckInUserModel)(nil)

type (
	// CheckInUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCheckInUserModel.
	CheckInUserModel interface {
		checkInUserModel
		withSession(session sqlx.Session) CheckInUserModel
	}

	customCheckInUserModel struct {
		*defaultCheckInUserModel
	}
)

// NewCheckInUserModel returns a model for the database table.
func NewCheckInUserModel(conn sqlx.SqlConn) CheckInUserModel {
	return &customCheckInUserModel{
		defaultCheckInUserModel: newCheckInUserModel(conn),
	}
}

func (m *customCheckInUserModel) withSession(session sqlx.Session) CheckInUserModel {
	return NewCheckInUserModel(sqlx.NewSqlConnFromSession(session))
}
