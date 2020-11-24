package service

import (
	"dubbo-go-producer/app/model"
	"strconv"
)

/**
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 *
 * 用户服务层
 *
 * @author  m.y <j01428@kok.work>
 * @date
 * @package service
 * @description
 *
 * 说明:
 *
 */

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (this *UserService) User(req []interface{}) map[string]string {
	return new(model.UserSchema).TransForm(model.NewUsersModel().Find(map[string]string{
		"Id": req[0].(string),
	}))
}

func (this *UserService) Users(req []interface{}) (int64, int64, int, map[int]map[string]string) {
	page, _ := strconv.ParseInt(req[0].(string), 10, 64)
	limit, _ := strconv.ParseInt(req[1].(string), 10, 64)
	count, uMaps := model.NewUsersModel().FindAll(map[string]string{"status": "1"}, page, limit)
	return page, limit, count, new(model.UserSchema).BatchTransForm(uMaps)
}

func (this *UserService) Update(req []interface{}) (err error) {
	id, _ := strconv.ParseInt(req[0].(string), 10, 64)
	if err = model.NewUsersModel().Update(id,
		new(model.UserSchema).MakeUpData(req)); err != nil {
		return
	}
	return
}

func (this *UserService) Destroy(req []interface{}) (err error) {
	id, _ := strconv.ParseInt(req[0].(string), 10, 64)
	if err = model.NewUsersModel().Delete(id); err != nil {
		return
	}
	return
}

func (this *UserService) Store(req []interface{}) (_insterID int, _rows int, err error) {
	if _insterID, _rows, err = model.NewUsersModel().Add(req); err != nil {
		return
	}
	return
}
