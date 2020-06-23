package controllers

import (
	"github.com/astaxie/beego"
	"log"
	"path"
	"strings"
)

type UploadController struct {
	beego.Controller
}



func (c *UploadController) Get() {
	c.Data["Title"] = beego.AppConfig.DefaultString("title","书库")
	c.TplName = "upload.html"
}

func (c *UploadController) Post() {
	var AllowExtMap map[string]bool = map[string]bool{
		".zip": true,
		".gz": true,
	}

	file, information, err  := c.GetFile("filename")
	if err != nil {
		log.Fatal("getfile err ", err)
		return
	}

	if _, ok := AllowExtMap[strings.ToLower(path.Ext(information.Filename))]; !ok {
		c.Ctx.WriteString("后缀名不符合上传要求")
		return
	}

	defer file.Close()
	err = c.SaveToFile("filename", "./upload/" + information.Filename)

	if err != nil {
		c.Ctx.WriteString("File upload failed！")
	} else {
		c.Ctx.WriteString("File upload succeed!")
	}
}