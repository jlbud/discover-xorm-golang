package main

import (
	"testing"
)

type Base struct {
    Id int `xorm:"pk autoincr 'id'" json:"-"`
}

// 样例库
type User struct {
	Base  `xorm:"extends"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

// 插入一行
func TestInsert_One(t *testing.T) {
	u := &User{}
	u.Name = "yang"
	u.Phone = "11111111111"
	_, err := Engine.InsertOne(u)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("affected id is: ", u.Id)
}