package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// "用户名:密码@[连接方式](主机名:端口号)/数据库名"
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang-db")
	if err != nil {
		fmt.Println("connect mysql error: ", err)
		return
	}
	//关闭数据库连接
	defer db.Close()

	// 操作四：多行查询
	rows, _ := db.Query("select id, user_name as userName from user") //获取所有数据
	var id, userName string
	for rows.Next() { //循环显示所有的数据
		rows.Scan(&id, &userName)
		fmt.Println(id, "--", userName)
	}
}
