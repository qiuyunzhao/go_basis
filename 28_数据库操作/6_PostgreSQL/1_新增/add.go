/*
@ Time : 2020/11/3 15:31
@ Author : qyz
@ File : add
@ Software: GoLand
@ Description:
*/

package add

import (
	"database/sql"
	"go_basis/28_数据库操作/6_PostgreSQL/entites"
	"go_basis/28_数据库操作/6_PostgreSQL/pgUtils"
)

// ============================================ ORM方式插入数据 ========================================================

// INSERT INTO "user" ("id","user_name","age","password","phone","e_mail") VALUES ($1,$2,$3,$4,$5,$6)
func addUser_ORM(user entites.User) (int64, error) {
	pg := pgUtils.PGConnection{}
	pg.ConnDB()
	defer pg.CloseDB()

	if affected, err := pg.Engine.Insert(user); err != nil {
		return 0, err
	} else {
		return affected, nil
	}
}

// ============================================ SQL方式插入数据 ========================================================

func addUser_SQL(user entites.User) (sql.Result, error) {
	pg := pgUtils.PGConnection{}
	pg.ConnDB()
	defer pg.CloseDB()

	// 巨坑！ 表名要用""包含
	sql_1 := `INSERT INTO "user"(id,user_name,age,password,phone,e_mail) VALUES (?,?,?,?,?,?)`
	if res, err := pg.Engine.Exec(sql_1, user.Id, user.UserName, user.Age, user.Password, user.Phone, user.Email); err != nil {
		return nil, err
	} else {
		return res, nil
	}
}
