package services

import (
	"github.com/liukunxin/superstar-db/dao"
	"github.com/liukunxin/superstar-db/datasource"
	"github.com/liukunxin/superstar-db/model"
)

type SuperstarService interface {
	GetAll() []model.StarInfo
	Search(country string) []model.StarInfo
	Get(id int) *model.StarInfo
	Delete(id int) error
	Update(user *model.StarInfo, columns []string) error
	Create(user *model.StarInfo) error
}

type superstarService struct {
	dao *dao.SuperstarDao
}

func NewSuperstarService() *superstarService {
	return &superstarService{
		dao : dao.NewSuperstarDao(datasource.InstanceMaster()),
	}
}

func (s *superstarService)GetAll() []model.StarInfo  {
	return s.dao.GetAll()
}

func (s *superstarService)Search(country string) []model.StarInfo  {
	return s.dao.Search(country)
}

func (s *superstarService)Get(id int) *model.StarInfo  {
	return s.dao.Get(id)
}

func (s *superstarService)Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *superstarService)Update(user *model.StarInfo, columns []string) error {
	return s.dao.Update(user, columns)
}
func (s *superstarService)Create(user *model.StarInfo) error {
	return s.dao.Create(user)
}