package sequence

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const sqlReplaceIntoStub =`REPLACE INTO sequence (stub) VALUES ('a')`

type MySQL struct {
	conn sqlx.SqlConn
}
func NewMySQL(dsn string)*MySQL{

	return  &MySQL{
		conn: sqlx.NewMysql(dsn),
	}
}
// Next 取下一个自增 ID
func (m *MySQL) Next() (seq uint64, err error) {
	// 1. prepare
	stmt, err := m.conn.Prepare(sqlReplaceIntoStub)
	if err != nil {
		logx.Errorw("conn.Prepare failed", logx.LogField{Key: "err", Value: err.Error()})
		return 0, err
	}
	defer stmt.Close()

	// 2. 执行
	result, err := stmt.Exec()
	if err != nil {
		logx.Errorw("stmt.Exec() failed", logx.LogField{Key: "err", Value: err.Error()})
		return 0, err
	}

	// 3. 获取插入的 id
	lid, err := result.LastInsertId()
	if err != nil {
		logx.Errorw("result.LastInsertId() failed", logx.LogField{Key: "err", Value: err.Error()})
		return 0, err
	}

	return uint64(lid), nil
}