package service

import (
	"lncios.cn/resume/cs"
	"lncios.cn/resume/model"
	"time"
)

var WorkExprience workExprience

type workExprience int

func (workExprience) Create(form *model.WorkExperience) error {
	form.BeforeInsert()
	if _, err := cs.Sql.Insert(form); err != nil {
		return err
	} else {
		return nil
	}
}
func (workExprience) Updata(form *model.WorkExperience) error {
	form.Lut = time.Now()
	if _, err := cs.Sql.Update(form); err != nil {
		return err
	}
	return nil
}
func (workExprience) List(page *model.Page, workExprience *model.WorkExperience, users *[]model.WorkExperience) error {
	if cnt, err := cs.Sql.Where("dtd=false").Limit(page.Limit(), page.Skip()).FindAndCount(users, workExprience); err != nil {
		return err
	} else {
		page.Cnt = cnt
	}
	return nil
}
func (workExprience) Delete(workExprience *model.WorkExperience) error {
	workExprience.Dtd = true
	if _, err := cs.Sql.Update(workExprience); err != nil {
		return err
	}
	return nil
}
func (workExprience) Get(workExprience *model.WorkExperience) error {

	if _, err := cs.Sql.ID(workExprience.Id).Get(workExprience); err != nil {
		return err
	}
	return nil
}
