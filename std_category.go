package decibel_category

import (
	"github.com/go-xorm/xorm"
)

type DefaultCategory struct {
	//Data map[CateType][]*Category
	//mutex sync.Mutex
	session *xorm.Session
}

func (self *DefaultCategory)Init()(err error){

	return
}

func (self *DefaultCategory)Add(category *Category)(id int64,err error){
	//self.mutex.Lock()
	//defer self.mutex.Unlock()
	_,err=self.session.Insert(category)
	return
}
func (self *DefaultCategory)Delete(id int64)(err error){
	_,err=self.session.ID(id).Update(&Category{State:StateDeleted})
	return
}
func (self *DefaultCategory)Update(id int64,category *Category)(err error){
	_,err=self.session.ID(id).AllCols().Update(category)
	return
}
func (self *DefaultCategory)Get(category *Category)(has bool,err error){
	has,err=self.session.Get(category)
	if err != nil {
		return
	}
	return
}

func (self *DefaultCategory)GetChild(categoryType CateType,id int64)(result []*Category,err error){
	session:=self.session.Where("cate_type=?",categoryType).And("state=?",StateOk).OrderBy("list_order")
	if id>0{
		session.And("parent_id=?",id)
	}
	err=session.Find(&session)
	return
}



