package service

import (
	"lncios.cn/resume/cs"
	"lncios.cn/resume/model"
	"time"
)

var Message message

type message int

func (message) Create(form *model.Message) error {
	form.BeforeInsert()
	if _, err := cs.Sql.Table("project").Insert(form); err != nil {
		return err
	} else {
		return nil
	}
}
func (message) Updata(form *model.Message) error {
	form.Lut = time.Now()
	if _, err := cs.Sql.Table("project").Update(form); err != nil {
		return err
	}
	return nil
}
func (message) List(page *model.Page, message *model.Message, users *[]model.Message) error {
	if err := cs.Sql.Table("project").Where("dtd=false").Limit(page.Limit(), page.Skip()).Find(users, message); err != nil {
		return err
	}
	return nil
}
func (message) Delete(message *model.Message) error {
	message.Dtd = true
	if _, err := cs.Sql.Table("project").Update(message); err != nil {
		return err
	}
	return nil
}
func (message) Get(message *model.Message) error {

	if _, err := cs.Sql.Table("project").ID(message.Id).Update(message); err != nil {
		return err
	}
	return nil
}
