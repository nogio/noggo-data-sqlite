package data_sqlite

import (
	. "github.com/nogio/noggo/base"
	"github.com/nogio/noggo/driver"
	"database/sql"
	"errors"
)

type (
	//数据库连接
	SqliteConnect struct {
		config Map

		//数据库对象
		db  *sql.DB
		//文件路由
		file string
		//模型
		models  map[string]Map
	}
)

//打开连接
func (conn *SqliteConnect) Open() error {
	db, err := sql.Open(SQLDRIVER, conn.file)
	if err != nil {
		return errors.New("数据库连接失败：" + err.Error())
	} else {
		conn.db = db
		return nil
	}
}
//关闭连接
func (conn *SqliteConnect) Close() error {
	if conn.db != nil {
		err := conn.db.Close()
		conn.db = nil
		return err
	}
	return nil
}


//注册模型
func (conn *SqliteConnect) Model(name string, config Map) {
	conn.models[name] = config
}

func (conn *SqliteConnect) Base(name string, cache driver.CacheBase) (driver.DataBase,error) {
	return &SqliteBase{name, conn, conn.models, conn.db, nil, cache, false},nil
}
