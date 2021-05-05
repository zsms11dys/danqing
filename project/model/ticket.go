package model

import "time"

type Ticket struct {
	Id         int64     `gorm:"column:id" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`
	ProductId  int64     `gorm:"column:product_id" json:"product_id"`
	ProductKey int64     `gorm:"column:product_key" json:"product_key"`
	Version    int64     `gorm:"column:product_version" json:"version_id"`
	TicketKey  int64     `gorm:"column:ticket_key" json:"ticket_key"`
	Count      int       `gorm:"column:product_count" json:"product_count"`
	Note       string    `gorm:"column:note" json:"note"`
	Extra      string    `gorm:"column:extra" json:"extra"`
	CreateTime time.Time `gorm:"column:create_time;default:'1970-01-02 00:00:00'" json:"create_time"`
	ModifyTime time.Time `gorm:"column:modify_time;default:'1970-01-02 00:00:00'" json:"modify_time"`
}

func (Ticket) TableName() string {
	return "ticket"
}
