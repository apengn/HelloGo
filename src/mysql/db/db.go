package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"encoding/json"
	"fmt"
)

type BaseInfo struct {
	Id          int    `json:"id"`
	CompanyName string `json:"company_name"`
}

var db *sql.DB

func init() {
	db, erro := sql.Open("mysql", "root:root@tcp(localhost:3306)/test?charset=utf8")

	if erro != nil {
		panic(erro.Error())
	}
	defer db.Close()
	erro = db.Ping()

	if erro != nil {
		panic(erro.Error())
	}


	rows, erro := db.Query("select id,company_name FROM baseinfo")

	defer rows.Close()
	baseinfo := make([]BaseInfo, 0)
	for rows.Next() {
		var id int
		var name string
		erro = rows.Scan(&id, &name)
		baseinfo = append(baseinfo, BaseInfo{Id: id, CompanyName: name})
	}

	fmt.Println(os.Getwd())

	f, erro := os.OpenFile("./src/mysql/db/base.json", os.O_RDWR, 0666)

	defer f.Close()
	if erro != nil {
		panic(erro.Error())
	}
	jsonByte, erro := json.Marshal(baseinfo)

	if erro != nil {
		panic(erro.Error())
	}

	n, erro := f.WriteString(string(jsonByte))

	fmt.Println(n, "\n", string(jsonByte))

	erro = rows.Err()
	if erro != nil {
		panic(erro.Error())
	}

}

func main() {

}

func Inser() {
	r, e := db.Exec(`INSERT INTO baseinfo (company_name) VALUES ('规划师跨国收购绝对时空观')`)
	if e != nil {
		fmt.Println(e.Error())
	}
	fmt.Println(r)
	Update()
}

func Update() {
	stmt, _ := db.Prepare(`UPDATE baseinfo SET company_name=? WHERE id=?`)
	stmt.Exec("供货商快捷方式看见过", 2)
	//db.Exec(`UPDATE baseinfo set company_name="54444444545" WHERE id=2`)
}

func Delete(){
	db.Exec(`DELETE FROM baseinfo WHERE id=1`)
	stmt,_:=db.Prepare(`DELETE FROM baseinfo WHERE id=?`)
	stmt.Exec(2)
}

//mysql事务处理
func Business(){

	tx,erro:=db.Begin()
	_,erro=tx.Exec(`DELETE FROM baseinfo WHERE id=1`)
	if erro!=nil {
      tx.Rollback()
	}
	tx.Commit()

}
