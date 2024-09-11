package dbmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ MerchantAppModel = (*customMerchantAppModel)(nil)

type (
	// MerchantAppModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMerchantAppModel.
	MerchantAppModel interface {
		merchantAppModel
		withSession(session sqlx.Session) MerchantAppModel
	}

	customMerchantAppModel struct {
		*defaultMerchantAppModel
	}
)

// NewMerchantAppModel returns a model for the database table.
func NewMerchantAppModel(conn sqlx.SqlConn) MerchantAppModel {
	return &customMerchantAppModel{
		defaultMerchantAppModel: newMerchantAppModel(conn),
	}
}

func (m *customMerchantAppModel) withSession(session sqlx.Session) MerchantAppModel {
	return NewMerchantAppModel(sqlx.NewSqlConnFromSession(session))
}
