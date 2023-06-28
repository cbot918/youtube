package api

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Dao struct {
	DB *sql.DB
}

func NewDao(db *sql.DB) *Dao {
	return &Dao{
		DB: db,
	}
}

func (d *Dao) Ping() error {
	err := d.DB.Ping()
	if err != nil {
		fmt.Println("db connect failed:", err)
		return err
	}
	fmt.Println("db connect success")
	return nil
}

func (d *Dao) ShowSelectAll(table string) error {
	sqlstr := fmt.Sprintf("select email,password from %s", table)

	rows, err := d.DB.Query(sqlstr, 0)
	if err != nil {
		fmt.Println("多行查询失败", err)
		return err
	}
	defer rows.Close() // 一定要释放rows的连接池

	for rows.Next() { //只要有，就会一直循环下去
		var u User
		err := rows.Scan(&u.Email, &u.Password)
		if err != nil {
			fmt.Println("scan row failed: ", err)
		}
		fmt.Println(u)
	}
	return nil
}

func (d *Dao) EmailExists(email string, table string) bool {
	sqlstr := fmt.Sprintf("SELECT email FROM %s WHERE email='%s'", table, email)

	email = ""
	row := d.DB.QueryRow(sqlstr)
	err := row.Scan(&email)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		fmt.Println("query email failed: ", err)
		return true
	}
	return true
}

func (d *Dao) PasswordCorrect(email string, pass string, table string) bool {

	sqlstr := fmt.Sprintf("SELECT password FROM %s WHERE email='%s'", table, email)

	pass_in_db := ""
	row := d.DB.QueryRow(sqlstr)
	err := row.Scan(&pass_in_db)
	if err != nil {
		fmt.Println("query email failed: ", err)
		return false
	}

	if pass != pass_in_db {
		return false
	}

	return true
}

func (d *Dao) AddUser(email string, password string, table string) bool {
	sqlstr := fmt.Sprintf("INSERT INTO %s (email, password) VALUES('%s','%s')", table, email, password)

	_, err := d.DB.Exec(sqlstr)
	if err != nil {
		fmt.Println("add user failed: ", err)
		return false
	}

	return true
}
