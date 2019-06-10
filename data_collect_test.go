package main

import (
	"testing"
	"time"
)

type DataCollectPlans struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	ProjectId int       `json:"project_id" xorm:"not null comment('项目id') INT(10)"`
	Name      string    `json:"name" xorm:"not null default '' comment('采集任务名称') VARCHAR(32)"`
	StartTime time.Time `json:"start_time" xorm:"not null comment('开始时间') DATETIME"`
	EndTime   time.Time `json:"end_time" xorm:"not null comment('结束时间') DATETIME"`
	Status    int       `json:"status" xorm:"not null default 0 comment('采集状态(1待采集，2采集中，3已完成，5暂停，6结束，7删除)') TINYINT(1)"`
	IsDeleted int       `json:"is_deleted" xorm:"not null TINYINT(1)"`
	UpdatedAt time.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' updated TIMESTAMP"`
	CreatedAt time.Time `json:"created_at" xorm:"not null default 'CURRENT_TIMESTAMP' created TIMESTAMP"`
}

func TestCollectPlanInsertOne(t *testing.T) {
	rl := &DataCollectPlans{
		Name: "一次测试",
	}

	_, err := Engine.InsertOne(rl)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("insert row num is: ", rl.Id)
}
