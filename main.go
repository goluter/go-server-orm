package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main()  {
	// sudo /usr/local/mysql/support-files/mysql.server start
	// https://medium.com/@amoebamach/go%EC%96%B8%EC%96%B4%EC%97%90%EC%84%9C-orm-gorm-83ab33ecdc98
	// Create the database handle, confirm driver is present
	//defer db.Close()
	// Connect and check the server version
	//var  version  string
	//db.QueryRow("SELECT test FROM TEST ").Scan(&version)
	//fmt.Println("Connected to:", version)
	createTable()
}

type Person struct {
	gorm.Model
	Name string
	Products []Product
}

type Product struct {
	gorm.Model
	Code string
	Price uint64
}

func createTable() {
	db, err  := gorm.Open("mysql", "r:r!@tcp(localhost:3306)/r?charset=utf8&parseTime=True&loc=Local")
	if err !=  nil {
		panic("failed to connect database")
	}
	defer db.Close()
	// 스키마를 마이그레이트(테이블 생성)
	db.AutoMigrate(&Person{})
	db.AutoMigrate(&Product{})

	// Create :  레코드 삽입/추가(insert)
	db.Create(&Product{Code: "L1212", Price: 1000})

	// Read : 읽기
	var  product Product
	db.First(&product, 1) // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Update - product의  price 를 2000 으로 갱신/수정
	db.Model(&product).Update("Price", 2000)
	// Delete - product를 삭제
	db.Delete(&product)
}
