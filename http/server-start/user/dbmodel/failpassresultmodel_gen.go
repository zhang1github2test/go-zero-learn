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
	failPassResultFieldNames          = builder.RawFieldNames(&FailPassResult{})
	failPassResultRows                = strings.Join(failPassResultFieldNames, ",")
	failPassResultRowsExpectAutoSet   = strings.Join(stringx.Remove(failPassResultFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	failPassResultRowsWithPlaceHolder = strings.Join(stringx.Remove(failPassResultFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	failPassResultModel interface {
		Insert(ctx context.Context, data *FailPassResult) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*FailPassResult, error)
		Update(ctx context.Context, data *FailPassResult) error
		Delete(ctx context.Context, id int64) error
	}

	defaultFailPassResultModel struct {
		conn  sqlx.SqlConn
		table string
	}

	FailPassResult struct {
		Id            int64     `db:"id"`              // 主键ID
		LivePhoto     string    `db:"live_photo"`      // 现场照片
		PassResult    int64     `db:"pass_result"`     // 通行记录结果 1-成功, 0-失败
		PassResultMsg string    `db:"pass_result_msg"` // 通行结果描述
		CreateTime    time.Time `db:"create_time"`     // 创建时间
		TraceId       string    `db:"trace_id"`        // 核验流水号
		Mhid          string    `db:"mhid"`            // 通行旅客的mhid
		RequestId     string    `db:"request_id"`
	}
)

func newFailPassResultModel(conn sqlx.SqlConn) *defaultFailPassResultModel {
	return &defaultFailPassResultModel{
		conn:  conn,
		table: "`fail_pass_result`",
	}
}

func (m *defaultFailPassResultModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultFailPassResultModel) FindOne(ctx context.Context, id int64) (*FailPassResult, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", failPassResultRows, m.table)
	var resp FailPassResult
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

func (m *defaultFailPassResultModel) Insert(ctx context.Context, data *FailPassResult) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, failPassResultRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.LivePhoto, data.PassResult, data.PassResultMsg, data.TraceId, data.Mhid, data.RequestId)
	return ret, err
}

func (m *defaultFailPassResultModel) Update(ctx context.Context, data *FailPassResult) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, failPassResultRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.LivePhoto, data.PassResult, data.PassResultMsg, data.TraceId, data.Mhid, data.RequestId, data.Id)
	return err
}

func (m *defaultFailPassResultModel) tableName() string {
	return m.table
}
