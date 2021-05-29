package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //只是用这个包力的init
)

func main() {
	// DSN:Data Source Name
	//数据库信息 'remoteuser',mysql:host=47.241.25.1;dbname=offline_oms;charset=UTF8;
	var goodsSn string

	dsn := "remoteuser:keSJAr0SfBxx5MYw@tcp(47.241.25.1:3306)/offline_oms?charset=utf8"
	db, err := sql.Open("mysql", dsn) //生成句柄
	if err != nil {
		panic(err)
	}
	defer db.Close() // 注意这行代码要写在上面err判断的下面
	err = db.Ping()  //检测是否出现错误
	if err != nil {
		fmt.Printf("open %s invaild ,err'=", err)
		return
	}
	fmt.Println("connect database succsed")

	sqlStr := "select goods_sn from ly_goods where goods_id=?" //占位符
	//2.调用查询函数
	rowObj := db.QueryRow(sqlStr, "2")

	columns, err := rowObj.Columns()

	fmt.Println(columns)

}
