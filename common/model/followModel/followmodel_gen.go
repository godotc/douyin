// Code generated by goctl. DO NOT EDIT.

package followModel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	followFieldNames          = builder.RawFieldNames(&Follow{})
	followRows                = strings.Join(followFieldNames, ",")
	followRowsExpectAutoSet   = strings.Join(stringx.Remove(followFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	followRowsWithPlaceHolder = strings.Join(stringx.Remove(followFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTiktokFollowIdPrefix = "cache:tiktok:follow:id:"
)

type (
	followModel interface {
		Insert(ctx context.Context, data *Follow) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Follow, error)
		FindAllByUserId(ctx context.Context, userId string) ([]*Follow, error)
		FindAllByFunId(ctx context.Context, funId string) ([]*Follow, error)
		CountByFollowRelation(ctx context.Context, id int64, field string) (int64, error)
		Update(ctx context.Context, data *Follow) error
		Delete(ctx context.Context, id int64) error
	}

	defaultFollowModel struct {
		sqlc.CachedConn
		table string
	}

	Follow struct {
		Id      int64          `db:"id"`
		UserId  sql.NullInt64  `db:"user_id"`
		FunId   int64          `db:"fun_id"`
		Removed int64          `db:"removed"`
		Msg     sql.NullString `db:"msg"`
	}
)

func newFollowModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultFollowModel {
	return &defaultFollowModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`follow`",
	}
}

func (m *defaultFollowModel) Delete(ctx context.Context, id int64) error {
	tiktokFollowIdKey := fmt.Sprintf("%s%v", cacheTiktokFollowIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, tiktokFollowIdKey)
	return err
}

func (m *defaultFollowModel) FindOne(ctx context.Context, id int64) (*Follow, error) {
	tiktokFollowIdKey := fmt.Sprintf("%s%v", cacheTiktokFollowIdPrefix, id)
	var resp Follow
	err := m.QueryRowCtx(ctx, &resp, tiktokFollowIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", followRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultFollowModel) Insert(ctx context.Context, data *Follow) (sql.Result, error) {
	tiktokFollowIdKey := fmt.Sprintf("%s%v", cacheTiktokFollowIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, followRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.FunId, data.Removed, data.Msg)
	}, tiktokFollowIdKey)
	return ret, err
}

func (m *defaultFollowModel) Update(ctx context.Context, data *Follow) error {
	tiktokFollowIdKey := fmt.Sprintf("%s%v", cacheTiktokFollowIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, followRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.FunId, data.Removed, data.Msg, data.Id)
	}, tiktokFollowIdKey)
	return err
}

func (m *defaultFollowModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheTiktokFollowIdPrefix, primary)
}

func (m *defaultFollowModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", followRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultFollowModel) tableName() string {
	return m.table
}
