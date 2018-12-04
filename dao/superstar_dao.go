package dao

import (
	"github.com/go-xorm/xorm"
	"github.com/liukunxin/superstar-db/model"
)

type SuperstarDao struct {
	engine *xorm.Engine
}


func NewSuperstarDao(engine *xorm.Engine) *SuperstarDao {
	return &SuperstarDao{
		engine:engine,
	}
}

func (d *SuperstarDao) Get(id int) *model.StarInfo  {
	data := &model.StarInfo{Id:id}
	if ok, err := d.engine.Get(data); ok && err == nil{
		return data
	}else {
		data.Id = 0
		return data
	}
}

func (d *SuperstarDao) GetAll() []model.StarInfo  {
	datalist := make([]model.StarInfo, 0)
	if err := d.engine.Desc("id").Find(&datalist);err != nil{
		return datalist
	}else {
		return datalist
	}
}

func (d *SuperstarDao) Search(country string) []model.StarInfo {
	datalist := make([]model.StarInfo, 0)
	err := d.engine.Where("country=?", country).
		Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *SuperstarDao) Delete(id int) error {
	data := &model.StarInfo{Id:id, SysStatus:1}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *SuperstarDao) Update(data *model.StarInfo, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *SuperstarDao) Create(data *model.StarInfo) error {
	_, err := d.engine.Insert(data)
	return err
}
