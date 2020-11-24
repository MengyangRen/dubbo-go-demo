package processor

import "time"

/**
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 *
 * user.wrap.go
 *
 * @author  m.y <j01428@kok.work>
 * @date
 * @package Model
 * @description
 *
 *  定义wrap.go文件的目地:
 *  1.由于业务数据，存在多与复杂结构的情况
 *  2.该文件负责生产端向消费端结构的包装
 *  3.保证代码可阅读与简洁感
 */

type UserWrap struct{}

func NewUserWrap() *UserWrap {
	return &UserWrap{}
}

func (this *UserWrap) BatchTransForm(uMaps map[int]map[string]string) []*User {
	_uMaps := []*User{}
	for _, item := range uMaps {
		_uMaps = append(_uMaps, this.TransForm(item))
	}
	return _uMaps
}

func (this *UserWrap) TransForm(uMap map[string]string) *User {
	//gxlog.CInfo("uMap:%v", uMap)
	return &User{
		Id:        uMap["Id"],
		Name:      uMap["Name"],
		City:      uMap["City"],
		Age:       18,
		Addr:      uMap["Addr"],
		Phone:     uMap["Phone"],
		Introduce: uMap["Introduce"],
		Time:      time.Now(),
	}
}
