package category

import (
	"github.com/go-xorm/xorm"
)

type CategoryService struct {
	session *xorm.Session
}

func NewCategoryService(s *xorm.Session) *CategoryService {
	return &CategoryService{session: s}
}

func (self *CategoryService) Add(category *CategoryEntity) (id int64, err error) {
	_, err = self.session.Insert(category)
	return
}
func (self *CategoryService) Delete(id int64) (err error) {
	_, err = self.session.ID(id).Update(&CategoryEntity{State: StateDeleted})
	return
}
func (self *CategoryService) Update(id int64, category *CategoryEntity) (err error) {
	_, err = self.session.ID(id).Update(category)
	return
}

func (self *CategoryService) GetChild(pid int64) (result []*CategoryEntity, err error) {
	session := self.session.Where("parent_id=?", pid).And("state=?", StateOk).OrderBy("list_order desc")
	err = session.Find(&result)
	if err != nil {
		return
	}
	return
}

/*
func recursiveSearchSubs(session *xorm.Session, cate *CategoryEntity) (result []*CategoryEntity, err error) {

	err = session.Where("cate_owner_id=?", cate.CateOwnerId).And("parent_id=?", cate.Id).Where("cate_type=?", cate.CateType).And("state=?", cate.State).OrderBy("list_order desc").Find(&result)
	if err != nil {
		return
	}
	if len(result) > 0 {
		for _, v := range result {
			v.Child, err = recursiveSearchSubs(session, v)
			if err != nil {
				return
			}
		}
	}
	return

}
func getFullParentInfo(session *xorm.Session, id int64) (result []*BasicCategory, err error) {
	t, err := recursiveSearchParent(session, &Category{Id: id})
	if err != nil {
		return
	}
	if t != nil || len(t) > 0 {
		for _, v := range t {
			result = append(result, &BasicCategory{Id: v.Id, Name: v.Name})
		}
	}
	return
}

func recursiveSearchParent(session *xorm.Session, cate *CategoryEntity) (result []*CategoryEntity, err error) {

	has, err := session.ID(cate.Id).Where("state=?", StateOk).Get(cate)
	if err != nil || !has {
		return
	}

	if cate.ParentId > 0 {
		var t []*CategoryEntity
		t, err = recursiveSearchParent(session, &Category{Id: cate.ParentId})
		if err != nil || !has {
			return
		}
		if t != nil && len(t) > 0 {
			result = append(result, t...)
		}
	}
	result = append(result, cate)
	return
}*/
