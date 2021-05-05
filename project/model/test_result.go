package model

import "time"

type TestResult struct {
	Id         int64     `gorm:"column:id" json:"id"`
	TicketId   int64     `gorm:"column:ticket_id" json:"ticket_id"`
	NodeKey    int64     `gorm:"column:node_key" json:"node_key"`
	NodeName   string    `gorm:"column:node_name" json:"node_name"`
	NodeNum    int64     `gorm:"column:node_num" json:"node_num"`
	Upper      float64   `gorm:"column:node_requirement_upper" json:"upper"`
	Lower      float64   `gorm:"column:node_requirement_lower" json:"lower"`
	Unit       string    `gorm:"column:node_requirement_unit" json:"unit"`
	Value      float64   `gorm:"column:node_test_value" json:"value"`
	Tester     int64     `gorm:"column:tester" json:"tester"`
	Note       string    `gorm:"column:note" json:"note"`
	Extra      string    `gorm:"column:extra" json:"extra"`
	CreateTime time.Time `gorm:"column:create_time;default:'1970-01-02 00:00:00'" json:"create_time"`
	ModifyTime time.Time `gorm:"column:modify_time;default:'1970-01-02 00:00:00'" json:"modify_time"`
}

func (TestResult) TableName() string {
	return "test_result"
}
