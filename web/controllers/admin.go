package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/liukunxin/superstar-db/model"
	"github.com/liukunxin/superstar-db/services"
	"log"
	"time"
)

type AdminController struct {
	Ctx iris.Context
	Service services.SuperstarService
}

func (c *AdminController) Get() mvc.Result  {
	datalist := c.Service.GetAll()

	// set the model and render the view template.
	return mvc.View{
		Name: "admin/index.html",
		Data: iris.Map{
			"Title":    "管理后台",
			"Datalist": datalist,
		},
		Layout: "admin/layout.html", // 不要跟前端的layout混用
	}
}

func (c *AdminController) GetEdit() mvc.Result  {
	var data *model.StarInfo
	id, err := c.Ctx.URLParamInt("id")
	if err != nil {
		data = c.Service.Get(id)
	}else {
		data = &model.StarInfo{
			Id: 0,
		}
	}
	// set the model and render the view template.
	return mvc.View{
		Name: "admin/edit.html",
		Data: iris.Map{
			"Title": "管理后台",
			"info":  data,
		},
		Layout: "admin/layout.html", // 不要跟前端的layout混用
	}
}

func (c *AdminController) GetDelete() mvc.Result {
	id, err := c.Ctx.URLParamInt("id")
	if err == nil {
		c.Service.Delete(id)
	}
	return mvc.Response{
		Path: "/admin/",
	}
}

func (c *AdminController) PostSave() mvc.Result {
	info := model.StarInfo{}
	err := c.Ctx.ReadForm(&info)
	//fmt.Printf("%v\n", info)
	if err != nil {
		log.Fatal(err)
	}
	if info.Id > 0 {
		info.SysUpdated = int(time.Now().Unix())
		c.Service.Update(&info, []string{"name_zh", "name_en", "avatar",
			"birthday", "height", "weight", "club", "jersy", "coutry",
			"birthaddress", "feature", "moreinfo", "sys_updated"})
	} else {
		info.SysCreated = int(time.Now().Unix())
		c.Service.Create(&info)
	}
	return mvc.Response{
		Path: "/admin/",
	}
}