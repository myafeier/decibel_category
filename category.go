package decibel_category

import (
	"github.com/go-xorm/xorm"
)

const (
	StateOk      State = 1
	StateDeleted State = -1
)

type CateType int8 //分类类型
type State int8    //状态

type ICategory interface {
	Init() error //初始化
	Add(category *Category) (id int64, err error)
	Delete(id int64) (err error)
	Update(id int64, category *Category) (err error)
	Get(category *Category) (has bool, err error)
	GetChild(categoryType CateType, withChildList bool, mid, pid int64) (result []*Category, err error)
}

var stdCategory ICategory

func Init(name string, session *xorm.Session) ICategory {
	switch name {

	default:
		stdCategory = &DefaultCategory{
			session: session,
		}
		stdCategory.Init()
	}
	return stdCategory
}

func Add(category *Category) (id int64, err error) {
	return stdCategory.Add(category)
}
func Delete(id int64) (err error) {
	return stdCategory.Delete(id)
}
func Update(id int64, category *Category) (err error) {
	return stdCategory.Update(id, category)
}
func GetChild(categoryType CateType, withChildList bool, mid, pid int64) (result []*Category, err error) {
	return stdCategory.GetChild(categoryType, withChildList, mid, pid)
}

type Category struct {
	Id          int64       `json:"id"`
	CateOwnerId int64       `json:"-" xorm:"default 0 index"` //
	CateType    CateType    `json:"-" xorm:"tinyint(2) default 0 index"`
	State       State       `json:"state" xorm:"tinyint(2) default 1 index"`
	ListOrder   int         `json:"list_order" xorm:"default 10000 index"` //排序，越小越靠前，默认10000
	ParentId    int64       `json:"parent_id" xorm:"default 0 index"`
	Name        string      `json:"name" xorm:"varchar(200) default ''"`
	Icon        string      `json:"icon" xorm:"varchar(500) default ''"`
	Child       []*Category `json:"child" xorm:"-"`
	HasChild    bool        `json:"has_child" xorm:"-"`
}
