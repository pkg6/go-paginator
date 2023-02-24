package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/pkg6/go-paginator"
	"github.com/pkg6/go-paginator/adapter"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Post struct {
	ID     uint `gorm:"primarykey" json:"id"`
	Number int  `json:"number"`
}

var db, _ = gorm.Open(sqlite.Open("gorm.db?cache=shared"), &gorm.Config{
	Logger: logger.Default.LogMode(logger.Info),
})

func initDb() {
	_ = db.AutoMigrate(&Post{})
	for i := 1; i <= 100; i++ {
		p := Post{
			Number: i,
		}
		db.Save(&p)
	}
}
func simple() {
	q := db.Model(Post{})
	var dest []Post
	var adapt = adapter.NewGORMAdapter(q)
	paginator := paginator.SimplePaginator(adapt, 10, 1)
	_ = paginator.Get(&dest)
	//获取最后页码
	page, err := paginator.GetLastPage()
	fmt.Println(fmt.Sprintf("获取最后页码:%v", page))
	fmt.Println(fmt.Sprintf("获取最后页码错误信息:%v", err))
	//获取总数
	total, err := paginator.GetTotal()
	fmt.Println(fmt.Sprintf("获取总数:%v", total))
	fmt.Println(fmt.Sprintf("获取总数错误信息:%v", err))
	fmt.Println(fmt.Sprintf("当前页码:%v", paginator.GetCurrentPage()))
	fmt.Println(fmt.Sprintf("每页显示多少条数:%v", paginator.GetListRows()))
	fmt.Println(fmt.Sprintf("是否还可以进行分页:%v", paginator.HasPages()))
	fmt.Println(dest)
}

func Total() {
	q := db.Model(Post{}).Where([]int64{20, 21, 22}).Order("id desc")
	var dest []Post
	var adapt = adapter.NewGORMAdapter(q)
	t, _ := adapt.Length()
	paginator := paginator.TotalPaginator(adapt, 10, 1, t)
	//_ = paginator.Get(&dest)
	r := paginator.GetResult(&dest)
	marshal, _ := json.Marshal(r)
	xmlbytes, err := xml.Marshal(r)
	fmt.Println(string(marshal))
	fmt.Println(string(xmlbytes))
	//获取最后页码
	page, err := paginator.GetLastPage()
	fmt.Println(fmt.Sprintf("获取最后页码:%v", page))
	fmt.Println(fmt.Sprintf("获取最后页码错误信息:%v", err))
	//获取总数
	total, err := paginator.GetTotal()
	fmt.Println(fmt.Sprintf("获取总数:%v", total))
	fmt.Println(fmt.Sprintf("获取总数错误信息:%v", err))
	fmt.Println(fmt.Sprintf("当前页码:%v", paginator.GetCurrentPage()))
	fmt.Println(fmt.Sprintf("每页显示多少条数:%v", paginator.GetListRows()))
	fmt.Println(fmt.Sprintf("是否还可以进行分页:%v", paginator.HasPages()))
	fmt.Println(dest)
}

func main() {
	Total()
}
