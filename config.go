package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var Engine *xorm.Engine

func init() {
	var err error
	Engine, err = xorm.NewEngine("mysql", "root:12345678@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		panic(fmt.Sprintf("init xorm err: %s", err))
	}
}
