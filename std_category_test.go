package decibel_category

import (
	"github.com/go-xorm/xorm"
	"testing"
)
import _ "github.com/go-sql-driver/mysql"

const host= "127.0.0.1"
const port= "3306"
const user= "mall"
const passwd= "Mall3424"
const dbName= "mall"
var session *xorm.Session
func init()  {


	Db, err := xorm.NewEngine("mysql", user+":"+passwd+"@tcp("+host+":"+port+")/"+dbName+"?charset=utf8mb4")
	if err != nil {
		panic("error when connect database!,err:" + err.Error())
	}

	session=Db.NewSession()
	exist,err:=Db.IsTableExist(Category{})
	if err != nil {
		panic( err.Error())
	}
	if !exist{
		Db.CreateTables(Category{})
		Db.CreateIndexes(Category{})
	}else{
		Db.Sync2(Category{})
	}
}

func TestRecursiveSearchParent(t *testing.T)  {

	result,err:=recursiveSearchParent(session,&Category{Id:1})
	if err != nil {
		t.Error(err)
		return
	}
	for _,v:=range result{
		t.Logf("%+v",*v)
	}

}
func TestDefaultCategory_Add(t *testing.T) {
	cate:=&Category{}
	cate.ParentId=3
	cate.Name="细叶"
	cate.State=StateOk
	cate.Icon="http://test.com/1.png"
	cate.CateType=CateType(1)
	cate.CateOwnerId=1

	d:=&DefaultCategory{session:session}
	_,err:=d.Add(cate)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v",*cate)

}

func TestDefaultCategory_Update(t *testing.T) {
	cate:=&Category{}
	cate.Id=6
	d:=&DefaultCategory{session:session}
	has,err:=d.Get(cate)
	if err != nil {
		t.Error(err)
	}
	if !has{
		t.Errorf("category not exist")
	}
	cate.Name="大叶"
	err=d.Update(cate.Id,cate)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v",*cate)
}