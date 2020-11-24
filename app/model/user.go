package model

import (
	"fmt"
)

/**
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 *
 * 用户模型层
 *
 * @author  m.y <j01428@kok.work>
 * @date
 * @package model
 * @description
 *
 * 说明:
 *
 */

type Users struct {
	Id        int    `gorm:"id"`
	Name      string `gorm:"name"`
	RegNo     int    `gorm:"reg_no"`
	Address   string `gorm:"address"`
	City      string `gorm:"city"`
	Url       string `gorm:"url",default:'galeone'`
	Crc32     int    `gorm:"crc32"`
	Phone     string `gorm:"phone"`
	Email     string `gorm:"email"`
	Status    uint32 `gorm:"status"`
	Introduce string `gorm:"introduce"`
}

func NewUsersModel() *Users {
	return &Users{}
}

func (this *Users) FindAll(condition map[string]string, page int64, limit int64) (int, []*Users) {
	var count int
	uMaps := []*Users{}
	_query := G.DB.Model(this).Where("status=?", condition["status"])
	_query.Count(&count)
	_query.Offset((page - 1) * limit).Limit(limit).Find(&uMaps)
	return count, uMaps
}

func (this *Users) Find(sMap map[string]string) *Users {
	var u Users
	G.DB.Where("id=?", sMap["Id"]).First(&u)
	fmt.Println("address:", u.Address)
	return &u
}

func (this *Users) Delete(id int64) (err error) {
	err = G.DB.Model(this).Where("id=?", id).Delete(&Users{}).Error
	if err != nil {
		return
	}
	return nil
}

func (this *Users) Update(id int64, upData map[string]interface{}) (err error) {
	err = G.DB.Model(this).Where("id=?", id).Updates(upData).Error
	if err != nil {
		return
	}
	return nil
}

func (this *Users) Add(req []interface{}) (_insterID int, _rows int, err error) {
	Users := new(UserSchema).MakeCreateData(req)
	fmt.Println("_createData:", Users)
	result := G.DB.Table("users").Create(Users)
	err = result.Error
	if err != nil {
		return
	}

	_insterID = Users.Id
	_rows = int(result.RowsAffected)
	return _insterID, _rows, nil
}

// func (this *Users) Add(data map[string]interface{}) (err error) {
// 	fmt.Println("_createData:", data)
// 	err = G.DB.Model(this).Create(data).Error
// 	fmt.Println(err)
// 	if err != nil {
// 		return
// 	}
// 	return nil
// }
