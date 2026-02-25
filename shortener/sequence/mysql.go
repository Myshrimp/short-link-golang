package sequence

import (
	"database/sql"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

// establish a connection to mysql database and run REPLACE INTO
// REPLACE INTO sequence (stub) VALUES ('a')
// SELECT LAST_INSERT_ID() to get the id of the last inserted record

const sqlReplaceStub = `REPLACE INTO sequence (stub) VALUES ('a')`
type MySQL struct {
	conn sqlx.SqlConn
}

func NewMySQL(dsn string) Sequence {
	conn := sqlx.NewMysql(dsn)
	return &MySQL{
		conn: conn,
	}
}

func (m *MySQL) Next() (seq uint64, err error) {
	var stmt sqlx.StmtSession
	stmt, err = m.conn.Prepare(sqlReplaceStub)
	if err != nil {
		logx.Errorw("conn.Prepare failed", logx.LogField{Key: "errValue", Value: err.Error()})
		return 0, err
	}
	defer stmt.Close()
	var res sql.Result
	res, err = stmt.Exec()
	if err != nil {
		logx.Errorw("stmt.Exec failed", logx.LogField{Key: "errValue", Value: err.Error()})
		return 0, err
	}
	
	var lid int64
	lid, err = res.LastInsertId()
	if err != nil {
		logx.Errorw("res.LastInsertId failed", logx.LogField{Key: "errValue", Value: err.Error()})
		return 0, err
	}
	return uint64(lid), nil
}