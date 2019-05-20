package main

import (
	"testing"
	"time"
)

type Base struct {
	Id         int       `xorm:"pk autoincr 'id'" json:"-"`
	Status     int       `json:"-"` // enable=0, disable=1
	CreateTime time.Time `xorm:"created 'create_time'" json:"-"`
	UpdateTime time.Time `xorm:"updated 'update_time'" json:"-"`
}

type LotteryLog struct {
	Base       `xorm:"extends"`
	UserId     int       `json:"user_id"`
	CardId     int       `json:"card_id"`
	CardType   int       `json:"card_type"`
	GetCashNum int       `json:"get_cash_num"`
	GetCoinNum int       `json:"get_coin_num"`
	PickedNum  string    `json:"picked_num"`
	WinningNum string    `json:"winning_num"`
	Won        int       `json:"won"`
	EndTime    time.Time `json:"end_time"`
}

func TestRewardLogInsertOne(t *testing.T) {
	rl := &LotteryLog{
		UserId:     1,
		CardId:     1,
		CardType:   1,
		GetCashNum: 1,
		GetCoinNum: 1,
		PickedNum:  "123456",
		WinningNum: "123456",
		EndTime:    time.Now(),
	}

	_, err := Engine.InsertOne(rl)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("insert row num is: ", rl.Base.Id)
}

func TestRewardLogUpdateByID(t *testing.T) {
	rl := &LotteryLog{
		UserId: 50,
	}
	affected, err := Engine.ID(1).Update(rl)
	if err != nil {
		return
	}
	t.Log("affected num is: ", affected)
}

func TestRewardLogSearchOneByID(t *testing.T) {
	rl := make([]LotteryLog, 0)
	err := Engine.Where("id=?", 1).Find(&rl)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(rl)
}
