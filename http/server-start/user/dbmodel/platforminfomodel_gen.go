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
	platformInfoFieldNames          = builder.RawFieldNames(&PlatformInfo{})
	platformInfoRows                = strings.Join(platformInfoFieldNames, ",")
	platformInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(platformInfoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	platformInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(platformInfoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	platformInfoModel interface {
		Insert(ctx context.Context, data *PlatformInfo) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*PlatformInfo, error)
		Update(ctx context.Context, data *PlatformInfo) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPlatformInfoModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PlatformInfo struct {
		Id              int64         `db:"id"`               // 主键ID
		PlatformCode    string        `db:"platform_code"`    // 平台标识
		Mode            string        `db:"mode"`             // 运行模式（HTTP  FTP)
		FtpUser         string        `db:"ftp_user"`         // ftp用户名
		FtpAddress      string        `db:"ftp_address"`      // ftp地址
		FtpPwd          string        `db:"ftp_pwd"`          // ftp密码
		FtpInterval     sql.NullInt64 `db:"ftp_interval"`     // ftp读取间隔（秒）
		PdcAddress      string        `db:"pdc_address"`      // http模式必填，接口地址,业务地址
		OpmAddress      string        `db:"opm_address"`      // http模式必填，接口地址，运维地址
		StorageInterval sql.NullInt64 `db:"storage_interval"` // 存储状况反馈时间 分钟
		LogInterval     int64         `db:"log_interval"`     // 核验日志反馈时间间隔 分钟
		TimeSpan        sql.NullInt64 `db:"time_span"`        // 码图时效值（秒）
		FaceStatus      sql.NullInt64 `db:"face_status"`      // 刷脸状态(1正常 2删除 3停用)
		FeatureTotal    sql.NullInt64 `db:"feature_total"`    // 特征值容量阈值
		DbdayStatus     sql.NullInt64 `db:"dbday_status"`     // 是否启用人脸比对小库（1启用 2关闭）
		YajSwitch       int64         `db:"yaj_switch"`       // 易安检开关, 0 - 关闭，1 - 打开
		Status          int64         `db:"status"`           // 平台状态(1正常  2删除  3停用)
		CreateTime      time.Time     `db:"create_time"`      // 初始化时间
		UpdateTime      time.Time     `db:"update_time"`      // 更新时间
	}
)

func newPlatformInfoModel(conn sqlx.SqlConn) *defaultPlatformInfoModel {
	return &defaultPlatformInfoModel{
		conn:  conn,
		table: "`platform_info`",
	}
}

func (m *defaultPlatformInfoModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultPlatformInfoModel) FindOne(ctx context.Context, id int64) (*PlatformInfo, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", platformInfoRows, m.table)
	var resp PlatformInfo
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

func (m *defaultPlatformInfoModel) Insert(ctx context.Context, data *PlatformInfo) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, platformInfoRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.PlatformCode, data.Mode, data.FtpUser, data.FtpAddress, data.FtpPwd, data.FtpInterval, data.PdcAddress, data.OpmAddress, data.StorageInterval, data.LogInterval, data.TimeSpan, data.FaceStatus, data.FeatureTotal, data.DbdayStatus, data.YajSwitch, data.Status)
	return ret, err
}

func (m *defaultPlatformInfoModel) Update(ctx context.Context, data *PlatformInfo) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, platformInfoRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.PlatformCode, data.Mode, data.FtpUser, data.FtpAddress, data.FtpPwd, data.FtpInterval, data.PdcAddress, data.OpmAddress, data.StorageInterval, data.LogInterval, data.TimeSpan, data.FaceStatus, data.FeatureTotal, data.DbdayStatus, data.YajSwitch, data.Status, data.Id)
	return err
}

func (m *defaultPlatformInfoModel) tableName() string {
	return m.table
}
