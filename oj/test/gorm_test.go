package test

import (
	"cncyx.xyz/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
)

func TestGormTest(t *testing.T) {
	username := "root"
	password := "root"
	host := "127.0.0.1"
	port := 3306
	Dbname := "gin_gorm_oj"
	timeout := "10s"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	data := make([]*models.ProblemBasic, 0)
	err = db.Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("aa=>", data)
	for _, v := range data {
		fmt.Printf("Problem => %v\n", v)
	}

}
