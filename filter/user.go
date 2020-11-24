package filter

import (
	"cloudlive/exception"
	"regexp"

	"github.com/gin-gonic/gin"
)

/**
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 *
 * 用户过滤层
 *
 * @author  m.y <j01428@kok.work>
 * @date
 * @package filter
 * @description
 *
 * 说明:
 *  1.校验参数
 *  2.过滤参数
 *  3.可与middleware进行配合基于配置做自动化参数校验
 *
 */
type UserFilter struct {
	c *gin.Context
}

func NewUserFilter(c *gin.Context) *UserFilter {
	return &UserFilter{c: c}
}

func (this *UserFilter) Users() map[string]string {

	return map[string]string{
		"page":  this.c.DefaultQuery("page", "1"),
		"limit": this.c.DefaultQuery("limit", "10"),
	}
}

func (this *UserFilter) User() map[string]string {

	this.IsEmpty("id").IsInt("id")
	return map[string]string{
		"id": this.c.Query("id"),
	}
}

func (this *UserFilter) Update() map[string]string {

	this.IsEmpty("id").IsInt("id")
	name := this.c.PostForm("name")
	addr := this.c.PostForm("addr")

	if name == "" && addr == "" {
		exception.NewApiException().Throw(
			exception.EX_REQUEST_INVAL,
			exception.Messages[exception.EX_REQUEST_INVAL],
		)
	}

	return map[string]string{
		"id":   this.c.Query("id"),
		"name": name,
		"addr": addr,
	}
}

func (this *UserFilter) Store() (params []interface{}) {
	checkIndexNum := 0
	params = []interface{}{
		this.c.PostForm("name"),
		this.c.PostForm("addr"),
		this.c.PostForm("city"),
		this.c.PostForm("url"),
		this.c.PostForm("phone"),
		this.c.PostForm("email"),
		this.c.PostForm("introduce"),
	}

	for _, v := range params {
		if v.(string) == "" {
			checkIndexNum++
		}
	}

	if checkIndexNum == len(params) {
		exception.NewApiException().Throw(
			exception.EX_REQUEST_INVAL,
			exception.Messages[exception.EX_REQUEST_INVAL],
		)
	}

	return params
}

func (this *UserFilter) IsInt(key string) *UserFilter {
	_v := this.c.Query(key)
	if res, _ := regexp.MatchString("^(\\d+)$", _v); res != true {
		exception.NewApiException().Throw(
			exception.EX_REQUEST_INVAL,
			exception.Messages[exception.EX_REQUEST_INVAL],
		)
	}
	return this
}

func (this *UserFilter) IsEmpty(key string) *UserFilter {
	_v := this.c.Query(key)
	if _v == "" {
		exception.NewApiException().Throw(
			exception.EX_REQUEST_INVAL,
			exception.Messages[exception.EX_REQUEST_INVAL],
		)
	}
	return this
}
