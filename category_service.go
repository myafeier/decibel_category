package category

import (
	"fmt"
	"sync"

	"github.com/prometheus/common/log"
	"xorm.io/xorm"
)

var Daemon *CategoryService

type CategoryService struct {
	session *xorm.Session
	Cache   map[int64]*CategoryEntity
	mutex   sync.RWMutex
}

func InitDaemon(s *xorm.Session) {
	if Daemon == nil {
		Daemon = &CategoryService{
			session: s,
			Cache:   make(map[int64]*CategoryEntity),
			mutex:   sync.RWMutex{},
		}
		if err := Daemon.initCache(); err != nil {
			panic(err)
		}
	}

}
func (s *CategoryService) initCache() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	var cates []*CategoryEntity
	if err := s.session.Where("state=?", StateOk).Find(&cates); err != nil {
		log.Error(err.Error())
		return err
	}
	for _, v := range cates {
		s.Cache[v.Id] = v
	}
	return nil
}

func NewCategoryService(s *xorm.Session) *CategoryService {
	return &CategoryService{session: s}
}

func (self *CategoryService) GetFirstLevelCateId(cateId int64) (parentId int64, userErr error, err error) {
	cate := new(CategoryEntity)
	has, err := self.session.ID(cateId).Get(cate)
	if err != nil {
		return
	}
	if !has {
		userErr = fmt.Errorf("id :%d not exist", cateId)
		return
	}
	if cate.Pid > 0 {
		parentId, userErr, err = self.GetFirstLevelCateId(cate.Pid)
	} else {
		parentId = cateId
	}
	return

}
func (self *CategoryService) Add(form *PostForm) (userErr error, err error) {
	if form.Name == "" {
		userErr = fmt.Errorf("栏目名称不可为空")
		return
	}
	category := new(CategoryEntity)
	category.Icon = form.Icon
	category.Pid = form.Pid
	category.Name = form.Name
	category.State = StateOk
	category.ListOrder = form.ListOrder
	_, err = self.session.Insert(category)
	if err != nil {
		log.Error(err.Error())
		return
	}
	Daemon.initCache()
	return
}
func (self *CategoryService) Delete(id int64) (userErr error, err error) {
	if id < 1 {
		userErr = fmt.Errorf("栏目id不可为0")
		return

	}
	_, err = self.session.ID(id).Update(&CategoryEntity{State: StateDeleted})
	if err != nil {
		log.Error(err.Error())
		return
	}

	Daemon.initCache()
	return
}
func (self *CategoryService) Update(id int64, category *CategoryEntity) (userErr error, err error) {
	if id < 1 {
		userErr = fmt.Errorf("栏目id不可为0")
		return
	}
	_, err = self.session.ID(id).Update(category)
	if err != nil {
		log.Error(err.Error())
		return
	}
	Daemon.initCache()
	return
}

func (self *CategoryService) GetChild(pid int64) (result []*CategoryEntity, err error) {
	session := self.session.Where("pid=?", pid).And("state=?", StateOk).OrderBy("list_order desc")
	err = session.Find(&result)
	if err != nil {
		return
	}
	return
}
func (self *CategoryService) GetAll() (result []*CategoryEntity, err error) {
	session := self.session.And("state=?", StateOk).OrderBy("list_order desc")
	err = session.Find(&result)
	if err != nil {
		return
	}
	return
}
func (self *CategoryService) GetListByIds(ids []int64) (result []*CategoryEntity, err error) {
	session := self.session.And("state=?", StateOk).In("id", ids).OrderBy("list_order desc")
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
