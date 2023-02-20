// Code generated by goctl. DO NOT EDIT.

package commentModel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	commentFieldNames          = builder.RawFieldNames(&Comment{})
	commentRows                = strings.Join(commentFieldNames, ",")
	commentRowsExpectAutoSet   = strings.Join(stringx.Remove(commentFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	commentRowsWithPlaceHolder = strings.Join(stringx.Remove(commentFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTiktokCommentIdPrefix = "cache:tiktok:comment:id:"
)

type (
	commentModel interface {
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		Insert(ctx context.Context, data *Comment) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Comment, error)
		FindAll(ctx context.Context, id int64) ([]*Comment, error)
		Update(ctx context.Context, data *Comment) error
		Delete(ctx context.Context, id int64) error
	}

	defaultCommentModel struct {
		sqlc.CachedConn
		table string
	}

	Comment struct {
		Id         int64     `db:"id"`
		UserId     int64     `db:"user_id"`
		VideoId    int64     `db:"video_id"`
		CreateTime time.Time `db:"create_time"`
		Removed    int64     `db:"removed"`
		Content    string    `db:"content"`
	}
)

func newCommentModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultCommentModel {
	return &defaultCommentModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`comment`",
	}
}

func (m *defaultCommentModel) Delete(ctx context.Context, id int64) error {
	tiktokCommentIdKey := fmt.Sprintf("%s%v", cacheTiktokCommentIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, tiktokCommentIdKey)
	return err
}

func (m *defaultCommentModel) FindOne(ctx context.Context, id int64) (*Comment, error) {
	tiktokCommentIdKey := fmt.Sprintf("%s%v", cacheTiktokCommentIdPrefix, id)
	var resp Comment
	err := m.QueryRowCtx(ctx, &resp, tiktokCommentIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", commentRows, m.table)
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

func (m *defaultCommentModel) Insert(ctx context.Context, data *Comment) (sql.Result, error) {
	tiktokCommentIdKey := fmt.Sprintf("%s%v", cacheTiktokCommentIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, commentRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.VideoId, data.Removed, data.Content)
	}, tiktokCommentIdKey)
	return ret, err
}

func (m *defaultCommentModel) FindAll(ctx context.Context, id int64) ([]*Comment, error) {
	fmt.Printf("findall:::::::::::::::::::::::::")
	tiktokCommentIdKey := fmt.Sprintf("%s%v", cacheTiktokCommentIdPrefix, id)
	var resp []*Comment
	err := m.QueryRowCtx(ctx, &resp, tiktokCommentIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `video_id` = ? and `removed` = 0 order by `create_time` DESC", commentRows, m.table)
		fmt.Printf("sql:------------->%v", query)
		return conn.QueryRowsCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCommentModel) Update(ctx context.Context, data *Comment) error {
	tiktokCommentIdKey := fmt.Sprintf("%s%v", cacheTiktokCommentIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, commentRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.VideoId, data.Removed, data.Content, data.Id)
	}, tiktokCommentIdKey)
	return err
}

func (m *defaultCommentModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheTiktokCommentIdPrefix, primary)
}

func (m *defaultCommentModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", commentRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultCommentModel) tableName() string {
	return m.table
}

// export logic
func (m *defaultCommentModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}
