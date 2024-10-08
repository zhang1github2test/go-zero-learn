// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.2

package dbmodel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	lfUserInfoFieldNames          = builder.RawFieldNames(&LfUserInfo{})
	lfUserInfoRows                = strings.Join(lfUserInfoFieldNames, ",")
	lfUserInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(lfUserInfoFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	lfUserInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(lfUserInfoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	lfUserInfoModel interface {
		Insert(ctx context.Context, data *LfUserInfo) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*LfUserInfo, error)
		Update(ctx context.Context, data *LfUserInfo) error
		Delete(ctx context.Context, id int64) error
	}

	defaultLfUserInfoModel struct {
		conn  sqlx.SqlConn
		table string
	}

	LfUserInfo struct {
		Id         int64     `db:"id"`          // 主键
		OpenId     string    `db:"open_id"`     // 旅服平台自身的标识ID
		CertName   string    `db:"cert_name"`   // 姓名
		CertNum    string    `db:"cert_num"`    // 身份证号码
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 更新时间
	}
)

func newLfUserInfoModel(conn sqlx.SqlConn) *defaultLfUserInfoModel {
	return &defaultLfUserInfoModel{
		conn:  conn,
		table: "`lf_user_info`",
	}
}

func (m *defaultLfUserInfoModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultLfUserInfoModel) FindOne(ctx context.Context, id int64) (*LfUserInfo, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", lfUserInfoRows, m.table)
	var resp LfUserInfo
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultLfUserInfoModel) Insert(ctx context.Context, data *LfUserInfo) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, lfUserInfoRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.OpenId, data.CertName, data.CertNum)
	return ret, err
}

func (m *defaultLfUserInfoModel) Update(ctx context.Context, data *LfUserInfo) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, lfUserInfoRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.OpenId, data.CertName, data.CertNum, data.Id)
	return err
}

func (m *defaultLfUserInfoModel) tableName() string {
	return m.table
}
