package model

import "time"

type Product struct {
	Id         int64     `gorm:"column:id" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`
	ProductKey int64     `gorm:"column:product_key" json:"product_key"`
	Version    int64     `gorm:"column:version_id" json:"version_id"`
	Processes  string    `gorm:"column:processes" json:"processes"`
	Note       string    `gorm:"column:note" json:"note"`
	Extra      string    `gorm:"column:extra" json:"extra"`
	CreateTime time.Time `gorm:"column:create_time;default:'1970-01-02 00:00:00'" json:"create_time"`
	ModifyTime time.Time `gorm:"column:modify_time;default:'1970-01-02 00:00:00'" json:"modify_time"`
}

func (Product) TableName() string {
	return "product"
}
