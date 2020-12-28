package pgUtils

import (
	_ "github.com/lib/pq"
	"github.com/xormplus/xorm"
	xormLog "github.com/xormplus/xorm/log"
	"log"
)

// 官方文档： https://www.kancloud.cn/xormplus/xorm/167077

// PostgreSQL支持openSSL,通过sslmode设置:
// disable         : 只尝试非SSL连接
// allow           ：首先尝试非SSL连接，若失败再尝试SSL连接
// prefer (default)：首先尝试SSL连接，若失败再尝试非SSL连接
// require         ：只尝试SSL连接，若有根证书存在，等同于verify-ca
// verify-ca       ：只尝试SSL连接，并用根证书验证服务器证书是不是根CA签发的
// verify-full     ：只尝试SSL连接，并用根证书验证服务器证书是不是根CA签发的，且主题必须匹配连接域名或IP地址
var (
	POSTGRE_URL    = "postgres://postgres:aaaaaa@10.24.12.150:5432/mytest?sslmode=disable"
	SHOW_SQL       = true
	LOG_LEVEL      = xormLog.LOG_DEBUG
	MAX_IDLE_CONNS = 0
	MAX_OPEN_CONNS = 0

	globalEngine *xorm.Engine
)

type PGConnection struct {
	Engine *xorm.Engine
}

// 设置是否显示sql语句
func enableLog(pg *PGConnection) {
	if pg.Engine != nil && SHOW_SQL {
		pg.Engine.ShowSQL(SHOW_SQL)
		pg.Engine.Logger().SetLevel(LOG_LEVEL)
	}
}

// 设置连接属性
func setOptions(pg *PGConnection) {
	if pg.Engine != nil {
		if MAX_IDLE_CONNS != 0 {
			pg.Engine.SetMaxIdleConns(MAX_IDLE_CONNS)
		}
		if MAX_OPEN_CONNS != 0 {
			pg.Engine.SetMaxOpenConns(MAX_OPEN_CONNS)
		}
	}
}

// 获取Engine，已经连接成功就复制一份，否则创建新的
func (pg *PGConnection) ConnDB() error {
	if globalEngine == nil {
		var err error
		globalEngine, err = xorm.NewPostgreSQL(POSTGRE_URL)
		if err != nil {
			log.Println("create postgres connection failed!!! error:", err)
			return err
		}
	}
	engine, err := globalEngine.Clone() // 复制不会拷贝配置的属性
	if err != nil {
		log.Println("clone engine from globalEngine failed!!! error:", err)
		return err
	}
	pg.Engine = engine
	enableLog(pg)
	setOptions(pg)
	return nil
}

// 关闭（将用完的engine放回池中，不是断开与数据库的连接）
func (pg *PGConnection) CloseDB() {
	if pg.Engine != nil {
		err := pg.Engine.Close()
		if err != nil {
			log.Println("close pg database sources failed!!! error:", err)
		}
	}
}

// 切换Schema
func (pg *PGConnection) ChangeSchema(schema string) {
	if pg.Engine != nil {
		pg.Engine.SetSchema(schema)
	}
}

// ping
func (pg *PGConnection) Ping() error {
	if pg.Engine != nil {
		err := pg.Engine.Ping()
		if err != nil {
			log.Println("postgres Ping failed!!! error:", err)
			return err
		}
	}
	return nil
}
