package model

/**
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 *
 * user.schema.go
 *
 * @author  m.y <j01428@kok.work>
 * @date
 * @package Model
 * @description
 *
 * 定义schema.go文件的目地:
 *  1.由于业务数据，存在多与复杂结构的情况
 *  2.该文件负责生产端服务层向控制层的结构包装
 *  3.保证代码可读与简洁
 *
 */
import (
	"fmt"
)

type UserSchema struct{}

func (this *UserSchema) Make() {}

func (this *UserSchema) BatchTransForm(uMaps []*Users) (_uMaps map[int]map[string]string) {
	_uMaps = make(map[int]map[string]string)
	for i, u := range uMaps {
		_uMaps[i] = this.TransForm(u)
	}
	fmt.Println("UserSchema.BatchTransForm->", _uMaps)
	return _uMaps
}
func (this *UserSchema) TransForm(u *Users) map[string]string {
	fmt.Println("address:", u.Address)
	return map[string]string{
		"Id":        fmt.Sprintf("%d", u.Id),
		"Name":      u.Name,
		"Addr":      u.Address,
		"City":      u.City,
		"Url":       u.Url,
		"Crc32":     fmt.Sprintf("%d", u.Crc32),
		"Phone":     u.Phone,
		"Email":     u.Email,
		"Introduce": u.Introduce,
	}
}

func (this *UserSchema) MakeUpData(req []interface{}) (_upData map[string]interface{}) {
	_upData = map[string]interface{}{
		"name":    req[1],
		"address": req[2],
	}

	for k, v := range _upData {
		if v.(string) == "" {
			delete(_upData, k)
		}
	}

	return _upData
}

func (this *UserSchema) MakeCreateData(req []interface{}) *Users {
	return &Users{
		Name:      req[0].(string),
		Address:   req[1].(string),
		City:      req[2].(string),
		Url:       req[3].(string),
		Phone:     req[4].(string),
		Email:     req[5].(string),
		Introduce: req[6].(string),
	}
}

/**
*	Id        string `gorm:"id"`
	Name      string `gorm:"name"`
	RegNo     int    `gorm:"regNo"`
	Address   string `gorm:"address"`
	City      string `gorm:"city"`
	Url       string `gorm:"url",default:'galeone'`
	Crc32     int    `gorm:"crc32"`
	Phone     string `gorm:"phone"`
	Email     string `gorm:"email"`
	Status    uint32 `gorm:"status"`
	Introduce string `gorm:"introduce"`
*/
// func (this *UserSchema) MakeCreateData(req []interface{}) (_createData map[string]interface{}) {
// 	_createData = map[string]interface{}{
// 		"name":      req[0],
// 		"address":   req[1],
// 		"city":      req[2],
// 		"url":       req[3],
// 		"phone":     req[4],
// 		"email":     req[5],
// 		"introduce": req[6],
// 	}

// 	for k, v := range _createData {
// 		if v.(string) == "" {
// 			delete(_createData, k)
// 		}
// 	}
// 	return _createData
// }
