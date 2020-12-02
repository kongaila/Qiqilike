package domain

import "time"

// 公告栏模型
type TbBulletin struct {
	Id         int32     `json:"id" gorm:"type:INT(11);NOT NULL"`
	Title      string    `json:"title" gorm:"type:VARCHAR(50);"`
	Content    string    `json:"content" gorm:"type:TEXT;"`
	CreateAt   time.Time `json:"createAt" gorm:"type:DATETIME(0);"`
	CreateUuid string    `json:"createUuid" gorm:"type:CHAR(32);"`
}
