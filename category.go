package category

import "time"

const (
	StateOk      State = 1  //正常
	StateDeleted State = -1 //已删
)

type State int8 //状态

type CategoryEntity struct {
	Id        int64     `json:"id"`
	State     State     `json:"state" xorm:"tinyint(2) default 1 index"`
	ListOrder int       `json:"list_order" xorm:"default 10000 index"` //排序，越小越靠前，默认10000
	Pid       int64     `json:"pid" xorm:"default 0 index"`
	Name      string    `json:"name" xorm:"varchar(200) default ''"`
	Icon      string    `json:"icon,omitempty" xorm:"varchar(500) default ''"` // icon
	Created   time.Time `json:"-" xorm:"created"`
	Updated   time.Time `json:"-" xorm:"updated"`
}

func (e *CategoryEntity) TableName() string {
	return "category"
}
