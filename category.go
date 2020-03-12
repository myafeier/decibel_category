package category

import "time"

const (
	StateOk      State = 1
	StateDeleted State = -1
)

type State int8 //状态

type CategoryEntity struct {
	Id        int64     `json:"id"`
	State     State     `json:"state" xorm:"tinyint(2) default 1 index"`
	ListOrder int       `json:"list_order" xorm:"default 10000 index"` //排序，越小越靠前，默认10000
	ParentId  int64     `json:"parent_id" xorm:"default 0 index"`
	Name      string    `json:"name" xorm:"varchar(200) default ''"`
	Icon      string    `json:"icon" xorm:"varchar(500) default ''"` // icon
	Created   time.Time `json:"created" xorm:"created"`
	Updated   time.Time `json:"updated" xorm:"updated"`
}

func (e *CategoryEntity) TableName() string {
	return "category"
}
