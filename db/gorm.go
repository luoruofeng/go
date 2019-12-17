package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)
import _ "github.com/jinzhu/gorm/dialects/mysql"

type Student struct {
	// gorm.Model
	Name  string
	Sex   bool
	Birth time.Time
}

func GormApp() {
	db, err := gorm.Open("mysql", "root:luoruofeng@tcp(localhost:8888)/school?charset=utf8&parseTime=true")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Student{})

	db.Create(&Student{Name: "gorm", Sex: true, Birth: time.Now()})

	var s Student
	db.First(&s, "name = ?", "gorm")

	fmt.Println(s)

	db.Model(&s).Update("birth", time.Now())

	fmt.Println("update birth " + s.Birth.Format("2006-01-02 15:04:05"))

	db.Delete(&s)
	fmt.Println("delete")
}
