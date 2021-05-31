package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" //只是用这个包力的init
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DemandPlanSearch struct {
	gorm.Model
	ProductLine string
	SaleStatus  int
	GoodsSn     string
	MadeType    int
}

var Config = struct[]string

var Config = struct[string] {
	mysql: "remoteuser:keSJAr0SfBxx5MYw@tcp(47.241.25.1:3306)/offline_oms?charset=utf8&parseTime=True"
}

func main() {

	dsn := "remoteuser:keSJAr0SfBxx5MYw@tcp(47.241.25.1:3306)/offline_oms?charset=utf8&parseTime=True"
	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "ly_", // 表名前缀，`User` 的表名应该是 `t_users`
				SingularTable: true,  // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
			},
		},
	)

	if err != nil {
		panic("failed to connect database")
	}

	lyDemandPlanSearch := DemandPlanSearch{}

	db.Where(16).First(&lyDemandPlanSearch)

	fmt.Println(lyDemandPlanSearch)

	// defer db.Close()

	// gorm.Open(mysql.Open("mysql", dsn), &gorm.Config{})
	//db, err := gorm.Open(sql.Open("mysql", dsn), &gorm.Config{})
}

// func main() {
// 	// DSN:Data Source Name
// 	//数据库信息 'remoteuser',mysql:host=47.241.25.1;dbname=offline_oms;charset=UTF8;

// 	dsn := "remoteuser:keSJAr0SfBxx5MYw@tcp(47.241.25.1:3306)/offline_oms?charset=utf8"
// 	db, err := sql.Open("mysql", dsn) //生成句柄
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close() // 注意这行代码要写在上面err判断的下面
// 	err = db.Ping()  //检测是否出现错误
// 	if err != nil {
// 		fmt.Printf("open %s invaild ,err'=", err)
// 		return
// 	}
// 	fmt.Println("connect database succsed")

// 	sqlStr := "select goods_sn from ly_goods where goods_id=?" //占位符
// 	//2.调用查询函数

// 	rowObj := db.QueryRow(sqlStr, "2")

// 	//	str, ok := data.(string)
// 	goodsSn := goodsModel(rowObj).(string)

// 	fmt.Println(goodsSn)

// }

// /**
//  *
//  * 很多很多很多的
//  *
//  *
//  *
//  */
// func goodsModel(row *sql.Row) interface{} {

// 	type Goods struct {
// 		goodsSn string
// 		author  string
// 		book_id int
// 	}

// 	var goods = Goods{
// 		goodsSn: "",
// 		author:  "",
// 		book_id: 0,
// 	}

// 	row.Scan(&goods.goodsSn)
// 	return goods
// }
